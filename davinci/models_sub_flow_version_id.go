package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = SubFlowVersionID{}
)

type _SubFlowVersionID SubFlowVersionID
type SubFlowVersionID struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowVersionIDValue `davinci:"value,unmapped,omitempty"`
}

func (o SubFlowVersionID) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowVersionID) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SubFlowVersionID) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowVersionID := _SubFlowVersionID{}

	if err = json.Unmarshal(bytes, &varSubFlowVersionID); err == nil {
		*o = SubFlowVersionID(varSubFlowVersionID)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o SubFlowVersionID) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o SubFlowVersionID) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o SubFlowVersionID) FlowConfigFields() []string {
	return []string{
		"Value",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o SubFlowVersionID) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o SubFlowVersionID) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o SubFlowVersionID) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
