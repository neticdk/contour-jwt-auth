load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["root.go"],
    importpath = "github.com/neticdk/contour-jwt-auth/pkg/cmd",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/server:go_default_library",
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials:go_default_library",
    ],
)
