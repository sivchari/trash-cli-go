#!/bin/bash

# Test data creation script for benchmarking

set -e

TESTDATA_DIR="$(dirname "$0")/testdata"
rm -rf "$TESTDATA_DIR"
mkdir -p "$TESTDATA_DIR"

echo "Creating test files..."

# Small files (1KB x 1000)
echo "Creating 1000 small files (1KB each)..."
for i in {1..1000}; do
    head -c 1024 /dev/urandom > "$TESTDATA_DIR/small_${i}.txt"
done

# Medium files (100KB x 100)
echo "Creating 100 medium files (100KB each)..."
for i in {1..100}; do
    head -c 102400 /dev/urandom > "$TESTDATA_DIR/medium_${i}.dat"
done

# Large files (10MB x 10)
echo "Creating 10 large files (10MB each)..."
for i in {1..10}; do
    head -c 10485760 /dev/urandom > "$TESTDATA_DIR/large_${i}.bin"
done

# Directory structure
echo "Creating directory structure..."
for i in {1..50}; do
    mkdir -p "$TESTDATA_DIR/dir_${i}"
    for j in {1..10}; do
        echo "content_${i}_${j}" > "$TESTDATA_DIR/dir_${i}/file_${j}.txt"
    done
done

echo "Test data created successfully!"
echo "Total files: $(find "$TESTDATA_DIR" -type f | wc -l)"
echo "Total size: $(du -sh "$TESTDATA_DIR" | cut -f1)"