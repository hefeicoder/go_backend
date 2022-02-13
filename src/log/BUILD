load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log",
    srcs = ["logger.go"],
    importpath = "log",
    visibility = ["//visibility:public"],
    deps = [
        "@org_uber_go_zap//:go_default_library",
    ],
)
