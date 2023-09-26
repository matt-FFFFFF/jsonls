# jsonls

Simple utility to get current directory listing as JSON array.

Can be used to generate matrix builds for GitHub Actions.

## Installation

```bash
go install github.com/matt-FFFFFF/jsonls@latest
```

## Usage

```bash
$ jsonls -h
Provides a JSON representation of the directory listing.

Usage:
  jsonls [flags]

Flags:
  -d, --dirs   List directories only
  -h, --help   Display help
```
