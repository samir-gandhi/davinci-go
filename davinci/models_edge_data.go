package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = EdgeData{}
)

type _EdgeData EdgeData
type EdgeData struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	CapabilityName       *string                `json:"capabilityName,omitempty"`
	ConnectionID         *string                `json:"connectionId,omitempty"`
	ConnectorID          *string                `json:"connectorId,omitempty"`
	ID                   *string                `json:"id,omitempty"`
	Label                *string                `json:"label,omitempty"`
	MultiValueSourceId   *string                `json:"multiValueSourceId,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	NodeType             *string                `json:"nodeType,omitempty"`
	Properties           *Properties            `json:"properties,omitempty"`
	Source               *string                `json:"source,omitempty"`
	Status               *string                `json:"status,omitempty"`
	Target               *string                `json:"target,omitempty"`
	Type                 *string                `json:"type,omitempty"`
}

func (o EdgeData) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o EdgeData) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.CapabilityName != nil {
		result["capabilityName"] = o.CapabilityName
	}

	if o.ConnectionID != nil {
		result["connectionId"] = o.ConnectionID
	}

	if o.ConnectorID != nil {
		result["connectorId"] = o.ConnectorID
	}

	if o.ID != nil {
		result["id"] = o.ID
	}

	if o.Label != nil {
		result["label"] = o.Label
	}

	if o.MultiValueSourceId != nil {
		result["multiValueSourceId"] = o.MultiValueSourceId
	}

	if o.Name != nil {
		result["name"] = o.Name
	}

	if o.NodeType != nil {
		result["nodeType"] = o.NodeType
	}

	if o.Properties != nil {
		result["properties"] = o.Properties
	}

	if o.Source != nil {
		result["source"] = o.Source
	}

	if o.Status != nil {
		result["status"] = o.Status
	}

	if o.Target != nil {
		result["target"] = o.Target
	}

	if o.Type != nil {
		result["type"] = o.Type
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *EdgeData) UnmarshalJSON(bytes []byte) (err error) {
	varEdgeData := _EdgeData{}

	if err = json.Unmarshal(bytes, &varEdgeData); err == nil {
		*o = EdgeData(varEdgeData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilityName")
		delete(additionalProperties, "connectionId")
		delete(additionalProperties, "connectorId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "multiValueSourceId")
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

// DesignerCuesFields implements DaVinciExportModel.
func (o EdgeData) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o EdgeData) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o EdgeData) FlowConfigFields() []string {
	return []string{
		"CapabilityName",
		"ConnectionID",
		"ConnectorID",
		"ID",
		"Label",
		"MultiValueSourceId",
		"Name",
		"NodeType",
		"Source",
		"Status",
		"Target",
		"Type",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o EdgeData) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o EdgeData) VersionMetadataFields() []string {
	return []string{}
}
