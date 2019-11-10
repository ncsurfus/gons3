// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/schemas/node.py
// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/handlers/api/controller/node_handler.py

package gons3

// Node models a GNS3 node.
type Node struct {
	ComputeID        string                 `json:"compute_id"`
	ProjectID        string                 `json:"project_id"`
	NodeID           string                 `json:"node_id"`
	TemplateID       string                 `json:"template_id"`
	NodeType         string                 `json:"node_type"`
	NodeDirectory    string                 `json:"node_directory"`
	CommandLine      string                 `json:"command_line"`
	Name             string                 `json:"name"`
	Console          int                    `json:"console"`
	ConsoleHost      string                 `json:"console_host"`
	ConsoleType      string                 `json:"console_type"`
	ConsoleAutoStart bool                   `json:"console_auto_start"`
	Properties       map[string]interface{} `json:"properties"`
	Status           string                 `json:"status"`
	Label            Label                  `json:"label"`
	Symbol           string                 `json:"symbol"`
	Width            int                    `json:"width"`
	Height           int                    `json:"height"`
	X                int                    `json:"x"`
	Y                int                    `json:"y"`
	Z                int                    `json:"z"`
	Locked           bool                   `json:"locked"`
	PortNameFormat   string                 `json:"port_name_format"`
	PortSegmentSize  int                    `json:"port_segment_size"`
	FirstPortName    string                 `json:"first_port_name"`
	CustomAdapters   []CustomAdapter        `json:"custom_adapters"`
	Ports            []NodePort             `json:"ports"`
}

// NodePort models a GNS3 node's port.
type NodePort struct {
	Name          string                 `json:"name"`
	ShortName     string                 `json:"short_name"`
	AdapterNumber int                    `json:"adapter_number"`
	AdapterType   string                 `json:"adapter_type"`
	PortNumber    int                    `json:"port_number"`
	LinkType      string                 `json:"link_type"`
	DataLinkTypes map[string]interface{} `json:"data_link_types"`
	MACAddress    string                 `json:"mac_address"`
}

// CustomAdapter models a GNS3 custom adapters.
type CustomAdapter struct {
	AdapterNumber int    `json:"adapter_number"`
	PortName      string `json:"port_name"`
	AdapterType   string `json:"adapter_type"`
	MACAddress    string `json:"mac_address"`
}

// Label models a GNS3 label.
type Label struct {
	Text     string `json:"text"`
	Style    string `json:"style"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Rotation int    `json:"rotation"`
}
