# trash-cli-go

[![Test](https://github.com/sivchari/trash-cli-go/actions/workflows/test.yml/badge.svg)](https://github.com/sivchari/trash-cli-go/actions/workflows/test.yml)
[![Release](https://github.com/sivchari/trash-cli-go/actions/workflows/release.yml/badge.svg)](https://github.com/sivchari/trash-cli-go/actions/workflows/release.yml)

A Go reimplementation of [trash-cli](https://github.com/andreafrancia/trash-cli) that follows the FreeDesktop.org trash standard.

## Features

- **trash put** - Move files and directories to trash
- **trash list** - List trashed files  
- **trash restore** - Restore files from trash (interactive UI)
- **trash empty** - Empty the trash (with optional age filter)
- **trash rm** - Remove specific files from trash (pattern matching)
- **trash version** - Show version information

## Installation

### Download prebuilt binaries

Download the latest release from [GitHub Releases](https://github.com/sivchari/trash-cli-go/releases).

### Using Go

```bash
go install github.com/sivchari/trash-cli-go@latest
```

### Build from source

```bash
git clone https://github.com/sivchari/trash-cli-go
cd trash-cli-go
go build -o trash
```


## Usage

### Basic Operations

```bash
# Move files to trash
./trash put file.txt directory/

# List trashed files
./trash list

# Restore files (interactive selection)
./trash restore

# Empty trash completely
./trash empty

# Remove files older than 30 days
./trash empty 30

# Remove specific files from trash
./trash rm "*.txt"
./trash rm "old_file"

# Show version
./trash version

# Get help
./trash --help
./trash [command] --help
```

### Advanced Features

- **Interactive restore** - User-friendly TUI for selecting files to restore
- **Pattern matching** - Use wildcards with `trash rm` command
- **Age-based cleanup** - Remove old files with `trash empty <days>`
- **FreeDesktop.org compliance** - Compatible with Linux desktop environments
- **Cross-platform** - Works on Linux, macOS, and Windows

## Trash Location

Files are stored in `~/.local/share/Trash/` following the FreeDesktop.org standard, making it compatible with desktop environments like GNOME, KDE, and XFCE.

## Performance Comparison

Benchmark results comparing trash-cli-go with the original Python implementation:

| Operation | Original (Python) | Go Implementation | Improvement |
|-----------|------------------|-------------------|-------------|
| **Single file trash** | 0.488s | 0.305s | **37% faster** |
| **100 files trash** | 0.130s | 0.068s | **48% faster** |
| **List trash contents** | 0.377s | 0.031s | **92% faster** |
| **Startup time** | 0.041s | 0.004s | **90% faster** |
| **Memory usage (10MB file)** | 14.2MB | 5.2MB | **63% less** |

### Key Advantages

- **Lightning-fast startup** - No Python interpreter overhead
- **Efficient file operations** - Native system calls
- **Low memory footprint** - Ideal for resource-constrained environments
- **Better performance at scale** - Handles large numbers of files efficiently

## Compatibility

This implementation aims to be compatible with the original Python trash-cli, supporting the same command structure and .trashinfo file format.

## Development

### Prerequisites

- Go 1.24.4 or later
- make (optional, for using Makefile)

### Building

```bash
# Using make
make build

# Or directly with go
go build -o trash
```

### Testing

```bash
# Run tests
make test

# Generate coverage report
make coverage
```

### Running Benchmarks

```bash
make benchmark
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.