# go_backend
bootstrap template for golang backend with bazel

The main purpose of this code is to serve as template to booststrap your __golang__ based backend with __bazel__ build system



## Library versions

### Prerequisite

#### [bazelisk](https://github.com/bazelbuild/bazelisk): (OSX) brew install bazelisk
#### [pre-commit](https://pre-commit.com/) : (OSX) brew install pre-commit

### Build system and compiler
| Syntax                                             |        Version        | Description                                           |
|:---------------------------------------------------|:---------------------:|:------------------------------------------------------|
| [Bazel](https://bazel.build/)                      | 4.2.2  | Build system, fast and cross language (.bazelversion) |
| [bazelisk](https://github.com/bazelbuild/bazelisk) |          N/A          | Manage bazel version for the project                  |
| [rules_go](https://go.dev/)                        |  v0.30.0    | Golang build rules for bazel  (version in WORKSPACE)  |
| [golang](https://go.dev/)                          |  1.17.7    | Golang (version in WORKSPACE)                         |

### Server backend
| Syntax                           |        Version        | Description                                           |
|:---------------------------------|:---------------------:|:------------------------------------------------------|
| [fasthttp](https://github.com/valyala/fasthttp) | 4.2.2  | Build system, fast and cross language (.bazelversion) |

### Lint

### CI/CD

### Test
