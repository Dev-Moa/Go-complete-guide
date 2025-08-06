# Go course
- Getting Started
- Essentails (basic programming)
- Packages
- Pointers
- Structs
- Interfaces
- Arrays,slices,maps
- Functions deep dive
- Practice project
- Concurrency
- Main project: Rest API

# Getting Started

- Go is opensource programming language developed by google.

- Focus on simplicity clarity and scalability
- High performance & focus on concurency
- Batteries included
- Static Typing

- Usecases :
    - Networking & APIs
    - Microservices
    - CLI tools

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello world")
}

```

# Essentials 
- key component of go program
- working with values and types
- functions
- control structures

# key component of go program

- Every Go file must start with a package clause.
- This organizes your code.
- package main is special — it marks the entry point of an executable Go program.
- You can have multiple .go files in the same package.
- import is used to include code from other packages.
- For example, import "fmt" pulls in the formatting package from Go’s standard library.
- func main() is the starting point of execution.
- Go does not start running code from the top of the file — it starts by calling main().
- There can only be one main() function per main package.
- A module is a bigger structure than a package.
- It’s like your whole Go project.
- You create one with go mod init [module-name].
- This generates a go.mod file that marks the folder as a Go module.
- To build an executable:
- Use go build — this compiles the program into a runnable file (e.g. .exe on Windows).
- This is how you share Go apps without needing Go installed on other systems.