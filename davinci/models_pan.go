package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = Pan{}
)

type _Pan Pan
type Pan struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	X                    *float64               `json:"x,omitempty"`
	Y                    *float64               `json:"y,omitempty"`
}

func (o Pan) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Pan) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.X != nil {
		result["x"] = o.X
	}

	if o.Y != nil {
		result["y"] = o.Y
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Pan) UnmarshalJSON(bytes []byte) (err error) {
	varPan := _Pan{}

	if err = json.Unmarshal(bytes, &varPan); err == nil {
		*o = Pan(varPan)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o Pan) DesignerCuesFields() []string {
	return []string{
		"X",
		"Y",
	}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Pan) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Pan) FlowConfigFields() []string {
	return []string{}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o Pan) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Pan) VersionMetadataFields() []string {
	return []string{}
}
