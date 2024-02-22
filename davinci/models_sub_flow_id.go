package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = SubFlowID{}
)

type _SubFlowID SubFlowID
type SubFlowID struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowValue          `davinci:"value,unmapped,omitempty"`
}

func (o SubFlowID) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowID) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SubFlowID) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowID := _SubFlowID{}

	if err = json.Unmarshal(bytes, &varSubFlowID); err == nil {
		*o = SubFlowID(varSubFlowID)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o SubFlowID) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o SubFlowID) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o SubFlowID) FlowConfigFields() []string {
	return []string{
		"Value",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o SubFlowID) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o SubFlowID) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o SubFlowID) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
