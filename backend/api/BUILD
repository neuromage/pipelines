load("@com_google_protobuf//:protobuf.bzl", "cc_proto_library")

proto_library(
    name = "resource_reference_proto",
    srcs = ["resource_reference.proto"],
)

load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "resource_reference_go_proto",
    importpath = "github.com/kubeflow/pipelines/backend/api",
    proto = ":resource_reference_proto",
    visibility = ["//visibility:public"],
)
