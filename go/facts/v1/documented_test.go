// Package v1_test checks that every element of the facts schemas has a doc
// comment.
//
// docs/facts.md defines TaxYearFacts for any tax product to store,
// so each message, field, oneof, enum and enum value says what it holds. A
// compiled descriptor carries no comments, so the test parses the .proto
// sources. Their imports resolve from the descriptors the generated Go
// packages register, so only the files checked need to be on disk.
//
// The check mirrors the one protodiffgen runs on its own schemas, in
// generate/generate_test.go at github.com/nexteggapp/protodiffgen.
package v1_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/bufbuild/protocompile"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	// Registers the facts descriptors and, through them, every file they
	// import.
	_ "github.com/nexteggapp/taxprotos/go/facts/v1"
)

// runfilesPrefix is the workspace name every runfiles path starts with.
const runfilesPrefix = "nexteggapp_taxprotos"

// importPrefix is what each proto_library's import_prefix puts before the
// module-relative path.
const importPrefix = "taxprotos/"

// checked lists the files whose elements must all be documented, by import
// path: the facts schemas and the three forms added with them.
var checked = []string{
	"taxprotos/facts/v1/documents.proto",
	"taxprotos/facts/v1/tax_year_facts.proto",
	"taxprotos/us/federal/v1/form_1095_a/form_1095_a.proto",
	"taxprotos/us/federal/v1/form_1098_e/form_1098_e.proto",
	"taxprotos/us/federal/v1/form_1098_t/form_1098_t.proto",
	"taxprotos/us/federal/v1/form_1099_g/form_1099_g.proto",
	"taxprotos/us/federal/v1/form_1099_misc/form_1099_misc.proto",
	"taxprotos/us/federal/v1/form_1099_sa/form_1099_sa.proto",
	"taxprotos/us/federal/v1/form_5498/form_5498.proto",
	"taxprotos/us/federal/v1/form_5498_sa/form_5498_sa.proto",
	"taxprotos/us/federal/v1/form_8889/form_8889.proto",
	"taxprotos/us/federal/v1/form_ssa_1099/form_ssa_1099.proto",
	"taxprotos/us/federal/v1/form_w2g/form_w2g.proto",
	"taxprotos/us/federal/v1/schedule_k1/schedule_k1_1041.proto",
	"taxprotos/us/federal/v1/schedule_k1/schedule_k1_1065.proto",
	"taxprotos/us/federal/v1/schedule_k1/schedule_k1_1120_s.proto",
}

// TestEveryElementIsDocumented fails for each element of a checked file
// without a leading comment.
func TestEveryElementIsDocumented(t *testing.T) {
	sources := map[string][]byte{}
	for _, p := range checked {
		resolved, err := runfiles.Rlocation(path.Join(runfilesPrefix, strings.TrimPrefix(p, importPrefix)))
		if err != nil {
			t.Fatalf("resolve %s from runfiles: %v", p, err)
		}
		b, err := os.ReadFile(resolved)
		if err != nil {
			t.Fatal(err)
		}
		sources[p] = b
	}
	c := protocompile.Compiler{
		SourceInfoMode: protocompile.SourceInfoStandard,
		Resolver: protocompile.ResolverFunc(func(p string) (protocompile.SearchResult, error) {
			if b, ok := sources[p]; ok {
				return protocompile.SearchResult{Source: io.NopCloser(bytes.NewReader(b))}, nil
			}
			fd, err := protoregistry.GlobalFiles.FindFileByPath(p)
			if err != nil {
				return protocompile.SearchResult{}, err
			}
			return protocompile.SearchResult{Desc: fd}, nil
		}),
	}
	files, err := c.Compile(context.Background(), checked...)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		checkDocumented(t, f)
	}
}

// checkDocumented fails for each message, field, oneof, enum, enum value or
// extension in fd without a leading comment.
func checkDocumented(t *testing.T, fd protoreflect.FileDescriptor) {
	t.Helper()
	locs := fd.SourceLocations()
	check := func(d protoreflect.Descriptor) {
		if strings.TrimSpace(locs.ByDescriptor(d).LeadingComments) == "" {
			t.Errorf("%s: %s has no doc comment", fd.Path(), d.FullName())
		}
	}
	var enums func(protoreflect.EnumDescriptors)
	enums = func(es protoreflect.EnumDescriptors) {
		for i := 0; i < es.Len(); i++ {
			check(es.Get(i))
			for j := 0; j < es.Get(i).Values().Len(); j++ {
				check(es.Get(i).Values().Get(j))
			}
		}
	}
	exts := func(xs protoreflect.ExtensionDescriptors) {
		for i := 0; i < xs.Len(); i++ {
			check(xs.Get(i))
		}
	}
	var messages func(protoreflect.MessageDescriptors)
	messages = func(ms protoreflect.MessageDescriptors) {
		for i := 0; i < ms.Len(); i++ {
			md := ms.Get(i)
			if md.IsMapEntry() {
				continue
			}
			check(md)
			for j := 0; j < md.Fields().Len(); j++ {
				check(md.Fields().Get(j))
			}
			for j := 0; j < md.Oneofs().Len(); j++ {
				if !md.Oneofs().Get(j).IsSynthetic() {
					check(md.Oneofs().Get(j))
				}
			}
			enums(md.Enums())
			exts(md.Extensions())
			messages(md.Messages())
		}
	}
	messages(fd.Messages())
	enums(fd.Enums())
	exts(fd.Extensions())
}
