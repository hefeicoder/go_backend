load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "server",
    srcs = ["server.go"],
    importpath = "demo/server",
    visibility = ["//visibility:public"],
    deps = [
        "//src/demo/proto:demo_service_go_proto",
    ],
)
