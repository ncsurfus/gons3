package gons3

// Label models a GNS3 label.
type Label struct {
	Text     string `json:"text"`
	Style    string `json:"style"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Rotation int    `json:"rotation"`
}

// LabelCreator models a new GNS3 label.
// GNS3 schema requires values: Text
type LabelCreator struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (l *LabelCreator) SetProperty(name string, value interface{}) {
	if l.values == nil {
		l.values = map[string]interface{}{}
		l.values["text"] = ""
	}
	l.values[name] = value
}

// SetStyle sets the SVG style attribute.
func (l *LabelCreator) SetStyle(style string) {
	l.SetProperty("style", style)
}

// SetX sets the relative X position.
func (l *LabelCreator) SetX(x int) {
	l.SetProperty("x", x)
}

// SetY sets the relative y position.
func (l *LabelCreator) SetY(y int) {
	l.SetProperty("y", y)
}

// SetRotation sets the relative rotation: Between -359 to 360
func (l *LabelCreator) SetRotation(rotation int) {
	l.SetProperty("rotation", rotation)
}
