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

const helpMessage = `
Usage: run_bench [OPTIONS]

Options:
  -help                           Show this help message and exit.

  -cooldown=value                 Set the cooldown period. Valid values:
                                  - "disabled" or "0" to disable cooldown.
                                  - An integer between 1 and 300 (inclusive).
                                  Default: 20 seconds.

  -benchtime=value                Set the benchmark time duration. Valid value:
                                  - An integer between 1 and 30 (inclusive).
                                  Default: 5 seconds.

  -save=value                     Specify the format to save benchmark results. Valid values:
                                  - "json" to save as JSON format.
                                  - "csv" to save as CSV format.
                                  Default: json.

  -wd=value,                      Set the working directory. Valid value:
  -workingdirectory=value         - An absolute path to a directory that exists.
                                  Default: project working directory.

Example:
  go run run_bench.go -cooldown=10 -benchtime=5 -save=csv -wd=/absolute/path/to/directory

  or

  go run run_bench.go (it will run with default settings)`

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

type Parameters struct {
	CooldownFunc     func()
	Benchtime        uint
	SaveFunc         func([]*BenchmarkResult, string) error
	WorkingDirectory string
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

func generateCommand(projectDirectory, testFilePath string, benchtimeAmount uint) *exec.Cmd {
	benchtime := fmt.Sprintf("-benchtime=%ds", benchtimeAmount)
	cmd := exec.Command("go", "test", "-bench=.", benchtime, testFilePath)
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

func saveResultsAsJSON(results []*BenchmarkResult, fileName string) error {
	fileName = fmt.Sprintf("%s.json", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(results); err != nil {
		return fmt.Errorf("failed to encode results to JSON: %w", err)
	}
	return nil
}

func saveResultsAsCSV(results []*BenchmarkResult, fileName string) error {
	fileName = fmt.Sprintf("%s.csv", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	defer file.Close()

	var sb strings.Builder
	sb.WriteString("Name;Score;ns/op\n")
	for _, result := range results {
		for _, run := range result.Runs {
			fmt.Fprintf(&sb, "%s;%.1f;%.1f\n", run.Name, run.Score, run.NsPerOp)
		}
	}
	_, err = file.WriteString(sb.String())
	if err != nil {
		return fmt.Errorf("failed to write results to CSV: %w", err)
	}
	return nil
}

func formatTimeAsFileName(t time.Time) string {
	return t.Format("2006-01-02_15-04-05")
}

func formatTimeAsStamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func runBenchmarkSuite(parameters Parameters) error {
	projectDirectory := parameters.WorkingDirectory
	benchmarkTestFiles, err := findBenchmarkTestFiles(projectDirectory)
	if err != nil {
		return fmt.Errorf("could not find benchmark test files: %w", err)
	}
	results := make([]*BenchmarkResult, 0)
	startTime := time.Now()
	fmt.Printf("-- Starting Benchmark Suite at %s\n", formatTimeAsStamp(startTime))
	for _, testFile := range benchmarkTestFiles {
		fmt.Printf("%s Running benchmark for %s\n", formatTimeAsStamp(time.Now()), testFile)
		cmd := generateCommand(projectDirectory, testFile, parameters.Benchtime)
		result, err := runCommand(cmd)
		if err != nil {
			fmt.Printf("Error! Failed to run benchmark for %s. err: %s", testFile, err.Error())
		}
		results = append(results, result)
		parameters.CooldownFunc()
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("-- Benchmark Suite completed at %s\n", formatTimeAsStamp(endTime))
	fmt.Printf("-- Benchmarks ran for %.2fs\n", elapsedTime.Seconds())
	resultFileName := fmt.Sprintf("benchmark_results_%s", formatTimeAsFileName(startTime))
	if err = parameters.SaveFunc(results, resultFileName); err != nil {
		return fmt.Errorf("could not save benchmark results to file %s: %w", resultFileName, err)
	}
	return nil
}

func printHelpAndExit() {
	fmt.Println(helpMessage)
	os.Exit(0)
}

func isHelp(arg string) bool {
	arg = strings.ToLower(arg)
	return arg == "help" || arg == "-help" || arg == "--help"
}

func cooldownArg(value string) (func(), error) {
	switch value {
	case "disabled", "0":
		return func() {}, nil
	}
	cooldownAmount, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("illegal value for -cooldown parameter")
	}
	if cooldownAmount < 1 {
		return nil, fmt.Errorf("cooldown could not be a negative number")
	}
	if cooldownAmount > 300 {
		return nil, fmt.Errorf("cooldown could not be greater than 300")
	}
	return func() {
		time.Sleep(time.Duration(cooldownAmount) * time.Second)
	}, nil
}

func benchtimeArg(value string) (uint, error) {
	benchtime, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("illegal value for -benchtime parameter")
	}
	if benchtime < 1 {
		return 0, fmt.Errorf("benchtime must be a possitive number")
	}
	if benchtime > 30 {
		return 0, fmt.Errorf("benchtime could not be greater than 30")
	}
	return uint(benchtime), nil
}

func saveArg(value string) (func([]*BenchmarkResult, string) error, error) {
	switch value {
	case "json":
		return saveResultsAsJSON, nil
	case "csv":
		return saveResultsAsCSV, nil
	}
	return nil, fmt.Errorf("illegal value for -save parameter")
}

func workingDirectoryArg(path string) (string, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("%s does not exists on the system", path)
	}
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("%s not a directory path", path)
	}
	if !filepath.IsAbs(path) {
		return "", fmt.Errorf("path must be absolute and start with /")
	}
	return path, nil
}

func ensureResult[T any](result T, err error) T {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		printHelpAndExit()
	}
	return result
}

func normalizeString(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

func parseParameters(args []string) Parameters {
	projectDirectory, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("could not read project directory: %w", err))
	}
	parameters := Parameters{
		CooldownFunc:     func() { time.Sleep(20 * time.Second) },
		Benchtime:        5,
		SaveFunc:         saveResultsAsJSON,
		WorkingDirectory: projectDirectory,
	}
	for _, arg := range args {
		if isHelp(arg) {
			printHelpAndExit()
		}
		// arg format: -key=value
		if arg[0] != '-' {
			fmt.Printf("Error: All arguments must start with a single hyphen (-)\n")
			printHelpAndExit()
		}
		keyValue := strings.Split(arg[1:], "=")
		if len(keyValue) != 2 {
			fmt.Printf("Error: Arguments must be in the format -key=value\n")
			printHelpAndExit()
		}
		key, value := normalizeString(keyValue[0]), normalizeString(keyValue[1])
		switch key {
		case "cooldown":
			parameters.CooldownFunc = ensureResult(cooldownArg(value))
		case "benchtime":
			parameters.Benchtime = ensureResult(benchtimeArg(value))
		case "save":
			parameters.SaveFunc = ensureResult(saveArg(value))
		case "wd", "workingdirectory":
			parameters.WorkingDirectory = ensureResult(workingDirectoryArg(value))
		default:
			fmt.Printf("Error: unknown parameter <%s>. please read the help message 🙏\n", key)
			printHelpAndExit()
		}
		// fmt.Printf("[DEBUG] key: %s, value: %s\n", key, value)
	}
	return parameters
}

func main() {
	// -help
	// -cooldown = disable, number[1,300] default 20
	// -benchtime = number[1, 30] default 5
	// -save = json, csv default json
	// -wd, -workingdirectory = string directory default project working directory later

	parameters := parseParameters(os.Args[1:])
	if err := runBenchmarkSuite(parameters); err != nil {
		panic(err)
	}
}
