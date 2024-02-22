package davinci

import (
	"encoding/json"
)

var (
	_ DaVinciExportModel = Flow{}
)

type _Flow Flow
type Flow struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	FlowConfiguration
	FlowEnvironmentMetadata
	FlowMetadata
	FlowVersionMetadata
}

type FlowConfiguration struct {
	FlowUpdateConfiguration
	FlowColor            *string       `davinci:"flowColor,unmapped,omitempty"`
	InputSchemaCompiled  interface{}   `davinci:"inputSchemaCompiled,unmapped,omitempty"`
	IsInputSchemaSaved   *bool         `davinci:"isInputSchemaSaved,unmapped,omitempty"`
	IsOutputSchemaSaved  *bool         `davinci:"isOutputSchemaSaved,unmapped,omitempty"`
	OutputSchemaCompiled *OutputSchema `davinci:"outputSchemaCompiled,unmapped,omitempty"` //compiled is used in exported flow json, must be converted to JUST output when updating flow.
}

type FlowUpdateConfiguration struct {
	GraphData    *GraphData    `davinci:"graphData,unmapped,omitempty"`
	InputSchema  []interface{} `davinci:"inputSchema,unmapped,omitempty"`
	OutputSchema *OutputSchema `davinci:"outputSchema,unmapped,omitempty"`
	Settings     interface{}   `davinci:"settings,unmapped,omitempty"`
	Trigger      *Trigger      `davinci:"trigger,unmapped,omitempty"`
}

type FlowEnvironmentMetadata struct {
	CompanyID   string    `davinci:"companyId,unmapped"`
	CreatedDate EpochTime `davinci:"createdDate,unmapped"`
	CustomerID  string    `davinci:"customerId,unmapped"`
	FlowID      string    `davinci:"flowId,unmapped"`
}

type FlowMetadata struct {
	AuthTokenExpireIds   []interface{}  `davinci:"authTokenExpireIds,unmapped,omitempty"`
	Connections          []interface{}  `davinci:"connections,unmapped,omitempty"`
	ConnectorIds         []string       `davinci:"connectorIds,unmapped"`
	Description          *string        `davinci:"description,unmapped,omitempty"`
	EnabledGraphData     interface{}    `davinci:"enabledGraphData,unmapped,omitempty"`
	FunctionConnectionID interface{}    `davinci:"functionConnectionId,unmapped,omitempty"`
	Name                 string         `davinci:"name,unmapped"`
	Orx                  *string        `davinci:"orx,unmapped,omitempty"`
	Timeouts             interface{}    `davinci:"timeouts,unmapped,omitempty"`
	Variables            []FlowVariable `davinci:"variables,unmapped,omitempty"`
}

type FlowVersionMetadata struct {
	CurrentVersion   *int32     `davinci:"currentVersion,unmapped,omitempty"`
	DeployedDate     *EpochTime `davinci:"deployedDate,unmapped,omitempty"`
	FlowStatus       string     `davinci:"flowStatus,unmapped"`
	PublishedVersion *int32     `davinci:"publishedVersion,unmapped,omitempty"`
	SavedDate        EpochTime  `davinci:"savedDate,unmapped"`
	UpdatedDate      *EpochTime `davinci:"updatedDate,unmapped,omitempty"`
	VersionID        int32      `davinci:"versionId,unmapped"`
}

func (o Flow) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Flow) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	result["authTokenExpireIds"] = o.AuthTokenExpireIds
	result["companyId"] = o.CompanyID
	result["connections"] = o.Connections
	result["connectorIds"] = o.ConnectorIds
	result["createdDate"] = o.CreatedDate

	if o.CurrentVersion != nil {
		result["currentVersion"] = o.CurrentVersion
	}

	result["customerId"] = o.CustomerID

	if o.DeployedDate != nil {
		result["deployedDate"] = o.DeployedDate
	}

	if o.Description != nil {
		result["description"] = o.Description
	}

	result["enabledGraphData"] = o.EnabledGraphData

	if o.FlowColor != nil {
		result["flowColor"] = o.FlowColor
	}

	result["flowId"] = o.FlowID
	result["flowStatus"] = o.FlowStatus
	result["functionConnectionId"] = o.FunctionConnectionID
	result["graphData"] = o.GraphData
	result["inputSchema"] = o.InputSchema
	result["inputSchemaCompiled"] = o.InputSchemaCompiled

	if o.IsInputSchemaSaved != nil {
		result["isInputSchemaSaved"] = o.IsInputSchemaSaved
	}

	if o.IsOutputSchemaSaved != nil {
		result["isOutputSchemaSaved"] = o.IsOutputSchemaSaved
	}

	result["name"] = o.Name

	if o.Orx != nil {
		result["orx"] = o.Orx
	}

	if o.OutputSchema != nil {
		result["outputSchema"] = o.OutputSchema
	}

	if o.OutputSchemaCompiled != nil {
		result["outputSchemaCompiled"] = o.OutputSchemaCompiled
	}

	if o.PublishedVersion != nil {
		result["publishedVersion"] = o.PublishedVersion
	}

	result["savedDate"] = o.SavedDate
	result["settings"] = o.Settings
	result["timeouts"] = o.Timeouts

	if o.Trigger != nil {
		result["trigger"] = o.Trigger
	}

	if o.UpdatedDate != nil {
		result["updatedDate"] = o.UpdatedDate
	}

	result["variables"] = o.Variables
	result["versionId"] = o.VersionID

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Flow) UnmarshalJSON(bytes []byte) (err error) {
	varFlow := _Flow{}

	if err = json.Unmarshal(bytes, &varFlow); err == nil {
		*o = Flow(varFlow)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authTokenExpireIds")
		delete(additionalProperties, "companyId")
		delete(additionalProperties, "connections")
		delete(additionalProperties, "connectorIds")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "currentVersion")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "deployedDate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabledGraphData")
		delete(additionalProperties, "flowColor")
		delete(additionalProperties, "flowId")
		delete(additionalProperties, "flowStatus")
		delete(additionalProperties, "functionConnectionId")
		delete(additionalProperties, "graphData")
		delete(additionalProperties, "inputSchema")
		delete(additionalProperties, "inputSchemaCompiled")
		delete(additionalProperties, "isInputSchemaSaved")
		delete(additionalProperties, "isOutputSchemaSaved")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orx")
		delete(additionalProperties, "outputSchema")
		delete(additionalProperties, "outputSchemaCompiled")
		delete(additionalProperties, "publishedVersion")
		delete(additionalProperties, "savedDate")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "timeouts")
		delete(additionalProperties, "trigger")
		delete(additionalProperties, "updatedDate")
		delete(additionalProperties, "variables")
		delete(additionalProperties, "versionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o Flow) DesignerCuesFields() []string {
	return []string{
		"FlowColor",
	}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Flow) EnvironmentMetadataFields() []string {
	return []string{
		"CompanyID",
		"CreatedDate",
		"CustomerID",
		"FlowID",
	}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Flow) FlowConfigFields() []string {
	return []string{
		"InputSchema",
		"Settings",
		"InputSchemaCompiled",
		"IsInputSchemaSaved",
		"IsOutputSchemaSaved",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o Flow) FlowMetadataFields() []string {
	return []string{
		"AuthTokenExpireIds",
		"Connections",
		"ConnectorIds",
		"Description",
		"EnabledGraphData",
		"FunctionConnectionID",
		"Name",
		"Orx",
		"Timeouts",
	}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Flow) VersionMetadataFields() []string {
	return []string{
		"CurrentVersion",
		"DeployedDate",
		"FlowStatus",
		"PublishedVersion",
		"SavedDate",
		"UpdatedDate",
		"VersionID",
	}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o Flow) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
