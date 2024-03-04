package davinci

type SubFlowProperties struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	SubFlowID            *SubFlowID             `davinci:"subFlowId,*,omitempty"`
	SubFlowVersionID     *SubFlowVersionID      `davinci:"subFlowVersionId,*,omitempty"`
}
