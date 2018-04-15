load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "fila_lib",
    srcs = ["fila.go"],
    importpath = "github.com/patrickrodee/fila",
    visibility = ["//visibility:public"],
)

go_test(
    name = "fila_lib_test",
    srcs = ["fila_test.go"],
    embed = [":fila_lib"],
)