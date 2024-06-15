# Directory Mirror

## Description

Mirrors the contents of a directory to another directory. This is a simple
utility meant to copy the contents from one directory to another, and to
remove any files that are not present in the source directory from the destination
directory.

This utility was created to copy music files from a computer to a SD card used
in a DAP(Digital Audio Player). Meant to break the dependency on a music player
application to manage the music library. So it is relatively simple and does
not have many features.

## Installation

### Prerequisites

Have Go installed on your system. You can download it from the [official website](https://golang.org/).

### Build

Build the executable with:

```bash
go build -o dir-mirror main.go
```

Place the executable in a directory that is in your `PATH`.

## Usage

### Basic Usage

Build and call the executable with the source and destination directories as:

```bash
./dir-mirror --source="/Abs/Src/Path/" --destination="/Abs/Dest/Path/"
```

`--source` and `--destination` are required flags and must be absolute paths.
Both directories must exist and the source directory must be readable while
the destination directory must be writable.

### Additional Flags

```bash
--include-hidden [True/False]
```

Copy hidden files and directories as well. Default is `False`.

```bash
--copy-existing [True/False]
```

Copy files that already exist in the destination directory. Default is `False`.

```bash
--delete-missing [True/False]
```

Delete files and directories that are missing in the source directory from the
destination directory. Default is `False`.

## Caution

Not responsible for any data loss or corruption. This tool is meant to be used
in non-critical situations. It is recommended that you backup the destination directory.

The utility has only been tested for copying directories/files from a computer to
a SD card used in a DAP(Digital Audio Player), so a relatively risk free use-case.

## Contributing

If you find any bugs or have any suggestions, please open a pull request or an issue.
