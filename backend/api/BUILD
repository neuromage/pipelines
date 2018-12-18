proto_library(
    name = "resource_reference_proto",
    srcs = ["resource_reference.proto"],
)

load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "resource_reference_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/kubeflow/pipelines/backend/api",
    proto = ":api_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//:go_default_library",
        "//protoc-gen-swagger/options:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)

proto_library(
    name = "api_proto",
    srcs = [
        "error.proto",
        "experiment.proto",
        "filter.proto",
        "job.proto",
        "parameter.proto",
        "pipeline.proto",
        "pipeline_spec.proto",
        "report.proto",
        "resource_reference.proto",
        "run.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//:root_proto",
        "//protoc-gen-swagger/options:options_proto",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@go_googleapis//google/api:annotations_proto",
    ],
)
