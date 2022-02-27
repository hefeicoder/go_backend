load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "client",
    srcs = ["main.go"],
    deps = [
        "//src/demo/proto:demo_service_go_proto",
        "//src/logger",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)
