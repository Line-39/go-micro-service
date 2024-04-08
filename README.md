---
marp: true
theme: gaia
backgroundColor: #fff
title: Simple microservice with Go
---

# Setup

To clone the project from GitHub, follow this [link](https://github.com/Line-39/go-micro-service.git) and copy the project url. Go back to your terminal window and execute following command:

```Bash
git clone https://github.com/Line-39/go-micro-service.git
cd go-micro-service
ls -ahl
```

---

# Project structure
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

---

# Project structure
If you the output differs from what you see above, make sure you are on the branch `00-setup`, and switch to this branch if required:

```Bash
git switch 00-setup
```

---

# Alternative: create your own folder
In your *working directory* create the project folder, and change into it:

```Bash
mkdir go-micro-service
cd go-micro-service
```

Initialize *git repository*, add `README.md`:

```Bash
git init
echo -e "# Simple microservice with Golang\nAdd your description..." > README.md 
```

---

# Alternative: commit your changes
If you are working alone, on manually created project, commit your changes as shown below:

```Bash
git add --all
git commit -m 'Initializing repository'
git switch -b 00-setup
```

---

## Installation: Prerequests
If you are working with clonned project, switch to the next branch (`01-installation`):
```Bash
git switch 01-installation
```

---

### Installation: Instructions
If not yet installed, follow the [official guide](https://go.dev/doc/install) to install Go on your system.

Verify the installation:
```Bash
go version
# go version go1.22.1 linux/amd64
```

You are ready to Go!

---

# Setup: Prerequests

If you are working with clonned project, switch to the next branch (`02-first-programm`):
```Bash
git switch 02-first-programm
```

---

# Setup: Package, module, repository
Go *package* is a collection of functions, types, variables and constants defined in source files located at the same directory, functionally related and *visible* to each other.

Go *module* is a collection of one or more related *packages*. The `go.mod` located in the module directory, defines the *paths* for all packages used by module.

Go *repository* is composed from the different modules and can be compiled into the *application* providing the required functionality.

Read more about Go code organisation [here](https://go.dev/doc/code).

---

# `go.mod` file

`go.mod` contains all the paths for your module. The naming convention for the modules, requires name of the module to be composed of the name of your organisation plus module parent directory plus module name. E.g.: `github.com/Line-39/go-micro-service`. Note that it is not required, and it is not required for the module to be published online - if you do not follow this convention, or you do not sotre your code at public repository, the module will be still available locally for other modules within your *workspace*.

---

# Main function
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
---

# Create a `main.go` file
In the project directory, create a `main.go` function and open it in your online editor.

```Bash
touch main.go
# vim main.go   # if you ran this command - unplag you PC and reboot
code main.go
```

---

# Create `main()` function
Add following lines to the file you've just opened:

```Go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world 🪄🪄🪄")
}
```

---

# VSCode warnings

If your are working in VScode and you have Go extensions installed, you will notice VSCode complains that no packages found for the `main`. Ignore it for now.

---

# Run the programm (first attempt)
From you project directory, run folowwing command from the terminal:

```Bash
go run .
```

Unfortunataly it doesn't work. You should see following output in your console:

```Text
current directory outside modules listed in go.work or their selected dependencies
```

There are two particular problems that has to be solved yet.

---

# What is the `go.mod` file?
We are about to write a Go module, to combine useful functionality from different packages in order to create a tool we need. In order to do so, we need to specify the dependencies and the version of Go we are using for this specific module. 

---

# Initialize your module
As scary as it sounds, in practice we only need to run following command from the terminal in our projects directory:

```bash
go mod init github.com/line-39/go-microservice
# go: creating new go.mod: module github.com/line-39/go-microservice
# go: to add module requirements and sums: go mod tidy
```

This command creates a `go.mod` file in the current directory, and collects all the required information for us. Read more about modules [here](https://go.dev/doc/tutorial/create-module).

---

# Inspect the `go.mod` file

```Bash
ls .
# go.mod  LICENSE  main.go  README.md

cat go.mod
# module github.com/line-39/go-microservice
#
# go 1.22.1
```

---

# Run the programm (second attempt)
We are going to run our program again:

```Bash
go run .
# current directory is contained in a module that is not one of the workspace 
# modules listed in go.work. You can add the module to the workspace using:
# go work use .
```
Still doesn't work! But we can see that output differs from what we've seen before.

---

# Go workspace
In a nutshell Go complains that the module we are trying to run is not associated with the *go workspace*. Go workspace is group of modules / packages defined in `go.work` file which should be located in the directory containing the *modules*. For the workspace initialization we use `go work init <directory to use>` command. You can read more about it [here](https://go.dev/doc/tutorial/workspaces).

---

# Project directory
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
And you are going to use all the modules located under `github.com/...` and `sequery.de` dirs. 

---

# Initializing the Go workspace
From the parent directory containing all of your modules (in this specific case `src`), initialize go workspace with following command:

```Bash
go work init .
```

--- 

# Parent directory structure
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

---

# Add your module to the Go workspace
In case you created your module within the directory with *initialized* `go.work`, assuming you are located in the *project directory* just add your module to the workspace as shown below:

```Bash
go work use .
```

---

# Run the programm (third attempt)
Run the program one more time

```Bash
go run .
# Hello, world 🪄🪄🪄
```
Finally, the magick works 🪄🪄🪄

---

# Building your program
The `go run` compiles your programm on background, saves the binary to your `/temp` dir, and run it. You can build it yourself with `go build` command. Build and run your programm from the terminal:

```Bash
go build .
ls
# go-microservice  go.mod  LICENSE  main.go  README.md
chmod 776 go-microservice
./go-microservice
# Hello, world 🪄🪄🪄
rm go-microservice
```

---

# First API service: Intro
Now let's make a big jump from iconic "Hello, world!" app to the web server. Our service will respond to requests to it's `/` endpoint with simple "Hi there 👋" message.

We must consider three main components of our service:

1. A *handler* executing the logic of our app in response to the *specific* request;
2. A *router / servermux* which maps *URL patterns* to corresponding *handlers*;
3. A webserver listening to the requests; 

For building our webserver, we are going to use `net/http` package from Go's *standard library*. It will help us to create simple, light-weight, fully functional app.

Switch to the branch `03-first-api` or modify your `main.go` as shown below:

---

### First API: Implementation

```Go
package main

import (
    "log"
    "net/http"
)

func hello(w http.ResponseWriter, r http.Request) {
    w.Write([]byte("Hi there 👋"))
}

func main() {
    // create a new servermux
    mux := http.NewServeMux()
    // register hello() as a handler for "/" pattern
    mux.HandleFunc("/", hello)

    // log the service startup
    log.Print("starting service on :4000")

    // start http server on :4000
    err := http.ListenAndServe(":4000", mux)
    // log the error message if ListenAndServe() encounters error
    log.Fatal(err)
}
```

---

### Run the service
From the command line in your project folder run the command below:

```Bash
go run .
```

If everything goes right, you will see the log message printed into your terminal, saying that service is starting on port 4000.

---

### Call the API endpoint
Open your browther at [http://localhost:4000](http://localhost:4000) to see the response from the service. 

Alternatively, you can query the service with `curl`:

```Bash
curl http://localhost:4000/
# Hi there 👋
```

---

### Adding more endpoints
Switch to the new branch `04-first-api` by running `git switch 04-first-api` if you are working on clonned project, or modify yor `main.go` as shown on the next slide.

---

### Adding new handlers

First, define new handlers (definitions go before `main()` definition):

```Go
// ...

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there 👋"))
}

func viewData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display user data 📁"))
}

func uploadData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload user data 📤"))
}
```

---

### Registering new handlers

Second register your new handlers:

```Go
func main() {
	// create a new servermux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/data/view", viewData)
	mux.HandleFunc("/data/upload", uploadData)

	// ... the rest of the code
}
```

---

#### Run the service
Now, run the service locally, executing following command in your terminal.

```Bash
go run .
# 1970/01/01 00:00:00 starting service on :4000
```

You are ready to query the API

---

### Query the API (browser)
While service is running, open your browser and type the following addresses to see the response from the service:
- [localhost:4000/](localhost:4000/hellow) our `hello` endpoint;
- [localhost:4000/data/view](localhost:4000/data/view) our `viewData` endpoint;
- [localhost:4000/data/upload](localhost:4000/data/upload) our `uploadData` endpoint;

---

### Query the API (CLI)

Type following commands to see the response from the endpoints we just created.

```Bash
curl -i localhost:4000/
curl -i localhost:4000/data/view
curl -i localhost:4000/data/upload
```

---

### Restricting subtree path
Every path that deos not end with the trailing slash will be matched *exectly* by the router. However, any path with *trailing slash* is considered to be a *subtree path pattern*, and it matches for any path matching subtree pattern.

Run the service with `go run .`, and navigate to [localhost:4000/foo](localhost:4000/foo). This endpoint does not exist in our `servermux` definition. But server responds with greetings, it is because *subtree path pattern* match. I.e [localhost:4000/](localhost:4000) will call `hello()` handler, sending the response to the user. In other words, trailing slash can be red as `/**` i.e. wildcard pattern.

---

### Restricting subtree path
To *restrict subtree pattern matching* we can add a *special character* `{$}` to the end of the path, after the trailing slah. Modify your router definition as shown below and restart the server:

```Go
// ...
mux.HandleFunc("/{$}", hello)
// ...
```

Now restart the service, and navigate to [localhost:4000/foo](localhost:4000/foo) again, you should receive `404 page not found` response.

