# How to contribute

- Fork this repository

    ```sh
    $ git clone https://github.com/YOUR_USERNAME/Restobook.git
    > Cloning into `Restobook`...
    > remote: Counting objects: 10, done.
    > remote: Compressing objects: 100% (8/8), done.
    > remove: Total 10 (delta 1), reused 10 (delta 1)
    > Unpacking objects: 100% (10/10), done.
    ```

    ```sh
    cd Restobook
    ```

- Simple run  

    ```sh
    go mod init Restobook
    ```

    ```sh
    touch main.go    
    ```

    ```sh
    echo 'package main 
    
    import "fmt"
    
    func main(){
    
        fmt.Println("Hello World")
    
    }' >> main.go
    ```

    ```sh
    go run main.go
    ```

- Important

    ```sh
    git checkout -b feature-name 
    ```

    Always create new branch when develop something

    ```sh
    git add .    
    ```

    ```sh
    git commit -m "feature description"
    ```

    ```sh
    $ git remote -v
    > origin  https://github.com/YOUR_USERNAME/Restobook.git (fetch)
    > origin  https://github.com/YOUR_USERNAME/Restobook.git (push)
    ```

    ```sh
    git remote add upstream https://github.com/herlianto-github/Restobook.git
    ```

    ```sh
    $ git remote -v
    > origin    https://github.com/YOUR_USERNAME/Restobook.git (fetch)
    > origin    https://github.com/YOUR_USERNAME/Restobook.git (push)
    > upstream  https://github.com/herlianto-github/Restobook.git (fetch)
    > upstream  https://github.com/herlianto-github/Restobook.git (push)
    ```

    ```sh
    git push -u origin feature-name    
    ```
