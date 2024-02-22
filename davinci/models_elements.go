package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = Elements{}
)

type _Elements Elements
type Elements struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Nodes                []Node                 `davinci:"nodes,unmapped,omitempty"`
	Edges                []Edge                 `davinci:"edges,unmapped,omitempty"`
}

// EmptyObject implements DaVinciExportModel.
func (Elements) EmptyObject() any {
	return Elements{
		Nodes: func() []Node {
			r := make([]Node, 0)
			return r
		}(),
		Edges: func() []Edge {
			r := make([]Edge, 0)
			return r
		}(),
	}
}

func (o Elements) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Elements) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Nodes != nil {
		result["nodes"] = o.Nodes
	} else {
		result["nodes"] = make([]Node, 0)
	}

	if o.Edges != nil {
		result["edges"] = o.Edges
	} else {
		result["edges"] = make([]Edge, 0)
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Elements) UnmarshalJSON(bytes []byte) (err error) {
	varElements := _Elements{}

	if err = json.Unmarshal(bytes, &varElements); err == nil {
		*o = Elements(varElements)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nodes")
		delete(additionalProperties, "edges")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o Elements) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Elements) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Elements) FlowConfigFields() []string {
	return []string{}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o Elements) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Elements) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o Elements) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
