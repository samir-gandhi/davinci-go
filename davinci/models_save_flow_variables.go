package davinci

type SaveFlowVariables struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                []FlowVariable         `davinci:"value,*,omitempty"`
}
