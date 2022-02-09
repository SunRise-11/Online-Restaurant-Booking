<div id="top"></div>
[![Contributors][contributors-shield]][contributors-url]

# Restobook

## Description

<!-- Description -->
<br />
<div align="center">
  <a href="https://github.com/herlianto-github/Restobook">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Restobook</h3>

  <p align="center">
    coming soon
    <br />
    <a href="https://github.com/herlianto-github/Restobook"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/herlianto-github/Restobook">Wireframe</a>
    ·
    <a href="https://github.com/herlianto-github/Restobook/issues">Report Bug</a>
    ·
    <a href="https://github.com/herlianto-github/Restobook/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details open>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#description">About The Project</a>
      <ul>
        <li><a href="#depedencies">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#how_to_contribute">Contributing</a></li>    
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## Getting Started

### Dependencies

- [Git](https://git-scm.com)
- [Golang](https://go.dev)
- [Visual Studio Code](https://code.visualstudio.com)

### How To Contribute

- Fork this repository

    ```console
    $ git clone https://github.com/YOUR_USERNAME/Restobook.git
    > Cloning into `Restobook`...
    > remote: Counting objects: 10, done.
    > remote: Compressing objects: 100% (8/8), done.
    > remove: Total 10 (delta 1), reused 10 (delta 1)
    > Unpacking objects: 100% (10/10), done.
    ```

    ```console
    cd Restobook
    ```

- Simple run  

    ```console
    go mod init Restobook
    ```

    ```console
    touch main.go    
    ```

    ```console
    echo 'package main 
    
    import "fmt"
    
    func main(){
    
        fmt.Println("Hello World")
    
    }' >> main.go
    ```

    ```console
    go run main.go
    ```

- Important

    ```console
    git checkout -b feature-name 
    ```

    Always create new branch when develop something

    ```console
    git add .    
    ```

    ```console
    git commit -m "feature description"
    ```

    ```console
    $ git remote -v
    > origin  https://github.com/YOUR_USERNAME/Restobook.git (fetch)
    > origin  https://github.com/YOUR_USERNAME/Restobook.git (push)
    ```

    ```console
    git remote add upstream https://github.com/herlianto-github/Restobook.git
    ```

    ```console
    $ git remote -v
    > origin    https://github.com/YOUR_USERNAME/Restobook.git (fetch)
    > origin    https://github.com/YOUR_USERNAME/Restobook.git (push)
    > upstream  https://github.com/herlianto-github/Restobook.git (fetch)
    > upstream  https://github.com/herlianto-github/Restobook.git (push)
    ```

    ```console
    git push -u origin feature-name    
    ```

### Executing program

- How to run the program

    ```console
    go run main.go    
    ```

## Help

- **Configs**<br/>Contain database and http configuration
- **Delivery (API)**<br/>API http handlers or controllers
- **Entities**<br/>Contain database model
- **Repository** <br/> Contain implementation entities database anq query with ORM.
- **Utils**<br/>Contain database driver (mySQL)

## Authors

[Andrew Prasetyo](https://github.com/andrewptjio)

[Herlianto](https://github.com/herlianto-github)

[Ilham Junius](https://github.com/ilhamjunius)

## Version History

- 0.0.1
  - Initial Release

## Acknowledgments

- [Layered Architecture](https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html)

[controbutors-shield]: https://img.shields.io/github/contributors/herlianto-github/Restobook.svg?style=for-the-badge
[contributors-url]: https://github.com/herlianto-github/Restobook/graphs/contributors
