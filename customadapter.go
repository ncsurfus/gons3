package gons3

// CustomAdapter models a GNS3 custom adapter.
type CustomAdapter struct {
	AdapterNumber int    `json:"adapter_number"`
	PortName      string `json:"port_name"`
	AdapterType   string `json:"adapter_type"`
	MACAddress    string `json:"mac_address"`
}

// NewCustomAdapter models a new GNS3 label.
// GNS3 schema requires values: AdapterNumber
type NewCustomAdapter struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (c *NewCustomAdapter) SetProperty(name string, value interface{}) {
	if c.values == nil {
		c.values = map[string]interface{}{}
	}
	c.values[name] = value
}

// SetAdapterNumber sets the adapter number.
func (c *NewCustomAdapter) SetAdapterNumber(number int) {
	c.SetProperty("adapter_number", number)
}

// SetPortName sets the name of the port.
func (c *NewCustomAdapter) SetPortName(portName string) {
	c.SetProperty("port_name", portName)
}

// SetAdapterType sets the type of adapter.
func (c *NewCustomAdapter) SetAdapterType(adapterType string) {
	c.SetProperty("adapter_type", adapterType)
}

// SetMACAddress sets a custom MAC address for the port.
func (c *NewCustomAdapter) SetMACAddress(macAddress string) {
	c.SetProperty("mac_address", macAddress)
}
