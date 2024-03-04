package davinci

type OutputSchema struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Output               interface{}            `davinci:"output,config,omitempty"`
}
