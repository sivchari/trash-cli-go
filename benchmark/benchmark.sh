#!/bin/bash

# trash-cli vs trash-cli-go benchmark script

set -e

# Command paths
ORIGINAL_TRASH_PUT="/Users/takuma.shibuya/Library/Python/3.9/bin/trash-put"
ORIGINAL_TRASH_LIST="/Users/takuma.shibuya/Library/Python/3.9/bin/trash-list"
ORIGINAL_TRASH_EMPTY="/Users/takuma.shibuya/Library/Python/3.9/bin/trash-empty"

GO_TRASH="$(dirname "$0")/../trash"
TESTDATA_DIR="$(dirname "$0")/testdata"

# Results file
RESULTS_FILE="$(dirname "$0")/benchmark_results.txt"

echo "trash-cli vs trash-cli-go Benchmark" > "$RESULTS_FILE"
echo "====================================" >> "$RESULTS_FILE"
echo "Date: $(date)" >> "$RESULTS_FILE"
echo "" >> "$RESULTS_FILE"

# Benchmark function
benchmark_command() {
    local name="$1"
    local command="$2"
    echo "Running: $name"
    echo "Command: $command" >> "$RESULTS_FILE"
    
    # Time measurement
    local time_output
    time_output=$( { time eval "$command" > /dev/null 2>&1; } 2>&1 )
    
    echo "$name:" >> "$RESULTS_FILE"
    echo "$time_output" >> "$RESULTS_FILE"
    echo "" >> "$RESULTS_FILE"
    
    echo "  $time_output"
}

# Memory usage measurement function
benchmark_memory() {
    local name="$1"
    local command="$2"
    echo "Memory test: $name"
    
    if command -v /usr/bin/time >/dev/null; then
        local mem_output
        mem_output=$(/usr/bin/time -l sh -c "$command > /dev/null 2>&1" 2>&1 | grep "maximum resident set size")
        echo "Memory - $name:" >> "$RESULTS_FILE"
        echo "$mem_output" >> "$RESULTS_FILE"
        echo "" >> "$RESULTS_FILE"
        echo "  $mem_output"
    else
        echo "  /usr/bin/time not available for memory measurement"
    fi
}

# Test preparation
echo "Preparing benchmark..."
echo "Clearing existing trash..."
"$ORIGINAL_TRASH_EMPTY" > /dev/null 2>&1 || true

# 1. Single file trash test
echo ""
echo "=== Single File Trash Test ==="
echo "Single File Trash Test" >> "$RESULTS_FILE"
echo "=========================" >> "$RESULTS_FILE"

cp "$TESTDATA_DIR/small_1.txt" "/tmp/test_single_original.txt"
cp "$TESTDATA_DIR/small_1.txt" "/tmp/test_single_go.txt"

benchmark_command "Original trash-put (single file)" "$ORIGINAL_TRASH_PUT /tmp/test_single_original.txt"
benchmark_command "Go trash put (single file)" "$GO_TRASH put /tmp/test_single_go.txt"

# 2. Multiple files trash test
echo ""
echo "=== Multiple Files Trash Test ==="
echo "Multiple Files Trash Test" >> "$RESULTS_FILE"
echo "============================" >> "$RESULTS_FILE"

# Copy 100 small files
mkdir -p /tmp/test_original /tmp/test_go
for i in {1..100}; do
    cp "$TESTDATA_DIR/small_${i}.txt" "/tmp/test_original/"
    cp "$TESTDATA_DIR/small_${i}.txt" "/tmp/test_go/"
done

benchmark_command "Original trash-put (100 files)" "$ORIGINAL_TRASH_PUT /tmp/test_original/small_*.txt"
benchmark_command "Go trash put (100 files)" "$GO_TRASH put /tmp/test_go/small_*.txt"

# 3. Trash list test
echo ""
echo "=== Trash List Test ==="
echo "Trash List Test" >> "$RESULTS_FILE"
echo "================" >> "$RESULTS_FILE"

benchmark_command "Original trash-list" "$ORIGINAL_TRASH_LIST"
benchmark_command "Go trash list" "$GO_TRASH list"

# 4. Memory usage test
echo ""
echo "=== Memory Usage Test ==="
echo "Memory Usage Test" >> "$RESULTS_FILE"
echo "==================" >> "$RESULTS_FILE"

# Create new test files
cp "$TESTDATA_DIR/large_1.bin" "/tmp/test_mem_original.bin"
cp "$TESTDATA_DIR/large_1.bin" "/tmp/test_mem_go.bin"

benchmark_memory "Original trash-put (10MB file)" "$ORIGINAL_TRASH_PUT /tmp/test_mem_original.bin"
benchmark_memory "Go trash put (10MB file)" "$GO_TRASH put /tmp/test_mem_go.bin"

# 5. Startup time test
echo ""
echo "=== Startup Time Test ==="
echo "Startup Time Test" >> "$RESULTS_FILE"
echo "==================" >> "$RESULTS_FILE"

benchmark_command "Original trash-put --help" "$ORIGINAL_TRASH_PUT --help"
benchmark_command "Go trash put --help" "$GO_TRASH put --help"

echo ""
echo "Benchmark completed! Results saved to: $RESULTS_FILE"

# Cleanup test data
echo ""
echo "Cleaning up test data..."
rm -rf "$TESTDATA_DIR"
rm -rf /tmp/test_original /tmp/test_go
rm -f /tmp/test_single_*.txt /tmp/test_mem_*.bin
echo "Test data cleaned up."

echo ""
echo "Summary:"
cat "$RESULTS_FILE"

# Generate formatted markdown table
echo ""
echo "Generating formatted results..."
"$(dirname "$0")/format_results.sh" "$RESULTS_FILE"