package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = SaveFlowVariables{}
)

type _SaveFlowVariables SaveFlowVariables
type SaveFlowVariables struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Value                []FlowVariable         `json:"value,omitempty"`
}

func (o SaveFlowVariables) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SaveFlowVariables) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	result["value"] = o.Value
	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SaveFlowVariables) UnmarshalJSON(bytes []byte) (err error) {
	varSaveFlowVariables := _SaveFlowVariables{}

	if err = json.Unmarshal(bytes, &varSaveFlowVariables); err == nil {
		*o = SaveFlowVariables(varSaveFlowVariables)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o SaveFlowVariables) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o SaveFlowVariables) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o SaveFlowVariables) FlowConfigFields() []string {
	return []string{
		"Value",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o SaveFlowVariables) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o SaveFlowVariables) VersionMetadataFields() []string {
	return []string{}
}
