package davinci

type FlowInfo struct {
	Flow Flow `json:"flowInfo"`
}
type FlowsInfo struct {
	Flow []Flow `json:"flowsInfo,omitempty"`
}

type FlowImport struct {
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
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

// Used specifically for PUTs to existing flows.
type FlowUpdate struct {
	CurrentVersion *int32        `json:"currentVersion,omitempty"`
	Name           *string       `json:"name,omitempty"`
	Description    *string       `json:"description,omitempty"`
	Settings       interface{}   `json:"settings,omitempty"`
	Trigger        *Trigger      `json:"trigger,omitempty"`
	GraphData      *GraphData    `json:"graphData,omitempty"`
	InputSchema    []interface{} `json:"inputSchema,omitempty"`
	OutputSchema   *OutputSchema `json:"outputSchema,omitempty"`
}

type OutputSchema struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Output               interface{}            `json:"output,omitempty"`
}

//	type ShowContinueButton struct {
//		Value bool `json:"value,omitempty"`
//	}
type LabelValue struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Label                *string                `json:"label,omitempty"`
	Value                *string                `json:"value,omitempty"`
}

type SubFlowValue LabelValue

type SubFlowID struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowValue          `json:"value,omitempty"`
}

// Used for type assertion on Properties of a Node Data
type SubFlowProperties struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	SubFlowID            *SubFlowID             `json:"subFlowId,omitempty"`
	SubFlowVersionID     *SubFlowVersionID      `json:"subFlowVersionId,omitempty"`
}

type SaveFlowVariables struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Value                []FlowVariable         `json:"value,omitempty"`
}

type AdditionalProperties map[string]interface{}
