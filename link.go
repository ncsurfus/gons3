package gons3

import "net/url"

// LinkNode models the node object within the GNS3 link.
type LinkNode struct {
	NodeID        string `json:"node_id"`
	AdapterNumber int    `json:"adapter_number"`
	PortNumber    int    `json:"port_number"`
	Label         Label  `json:"label"`
}

// Link models a GNS3 link between two or more nodes.
type Link struct {
	LinkID           string                 `json:"link_id"`
	ProjectID        string                 `json:"project_id"`
	Nodes            []LinkNode             `json:"nodes"`
	Filter           map[string]interface{} `json:"filter"`
	Suspended        bool                   `json:"suspend"`
	Capturing        bool                   `json:"capturing"`
	CaptureFileName  string                 `json:"capture_file_name"`
	CaptureFilePath  string                 `json:"capture_file_path"`
	CaptureComputeID string                 `json:"capture_compute_id"`
	LinkType         string                 `json:"link_type"`
}

// GetLinks gets the links associated with the GNS3 node.
func GetLinks(g GNS3Client, projectID, nodeID string) ([]Link, error) {
	if projectID == "" {
		return []Link{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return []Link{}, ErrEmptyNodeID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/links"
	link := []Link{}
	if err := get(g, path, 200, &link); err != nil {
		return []Link{}, err
	}
	return link, nil
}
