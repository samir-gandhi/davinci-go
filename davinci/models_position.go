package davinci

type Position struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	X                    *float64               `davinci:"x,designercue,omitempty"`
	Y                    *float64               `davinci:"y,designercue,omitempty"`
}
