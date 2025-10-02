# Golang
- Language created by **Google**.
- Compiled, statically typed, but with the simplicity of Python.

## Uses
- Backend
- Distributed system
- Cli tools
- Microservices

## Commands
go version        # View versión
go run main.go    # Execute file
go build main.go  # Compile
go fmt main.go    # Format code
go mod init miapp # Created módule (package manager)
go get <lib>      # Install library external
go run . # when you are inside the directory where the file is located

# Modules use in local in not internet
1.(install) go mod edit -replace <name_module>=<ubication_file> // go mod edit -replace example.com/greetings=../greetings
2.(update) go mod tiny