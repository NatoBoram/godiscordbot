# How to create a Discord bot

[![Go Report Card](https://goreportcard.com/badge/github.com/NatoBoram/godiscordbot)](https://goreportcard.com/report/github.com/NatoBoram/godiscordbot)
[![Go](https://github.com/NatoBoram/godiscordbot/workflows/Go/badge.svg?branch=master)](https://github.com/NatoBoram/godiscordbot/actions?query=workflow%3AGo)

## Requirements

1. Install required software
    1. [Git](https://git-scm.com/)
    2. [Go](https://golang.org/)
    3. [Visual Studio Code](https://code.visualstudio.com/)
        1. [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
2. Create a [new repository](https://github.com/new) on GitHub
    - Please refer to [package-names](https://blog.golang.org/package-names) to
name the repository
    - Add `.gitignore`: **Go**
3. [Connect to GitHub with SSH](https://help.github.com/github/authenticating-to-github/connecting-to-github-with-ssh)
by reading the following documents
    - [About SSH](https://help.github.com/github/authenticating-to-github/about-ssh)
    - [Checking for existing SSH keys](https://help.github.com/github/authenticating-to-github/checking-for-existing-ssh-keys)
    - [Generating a new SSH key and adding it to the ssh-agent](https://help.github.com/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)
    - [Adding a new SSH key to your GitHub account](https://help.github.com/github/authenticating-to-github/adding-a-new-ssh-key-to-your-github-account)
    - [Testing your SSH connection](https://help.github.com/github/authenticating-to-github/testing-your-ssh-connection)
    - [Working with SSH key passphrases](https://help.github.com/github/authenticating-to-github/working-with-ssh-key-passphrases)
4. [Understand the GitHub flow](https://guides.github.com/introduction/flow/)
5. [Learn Go](https://golang.org/doc/) by reading the following documents
    - [A Tour of Go](https://tour.golang.org/)
    - [How to write Go code](https://golang.org/doc/code.html)
    - [Editor plugins and IDEs](https://golang.org/doc/editors.html)
    - [Effective Go](https://golang.org/doc/effective_go.html)
6. Discord
    1. Read the introduction to the [Discord Developer Portal](https://discordapp.com/developers/docs/intro)
    2. [Create an application](https://discordapp.com/developers/applications)
        - Add a bot user

## Setup the repository

1. Download the newly created GitHub repository using `git clone` via SSH
2. Open the folder with VSCode
3. Create a file named `main.go`

In `main.go`, paste the following.

```go
package main

func main() {

}
```

In the following commands, you'll need to use the URL of the GitHub repository
you've just created in the format `github.com/username/repo`.

```sh
go mod init <url of repository>
go get -u -v github.com/bwmarrin/discordgo
```

Once you have something in the repository, commit it and send it to the GitHub
repository. Use VSCode's _Source Control_ tab to help you out.

## Coding something

Use [discordgo](https://github.com/bwmarrin/discordgo) to get started.
The full documentation is at [pkg.go.dev/github.com/bwmarrin/discordgo](https://pkg.go.dev/github.com/bwmarrin/discordgo).

Remember that your token should not be stored inside your source code, but
instead be in a configuration file.
