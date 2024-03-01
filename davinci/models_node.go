package davinci

type Node struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	Data                 *NodeData              `davinci:"data,*,omitempty"`
	Position             *Position              `davinci:"position,*,omitempty"`
	Group                string                 `davinci:"group,designercue"`
	Removed              bool                   `davinci:"removed,designercue"`
	Selected             bool                   `davinci:"selected,designercue"`
	Selectable           bool                   `davinci:"selectable,designercue"`
	Locked               bool                   `davinci:"locked,designercue"`
	Grabbable            bool                   `davinci:"grabbable,designercue"`
	Pannable             bool                   `davinci:"pannable,designercue"`
	Classes              string                 `davinci:"classes,config"`
}
