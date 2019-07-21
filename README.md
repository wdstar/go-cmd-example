# go-cmd-example
Golang CLI implementation example

## Examples

- [Display version](./cmd/version.go)
- [Parse flags](./cmd/test.go)
- [Read in stdin](./cmd/stdio.go)
- [Execute an external command](./cmd/exec.go)

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
1. Add goreleaser configurations.
    ```bash
    $ curl -sL -o /usr/local/bin/goreleaser https://git.io/goreleaser
    $ chmod 755 /usr/local/bin/goreleaser
    $ goreleaser init
    # Snapshot build
    $ goreleaser --snapshot --skip-publish --rm-dist
    $ ./dist/go-cmd-example_linux_amd64/go-cmd-example version
    # Production release
    $ export GITHUB_TOKEN=**********************************
    $ git tag -a vX.X.X -m 'initial release.'
    $ git push --tags
    $ goreleaser --rm-dist
    ```
