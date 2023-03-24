package davinci

type FlowInfo struct {
	Flow Flow `json:"flowInfo,omitempty"`
}
type FlowsInfo struct {
	Flow []Flow `json:"flowsInfo,omitempty"`
}

type FlowImport struct {
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description"`
	FlowInfo        Flow              `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}
type FlowsImport struct {
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	FlowInfo        Flows             `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}

type Flows struct {
	Flow []Flow `json:"flows,omitempty"`
}

type Flow struct {
	CustomerID         string        `json:"customerId,omitempty"`
	FlowStatus         string        `json:"flowStatus,omitempty"`
	CurrentVersion     int           `json:"currentVersion,omitempty"`
	PublishedVersion   int           `json:"publishedVersion,omitempty"`
	Name               string        `json:"name,omitempty"`
	Description        string        `json:"description,omitempty"`
	CreatedDate        int64         `json:"createdDate,omitempty"`
	UpdatedDate        int64         `json:"updatedDate,omitempty"`
	AuthTokenExpireIds []interface{} `json:"authTokenExpireIds,omitempty"`
	DeployedDate       int64         `json:"deployedDate,omitempty"`
	// Edited, removed struct. Staying one level only
	EnabledGraphData     interface{} `json:"enabledGraphData,omitempty"`
	FunctionConnectionID interface{} `json:"functionConnectionId,omitempty"`
	// edited
	InputSchemaCompiled interface{}   `json:"inputSchemaCompiled,omitempty"`
	IsInputSchemaSaved  bool          `json:"isInputSchemaSaved,omitempty"`
	IsOutputSchemaSaved bool          `json:"isOutputSchemaSaved,omitempty"`
	Orx                 string        `json:"orx,omitempty"`
	Settings            interface{}   `json:"settings,omitempty"`
	Trigger             Trigger       `json:"trigger"`
	Timeouts            interface{}   `json:"timeouts,omitempty"`
	FlowID              string        `json:"flowId,omitempty"`
	CompanyID           string        `json:"companyId,omitempty"`
	GraphData           GraphData     `json:"graphData,omitempty"`
	InputSchema         []interface{} `json:"inputSchema"`
	OutputSchema        OutputSchema  `json:"outputSchema"`
	//compiled is used in exported flow json, must be converted to JUST output when updating flow.
	OutputSchemaCompiled OutputSchema `json:"outputSchemaCompiled"`
	FlowColor            string       `json:"flowColor,omitempty"`
	ConnectorIds         []string     `json:"connectorIds,omitempty"`
	SavedDate            int64        `json:"savedDate,omitempty"`
}

// Used specifically for PUTs to existing flows.
type FlowUpdate struct {
	CurrentVersion int           `json:"currentVersion,omitempty"`
	Name           string        `json:"name,omitempty"`
	Description    string        `json:"description,omitempty"`
	Settings       interface{}   `json:"settings,omitempty"`
	Trigger        Trigger       `json:"trigger"`
	GraphData      GraphData     `json:"graphData,omitempty"`
	InputSchema    []interface{} `json:"inputSchema"`
	OutputSchema   OutputSchema  `json:"outputSchema"`
}

type OutputSchema struct {
	Output interface{} `json:"output,omitempty"`
}

type Trigger struct {
	TriggerType string `json:"type,omitempty"`
}

//	type ShowContinueButton struct {
//		Value bool `json:"value,omitempty"`
//	}
type SubFlowValue struct {
	Label string `json:"label,omitempty" mapstructure:"label"`
	Value string `json:"value,omitempty" mapstructure:"value"`
}
type SubFlowID struct {
	Value SubFlowValue `json:"value,omitempty" mapstructure:"value"`
}
type SubFlowVersionID struct {
	Value string `json:"value,omitempty" mapstructure:"value"`
}

// Used for type assertion on Properties of a Node Data
type SubFlowProperties struct {
	SubFlowID        SubFlowID        `json:"subFlowId,omitempty" mapstructure:"subFlowId"`
	SubFlowVersionID SubFlowVersionID `json:"subFlowVersionId,omitempty" mapstructure:"subFlowVersionId"`
}

type Data struct {
	ID             string `json:"id,omitempty"`
	NodeType       string `json:"nodeType,omitempty"`
	ConnectionID   string `json:"connectionId,omitempty"`
	ConnectorID    string `json:"connectorId,omitempty"`
	Name           string `json:"name,omitempty"`
	Label          string `json:"label,omitempty"`
	Status         string `json:"status,omitempty"`
	CapabilityName string `json:"capabilityName,omitempty"`
	Type           string `json:"type,omitempty"`
	// have not removed omitempty on general Data struct, not sure if it is needed
	Properties         Properties `json:"properties,omitempty"`
	Source             string     `json:"source,omitempty"`
	Target             string     `json:"target,omitempty"`
	MultiValueSourceId string     `json:"multiValueSourceId,omitempty"`
}

type NodeData struct {
	ID             string     `json:"id,omitempty"`
	NodeType       string     `json:"nodeType,omitempty"`
	ConnectionID   string     `json:"connectionId,omitempty"`
	ConnectorID    string     `json:"connectorId,omitempty"`
	Name           string     `json:"name,omitempty"`
	Label          string     `json:"label,omitempty"`
	Status         string     `json:"status,omitempty"`
	CapabilityName string     `json:"capabilityName,omitempty"`
	Type           string     `json:"type,omitempty"`
	Properties     Properties `json:"properties"`
	Source         string     `json:"source,omitempty"`
	Target         string     `json:"target,omitempty"`
}

type EdgeData struct {
	ID                 string     `json:"id,omitempty"`
	NodeType           string     `json:"nodeType,omitempty"`
	ConnectionID       string     `json:"connectionId,omitempty"`
	ConnectorID        string     `json:"connectorId,omitempty"`
	Name               string     `json:"name,omitempty"`
	Label              string     `json:"label,omitempty"`
	Status             string     `json:"status,omitempty"`
	CapabilityName     string     `json:"capabilityName,omitempty"`
	Type               string     `json:"type,omitempty"`
	Properties         Properties `json:"properties,omitempty"`
	Source             string     `json:"source,omitempty"`
	Target             string     `json:"target,omitempty"`
	MultiValueSourceId string     `json:"multiValueSourceId,omitempty"`
}

type Position struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type Nodes struct {
	Data       NodeData `json:"data,omitempty"`
	Position   Position `json:"position,omitempty"`
	Group      string   `json:"group"`
	Removed    bool     `json:"removed"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbable  bool     `json:"grabbable"`
	Pannable   bool     `json:"pannable"`
	Classes    string   `json:"classes"`
}

type Edges struct {
	Data       EdgeData `json:"data,omitempty"`
	Position   Position `json:"position,omitempty"`
	Group      string   `json:"group"`
	Removed    bool     `json:"removed"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbable  bool     `json:"grabbable"`
	Pannable   bool     `json:"pannable"`
	Classes    string   `json:"classes"`
}

type Elements struct {
	Nodes []Nodes `json:"nodes,omitempty"`
	Edges []Edges `json:"edges,omitempty"`
}

type Pan struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type Renderer struct {
	Name string `json:"name,omitempty"`
}

type GraphData struct {
	Elements            Elements `json:"elements,omitempty"`
	Data                Data     `json:"data,omitempty"`
	ZoomingEnabled      bool     `json:"zoomingEnabled,omitempty"`
	UserZoomingEnabled  bool     `json:"userZoomingEnabled,omitempty"`
	Zoom                int      `json:"zoom,omitempty"`
	MinZoom             float64  `json:"minZoom,omitempty"`
	MaxZoom             float64  `json:"maxZoom,omitempty"`
	PanningEnabled      bool     `json:"panningEnabled,omitempty"`
	UserPanningEnabled  bool     `json:"userPanningEnabled,omitempty"`
	Pan                 Pan      `json:"pan,omitempty"`
	BoxSelectionEnabled bool     `json:"boxSelectionEnabled,omitempty"`
	Renderer            Renderer `json:"renderer,omitempty"`
}

// type GraphData struct {
// 	Elements            interface{} `json:"elements"`
// 	Data                interface{} `json:"data"`
// 	ZoomingEnabled      bool        `json:"zoomingEnabled"`
// 	UserZoomingEnabled  bool        `json:"userZoomingEnabled"`
// 	Zoom                int         `json:"zoom"`
// 	MinZoom             float64     `json:"minZoom"`
// 	MaxZoom             float64     `json:"maxZoom"`
// 	PanningEnabled      bool        `json:"panningEnabled"`
// 	UserPanningEnabled  bool        `json:"userPanningEnabled"`
// 	Pan                 interface{} `json:"pan"`
// 	BoxSelectionEnabled bool        `json:"boxSelectionEnabled"`
// 	Renderer            interface{} `json:"renderer"`
// }
