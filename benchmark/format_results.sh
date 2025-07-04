#!/bin/bash

# Script to format benchmark results into a markdown table

RESULTS_FILE="${1:-benchmark_results.txt}"

if [ ! -f "$RESULTS_FILE" ]; then
    echo "Results file not found: $RESULTS_FILE"
    exit 1
fi

# Extract values from benchmark results
# Single file test
ORIG_SINGLE=$(grep -A3 "Original trash-put (single file):" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')
GO_SINGLE=$(grep -A3 "Go trash put (single file):" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')

# Multiple files test
ORIG_MULTI=$(grep -A3 "Original trash-put (100 files):" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')
GO_MULTI=$(grep -A3 "Go trash put (100 files):" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')

# List test
ORIG_LIST=$(grep -A3 "Original trash-list:" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')
GO_LIST=$(grep -A3 "Go trash list:" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')

# Startup test
ORIG_STARTUP=$(grep -A3 "Original trash-put --help:" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')
GO_STARTUP=$(grep -A3 "Go trash put --help:" "$RESULTS_FILE" | grep "real" | awk '{print $2}' | sed 's/0m//' | sed 's/s//')

# Memory test (convert to MB)
ORIG_MEM=$(grep -A1 "Memory - Original trash-put" "$RESULTS_FILE" | grep "maximum" | awk '{print $1}')
GO_MEM=$(grep -A1 "Memory - Go trash put" "$RESULTS_FILE" | grep "maximum" | awk '{print $1}')
ORIG_MEM_MB=$(echo "scale=1; $ORIG_MEM / 1048576" | bc)
GO_MEM_MB=$(echo "scale=1; $GO_MEM / 1048576" | bc)

# Calculate improvements
SINGLE_IMP=$(echo "scale=0; (1 - $GO_SINGLE / $ORIG_SINGLE) * 100" | bc)
MULTI_IMP=$(echo "scale=0; (1 - $GO_MULTI / $ORIG_MULTI) * 100" | bc)
LIST_IMP=$(echo "scale=0; (1 - $GO_LIST / $ORIG_LIST) * 100" | bc)
STARTUP_IMP=$(echo "scale=0; (1 - $GO_STARTUP / $ORIG_STARTUP) * 100" | bc)
MEM_IMP=$(echo "scale=0; (1 - $GO_MEM / $ORIG_MEM) * 100" | bc)

# Output markdown table
echo "## Performance Comparison"
echo ""
echo "| Operation | Original (Python) | Go Implementation | Improvement |"
echo "|-----------|------------------|-------------------|-------------|"
echo "| **Single file trash** | ${ORIG_SINGLE}s | ${GO_SINGLE}s | **${SINGLE_IMP}% faster** |"
echo "| **100 files trash** | ${ORIG_MULTI}s | ${GO_MULTI}s | **${MULTI_IMP}% faster** |"
echo "| **List trash contents** | ${ORIG_LIST}s | ${GO_LIST}s | **${LIST_IMP}% faster** |"
echo "| **Startup time** | ${ORIG_STARTUP}s | ${GO_STARTUP}s | **${STARTUP_IMP}% faster** |"
echo "| **Memory usage (10MB file)** | ${ORIG_MEM_MB}MB | ${GO_MEM_MB}MB | **${MEM_IMP}% less** |"