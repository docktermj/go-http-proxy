# go-http-proxy development

## Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

## Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=docktermj
    export GIT_REPOSITORY=go-http-proxy
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow
   steps in [clone-repository] to install the Git repository.

## Dependencies

1. A one-time command to install dependencies needed for `make` targets.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make dependencies-for-development

    ```

1. Install dependencies needed for [Go] code.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make dependencies

    ```

## Lint

1. Run linting.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make lint

    ```

## Build

1. Build the binaries.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean build

    ```

1. The binaries will be found in the `${GIT_REPOSITORY_DIR}/target` directory.
   Example:

    ```console
    tree ${GIT_REPOSITORY_DIR}/target

    ```

## Run

1. Run program.
   Examples:

    1. Linux

        ```console
        ${GIT_REPOSITORY_DIR}/target/linux-amd64/go-http-proxy

        ```

    1. macOS

        ```console
        ${GIT_REPOSITORY_DIR}/target/darwin-amd64/go-http-proxy

        ```

    1. Windows

        ```console
        ${GIT_REPOSITORY_DIR}/target/windows-amd64/go-http-proxy

        ```

1. Clean up.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

## Test

1. Run tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean setup test

    ```

## Coverage

Create a code coverage map.

1. Run Go tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean setup coverage

    ```

## Documentation

1. View documentation.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean documentation

    ```

1. If a web page doesn't appear, visit [localhost:6060].
1. Documentation will be in the "Third party" section.
   `github.com` > `docktermj` > `go-http-proxy`

1. When a versioned release is published with a `v0.0.0` format tag,
the reference can be found by clicking on the following badge at the top of the README.md page.
Example:

    [![Go Reference Badge]][Go Reference]

1. To stop the `godoc` server, run

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

## References

[clone-repository]: https://github.com/docktermj/knowledge-base/blob/main/HOWTO/clone-repository.md
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/docktermj/template-go.svg
[Go Reference]: https://pkg.go.dev/github.com/docktermj/template-go
[go]: https://github.com/docktermj/knowledge-base/blob/main/WHATIS/go.md
[localhost:6060]: http://localhost:6060/pkg/github.com/docktermj/template-go/
