# go_backend
bootstrap template for golang backend with bazel

The main purpose of this code is to serve as template to booststrap your __golang__ based backend with __bazel__ build system



## Library versions

### Prerequisite
install [golang](https://go.dev/doc/install)

run `bash install_prerequisite.sh` to install all prerequisite

### Build system and compiler
| Syntax                                             |        Version        | Description                                                                                                                 |
|:---------------------------------------------------|:---------------------:|:----------------------------------------------------------------------------------------------------------------------------|
| [Bazel](https://bazel.build/)                      | 4.2.2  | Build system, fast and cross language (.bazelversion)                                                                       |
| [bazelisk](https://github.com/bazelbuild/bazelisk) |          N/A          | Manage bazel version for the project                                                                                        |
| [rules_go](https://go.dev/)                        |  v0.30.0    | Golang build rules for bazel  (version in WORKSPACE)                                                                        |
| [golang](https://go.dev/)                          |  1.17.7    | Golang (version in WORKSPACE), note: this golang is for building this project and will not be conflicit with the system one |

### Server backend
| Syntax                           |        Version        | Description                                           |
|:---------------------------------|:---------------------:|:------------------------------------------------------|
| [fasthttp](https://github.com/valyala/fasthttp) | v1.33.0  | fasthttp |

### Lint with pre-commit hook
#### `./lint.sh` to lint everything: BUILD, json, golang, yaml etc
| Syntax                                | Version | Description               |
|:--------------------------------------|:-------:|:--------------------------|
| [pre-commit](https://pre-commit.com/) | latest  | precommit/lint management |
| buildifier                            | latest  | lint for BUILD file       |
| goimports                             | latest  | lint for golang file      |
| golangci-lint                         | latest  | Go linters aggregator that runs linters in parallel      |


### CI/CD

### Test
