package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = Properties{}
)

type _Properties Properties
type Properties struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Form                 *string                `davinci:"form,unmapped,omitempty"`
	SubFlowID            *SubFlowID             `davinci:"subFlowId,unmapped,omitempty"`
	SubFlowVersionID     *SubFlowVersionID      `davinci:"subFlowVersionId,unmapped,omitempty"`
	SaveFlowVariables    *SaveFlowVariables     `davinci:"saveFlowVariables,unmapped,omitempty"`
}

func (o Properties) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Properties) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Form != nil {
		result["form"] = o.Form
	}

	if o.SubFlowID != nil {
		result["subFlowId"] = o.SubFlowID
	}

	if o.SubFlowVersionID != nil {
		result["subFlowVersionId"] = o.SubFlowVersionID
	}

	if o.SaveFlowVariables != nil {
		result["saveFlowVariables"] = o.SaveFlowVariables
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Properties) UnmarshalJSON(bytes []byte) (err error) {
	varProperties := _Properties{}

	if err = json.Unmarshal(bytes, &varProperties); err == nil {
		*o = Properties(varProperties)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "form")
		delete(additionalProperties, "subFlowId")
		delete(additionalProperties, "subFlowVersionId")
		delete(additionalProperties, "saveFlowVariables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o Properties) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Properties) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Properties) FlowConfigFields() []string {
	return []string{
		"Form",
		"SubFlowID",
		"SubFlowVersionID",
		"SaveFlowVariables",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o Properties) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Properties) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o Properties) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
