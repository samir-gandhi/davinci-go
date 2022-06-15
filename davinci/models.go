package davinci

type AuthStruct struct {
	Username string `json:"email"`
	Password string `json:"password"`
}
type SkSdkToken struct {
	CapabilityName string `json:"capabilityName"`
	AccessToken    string `json:"access_token"`
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	Success        bool   `json:"success"`
}
type LoginResponse struct {
	AccessToken     string     `json:"access_token"`
	TokenType       string     `json:"token_type"`
	MfaRequired     bool       `json:"mfaRequired"`
	Status          string     `json:"status"`
	CustomerID      string     `json:"customerId"`
	AppConfig       bool       `json:"appConfig"`
	SkSdkToken      SkSdkToken `json:"skSdkToken"`
	FlowPolicyID    string     `json:"flowPolicyId"`
	CompanyID       string     `json:"companyId"`
	SelectedCompany string     `json:"selectedCompany"`
}

type Callback struct {
	InteractionID    string `json:"interactionId"`
	CompanyID        string `json:"companyId"`
	ConnectionID     string `json:"connectionId"`
	ConnectorID      string `json:"connectorId"`
	ID               string `json:"id"`
	CapabilityName   string `json:"capabilityName"`
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	IDToken          string `json:"id_token"`
	Success          bool   `json:"success"`
	InteractionToken string `json:"interactionToken"`
}

type AuthResponse struct {
	AccessToken     string `json:"access_token"`
	TokenType       string `json:"token_type"`
	Status          string `json:"status"`
	CustomerID      string `json:"customerId"`
	TryFlowToken    string `json:"tryFlowToken"`
	SelectedCompany string `json:"selectedCompany"`
}

type Companies struct {
	CompanyID string   `json:"companyId"`
	Roles     []string `json:"roles"`
	Name      string   `json:"name"`
	SvgIcon   string   `json:"svgIcon"`
}

type Environments struct {
	CustomerID  string    `json:"customerId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	CompanyID   string    `json:"companyId"`
	Companies   Companies `json:"companies"`
	ClientID    string    `json:"clientId"`
	CreatedDate int64     `json:"createdDate"`
}
