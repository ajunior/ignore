# ignore
**ignore** is a *really* simple CLI that helps you add [GitHub-based](https://github.com/github/gitignore) 
```.gitignore``` file to your local git repositories.

# Getting started
You can download [binaries](https://github.com/ajunior/ignore/releases) for Linux, macOS (x86-64 and M1), and Windows, 
but I also intend to make it to be installed through some package managers as soon as possible.

## Instalation 

### Linux / macOS
Open the terminal and follow the steps:

```bash
# Step 1 - Download the binary
$ wget https://github.com/ajunior/ignore/releases/download/v0.1.0/ignore_macos_x64.zip

# Unzip
$ unzip ignore_macos_x64.zip

# Change permission
$ sudo chmod +x ignore

# Move it
$ sudo mv ignore /usr/local/bin
```

For Linux, download the ```ignore_linux_x64.zip``` file.

### Windows
[Download](https://github.com/ajunior/ignore/releases/download/v0.1.0/ignore_win10_x64.zip) and unzip it to a folder 
named ```ignore``` into your user directory (for example, ```C:\Users\username\ignore```), then add that directory
into your path variable.

### Building from source
Go through the following steps if you want to build it by yourself.

```bash
# Step 1 - Clone the git repository
$ git clone https://github.com/ajunior/ignore.git

# Step 2 - Change directory
$ cd ignore

# Step 3 - Build it
$ go build -o ignore src/*.go

# Step 4 - Move it
$ sudo mv ignore /usr/local/bin/

# Step 5 - Change permission
$ sudo chmod +x /usr/local/bin/ignore
```

> If you are using Windows skip the steps 4 and 5 above. Instead, create a folder named ```ignore``` on your user directory (```C:\Users\username\local```), then move the executable and add that directory to your PATH.

### Testing
After installation run ```ignore --version```. If everything is ok you should see something like:

```bash
$ ignore --version
ignore v0.1.0 - https://github.com/ajunior/ignore
```

## Usage

### List the available templates
```bash
$ ignore --templates
```

### Create a .gitignore file from the Java template
```bash
$ ignore --create /path/to/project/directory java
```

### Create a .gitignore file from multiple templates
```bash
$ ignore --create /path/to/project/directory java go python
```
- It combines every template you inputted into the .gitignore file separating them by a newline.
- There's no limit of templates that can be added, so you can create a file with all available templates 
  (but it makes no sense).

# Contributing
If you want to contribute for this project you should fork this repository, clone your fork and modify it, then 
commit your changes and send me a pull request, if it makes sense to the project I'll merge it.

For any questions, suggestions or report of bugs, please open an [issue](https://github.com/ajunior/ignore/issues).

# Versioning
Version numbering follows the [Semantic versioning](https://semver.org/) approach.

# License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/ajunior/ignore/blob/master/LICENSE) file for details.
