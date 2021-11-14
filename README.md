go mod:\
The `go mod init` command initializes and writes a new go.mod file in the current directory, in effect creating a new module rooted at the current directory.

go mod tidy:\
The `go mod tidy` ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

go mod vendor:\
The `go mod vendor` command constructs a directory named vendor in the main module’s root directory that contains copies of all packages needed to support builds and tests of packages in the main module. When vendoring is enabled, the go command will load packages from the vendor directory instead of downloading modules from their sources into the module cache and using packages those downloaded copies.

go get <dependency_name>:\
The `go get` command add dependencies to current module and install them.

go run <file_name>:\
The `go run` command compile and run Go program
