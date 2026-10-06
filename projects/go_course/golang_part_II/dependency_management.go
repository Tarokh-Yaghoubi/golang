
package main

// GO_MODULES
// go modules are used for managing dependencies in your go projects
// modules are packages that group similar code together
// Go modules use go.mod and go.sum files to track dependencies of a GO project
// these files are located in the root dir of the project
// go supports multiple modules
// One go.mod or go.sum is required in each dir
// modules are also used for version control
// if you multiple modules in one repo, managing the versions of the modules will be cumbersome

// to create a new go module, you can use the `go mod init` command
// this command is used to initialize a new go module
// name of the module is as such if you are using github:
// "github.com/username/repo", if u are using bitbucket, it will be different

// U can keep modules private, and still allow go to fetch it

// we can add dependencies to the module using `go get`
// by default, go will pull the latest version of the module
// you can specify the version explicitly if u wish to use a specific version

