#!/bin/bash

benchmark_test() {
    local file=$1
    echo "$(date +"%Y-%m-%d %H:%M:%S") Running benchmark for $file"
    go test -bench=. -benchtime=5s "$file" 2>&1 | tee "$file.log"
    if [ $? -ne 0 ]; then
        echo "Error running benchmark for $file. Check $file.log for details."
    fi
}

test_files=$(find . -name "*_test.go" 2>/dev/null)

if [ $? -eq 0 ]; then
    start_time=$(date +%s)
    for file in $test_files;
    do
        benchmark_test "$file"
        sleep 30
    done
    end_time=$(date +%s)
    elapsed_time=$((end_time - start_time))
    echo "All benchmarks completed in $elapsed_time seconds."
else
    echo "Error finding test files."
fi