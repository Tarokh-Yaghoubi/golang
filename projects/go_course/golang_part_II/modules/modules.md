# Modules and Dependency Injection in Golang
## Creating a module
- `go mod init "github.com/username/repo"`
## Adding Dependencies
- `go get github.com/username/repo`
## If you wish you get a special version of a package, this is how to
- `go get github.com/username/repo@v1.2.9`
# `GO.SUM` File
## `go.sum` file serves as a lock file, it stores cryptographic hashes for specific modules version
## the `go.sum` and `go.mod` file together, ensure that same module versions are downloaded everytime
# SOME USEFUL COMMANDS
## list all imported packages in your modules, you can use the `-m` flag to only list the imported modules
- `go list all`
## tidy up your go project, if a module is not being used, it will be removed from your project after this command
## the dependency will be removed from `go.mod` and `go.sum` files, if you remove its `import`, and run this command
## if you import it again, and run this command, it will automatically import it again into `go.mod` and `go.sum`
- `go mod tidy`
