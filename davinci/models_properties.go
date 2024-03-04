package davinci

type Properties struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Form                 *string                `davinci:"form,config,omitempty"`
	SubFlowID            *SubFlowID             `davinci:"subFlowId,config,omitempty"`
	SubFlowVersionID     *SubFlowVersionID      `davinci:"subFlowVersionId,config,omitempty"`
	SaveFlowVariables    *SaveFlowVariables     `davinci:"saveFlowVariables,config,omitempty"`
}
