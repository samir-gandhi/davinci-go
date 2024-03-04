package davinci

type Elements struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Nodes                []Node                 `davinci:"nodes,*,omitempty"`
	Edges                []Edge                 `davinci:"edges,*,omitempty"`
}
