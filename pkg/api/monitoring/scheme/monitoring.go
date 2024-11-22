package scheme

// AgentScheme is for information of monitoring agent
type AgentScheme struct {
	Name     string `json:"name"        binding:"required"`
	Endpoint string `json:"endpoint"        binding:"required"`
}

// DataScheme is for information of data
type DataScheme struct {
	Type     []string `json:"type"        binding:"required"`
	Interval int      `json:"interval"        binding:"required"`
}

// SubscribeScheme is for information of MLApp
type SubscribeScheme struct {
	Agent string     `json:"agent"        binding:"required"`
	Name  string     `json:"name"        binding:"required"`
	Data  DataScheme `json:"data"        binding:"required"`
}

// UnsubscribeScheme is for information of MLApp
type UnsubscribeScheme struct {
	Agent string `json:"agent"        binding:"required"`
	Name  string `json:"name"        binding:"required"`
}
