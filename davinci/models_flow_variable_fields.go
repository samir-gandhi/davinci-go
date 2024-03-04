package davinci

type FlowVariableFields struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Type                 *string                `davinci:"type,config,omitempty"`
	DisplayName          *string                `davinci:"displayName,config,omitempty"`
	Mutable              *bool                  `davinci:"mutable,config,omitempty"`
	Value                *string                `davinci:"value,config,omitempty"`
	Min                  *int32                 `davinci:"min,config,omitempty"`
	Max                  *int32                 `davinci:"max,config,omitempty"`
}
