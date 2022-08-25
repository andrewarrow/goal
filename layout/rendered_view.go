package layout

import "fmt"

type RenderedView struct {
	Top        int
	Leading    int
	Width      int
	Height     int
	TopSet     bool
	LeadingSet bool
	WidthSet   bool
	HeightSet  bool
}

func (rv *RenderedView) isComplete() bool {
	return rv.TopSet && rv.LeadingSet && rv.WidthSet && rv.HeightSet
}

func (rv *RenderedView) setComplete() {
	rv.TopSet = true
	rv.LeadingSet = true
	rv.WidthSet = true
	rv.HeightSet = true
}

func (rv *RenderedView) String() string {
	return fmt.Sprintf("%d,%d|%d,%d", rv.Top, rv.Leading, rv.Width, rv.Height)
}
