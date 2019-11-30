package gons3

// Label models a GNS3 label.
type Label struct {
	Text     string `json:"text"`
	Style    string `json:"style"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Rotation int    `json:"rotation"`
}
