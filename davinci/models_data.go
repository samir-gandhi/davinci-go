package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = Data{}
)

type _Data Data
type Data struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	CapabilityName       *string                `davinci:"capabilityName,unmapped,omitempty"`
	ConnectionID         *string                `davinci:"connectionId,unmapped,omitempty"`
	ConnectorID          *string                `davinci:"connectorId,unmapped,omitempty"`
	ID                   *string                `davinci:"id,unmapped,omitempty"`
	Label                *string                `davinci:"label,unmapped,omitempty"`
	MultiValueSourceId   *string                `davinci:"multiValueSourceId,unmapped,omitempty"`
	Name                 *string                `davinci:"name,unmapped,omitempty"`
	NodeType             *string                `davinci:"nodeType,unmapped,omitempty"`
	Properties           *Properties            `davinci:"properties,unmapped,omitempty"`
	Source               *string                `davinci:"source,unmapped,omitempty"`
	Status               *string                `davinci:"status,unmapped,omitempty"`
	Target               *string                `davinci:"target,unmapped,omitempty"`
	Type                 *string                `davinci:"type,unmapped,omitempty"`
}

func (o Data) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Data) ToMap() (map[string]interface{}, error) {

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

func (o *Data) UnmarshalJSON(bytes []byte) (err error) {
	varData := _Data{}

	if err = json.Unmarshal(bytes, &varData); err == nil {
		*o = Data(varData)
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
func (o Data) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Data) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Data) FlowConfigFields() []string {
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
func (o Data) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Data) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o Data) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
