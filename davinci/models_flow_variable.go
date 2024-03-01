package davinci

type FlowVariable struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	CompanyID            *string                `davinci:"companyId,environmentmetadata,omitempty"`
	Context              *string                `davinci:"context,config,omitempty"`
	CreatedDate          *EpochTime             `davinci:"createdDate,versionmetadata,omitempty"`
	CustomerID           *string                `davinci:"customerId,environmentmetadata,omitempty"`
	Fields               *FlowVariableFields    `davinci:"fields,*,omitempty"`
	FlowID               *string                `davinci:"flowId,environmentmetadata,omitempty"`
	Key                  *float64               `davinci:"key,flowmetadata,omitempty"`
	Label                *string                `davinci:"label,config,omitempty"`
	Name                 string                 `davinci:"name,config"`
	Type                 string                 `davinci:"type,config"`
	UpdatedDate          *EpochTime             `davinci:"updatedDate,versionmetadata,omitempty"`
	Value                *string                `davinci:"value,config,omitempty"`
	Visibility           *string                `davinci:"visibility,flowmetadata,omitempty"`
}
