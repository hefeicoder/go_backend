load("@bazel_gazelle//:deps.bzl", "go_repository")

def load_external_go_repositories():
    go_repository(
        name = "com_github_valyala_fasthttp",
        importpath = "github.com/valyala/fasthttp",
        tag = "v1.33.0"
    )