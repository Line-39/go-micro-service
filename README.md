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
# vim main.go   # if you ran this command - burn your pc - your only path to freedom
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

---

### WIldcard patterns
The `net/http` `servemux` lets us to us whildcards in path patterns. Lets consider following scenario for our service:

1. All the data served based on specified user ID;
2. There are two types of data: *raw* data and *clean* data;

Lets change our `data/view` and `data/upload` routes:

```Go
// ...
mux.HandleFunc("/{user}/data/{datatype}/view", viewData)
mux.HandleFunc("/{user}/data/{datatype}/upload", uploadData)
// ...
```
On the next slide, we modify our handler (shown for `view`) to use the wildcards.

---

### Wildcard patterns (continued)

```Go
func viewData(w http.ResponseWriter, r *http.Request) {
    user := r.PathValue("user")
    if user == "" {
        http.NotFound(w, r)
        return
    }
    dtype := r.PathValue("datatype")
        if user == "" {
        http.NotFound(w, r)
        return
    }

    msg := fmt.Sprintf("📁 Display the %s data for user %s\n", dtype, user)
    w.Write([]byte(msg))

}
```

---

### Wildcard patterns (continued)
Now we can restart our service. Let say we want to see **raw** data for the user **ubot**:

```Bash
curl -i localhost:4000/ubot/data/raw/view
#HTTP/1.1 200 OK
#Date: Mon, 08 Apr 2024 14:42:04 GMT
#Content-Length: 40
#Content-Type: text/plain; charset=utf-8
#
#📁 Display the raw data for user ubot
```

Keep in mind, that user can send any kind of parameter as a wildcard. So checking the validity is entirely on you.

--- 

### Wildcard patterns (continued)
Let's try another request

```Bash
curl -i localhost:4000/jer/data/jobs/view
#HTTP/1.1 200 OK
#Date: Mon, 08 Apr 2024 14:44:53 GMT
#Content-Length: 40
#Content-Type: text/plain; charset=utf-8
#
#📁 Display the jobs data for user jer
```

---

### Wildcard patterns (continued)
Be aware that patterns defined with *wildcard* might overlap. E.g. `user/view` and `user/{data}` requests overlap (incoming `user/view` request is a valid match for `user/{data}` pattern). 

In such cases `servemux` applies following precedence rule: *The most specific pattern wins*. 

Since `user/view` matches only *one specific* request, and `user/{data}` matches infinit amount of possible requests `user/view` will take precedent.

---

### HTTP methods
We can introduce constraints, so our API responts *only* to the HTTP requests with *appropriate HTTP method*. To achive this we will edit the rout regisration in `main.go`:

```Go
// ...
// register handlers
mux.HandleFunc("GET /{$}", hello)
mux.HandleFunc("GET /{user}/data/{datatype}/view", viewData)
mux.HandleFunc("POST /{user}/data/{datatype}/upload", uploadData)
// ...
```

Now save your changes and restart the service to request API.

---

### HTTP methods (continued)
From CLI try following querries

```Bash
curl -i localhost:4000/                             # 200 
curl -i -X POST localhost:4000/ubot/data/raw/view   # 405 Method Not Allowed
curl -i -X GET  localhost:4000/ubot/data/raw/view   # 200
curl -i -X GET  localhost:4000/ubot/data/raw/upload # 405 Method Not Allowed
curl -i -X POST localhost:4000/ubot/data/raw/upload # 200
```
All the requests with unappropriate HTTP methods were rejected.

--- 

### HTTP methods (continued)
Using HTTP methods, we can simplify the pattern of our endpoints, getting rid of the last `view / upload` part, since it will be defined by the HTTP method we use.

```Go
// ...
// register handlers
mux.HandleFunc("GET /{$}", hello)
mux.HandleFunc("GET /{user}/data/{datatype}", viewData)
mux.HandleFunc("POST /{user}/data/{datatype}", uploadData)
// ...
```

---

### HTTP methods (continued)
Modifying our request, we can see that the same endpoint is now responds differently depending on the HTTP method we are using.

```Bash
curl -X GET localhost:4000/jer/data/raw
# 📁 Display the raw data for user jer

curl -X POST localhost:4000/jer/data/raw
# 📤 Upload the raw data for user jer
```

--- 

### HTTP status codes
First time the `w.Write()` has been called it wrights `200 OK` to the response *header*. Write `201 Created` status code to the header, as shown below: 

```Go
package main

// ...
func uploadData(w http.ResponseWriter, r *http.Request) {
    // ...
    w.WriteHeader(http.statusCreated)
    msg := fmt.Sprintf("📤 Upload the %s data for user %s\n", dtype, user)
    w.Write([]byte(msg))
}
```

---

### Http status codes (continued)
Onece writen, header's status code *can not be changed*. So the status code has to be modified *before* any subsequent `w.Write()` call. For example, the code below returns warning, and status code remains to be `200 OK`:

```Go
func someHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.statusProcessing)

    if ok := true; ok {
        w.WriteHeader(http.statusAccepted)
    }
    w.Write([]byte("Hell yeah!"))
}
```

---

### HTTP status codes (continued)
Run you service, and call `<user>/data/<datatype>` endpoint using `POST` request:

```Bash
go run .
# 1970/01/01 00:00:00 starting service on :4000
# another terminal window
curl -Xi POST localhost:4000/jer/data/raw
#HTTP/1.1 201 Created
#Date: Thu, 01 Jan 1970 00:00:00 GMT
#Content-Length: 38
#Content-Type: text/plain; charset=utf-8
#
#📤 Upload the raw data for user jer
```

---

### Modifying header map
We can modify header with `Header()`, `Add()`, `Set()`, `Del()`, `Get()`, `Values()` methods. For example, overwrite *content-type of the response, add the information about our server (key) name (value) for `/` handler:

```Go
func hello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Add("Server", "Simple Go Service")
    w.Write([]byte(`{"message":"Hi there 👋"}`))
}
```

---

### Modifying header map (continued)
Restart the service and call `/` endpoint:

```Bash
go run .
curl -i localhost:4000/
#HTTP/1.1 200 OK
#Content-Type: application/json
#Server: Simple Go Service
#Date: Thu, 01 Jan 1970 00:00:00 GMT
#Content-Length: 29
#
#{"message":"Hi there 👋\n"}
```

---

### Configuration settings: CLI arguments
Our service listens to the network address wich is hardcoded in `main()`. This is OK for early builds, but not a good way to configure the application.

```Go
// log the service startup
log.Print("starting service on :4000")

// start http server on :4000
err := http.ListenAndServe(":4000", mux)
```
The good idea is to make such things as network-address, name and version of the application, secrets  etc configurable at runtime.

---

### Configuration settings: CLI arguments (continued)
We can pass command-line arguments when starting the application:

```Bash
go run . -addr=":4000" -name="Simple Go Microservice" -version="0.0.1"
```

We can access CLI arguments with `flag` package `String` function:

```Go
func main() {
    // ...
    addr    := flag.String("addr", ":4000", "HTTP network address")
    name    := flag.String("name", "Simple Go Microservice", "The name of the app")
    version := flag.String("version", "0.0.0", "The version of the app")
    flag.Parse()
    // ...
}
```

---

### Configuration settings: CLI arguments (continued)
It is important to understand that first we *define* our variables `addr`, `name` and `version` as new vars of *CLI flag type*. When latter we call `flag.Parse()` it parses CLI arguments and passes the values to this variables:

```Go
func main() {
    //...
    // log on startup
    log.Printf("Starting %s, version %s on %s", *addr, *name, *version)

    // start http server on :4000
    err := http.ListenAndServe(*addr, mux)
    //...
}
```

---

### Configuration settings: CLI arguments (continued)
Modify `main()` and run the service without *any* command-line flags:

```Bash
go run .
# 1970/01/01 00:00:01 starting Simbple Go Microservice, version 0.0.0 on :4000
```
Run the service with some flags:

```Bash
go run . -addr="5000" -name="Einfacher Go-Service"
# 1970/01/01 00:00:01 starting Einfacher Go-Service, version Infinity on :5000
```

---

### Configuration settings: CLI arguments (continued)
The help for command-line arguments is automatically created for us:

```Bash
go run . -help
# Usage of /tmp/go-build2113936530/b001/exe/go-microservice:
#   -addr string
#         HTTP network address (default ":4000")
#   -name string
#         Service name (default "Simple Go Microservice")
#   -version string
#         Service version (default "0.0.0")
```
