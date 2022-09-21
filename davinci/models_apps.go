package davinci

type Apps struct {
	Apps []App `json:"apps,omitempty"`
}
type APIKeys struct {
	Prod string `json:"prod,omitempty"`
	Test string `json:"test,omitempty"`
}
type Metadata struct {
	RpName string `json:"rpName,omitempty"`
}
type UserPools struct {
	ConnectionID string `json:"connectionId,omitempty"`
	ConnectorID  string `json:"connectorId,omitempty"`
}
type Values struct {
	Enabled       bool          `json:"enabled,omitempty"`
	ClientSecret  string        `json:"clientSecret,omitempty"`
	RedirectUris  []string      `json:"redirectUris,omitempty"`
	LogoutUris    []interface{} `json:"logoutUris,omitempty"`
	AllowedScopes []string      `json:"allowedScopes,omitempty"`
	AllowedGrants []string      `json:"allowedGrants,omitempty"`
}
type CreatedApp struct {
	App App `json:"app"`
}
type App struct {
	CompanyID     string        `json:"companyId,omitempty"`
	Name          string        `json:"name"`
	CustomerID    string        `json:"customerId,omitempty"`
	APIKeys       *APIKeys      `json:"apiKeys,omitempty"`
	Metadata      *Metadata     `json:"metadata,omitempty"`
	UserPools     []UserPools   `json:"userPools,omitempty"`
	Oauth         *Oauth        `json:"oauth,omitempty"`
	Saml          *Saml         `json:"saml,omitempty"`
	Flows         []interface{} `json:"flows,omitempty"`
	Policies      []Policies    `json:"policies,omitempty"`
	CreatedDate   int64         `json:"createdDate,omitempty"`
	APIKeyEnabled bool          `json:"apiKeyEnabled,omitempty"`
	AppID         string        `json:"appId,omitempty"`
}

type Oauth struct {
	Enabled bool         `json:"enabled,omitempty"`
	Values  *OauthValues `json:"values,omitempty"`
}

type OauthValues struct {
	Enabled                    bool          `json:"enabled,omitempty"`
	ClientSecret               string        `json:"clientSecret,omitempty"`
	RedirectUris               []string      `json:"redirectUris,omitempty"`
	LogoutUris                 []interface{} `json:"logoutUris,omitempty"`
	AllowedScopes              []string      `json:"allowedScopes,omitempty"`
	AllowedGrants              []string      `json:"allowedGrants,omitempty"`
	EnforceSignedRequestOpenid bool          `json:"enforceSignedRequestOpenid,omitempty"`
	SpjwksUrl                  string        `json:"spjwksUrl,omitempty"`
	SpJwksOpenid               string        `json:"spJwksOpenid,omitempty"`
}

type Saml struct {
	Values *SamlValues `json:"values,omitempty"`
}

type SamlValues struct {
	Enabled              bool   `json:"enabled,omitempty"`
	RedirectURI          string `json:"redirectUri,omitempty"`
	Audience             string `json:"audience,omitempty"`
	EnforceSignedRequest bool   `json:"enforceSignedRequest,omitempty"`
	SpCert               string `json:"spCert,omitempty"`
}

type PolicyFlows struct {
	FlowID       string   `json:"flowId,omitempty"`
	VersionID    int      `json:"versionId,omitempty"`
	Weight       int      `json:"weight,omitempty"`
	SuccessNodes []string `json:"successNodes,omitempty"`
}
type Policies struct {
	PolicyFlows []PolicyFlows `json:"flows,omitempty"`
	Name        string        `json:"name,omitempty"`
	Status      string        `json:"status,omitempty"`
	PolicyID    string        `json:"policyId,omitempty"`
	CreatedDate int64         `json:"createdDate,omitempty"`
}

type AppUpdate struct {
	Name          string        `json:"name"`
	Oauth         *OauthUpdate  `json:"oauth,omitempty"`
	Saml          *SamlUpdate   `json:"saml,omitempty"`
	Flows         []interface{} `json:"flows,omitempty"`
	Policies      []Policies    `json:"policies,omitempty"`
	CreatedDate   int64         `json:"createdDate,omitempty"`
	APIKeyEnabled bool          `json:"apiKeyEnabled,omitempty"`
	AppID         string        `json:"appId,omitempty"`
}

type OauthUpdate struct {
	Enabled bool               `json:"enabled,omitempty"`
	Values  *OauthValuesUpdate `json:"values,omitempty"`
}

type OauthValuesUpdate struct {
	Enabled                    bool          `json:"enabled,omitempty"`
	RedirectUris               []string      `json:"redirectUris,omitempty"`
	LogoutUris                 []interface{} `json:"logoutUris,omitempty"`
	AllowedScopes              []string      `json:"allowedScopes,omitempty"`
	AllowedGrants              []string      `json:"allowedGrants,omitempty"`
	EnforceSignedRequestOpenid bool          `json:"enforceSignedRequestOpenid,omitempty"`
	SpjwksUrl                  string        `json:"spjwksUrl,omitempty"`
	SpJwksOpenid               string        `json:"spJwksOpenid,omitempty"`
}

type SamlUpdate struct {
	Values *SamlValuesUpdate `json:"values,omitempty"`
}

type SamlValuesUpdate struct {
	Enabled              bool   `json:"enabled,omitempty"`
	RedirectURI          string `json:"redirectUri,omitempty"`
	Audience             string `json:"audience,omitempty"`
	EnforceSignedRequest bool   `json:"enforceSignedRequest,omitempty"`
	SpCert               string `json:"spCert,omitempty"`
}
