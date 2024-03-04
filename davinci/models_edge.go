package davinci

type Edge struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Data                 *Data                  `davinci:"data,*,omitempty"`
	Position             *Position              `davinci:"position,*,omitempty"`
	Group                *string                `davinci:"group,designercue,omitempty"`
	Removed              *bool                  `davinci:"removed,designercue,omitempty"`
	Selected             *bool                  `davinci:"selected,designercue,omitempty"`
	Selectable           *bool                  `davinci:"selectable,designercue,omitempty"`
	Locked               *bool                  `davinci:"locked,designercue,omitempty"`
	Grabbable            *bool                  `davinci:"grabbable,designercue,omitempty"`
	Pannable             *bool                  `davinci:"pannable,designercue,omitempty"`
	Classes              *string                `davinci:"classes,config,omitempty"`
}
