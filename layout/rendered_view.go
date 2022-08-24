package layout

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
