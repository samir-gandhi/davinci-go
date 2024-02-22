package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = LabelValue{}
)

type _LabelValue LabelValue
type LabelValue struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Label                *string                `davinci:"label,unmapped,omitempty"`
	Value                *string                `davinci:"value,unmapped,omitempty"`
}

func (o LabelValue) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o LabelValue) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Label != nil {
		result["label"] = o.Label
	}

	if o.Value != nil {
		result["value"] = o.Value
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *LabelValue) UnmarshalJSON(bytes []byte) (err error) {
	varLabelValue := _LabelValue{}

	if err = json.Unmarshal(bytes, &varLabelValue); err == nil {
		*o = LabelValue(varLabelValue)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o LabelValue) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o LabelValue) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o LabelValue) FlowConfigFields() []string {
	return []string{
		"Label",
		"Value",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o LabelValue) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o LabelValue) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o LabelValue) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
