package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = SubFlowProperties{}
)

type _SubFlowProperties SubFlowProperties
type SubFlowProperties struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	SubFlowID            *SubFlowID             `davinci:"subFlowId,unmapped,omitempty"`
	SubFlowVersionID     *SubFlowVersionID      `davinci:"subFlowVersionId,unmapped,omitempty"`
}

func (o SubFlowProperties) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowProperties) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.SubFlowID != nil {
		result["subFlowId"] = o.SubFlowID
	}

	if o.SubFlowVersionID != nil {
		result["subFlowVersionId"] = o.SubFlowVersionID
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SubFlowProperties) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowProperties := _SubFlowProperties{}

	if err = json.Unmarshal(bytes, &varSubFlowProperties); err == nil {
		*o = SubFlowProperties(varSubFlowProperties)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "subFlowId")
		delete(additionalProperties, "subFlowVersionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o SubFlowProperties) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o SubFlowProperties) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o SubFlowProperties) FlowConfigFields() []string {
	return []string{}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o SubFlowProperties) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o SubFlowProperties) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o SubFlowProperties) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
