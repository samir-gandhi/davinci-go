package davinci

import "encoding/json"

type _NodeData NodeData
type NodeData struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	CapabilityName       string                 `json:"capabilityName,omitempty"`
	ConnectionID         string                 `json:"connectionId,omitempty"`
	ConnectorID          string                 `json:"connectorId,omitempty"`
	ID                   string                 `json:"id,omitempty"`
	Label                string                 `json:"label,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	NodeType             string                 `json:"nodeType,omitempty"`
	Properties           *Properties            `json:"properties,omitempty"`
	Source               string                 `json:"source,omitempty"`
	Status               string                 `json:"status,omitempty"`
	Target               string                 `json:"target,omitempty"`
	Type                 string                 `json:"type,omitempty"`
}

func (o NodeData) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o NodeData) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	result["capabilityName"] = o.CapabilityName
	result["connectionId"] = o.ConnectionID
	result["connectorId"] = o.ConnectorID
	result["id"] = o.ID
	result["label"] = o.Label
	result["name"] = o.Name
	result["nodeType"] = o.NodeType

	if o.Properties != nil {
		result["properties"] = o.Properties
	}

	result["source"] = o.Source
	result["status"] = o.Status
	result["target"] = o.Target
	result["type"] = o.Type

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *NodeData) UnmarshalJSON(bytes []byte) (err error) {
	varNodeData := _NodeData{}

	if err = json.Unmarshal(bytes, &varNodeData); err == nil {
		*o = NodeData(varNodeData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilityName")
		delete(additionalProperties, "connectionId")
		delete(additionalProperties, "connectorId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nodeType")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "source")
		delete(additionalProperties, "status")
		delete(additionalProperties, "target")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
