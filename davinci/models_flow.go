package davinci

import (
	"encoding/json"
)

type _Flow Flow
type Flow struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	AuthTokenExpireIds   []interface{}          `json:"authTokenExpireIds,omitempty"`
	CompanyID            string                 `json:"companyId"`
	Connections          []interface{}          `json:"connections,omitempty"`
	ConnectorIds         []string               `json:"connectorIds" dv:"config"`
	CreatedDate          int                    `json:"createdDate"`
	CurrentVersion       *int                   `json:"currentVersion,omitempty"`
	CustomerID           string                 `json:"customerId"`
	DeployedDate         *int                   `json:"deployedDate,omitempty"`
	Description          *string                `json:"description,omitempty" dv:"config"`
	EnabledGraphData     interface{}            `json:"enabledGraphData,omitempty" dv:"config"`
	FlowColor            *string                `json:"flowColor,omitempty" dv:"config"`
	FlowID               string                 `json:"flowId"`
	FlowStatus           string                 `json:"flowStatus" dv:"config"`
	FunctionConnectionID interface{}            `json:"functionConnectionId,omitempty" dv:"config"`
	GraphData            GraphData              `json:"graphData" dv:"config"`
	InputSchema          []interface{}          `json:"inputSchema,omitempty" dv:"config"`
	InputSchemaCompiled  interface{}            `json:"inputSchemaCompiled,omitempty" dv:"config"`
	IsInputSchemaSaved   *bool                  `json:"isInputSchemaSaved,omitempty" dv:"config"`
	IsOutputSchemaSaved  *bool                  `json:"isOutputSchemaSaved,omitempty" dv:"config"`
	Name                 string                 `json:"name" dv:"config"`
	Orx                  *string                `json:"orx,omitempty" dv:"config"`
	OutputSchema         *OutputSchema          `json:"outputSchema,omitempty" dv:"config"`
	OutputSchemaCompiled *OutputSchema          `json:"outputSchemaCompiled,omitempty" dv:"config"` //compiled is used in exported flow json, must be converted to JUST output when updating flow.
	PublishedVersion     *int                   `json:"publishedVersion,omitempty"`
	SavedDate            int                    `json:"savedDate"`
	Settings             interface{}            `json:"settings,omitempty" dv:"config"`
	Timeouts             interface{}            `json:"timeouts,omitempty" dv:"config"`
	Trigger              *Trigger               `json:"trigger,omitempty" dv:"config"`
	UpdatedDate          *int                   `json:"updatedDate,omitempty"`
	Variables            []FlowVariable         `json:"variables,omitempty" dv:"config"`
	VersionID            int64                  `json:"versionId"`
}

var marshalWithDVTags = false

func (o Flow) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Flow) MarshalJSONWithDVTag() ([]byte, error) {
	marshalWithDVTags = true
	result, err := json.Marshal(o)
	marshalWithDVTags = false
	return result, err
}

func (o Flow) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	flowExportMetadataJsonData, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(flowExportMetadataJsonData, &result); err != nil {
		return nil, err
	}

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
