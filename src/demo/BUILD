load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "demoserver",
    srcs = ["main.go"],
    deps = [
        "//src/demo/server",
        "//src/logger",
        "@com_github_buaazp_fasthttprouter//:go_default_library",
        "@com_github_valyala_fasthttp//:go_default_library",
    ],
)
