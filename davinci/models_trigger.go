package davinci

type Trigger struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	TriggerType          *string                `davinci:"type,config,omitempty"`
}
