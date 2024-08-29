# Tmux Compose (tcomp)

Tmux Compose is a wrapper around `tmux` to allow you to create workspace 
quickly. 

## Table of contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

### Go

```sh
go install github.com/bmehdi777/tcomp
```

### Manual

First, you will need to install [Go](https://go.dev/doc/install).

Then :

```sh
git clone https://github.com/bmehdi777/tcomp
cd tcomp
make build
```

The binary will be located in the `bin/` directory.

## Usage

### help

Every command can be seen with :

```sh
tcomp help
```

### list

You can list every repository located in your `$HOME/.config/tcomp/repository/` 
:

```sh
tcomp list
```

or

```sh
tcomp ls
```

or

```sh
tcomp see
```

### new

You can create manually your repository (they are simple yaml files) in your 
`$HOME/.config/tcomp/repository/` folder but you can also generate them with
`tcomp` : 


```sh
tcomp new <REPOSITORY_NAME>
```

### up

```sh
tcomp up <REPOSITORY_NAME>
```

or 

```sh
tcomp up -f <PATH_TO_REPOSITORY>
```

### down

```sh
tcomp down <REPOSITORY_NAME>
```

or 

```sh
tcomp down -f <PATH_TO_REPOSITORY>
```
