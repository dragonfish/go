# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/main/build-ref.html#workspace
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "go-commons",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

## GO
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "de7974538c31f76658e0d333086c69efdf6679dbc6a466ac29e65434bf47076d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.45.0/rules_go-v0.45.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.45.0/rules_go-v0.45.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "32938bda16e6700063035479063d9d24c60eda8d79fd4739563f50d331cb3209",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

load("//:go_deps.bzl", "go_dependencies")

# gazelle:repository_macro go_deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.21.6")

gazelle_dependencies()

gazelle_dependencies(go_repository_default_config = "//:WORKSPACE")
