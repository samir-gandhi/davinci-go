package davinci

import "encoding/json"

type _GraphData GraphData
type GraphData struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	BoxSelectionEnabled  bool                   `json:"boxSelectionEnabled,omitempty"`
	Data                 Data                   `json:"data,omitempty"`
	Elements             Elements               `json:"elements,omitempty"`
	MaxZoom              float64                `json:"maxZoom,omitempty"`
	MinZoom              float64                `json:"minZoom,omitempty"`
	Pan                  Pan                    `json:"pan,omitempty"`
	PanningEnabled       bool                   `json:"panningEnabled,omitempty"`
	Renderer             Renderer               `json:"renderer,omitempty"`
	UserPanningEnabled   bool                   `json:"userPanningEnabled,omitempty"`
	UserZoomingEnabled   bool                   `json:"userZoomingEnabled,omitempty"`
	Zoom                 int                    `json:"zoom,omitempty"`
	ZoomingEnabled       bool                   `json:"zoomingEnabled,omitempty"`
}

func (o GraphData) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o GraphData) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	jsonGraphData, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonGraphData, &result); err != nil {
		return nil, err
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