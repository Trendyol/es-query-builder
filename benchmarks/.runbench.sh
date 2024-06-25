#!/bin/bash

benchmark_test() {
    local file=$1
    echo "Running benchmark for $file"
    go test -bench=. $file -benchtime=5s
}

test_files=$(find . -type f -name "*_test.go")

for file in $test_files;
do
    benchmark_test $file
    echo "Waiting for 30 seconds before the next test..."
    sleep 30
done

echo "All benchmarks completed."
