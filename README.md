# go-cmd-example
Golang CLI implementation example

## Development

1. Generate a command template.
    ```bash
    $ go get -u github.com/spf13/cobra/cobra
    $ cobra init --pkg-name github.com/wdstar/go-cmd-example
    $ go mod tidy
    $ go build
    $ ./go-cmd-example 
    ```
1. Add sub commands
    ```bash
    $ cobra add version -p "rootCmd"
    $ go build
    $ ./go-cmd-example version
    ```
