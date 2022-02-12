load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "demo_service_proto",
    srcs = ["demo_service.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "demo_service_go_proto",
    importpath = "demo/proto/demo_service_proto",
    protos = [":demo_service_proto"],
    visibility = ["//visibility:public"],
)

