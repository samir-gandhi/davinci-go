package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = Renderer{}
)

type _Renderer Renderer
type Renderer struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Name                 *string                `davinci:"name,unmapped,omitempty"`
}

func (o Renderer) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Renderer) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Name != nil {
		result["name"] = o.Name
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Renderer) UnmarshalJSON(bytes []byte) (err error) {
	varRenderer := _Renderer{}

	if err = json.Unmarshal(bytes, &varRenderer); err == nil {
		*o = Renderer(varRenderer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o Renderer) DesignerCuesFields() []string {
	return []string{
		"Name",
	}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Renderer) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Renderer) FlowConfigFields() []string {
	return []string{}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o Renderer) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Renderer) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o Renderer) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
