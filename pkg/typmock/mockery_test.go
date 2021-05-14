package typmock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var (
	target1 = &typmock.Mock{Pkg: "pkg1", Dir: "dir1", Source: "target1"}
	target2 = &typmock.Mock{Pkg: "pkg1", Dir: "dir1", Source: "target2"}
	target3 = &typmock.Mock{Pkg: "pkg2", Dir: "dir2", Source: "target3"}
	target4 = &typmock.Mock{Pkg: "pkg1", Dir: "dir1", Source: "target4"}
	target5 = &typmock.Mock{Pkg: "pkg1", Dir: "dir1", Source: "target5"}
	target6 = &typmock.Mock{Pkg: "pkg2", Dir: "dir2", Source: "target6"}

	targets = []*typmock.Mock{
		target1,
		target2,
		target3,
		target4,
		target5,
		target6,
	}

	dir1 = []*typmock.Mock{
		target1,
		target2,
		target4,
		target5,
	}

	dir2 = []*typmock.Mock{
		target3,
		target6,
	}
)

func TestTargetMap(t *testing.T) {
	m := typmock.NewMockery("")
	for _, mock := range targets {
		m.Put(mock)
	}

	require.Equal(t, typmock.Map{"dir1": dir1, "dir2": dir2}, m.Filter("dir1", "dir2"))
	require.Equal(t, typmock.Map{"dir1": dir1}, m.Filter("dir1"))
	require.Equal(t, typmock.Map{}, m.Filter("not-found"))
}

func TestCreateMock(t *testing.T) {
	testcases := []struct {
		testName string
		annot    *typgen.Directive
		expected *typmock.Mock
	}{
		{
			annot: &typgen.Directive{
				Decl: &typgen.Decl{
					File: typgen.File{
						Package: "somePkg",
						Path:    "/path/folder/source.go",
					},
					Type: &typgen.InterfaceDecl{
						TypeDecl: typgen.TypeDecl{Name: "SomeInterface"},
					},
				},
				TagName: "mock",
			},
			expected: &typmock.Mock{
				Dir:          "/path/folder",
				Pkg:          "somePkg",
				Source:       "SomeInterface",
				TargetParent: "internal/generated/mock//path",
				MockPkg:      "somePkg_mock",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, typmock.CreateMock(tt.annot))
		})
	}
}

func TestGenTarget(t *testing.T) {
	testcases := []struct {
		TestName string
		Dir      string
		Expected string
	}{
		{
			Dir:      "internal/app/service",
			Expected: "internal/generated/mock/app",
		},
		{
			Dir:      "internal/service",
			Expected: "internal/generated/mock",
		},
		{
			Dir:      ".",
			Expected: "internal/generated/mock",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typmock.GenTarget(tt.Dir))
		})
	}
}
