# GoForge

Quickly bootstrap Go projects with your preferred framework.

---

## Features
- Fast project creation: `goforge new <project-name>`  
- Framework selection: Gin and more...
- Embed starter templates  
- Automatic `go.mod` generation  
- Single-binary installation via `go install` 

---

## Installation

```bash
# Install via Go
go install github.com/coderianx/goforge/cmd/goforge@latest
```

Make sure `$GOPATH/bin` or `$HOME/go/bin` is in your PATH

---

## Usage

# Create a new project
```bash
goforge new my-app
```

## Navigate into project
```bash
cd my-api
```

# Run the project
```bash
go run .
```

---

## Supported Frameworks
- Gin

---

## License
[MIT](./LICENSE)