package davinci

type SubFlowID struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowValue          `davinci:"value,*,omitempty"`
}
