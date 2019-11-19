package gons3

import (
	"errors"
	"net/url"
)

// ErrEmptyLinkID means that the link id cannot be empty.
var ErrEmptyLinkID = errors.New("linkID cannot be empty")

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

// GetNodeLinks gets the links associated with the GNS3 node.
func GetNodeLinks(g GNS3Client, projectID, nodeID string) ([]Link, error) {
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

// GetLinks gets the links associated with the GNS3 project.
func GetLinks(g GNS3Client, projectID string) ([]Link, error) {
	if projectID == "" {
		return []Link{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/links"
	link := []Link{}
	if err := get(g, path, 200, &link); err != nil {
		return []Link{}, err
	}
	return link, nil
}

// GetLink gets the link associated with the GNS3 project.
func GetLink(g GNS3Client, projectID, linkID string) (Link, error) {
	if projectID == "" {
		return Link{}, ErrEmptyProjectID
	}
	if linkID == "" {
		return Link{}, ErrEmptyLinkID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/links/" + url.PathEscape(linkID)
	link := Link{}
	if err := get(g, path, 200, &link); err != nil {
		return Link{}, err
	}
	return link, nil
}

// CreateLink creates the link associated with the GNS3 project.
func CreateLink(g GNS3Client, projectID string, l LinkCreator) (Link, error) {
	if projectID == "" {
		return Link{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/links"
	link := Link{}
	if err := post(g, path, 201, l.values, &link); err != nil {
		return Link{}, err
	}
	return link, nil
}

// DeleteLink deletes the link.
func DeleteLink(g GNS3Client, projectID, linkID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}
	if linkID == "" {
		return ErrEmptyLinkID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/links/" + url.PathEscape(linkID)
	if err := delete(g, path, 204, nil); err != nil {
		return err
	}
	return nil
}
