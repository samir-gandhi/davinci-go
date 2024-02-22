package davinci

type FlowInfo struct {
	Flow Flow `json:"flowInfo"`
}
type FlowsInfo struct {
	Flow []Flow `json:"flowsInfo,omitempty"`
}

type FlowImport struct {
	Name            *string           `json:"name,omitempty"`
	Description     *string           `json:"description,omitempty"`
	FlowInfo        *Flow             `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}
type FlowsImport struct {
	Name            *string           `json:"name,omitempty"`
	Description     *string           `json:"description,omitempty"`
	FlowInfo        *Flows            `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}

type Flows struct {
	Flow []Flow `json:"flows,omitempty"`
}

// Used specifically for PUTs to existing flows.
type FlowUpdate struct {
	FlowUpdateConfiguration
	CurrentVersion *int32  `json:"currentVersion,omitempty"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
}

//	type ShowContinueButton struct {
//		Value *bool `json:"value,omitempty"`
//	}

type AdditionalProperties map[string]interface{}
