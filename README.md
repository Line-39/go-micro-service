# Simple microservice with Golang

## Setup

### Clone a template project from GitHub

Follow this [link](https://github.com/Line-39/go-micro-service.git) and copy the project url. Go back to your terminal window and execute following command:

```Bash
git clone https://github.com/Line-39/go-micro-service.git
cd go-micro-service
ls -ahl
```

If everything went right, you will see following output:

```Text
total 52K
drwxrwxr-x 3 ubot ubot 4.0K Apr  3 17:11 .
drwxrwxr-x 3 ubot ubot 4.0K Apr  3 17:00 ..
drwxrwxr-x 8 ubot ubot 4.0K Apr  3 17:10 .git
-rw-rw-r-- 1 ubot ubot  478 Apr  3 17:00 .gitignore
-rw-rw-r-- 1 ubot ubot  35K Apr  3 17:00 LICENSE
-rw-rw-r-- 1 ubot ubot 1.4K Apr  3 17:37 README.md
```

If you the output differs from what you see above, make sure you are on the branch `00-setup`, and switch to this branch if required:

```Bash
git switch 00-setup
```

### Alternative: create your own folder
In your *working directory* create the project folder, and change into it:

```Bash
mkdir go-micro-service
cd go-micro-service
```

Initialize *git repository*:

```Bash
git init
```

Add `README.md` file, and commit your changes:

```Bash
echo -e "# Simple microservice with Golang\nAdd your description..." > README.md 
git add --all
git commit -m 'Initializing repository'
git switch -b 00-setup
```

Create and / or edit your `LICENSE` and `.gitignore` files accordingly.

### Next chapter
Switch to the next chapter [Installation](#installation) of the introduction:

```Bash
git switch 01-installation
```

## Installation

### Prerequests
If you are working with clonned project, switch to the next branch (`01-installation`):
```Bash
git switch 01-installation
```

### Instructions
If not yet installed, follow the [official guide](https://go.dev/doc/install) to install Go on your system.

Verify the installation:
```Bash
go version
# go version go1.22.1 linux/amd64
```

You are ready to Go!

## First programm

### Prerequests
If you are working with clonned project, switch to the next branch (`02-first-programm`):
```Bash
git switch 02-first-programm
```

### Intro
#### Package, module, repository
Go *package* is a collection of functions, types, variables and constants defined in source files located at the same directory, functionally related and *visible* to each other.

Go *module* is a collection of one or more related *packages*. The `go.mod` located in the module directory, defines the *paths* for all packages used by module.

Go *repository* is composed from the different modules and can be compiled into the *application* providing the required functionality.

Read more about Go code organisation [here](https://go.dev/doc/code).

#### `go.mod` file
`go.mod` contains all the paths for your module. The naming convention for the modules, requires name of the module to be composed of the name of your organisation plus module parent directory plus module name. E.g.: `github.com/Line-39/go-micro-service`. Note that it is not required, and it is not required for the module to be published online - if you do not follow this convention, or you do not sotre your code at public repository, the module will be still available locally for other modules within your *workspace*.

#### Main function
The `main` *package* in Go contains `main()` function, which serves as an entry point for the programm. It takes no arguments, returns no values and is not called directly:

```Go
package main          // package declaration

import "fmt"          // imports

func main() {         // main() declaration
    /*
      the magick lives here 🪄🪄🪄
    */
}
```

### Instructions
#### Create a `main.go` function
In the project directory, create a `main.go` function and open it in your online editor.

```Bash
touch main.go
# vim main.go   # if you ran this command - unplag you PC and reboot
code main.go
```

#### Write your first programm
Add following lines to the file you've just opened:

```Go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world 🪄🪄🪄")
}
```

If your are working in VScode and you have Go extensions installed, you will notice VSCode complains that no packages found for the `main`. Ignore it for now.

#### Run the programm (first attempt)
From you project directory, run folowwing command from the terminal:

```Bash
go run .
```

You will see following output:

```Text
current directory outside modules listed in go.work or their selected dependencies
```

There are two particular problems that has to be solved in this case.

#### Initialize your module
We are about to write a Go module, to combine useful functionality from different packages in order to create a tool we need. In order to do so, we need to specify the dependencies and the version of Go we are using for this specific module. As scary as it sounds, in practice we only need to run following command from the terminal in our projects directory:

```bash
go mod init github.com/line-39/go-microservice
# go: creating new go.mod: module github.com/line-39/go-microservice
# go: to add module requirements and sums: go mod tidy
```

This command creates a `go.mod` file in the current directory, and collects all the required information for us. Read more about modules [here](https://go.dev/doc/tutorial/create-module).

```Bash
ls .
# go.mod  LICENSE  main.go  README.md

cat go.mod
# module github.com/line-39/go-microservice
#
# go 1.22.1
```

#### Run the programm (second attempt)
We can attempt to run the programm again:
```Bash
go run .
# current directory is contained in a module that is not one of the workspace modules listed in go.work. You can add # the module to the workspace using:
#        go work use .
```

#### Go workspace
In a nutshell Go complains that the module we are trying to run is not associated with the *go workspace*. Go workspace is group of modules / packages defined in `go.work` file which should be located in the directory containing the *modules*. For the workspace initialization we use `go work init <directory to use>` command. You can read more about it [here](https://go.dev/doc/tutorial/workspaces).

Let say your Go code is organized this way:

```Text
go
└── src
    ├── github.com
    │   ├── line-39
    │   └── utubun
    └── seqquery.de
        └── s3
```
And you are going to use all the modules located under `github.com/...` and `sequery.de` dirs. Than go to the parent directory containing all of your modules (in this specific case `src`) and initialize go workspace with following command:

```Bash
go work init .
```

If you `tree` your parent directory, after that, you will get this output:

```Text
go
└── src
    ├── github.com
    │   ├── line-39
    │   └── utubun
    ├── go.code-workspace
    ├── go.work
    ├── go.work.sum
    └── seqquery.de
        └── s3
```

In case you created your module within the directory with *initialized* `go.work`, assuming you are located in the *project directory* just add your module to the workspace as shown below:

```Bash
go work use .
```

#### Run the programm (third attempt)
Let's run our program one more time:
```Bash
go run .
# Hello, world 🪄🪄🪄
```

Now, the magick finally works 🪄🪄🪄

#### Building your program
The `go run` compiles your programm on background and run it. You can build it yourself with `go build` command, which will produce *executable* file. Build and run your programm from the terminal:

```Bash
go build .
ls
# go-microservice  go.mod  LICENSE  main.go  README.md
chmod 776 go-microservice
./go-microservice
# Hello, world 🪄🪄🪄
rm go-microservice
```