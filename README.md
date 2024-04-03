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
Switch to the next chapter [Hello]() of the introduction:

```Bash
git switch 01-hello
```

## Installation
If not yet installed, follow the [official guide](https://go.dev/doc/install) to install Go on your system.

Verify the installation:
```Bash
go version
# go version go1.22.1 linux/amd64
```

You are ready to Go!
