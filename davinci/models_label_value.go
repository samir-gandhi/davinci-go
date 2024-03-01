package davinci

type LabelValue struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	Label                *string                `davinci:"label,config,omitempty"`
	Value                *string                `davinci:"value,config,omitempty"`
}
