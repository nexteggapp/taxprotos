// Package taxprotos_test compiles every schema in this repository with
// protocompile, the way a consumer that does not use Bazel would: each file is
// imported as taxprotos/<path from the repository root>.
//
// The Bazel build checks the same thing through proto_library, and the suites
// under go/ check the schemas against the forms. They build only with Bazel,
// so this is the check `go test ./...` runs.
package taxprotos_test

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bufbuild/protocompile"
	"google.golang.org/protobuf/types/descriptorpb"
)

// importPrefix is what each proto_library's import_prefix puts before the
// repository-relative path.
const importPrefix = "taxprotos/"

// goPackagePrefix is the start of every file's go_package option.
const goPackagePrefix = "github.com/nexteggapp/taxprotos/"

func TestEverySchemaCompiles(t *testing.T) {
	var paths []string
	err := filepath.WalkDir(".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && (p == "go" || strings.HasPrefix(p, "bazel-") || strings.HasPrefix(p, ".")) && p != "." {
			return filepath.SkipDir
		}
		if !d.IsDir() && strings.HasSuffix(p, ".proto") {
			paths = append(paths, importPrefix+filepath.ToSlash(p))
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		t.Fatal("found no .proto files")
	}

	c := protocompile.Compiler{
		Resolver: protocompile.WithStandardImports(&protocompile.SourceResolver{
			Accessor: func(p string) (io.ReadCloser, error) {
				rel, ok := strings.CutPrefix(p, importPrefix)
				if !ok {
					return nil, fs.ErrNotExist
				}
				return os.Open(filepath.FromSlash(rel))
			},
		}),
	}
	files, err := c.Compile(context.Background(), paths...)
	if err != nil {
		t.Fatal(err)
	}

	// Each go_package names the file's own directory, so bindings generated
	// with protoc-gen-go land in a package that matches the schema tree.
	for _, f := range files {
		dir := strings.TrimPrefix(filepath.ToSlash(filepath.Dir(f.Path())), importPrefix)
		goPkg, _, _ := strings.Cut(f.Options().(*descriptorpb.FileOptions).GetGoPackage(), ";")
		if want := goPackagePrefix + dir; goPkg != want {
			t.Errorf("%s: go_package is %q, want %q", f.Path(), goPkg, want)
		}
	}
}
