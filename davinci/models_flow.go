package davinci

import (
	"encoding/json"
)

type _Flow Flow
type Flow struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	FlowConfiguration
	FlowEnvironmentMetadata
	FlowMetadata
	FlowVersionMetadata
}

type FlowConfiguration struct {
	ConnectorIds         []string      `json:"connectorIds"`
	FlowColor            *string       `json:"flowColor,omitempty"`
	GraphData            GraphData     `json:"graphData"`
	InputSchema          []interface{} `json:"inputSchema,omitempty"`
	InputSchemaCompiled  interface{}   `json:"inputSchemaCompiled,omitempty"`
	IsInputSchemaSaved   *bool         `json:"isInputSchemaSaved,omitempty"`
	IsOutputSchemaSaved  *bool         `json:"isOutputSchemaSaved,omitempty"`
	OutputSchema         *OutputSchema `json:"outputSchema,omitempty"`
	OutputSchemaCompiled *OutputSchema `json:"outputSchemaCompiled,omitempty"` //compiled is used in exported flow json, must be converted to JUST output when updating flow.
	Settings             interface{}   `json:"settings,omitempty"`
	Trigger              *Trigger      `json:"trigger,omitempty"`
}

type FlowEnvironmentMetadata struct {
	CompanyID   string    `json:"companyId"`
	CreatedDate EpochTime `json:"createdDate"`
	CustomerID  string    `json:"customerId"`
	FlowID      string    `json:"flowId"`
}

type FlowMetadata struct {
	AuthTokenExpireIds   []interface{}  `json:"authTokenExpireIds,omitempty"`
	Connections          []interface{}  `json:"connections,omitempty"`
	Description          *string        `json:"description,omitempty"`
	EnabledGraphData     interface{}    `json:"enabledGraphData,omitempty"`
	FunctionConnectionID interface{}    `json:"functionConnectionId,omitempty"`
	Name                 string         `json:"name"`
	Orx                  *string        `json:"orx,omitempty"`
	Timeouts             interface{}    `json:"timeouts,omitempty"`
	Variables            []FlowVariable `json:"variables,omitempty"`
}

type FlowVersionMetadata struct {
	CurrentVersion   *int32     `json:"currentVersion,omitempty"`
	DeployedDate     *EpochTime `json:"deployedDate,omitempty"`
	FlowStatus       string     `json:"flowStatus"`
	PublishedVersion *int32     `json:"publishedVersion,omitempty"`
	SavedDate        EpochTime  `json:"savedDate"`
	UpdatedDate      *EpochTime `json:"updatedDate,omitempty"`
	VersionID        int32      `json:"versionId"`
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
