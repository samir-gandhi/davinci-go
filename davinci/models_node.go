package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = Node{}
)

type _Node Node
type Node struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Data                 *NodeData              `davinci:"data,unmapped,omitempty"`
	Position             *Position              `davinci:"position,unmapped,omitempty"`
	Group                string                 `davinci:"group,unmapped"`
	Removed              bool                   `davinci:"removed,unmapped"`
	Selected             bool                   `davinci:"selected,unmapped"`
	Selectable           bool                   `davinci:"selectable,unmapped"`
	Locked               bool                   `davinci:"locked,unmapped"`
	Grabbable            bool                   `davinci:"grabbable,unmapped"`
	Pannable             bool                   `davinci:"pannable,unmapped"`
	Classes              string                 `davinci:"classes,unmapped"`
}

func (o Node) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Node) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Data != nil {
		result["data"] = o.Data
	}

	if o.Position != nil {
		result["position"] = o.Position
	}

	result["group"] = o.Group
	result["removed"] = o.Removed
	result["selected"] = o.Selected
	result["selectable"] = o.Selectable
	result["locked"] = o.Locked
	result["grabbable"] = o.Grabbable
	result["pannable"] = o.Pannable
	result["classes"] = o.Classes

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Node) UnmarshalJSON(bytes []byte) (err error) {
	varNode := _Node{}

	if err = json.Unmarshal(bytes, &varNode); err == nil {
		*o = Node(varNode)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "position")
		delete(additionalProperties, "group")
		delete(additionalProperties, "removed")
		delete(additionalProperties, "selected")
		delete(additionalProperties, "selectable")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "grabbable")
		delete(additionalProperties, "pannable")
		delete(additionalProperties, "classes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o Node) DesignerCuesFields() []string {
	return []string{
		"Group",
		"Removed",
		"Selected",
		"Selectable",
		"Locked",
		"Grabbable",
		"Pannable",
	}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Node) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Node) FlowConfigFields() []string {
	return []string{
		"Classes",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o Node) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Node) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o Node) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
