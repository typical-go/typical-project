package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	dtorTags = []string{"dtor"}
)

type (
	// DtorAnnotation represent @dtor annotation
	DtorAnnotation struct{}
	// Dtor is destructor tag
	Dtor struct {
		*typast.Annot `json:"-"`
	}
)

var _ Action = (*DtorAnnotation)(nil)

// Compile @dtor
func (*DtorAnnotation) Execute(c *Context) error {
	var dtors []*typtmpl.Dtor
	for _, annot := range c.ASTStore.Annots {
		dtor := ParseDtor(annot)
		if dtor != nil {
			dtors = append(dtors, &typtmpl.Dtor{
				Def: fmt.Sprintf("%s.%s", dtor.Decl.Pkg, dtor.Decl.Name),
			})
		}
	}
	return writeGoSource(
		fmt.Sprintf("cmd/%s/dtor_annotated.go", c.Descriptor.Name),
		&typtmpl.DtorAnnotated{
			Package: "main",
			Imports: c.Imports,
			Dtors:   dtors,
		},
	)
}

// ParseDtor annotation
func ParseDtor(annot *typast.Annot) *Dtor {
	if !IsFuncTag(annot, dtorTags...) {
		return nil
	}

	return &Dtor{
		Annot: annot,
	}
}
