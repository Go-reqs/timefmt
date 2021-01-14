# timefmt

time format string from posix to golang timefmt

## command line usage

1. install

    ```shell
    go get github.com/go-reqs/timefmt/cmd/gotimefmt
    ```

1. time format string transfer from posix format

    [strftime posix format](https://man7.org/linux/man-pages/man3/strftime.3.html)

    ```shell
    gotimefmt "%y-%m-%d"
    ```

    will output:

    ```txt
    2006-01-02
    ```

## programing api usage

1. install

    ```shell
    go install github.com/go-reqs/timefmt
    ```

1. example

    ```golang
    package main

    import (
      "fmt"
      "github.com/go-reqs/timefmt"
    )

    func main() {
      fmt.Println(timefmt.FromPosix("%y-%m-%d"))
    }
    ```
