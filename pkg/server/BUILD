load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    importpath = "github.com/neticdk/contour-jwt-auth/pkg/server",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_envoyproxy_go_control_plane//envoy/service/auth/v3:go_default_library",
        "@com_github_sirupsen_logrus//:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//health/grpc_health_v1:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
    ],
)
