package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Run struct {
	Name    string  `json:"name"`
	Score   float64 `json:"score"`
	NsPerOp float64 `json:"nsop"`
}

type BenchmarkResult struct {
	Os      string  `json:"os"`
	Arch    string  `json:"arch"`
	Runs    []Run   `json:"runs"`
	RunTime float64 `json:"runTime"`
}

func parseBenchmarkOutput(benchmarkOutput string) (*BenchmarkResult, error) {
	benchmarkRegex := regexp.MustCompile(`goos:\s+(\w+)\ngoarch:\s+([\w]+)(.*\n)+PASS\nok\s*[a-z \-]+\s*(\d+(\.\d+)?)s`)
	matches := benchmarkRegex.FindStringSubmatch(benchmarkOutput)
	goos, goarch, runTime := "", "", 0.0
	if len(matches) > 0 {
		goos = matches[1]
		goarch = matches[2]
		runTimeSeconds, err := strconv.ParseFloat(matches[4], 64)
		if err != nil {
			return nil, err
		}
		runTime = runTimeSeconds
	}

	benchmarkRunRegex := regexp.MustCompile(`^(Benchmark\w+)-\d+\s+(\d+(\.\d+)?)\s+(\d+(\.\d+)?) ns/op`)
	lines := strings.Split(benchmarkOutput, "\n")
	runs := make([]Run, 0)
	for _, line := range lines {
		matches = benchmarkRunRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			score, err := strconv.ParseFloat(matches[2], 64)
			if err != nil {
				return nil, err
			}
			nsPerOp, err := strconv.ParseFloat(matches[4], 64)
			if err != nil {
				return nil, err
			}
			runs = append(runs, Run{
				Name:    matches[1],
				Score:   score,
				NsPerOp: nsPerOp,
			})
		}
	}

	return &BenchmarkResult{
		Os:      goos,
		Arch:    goarch,
		Runs:    runs,
		RunTime: runTime,
	}, nil
}

func generateCommand(projectDirectory, testFilePath string) *exec.Cmd {
	cmd := exec.Command("go", "test", "-bench=.", "-benchtime=5s", testFilePath)
	cmd.Dir = projectDirectory
	return cmd
}

func runCommand(cmd *exec.Cmd) (*BenchmarkResult, error) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("#combinedOutput - output: %s, err: %w", string(output), err)
	}
	return parseBenchmarkOutput(string(output))
}

func findBenchmarkTestFiles(projectDirectory string) ([]string, error) {
	var benchmarkFiles []string
	if err := filepath.Walk(projectDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			if matched, _ := filepath.Match("*_benchmark_test.go", info.Name()); matched {
				relativePath, _ := filepath.Rel(projectDirectory, path)
				benchmarkFiles = append(benchmarkFiles, relativePath)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return benchmarkFiles, nil
}

func saveResults(results []*BenchmarkResult, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		panic(fmt.Errorf("failed to create file %s: %w", fileName, err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(results); err != nil {
		panic(fmt.Errorf("failed to encode results to JSON: %w", err))
	}
	return nil
}

func formatTimeAsFileName(t time.Time) string {
	return t.Format("2006-01-02_15-04-05")
}

func formatTimeAsStamp(t time.Time) string {
	return t.Format("2006-01-02 03:04:05")
}

func main() {
	projectDirectory, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("could not read project directory: %w", err))
	}
	benchmarkTestFiles, err := findBenchmarkTestFiles(projectDirectory)
	if err != nil {
		panic(fmt.Errorf("could not find benchmark test files: %w", err))
	}
	results := make([]*BenchmarkResult, 0)
	startTime := time.Now()
	fmt.Printf("-- Benchmark Suite started at %s\n", formatTimeAsStamp(startTime))
	for _, testFile := range benchmarkTestFiles {
		fmt.Printf("%s Running benchmark for %s\n", formatTimeAsStamp(time.Now()), testFile)
		cmd := generateCommand(projectDirectory, testFile)
		result, err := runCommand(cmd)
		if err != nil {
			fmt.Printf("Error running benchmark for %s. err: %s", testFile, err.Error())
		}
		results = append(results, result)
		time.Sleep(20 * time.Second)
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("-- Benchmark Suite finisihed at %s\n", formatTimeAsStamp(endTime))
	fmt.Printf("-- Benchmarks run for %.2fs", elapsedTime.Seconds())
	resultFileName := fmt.Sprintf("benchmark_results_%s.json", formatTimeAsFileName(startTime))
	if err = saveResults(results, resultFileName); err != nil {
		panic(fmt.Errorf("could not save benchmark results to file %s: %w", resultFileName, err))
	}
}
