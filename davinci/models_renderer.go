package davinci

type Renderer struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Name                 *string                `davinci:"name,designercue,omitempty"`
}
