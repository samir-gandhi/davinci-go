package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = GraphData{}
)

type _GraphData GraphData
type GraphData struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	BoxSelectionEnabled  *bool                  `davinci:"boxSelectionEnabled,unmapped,omitempty"`
	Data                 *Data                  `davinci:"data,unmapped,omitempty"`
	Elements             *Elements              `davinci:"elements,unmapped,omitempty"`
	MaxZoom              *float64               `davinci:"maxZoom,unmapped,omitempty"`
	MinZoom              *float64               `davinci:"minZoom,unmapped,omitempty"`
	Pan                  *Pan                   `davinci:"pan,unmapped,omitempty"`
	PanningEnabled       *bool                  `davinci:"panningEnabled,unmapped,omitempty"`
	Renderer             *Renderer              `davinci:"renderer,unmapped,omitempty"`
	UserPanningEnabled   *bool                  `davinci:"userPanningEnabled,unmapped,omitempty"`
	UserZoomingEnabled   *bool                  `davinci:"userZoomingEnabled,unmapped,omitempty"`
	Zoom                 *int32                 `davinci:"zoom,unmapped,omitempty"`
	ZoomingEnabled       *bool                  `davinci:"zoomingEnabled,unmapped,omitempty"`
}

func (o GraphData) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o GraphData) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.BoxSelectionEnabled != nil {
		result["boxSelectionEnabled"] = o.BoxSelectionEnabled
	}

	if o.Data != nil {
		result["data"] = o.Data
	}

	if o.Elements != nil {
		result["elements"] = o.Elements
	}

	if o.MaxZoom != nil {
		result["maxZoom"] = o.MaxZoom
	}

	if o.MinZoom != nil {
		result["minZoom"] = o.MinZoom
	}

	if o.Pan != nil {
		result["pan"] = o.Pan
	}

	if o.PanningEnabled != nil {
		result["panningEnabled"] = o.PanningEnabled
	}

	if o.Renderer != nil {
		result["renderer"] = o.Renderer
	}

	if o.UserPanningEnabled != nil {
		result["userPanningEnabled"] = o.UserPanningEnabled
	}

	if o.UserZoomingEnabled != nil {
		result["userZoomingEnabled"] = o.UserZoomingEnabled
	}

	if o.Zoom != nil {
		result["zoom"] = o.Zoom
	}

	if o.ZoomingEnabled != nil {
		result["zoomingEnabled"] = o.ZoomingEnabled
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *GraphData) UnmarshalJSON(bytes []byte) (err error) {
	varGraphData := _GraphData{}

	if err = json.Unmarshal(bytes, &varGraphData); err == nil {
		*o = GraphData(varGraphData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "boxSelectionEnabled")
		delete(additionalProperties, "data")
		delete(additionalProperties, "elements")
		delete(additionalProperties, "maxZoom")
		delete(additionalProperties, "minZoom")
		delete(additionalProperties, "pan")
		delete(additionalProperties, "panningEnabled")
		delete(additionalProperties, "renderer")
		delete(additionalProperties, "userPanningEnabled")
		delete(additionalProperties, "userZoomingEnabled")
		delete(additionalProperties, "zoom")
		delete(additionalProperties, "zoomingEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o GraphData) DesignerCuesFields() []string {
	return []string{
		"BoxSelectionEnabled",
		"MaxZoom",
		"MinZoom",
		"PanningEnabled",
		"UserPanningEnabled",
		"UserZoomingEnabled",
		"Zoom",
		"ZoomingEnabled",
	}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o GraphData) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o GraphData) FlowConfigFields() []string {
	return []string{}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o GraphData) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o GraphData) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o GraphData) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
