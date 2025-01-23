# Healthchecker

Healthchecker is a tiny tool that checks whether a website is running or is down.

## Features

- Check the status of a website using its domain name and optional port.
- Simple and straightforward CLI interface.

## Installation

To install Healthchecker, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/FredSoares/go-health-checker.git
    ```
2. Navigate to the project directory:
    ```bash
    cd go-health-checker
    ```
3. Install the dependencies:
    ```bash
    go mod tidy
    ```

## Usage

To use Healthchecker, provide the domain name and, optionally, the port number:

```bash
go run main.go --domain example.com [--port 80]
```
or using aliases:
```bash
go run main.go -d example.com [-p 80]
```

### Examples

Check a website on the default port (80):

```bash
go run main.go --domain example.com
```

Check a website on a specific port (e.g., 443):

```bash
go run main.go --domain example.com --port 443
```

## Code Overview

- **main.go** - Contains the CLI setup and main application logic.
- **check.go** - Contains the function that checks the website status.

### Dependencies

- [urfave/cli/v2](https://github.com/urfave/cli/v2): A simple, fast, and fun package for building command line apps in Go.

## License

This project is licensed under the [MIT License](LICENSE).
