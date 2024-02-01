package davinci

import "encoding/json"

type _Node Node
type Node struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Data                 NodeData               `json:"data,omitempty"`
	Position             Position               `json:"position,omitempty"`
	Group                string                 `json:"group"`
	Removed              bool                   `json:"removed"`
	Selected             bool                   `json:"selected"`
	Selectable           bool                   `json:"selectable"`
	Locked               bool                   `json:"locked"`
	Grabbable            bool                   `json:"grabbable"`
	Pannable             bool                   `json:"pannable"`
	Classes              string                 `json:"classes"`
}

func (o Node) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Node) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	jsonNode, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonNode, &result); err != nil {
		return nil, err
	}

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