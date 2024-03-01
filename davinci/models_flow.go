package davinci

import "encoding/json"

type _Flow Flow
type Flow struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	FlowConfiguration
	FlowEnvironmentMetadata
	FlowMetadata
	FlowVersionMetadata
}

type FlowConfiguration struct {
	FlowUpdateConfiguration
	FlowColor            *string       `davinci:"flowColor,designercue,omitempty"`
	InputSchemaCompiled  interface{}   `davinci:"inputSchemaCompiled,config,omitempty"`
	IsInputSchemaSaved   *bool         `davinci:"isInputSchemaSaved,config,omitempty"`
	IsOutputSchemaSaved  *bool         `davinci:"isOutputSchemaSaved,config,omitempty"`
	OutputSchemaCompiled *OutputSchema `davinci:"outputSchemaCompiled,*,omitempty"` //compiled is used in exported flow json, must be converted to JUST output when updating flow.
}

type FlowUpdateConfiguration struct {
	GraphData    *GraphData    `davinci:"graphData,*,omitempty"`
	InputSchema  []interface{} `davinci:"inputSchema,config,omitempty"`
	OutputSchema *OutputSchema `davinci:"outputSchema,*,omitempty"`
	Settings     interface{}   `davinci:"settings,config,omitempty"`
	Trigger      *Trigger      `davinci:"trigger,*,omitempty"`
}

type FlowEnvironmentMetadata struct {
	CompanyID   string    `davinci:"companyId,environmentmetadata"`
	CreatedDate EpochTime `davinci:"createdDate,environmentmetadata"`
	CustomerID  string    `davinci:"customerId,environmentmetadata"`
	FlowID      string    `davinci:"flowId,environmentmetadata"`
}

type FlowMetadata struct {
	AuthTokenExpireIds   []interface{}  `davinci:"authTokenExpireIds,flowmetadata,omitempty"`
	Connections          []interface{}  `davinci:"connections,flowmetadata,omitempty"`
	ConnectorIds         []string       `davinci:"connectorIds,flowmetadata"`
	Description          *string        `davinci:"description,flowmetadata,omitempty"`
	EnabledGraphData     interface{}    `davinci:"enabledGraphData,flowmetadata,omitempty"`
	FunctionConnectionID interface{}    `davinci:"functionConnectionId,flowmetadata,omitempty"`
	Name                 string         `davinci:"name,flowmetadata"`
	Orx                  *string        `davinci:"orx,flowmetadata,omitempty"`
	Timeouts             interface{}    `davinci:"timeouts,flowmetadata,omitempty"`
	Variables            []FlowVariable `davinci:"variables,*,omitempty"`
}

type FlowVersionMetadata struct {
	CurrentVersion   *int32     `davinci:"currentVersion,versionmetadata,omitempty"`
	DeployedDate     *EpochTime `davinci:"deployedDate,versionmetadata,omitempty"`
	FlowStatus       string     `davinci:"flowStatus,versionmetadata"`
	PublishedVersion *int32     `davinci:"publishedVersion,versionmetadata,omitempty"`
	SavedDate        EpochTime  `davinci:"savedDate,versionmetadata"`
	UpdatedDate      *EpochTime `davinci:"updatedDate,versionmetadata,omitempty"`
	VersionID        int32      `davinci:"versionId,versionmetadata"`
}

func (o *Flow) UnmarshalDavinci(bytes []byte, opts ExportCmpOpts) (err error) {
	varFlow := _Flow{}

	if err = Unmarshal(bytes, &varFlow, opts); err != nil {
		return err
	}

	*o = Flow(varFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err != nil {
		return err
	}

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

	return nil
}
