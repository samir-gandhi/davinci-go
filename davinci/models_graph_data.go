package davinci

type GraphData struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	BoxSelectionEnabled  *bool                  `davinci:"boxSelectionEnabled,designercue,omitempty"`
	Data                 *Data                  `davinci:"data,*,omitempty"`
	Elements             *Elements              `davinci:"elements,*,omitempty"`
	MaxZoom              *float64               `davinci:"maxZoom,designercue,omitempty"`
	MinZoom              *float64               `davinci:"minZoom,designercue,omitempty"`
	Pan                  *Pan                   `davinci:"pan,*,omitempty"`
	PanningEnabled       *bool                  `davinci:"panningEnabled,designercue,omitempty"`
	Renderer             *Renderer              `davinci:"renderer,*,omitempty"`
	UserPanningEnabled   *bool                  `davinci:"userPanningEnabled,designercue,omitempty"`
	UserZoomingEnabled   *bool                  `davinci:"userZoomingEnabled,designercue,omitempty"`
	Zoom                 *int32                 `davinci:"zoom,designercue,omitempty"`
	ZoomingEnabled       *bool                  `davinci:"zoomingEnabled,designercue,omitempty"`
}
