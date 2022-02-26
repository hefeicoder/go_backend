load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "logger",
    srcs = ["logger.go"],
    importpath = "logger",
    visibility = ["//visibility:public"],
    deps = [
        "@org_uber_go_zap//:go_default_library",
    ],
)
