# Go bindings and tests

Everything in this repository that loads `rules_go` lives under `go/`. The
schema tree outside it loads only protobuf rules, so a consumer can depend on
`@nexteggapp_taxprotos//us/federal/v1/form_w2:form_w2_proto` without a Go
toolchain. Bazel loads a whole BUILD file to build any target in it, so a
single `rules_go` load in a schema package would break that consumer's build.

The bindings exist for the test suites here. No generated `.pb.go` file is
checked in, and the `go.mod` in this directory only keeps these packages out of
the root module's `go test ./...`. Run them with `bazel test //...`.

## Layout

`go/<path>` holds the Go targets for the schema package at `<path>`. A
`go_proto_library` names its schema by label:

```starlark
# us/federal/v1/form_w2/BUILD.bazel: schema only
proto_library(
    name = "form_w2_proto",
    import_prefix = "taxprotos",
    srcs = ["form_w2.proto"],
    visibility = ["//visibility:public"],
    deps = ["//common/v1:common_proto"],
)
```

```starlark
# go/us/federal/v1/form_w2/BUILD.bazel: bindings only
go_proto_library(
    name = "form_w2_go_proto",
    importpath = "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_w2",
    proto = "//us/federal/v1/form_w2:form_w2_proto",
    visibility = ["//visibility:public"],
    deps = ["//go/common/v1"],
)

go_library(
    name = "form_w2",
    embed = [":form_w2_go_proto"],
    importpath = "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_w2",
    visibility = ["//visibility:public"],
)
```

Import paths carry the `go/` segment because these packages belong to the
`github.com/nexteggapp/taxprotos/go` module in this directory.

A mapping file and the form extract it is checked against stay in the schema
package, because they are part of what the schema tree publishes. The test
under `go/` reads them through runfiles at their schema path. A package that
exists only to hold a test, such as `go/us/federal/v1/pub1220`, keeps its own
`testdata/`.

Because `rules_go` is a development dependency, `@nexteggapp_taxprotos//go/...`
builds only when this repository is the root module.
