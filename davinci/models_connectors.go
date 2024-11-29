package davinci

// CustomAuth is a field of the Properties struct
// but is kept separate because the structure is known.
// This is used to unmarshal the properties["customAuth"] field of a connection response.
type CustomAuth struct {
	Properties *CustomAuthProperties `json:"properties,omitempty" davinci:"properties,config,omitempty"`
}

type CustomAuthProperties struct {
	ProviderName                  *ProviderName                  `json:"providerName,omitempty" davinci:"providerName,config,omitempty"`
	AuthTypeDropdown              *AuthTypeDropdown              `json:"authTypeDropdown,omitempty" davinci:"authTypeDropdown,config,omitempty"`
	SkRedirectURI                 *SkRedirectURI                 `json:"skRedirectUri,omitempty" davinci:"skRedirectUri,config,omitempty"`
	IssuerURL                     *IssuerURL                     `json:"issuerUrl,omitempty" davinci:"issuerUrl,config,omitempty"`
	AuthorizationEndpoint         *AuthorizationEndpoint         `json:"authorizationEndpoint,omitempty" davinci:"authorizationEndpoint,config,omitempty"`
	TokenEndpoint                 *TokenEndpoint                 `json:"tokenEndpoint,omitempty" davinci:"tokenEndpoint,config,omitempty"`
	BearerToken                   *BearerToken                   `json:"bearerToken,omitempty" davinci:"bearerToken,config,omitempty"`
	UserInfoEndpoint              *UserInfoEndpoint              `json:"userInfoEndpoint,omitempty" davinci:"userInfoEndpoint,config,omitempty"`
	ClientID                      *ClientID                      `json:"clientId,omitempty" davinci:"clientId,config,omitempty"`
	ClientSecret                  *ClientSecret                  `json:"clientSecret,omitempty" davinci:"clientSecret,config,omitempty"`
	Scope                         *Scope                         `json:"scope,omitempty" davinci:"scope,config,omitempty"`
	Code                          *Code                          `json:"code,omitempty" davinci:"code,config,omitempty"`
	UserConnectorAttributeMapping *UserConnectorAttributeMapping `json:"userConnectorAttributeMapping,omitempty" davinci:"userConnectorAttributeMapping,config,omitempty"`
	CustomAttributes              *CustomAttributes              `json:"customAttributes,omitempty" davinci:"customAttributes,config,omitempty"`
	ReturnToURL                   *ReturnToURL                   `json:"returnToUrl,omitempty" davinci:"returnToUrl,config,omitempty"`
}

type ProviderName struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Required             *bool   `json:"required,omitempty" davinci:"required,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
}
type Options struct {
	Name  *string `json:"name,omitempty" davinci:"name,config,omitempty"`
	Value *string `json:"value,omitempty" davinci:"value,config,omitempty"`
}
type AuthTypeDropdown struct {
	DisplayName          *string   `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string   `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Required             *bool     `json:"required,omitempty" davinci:"required,config,omitempty"`
	Options              []Options `json:"options,omitempty" davinci:"options,config,omitempty"`
	Enum                 []string  `json:"enum,omitempty" davinci:"enum,config,omitempty"`
	Value                *string   `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string   `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type SkRedirectURI struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Disabled             *bool   `json:"disabled,omitempty" davinci:"disabled,config,omitempty"`
	InitializeValue      *string `json:"initializeValue,omitempty" davinci:"initializeValue,config,omitempty"`
	CopyToClip           *bool   `json:"copyToClip,omitempty" davinci:"copyToClip,config,omitempty"`
}
type IssuerURL struct {
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	Info                 *string `json:"info,omitempty" davinci:"info,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type AuthorizationEndpoint struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Required             *bool   `json:"required,omitempty" davinci:"required,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type TokenEndpoint struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Required             *bool   `json:"required,omitempty" davinci:"required,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type BearerToken struct {
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Type                 *string `json:"type,omitempty" davinci:"type,config,omitempty"`
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	Info                 *string `json:"info,omitempty" davinci:"info,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type UserInfoEndpoint struct {
	DisplayName          *string  `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string  `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Required             *bool    `json:"required,omitempty" davinci:"required,config,omitempty"`
	Value                []string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          []string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type ClientID struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Required             *bool   `json:"required,omitempty" davinci:"required,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type ClientSecret struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	HashedVisibility     *bool   `json:"hashedVisibility,omitempty" davinci:"hashedVisibility,config,omitempty"`
	Required             *bool   `json:"required,omitempty" davinci:"required,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type Scope struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Required             *bool   `json:"required,omitempty" davinci:"required,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type Code struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	Info                 *string `json:"info,omitempty" davinci:"info,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Language             *string `json:"language,omitempty" davinci:"language,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type Username struct {
	Value1 string `json:"value1,omitempty" davinci:"value1,config,omitempty"`
}
type FirstName struct {
	Value1 string `json:"value1,omitempty" davinci:"value1,config,omitempty"`
}
type LastName struct {
	Value1 string `json:"value1,omitempty" davinci:"value1,config,omitempty"`
}
type Name struct {
	Value1 string `json:"value1,omitempty" davinci:"value1,config,omitempty"`
}
type Email struct {
	Value1 string `json:"value1,omitempty" davinci:"value1,config,omitempty"`
}
type Mapping struct {
	Username  *Username  `json:"username,omitempty" davinci:"username,config,omitempty"`
	FirstName *FirstName `json:"firstName,omitempty" davinci:"firstName,config,omitempty"`
	LastName  *LastName  `json:"lastName,omitempty" davinci:"lastName,config,omitempty"`
	Name      *Name      `json:"name,omitempty" davinci:"name,config,omitempty"`
	Email     *Email     `json:"email,omitempty" davinci:"email,config,omitempty"`
}
type UserConnectorAttributeMappingValue struct {
	UserPoolConnectionID *string  `json:"userPoolConnectionId,omitempty" davinci:"userPoolConnectionId,config,omitempty"`
	Mapping              *Mapping `json:"mapping,omitempty" davinci:"mapping,config,omitempty"`
}
type UserConnectorAttributeMappingPlaceholder struct {
	UserPoolConnectionID *string  `json:"userPoolConnectionId,omitempty" davinci:"userPoolConnectionId,config,omitempty"`
	Mapping              *Mapping `json:"mapping,omitempty" davinci:"mapping,config,omitempty"`
}
type UserConnectorAttributeMapping struct {
	Type                 *string                                   `json:"type,omitempty" davinci:"type,config,omitempty"`
	DisplayName          interface{}                               `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string                                   `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	NewMappingAllowed    *bool                                     `json:"newMappingAllowed,omitempty" davinci:"newMappingAllowed,config,omitempty"`
	Title1               interface{}                               `json:"title1,omitempty" davinci:"title1,config,omitempty"`
	Title2               interface{}                               `json:"title2,omitempty" davinci:"title2,config,omitempty"`
	Sections             []string                                  `json:"sections,omitempty" davinci:"sections,config,omitempty"`
	Value                *UserConnectorAttributeMappingValue       `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *UserConnectorAttributeMappingPlaceholder `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type CustomAttributesValue struct {
	Name          *string     `json:"name,omitempty" davinci:"name,config,omitempty"`
	Description   *string     `json:"description,omitempty" davinci:"description,config,omitempty"`
	Type          *string     `json:"type,omitempty" davinci:"type,config,omitempty"`
	Value         interface{} `json:"value,omitempty" davinci:"value,config,omitempty"`
	MinLength     *string     `json:"minLength,omitempty" davinci:"minLength,config,omitempty"`
	MaxLength     *string     `json:"maxLength,omitempty" davinci:"maxLength,config,omitempty"`
	Required      *bool       `json:"required,omitempty" davinci:"required,config,omitempty"`
	AttributeType *string     `json:"attributeType,omitempty" davinci:"attributeType,config,omitempty"`
}
type Placeholder struct {
	Name          *string     `json:"name,omitempty" davinci:"name,config,omitempty"`
	Description   *string     `json:"description,omitempty" davinci:"description,config,omitempty"`
	Type          *string     `json:"type,omitempty" davinci:"type,config,omitempty"`
	Value         interface{} `json:"value,omitempty" davinci:"value,config,omitempty"`
	MinLength     *string     `json:"minLength,omitempty" davinci:"minLength,config,omitempty"`
	MaxLength     *string     `json:"maxLength,omitempty" davinci:"maxLength,config,omitempty"`
	Required      *bool       `json:"required,omitempty" davinci:"required,config,omitempty"`
	AttributeType *string     `json:"attributeType,omitempty" davinci:"attributeType,config,omitempty"`
}
type CustomAttributes struct {
	Type                 *string                 `json:"type,omitempty" davinci:"type,config,omitempty"`
	DisplayName          *string                 `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string                 `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Info                 *string                 `json:"info,omitempty" davinci:"info,config,omitempty"`
	Sections             []string                `json:"sections,omitempty" davinci:"sections,config,omitempty"`
	Value                []CustomAttributesValue `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          []Placeholder           `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}
type ReturnToURL struct {
	DisplayName          *string `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	PreferredControlType *string `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Info                 *string `json:"info,omitempty" davinci:"info,config,omitempty"`
	Value                *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Placeholder          *string `json:"placeholder,omitempty" davinci:"placeholder,config,omitempty"`
}

// Connector Config for read all CONNECTORS.
type ConnectorLoose struct {
	AccountConfigView   *ConnectorAccountConfigView       `json:"accountConfigView,omitempty" davinci:"accountConfigView,config,omitempty"`
	Capabilities        *ConnectorCapabilities            `json:"capabilities,omitempty" davinci:"capabilities,config,omitempty"`
	CompanyID           *string                           `json:"companyId,omitempty" davinci:"companyId,config,omitempty"`
	ConnectorCategories []ConnectorCategories             `json:"connectorCategories,omitempty" davinci:"connectorCategories,config,omitempty"`
	ConnectorType       *string                           `json:"connectorType,omitempty" davinci:"connectorType,config,omitempty"`
	CreatedDate         *EpochTime                        `json:"createdDate,omitempty" davinci:"createdDate,config,omitempty"`
	CustomerID          *string                           `json:"customerId,omitempty" davinci:"customerId,config,omitempty"`
	Description         *string                           `json:"description,omitempty" davinci:"description,config,omitempty"`
	FlowSections        []ConnectorFlowSections           `json:"flowSections,omitempty" davinci:"flowSections,config,omitempty"`
	ManifestVersion     *string                           `json:"manifestVersion,omitempty" davinci:"manifestVersion,config,omitempty"`
	Metadata            *ConnectorMetadata                `json:"metadata,omitempty" davinci:"metadata,config,omitempty"`
	Name                *string                           `json:"name,omitempty" davinci:"name,config,omitempty"`
	Properties          map[string]ConnectorPropertyLoose `json:"properties,omitempty" davinci:"properties,config,omitempty"`
	Sections            []ConnectorSections               `json:"sections,omitempty" davinci:"sections,config,omitempty"`
	Status              *string                           `json:"status,omitempty" davinci:"status,config,omitempty"`
	UpdatedDate         *EpochTime                        `json:"updatedDate,omitempty" davinci:"updatedDate,config,omitempty"`
	ConnectorID         *string                           `json:"connectorId,omitempty" davinci:"connectorId,config,omitempty"`
}
type Connector struct {
	AccountConfigView   *ConnectorAccountConfigView  `json:"accountConfigView,omitempty" davinci:"accountConfigView,config,omitempty"`
	Capabilities        *ConnectorCapabilities       `json:"capabilities,omitempty" davinci:"capabilities,config,omitempty"`
	CompanyID           *string                      `json:"companyId,omitempty" davinci:"companyId,environmentmetadata,omitempty"`
	ConnectorCategories []ConnectorCategories        `json:"connectorCategories,omitempty" davinci:"connectorCategories,config,omitempty"`
	ConnectorType       *string                      `json:"connectorType,omitempty" davinci:"connectorType,config,omitempty"`
	CreatedDate         *EpochTime                   `json:"createdDate,omitempty" davinci:"createdDate,config,omitempty"`
	CustomerID          *string                      `json:"customerId,omitempty" davinci:"customerId,environmentmetadata,omitempty"`
	Description         *string                      `json:"description,omitempty" davinci:"description,config,omitempty"`
	FlowSections        []ConnectorFlowSections      `json:"flowSections,omitempty" davinci:"flowSections,config,omitempty"`
	ManifestVersion     *string                      `json:"manifestVersion,omitempty" davinci:"manifestVersion,config,omitempty"`
	Metadata            *ConnectorMetadata           `json:"metadata,omitempty" davinci:"metadata,config,omitempty"`
	Name                *string                      `json:"name,omitempty" davinci:"name,config,omitempty"`
	Properties          map[string]ConnectorProperty `json:"properties,omitempty" davinci:"properties,config,omitempty"`
	Sections            []ConnectorSections          `json:"sections,omitempty" davinci:"sections,config,omitempty"`
	Status              *string                      `json:"status,omitempty" davinci:"status,config,omitempty"`
	UpdatedDate         *EpochTime                   `json:"updatedDate,omitempty" davinci:"updatedDate,config,omitempty"`
	ConnectorID         *string                      `json:"connectorId,omitempty" davinci:"connectorId,config,omitempty"`
}

type ConnectorAccountConfigView struct {
	Items []AccountConfigViewItems `json:"items,omitempty" davinci:"items,config,omitempty"`
}
type AccountConfigViewItems struct {
	PropertyName *string `json:"propertyName,omitempty" davinci:"propertyName,config,omitempty"`
}
type FlowConfigViewItems struct {
	PropertyName *string `json:"propertyName,omitempty" davinci:"propertyName,config,omitempty"`
}
type ConnectorFlowConfigView struct {
	Items []FlowConfigViewItems `json:"items,omitempty" davinci:"items,config,omitempty"`
}
type UpdateAnnotationProperties struct {
	Type           *string                  `json:"type,omitempty" davinci:"type,config,omitempty"`
	Title          *string                  `json:"title,omitempty" davinci:"title,config,omitempty"`
	SubTitle       *string                  `json:"subTitle,omitempty" davinci:"subTitle,config,omitempty"`
	UserViews      []interface{}            `json:"userViews,omitempty" davinci:"userViews,config,omitempty"`
	FlowConfigView *ConnectorFlowConfigView `json:"flowConfigView,omitempty" davinci:"flowConfigView,config,omitempty"`
}
type ConnectorCapabilities struct {
	UpdateAnnotationProperties *UpdateAnnotationProperties `json:"updateAnnotationProperties,omitempty" davinci:"updateAnnotationProperties,config,omitempty"`
}
type ConnectorCategories struct {
	Name  *string `json:"name,omitempty" davinci:"name,config,omitempty"`
	Value *string `json:"value,omitempty" davinci:"value,config,omitempty"`
}
type ConnectorFlowSections struct {
	Name  *string `json:"name,omitempty" davinci:"name,config,omitempty"`
	Value *string `json:"value,omitempty" davinci:"value,config,omitempty"`
}
type ConnectorMetadataColors struct {
	Canvas     *string `json:"canvas,omitempty" davinci:"canvas,config,omitempty"`
	CanvasText *string `json:"canvasText,omitempty" davinci:"canvasText,config,omitempty"`
	Dark       *string `json:"dark,omitempty" davinci:"dark,config,omitempty"`
}
type LogosCanvas struct {
	ImageFileName *string `json:"imageFileName,omitempty" davinci:"imageFileName,config,omitempty"`
}
type ConnectorMetadataLogos struct {
	Canvas *LogosCanvas `json:"canvas,omitempty" davinci:"canvas,config,omitempty"`
}
type ConnectorMetadata struct {
	SkType *string                  `json:"skType,omitempty" davinci:"skType,config,omitempty"`
	Colors *ConnectorMetadataColors `json:"colors,omitempty" davinci:"colors,config,omitempty"`
	Logos  *ConnectorMetadataLogos  `json:"logos,omitempty" davinci:"logos,config,omitempty"`
}
type ConnectorPropertyLoose struct {
	Type                 *string     `json:"type,omitempty" davinci:"type,config,omitempty"`
	DisplayName          *string     `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	CreatedDate          *EpochTime  `json:"createdDate,omitempty" davinci:"createdDate,config,omitempty"`
	CustomerID           *string     `json:"customerId,omitempty" davinci:"customerId,config,omitempty"`
	CompanyID            *string     `json:"companyId,omitempty" davinci:"companyId,config,omitempty"`
	PreferredControlType *string     `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Info                 interface{} `json:"info,omitempty" davinci:"info,config,omitempty"`
}

type ConnectorProperty struct {
	Type                 *string    `json:"type,omitempty" davinci:"type,config,omitempty"`
	DisplayName          *string    `json:"displayName,omitempty" davinci:"displayName,config,omitempty"`
	CreatedDate          *EpochTime `json:"createdDate,omitempty" davinci:"createdDate,config,omitempty"`
	CustomerID           *string    `json:"customerId,omitempty" davinci:"customerId,environmentmetadata,omitempty"`
	CompanyID            *string    `json:"companyId,omitempty" davinci:"companyId,environmentmetadata,omitempty"`
	PreferredControlType *string    `json:"preferredControlType,omitempty" davinci:"preferredControlType,config,omitempty"`
	Info                 *string    `json:"info,omitempty" davinci:"info,config,omitempty"`
}

type ConnectorSections struct {
	Name    *string `json:"name,omitempty" davinci:"name,config,omitempty"`
	Value   *string `json:"value,omitempty" davinci:"value,config,omitempty"`
	Default *bool   `json:"default,omitempty" davinci:"default,config,omitempty"`
}
