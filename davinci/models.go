package davinci

import "time"

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
	CustomerID  string      `json:"customerId"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Email       string      `json:"email"`
	PhoneNumber string      `json:"phoneNumber"`
	CompanyID   string      `json:"companyId"`
	Companies   []Companies `json:"companies"`
	ClientID    string      `json:"clientId"`
	CreatedDate int64       `json:"createdDate"`
}

type Environment struct {
		CreatedByCustomerID string `json:"createdByCustomerId"`
		CreatedByCompanyID  string `json:"createdByCompanyId"`
		Name                string `json:"name"`
		CompanyType         string `json:"companyType"`
		EntitlementTemplate string `json:"entitlementTemplate"`
		EntitlementProps    struct {
		} `json:"entitlementProps"`
		SecurityType string `json:"securityType"`
		JwtKeys      struct {
			Jwks struct {
				Keys []struct {
					Kty string `json:"kty"`
					Kid string `json:"kid"`
					N   string `json:"n"`
					E   string `json:"e"`
					Alg string `json:"alg"`
					Use string `json:"use"`
				} `json:"keys"`
			} `json:"jwks"`
		} `json:"jwtKeys"`
		SamlKeys struct {
			PublicKey string `json:"publicKey"`
			Cert      string `json:"cert"`
		} `json:"samlKeys"`
		Properties struct {
			PngIcon struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"pngIcon"`
			SvgIcon struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"svgIcon"`
			SvgViewBox struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"svgViewBox"`
			IconOpacity struct {
				DisplayName          string  `json:"displayName"`
				DataType             string  `json:"dataType"`
				PreferredControlType string  `json:"preferredControlType"`
				Value                float64 `json:"value"`
			} `json:"iconOpacity"`
			BackgroundColor struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"backgroundColor"`
			TextColor struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"textColor"`
			IconColor struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"iconColor"`
			ArcColor struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"arcColor"`
			ArcProgressColor struct {
				DisplayName          string `json:"displayName"`
				DataType             string `json:"dataType"`
				PreferredControlType string `json:"preferredControlType"`
				Value                string `json:"value"`
			} `json:"arcProgressColor"`
		} `json:"properties"`
		CreatedDate int64 `json:"createdDate"`
		Entitlement struct {
			Company struct {
				CreateAdditional bool `json:"createAdditional"`
			} `json:"company"`
			Connectors struct {
				Whitelist []string    `json:"whitelist"`
				//TODO
				// Blacklist []string `json:"blacklist"`
				Blacklist interface{} `json:"blacklist"`
			} `json:"connectors"`
			Connections struct {
				Total                              int `json:"total"`
				MaxNumberOfConnectionsPerConnector struct {
					ConnectorID string `json:"connectorId"`
					Total       int    `json:"total"`
				} `json:"maxNumberOfConnectionsPerConnector"`
			} `json:"connections"`
			Flows struct {
				Enabled bool  `json:"enabled"`
				Total   int   `json:"total"`
				Expires int64 `json:"expires"`
			} `json:"flows"`
			Attributes struct {
				Enabled bool `json:"enabled"`
				Total   int  `json:"total"`
			} `json:"attributes"`
			Apps struct {
				Total int `json:"total"`
			} `json:"apps"`
			Users struct {
				Total                   int `json:"total"`
				TotalCredentialsPerUser int `json:"totalCredentialsPerUser"`
			} `json:"users"`
			Expires int64 `json:"expires"`
		} `json:"entitlement"`
		CompanyID string `json:"companyId"`
}

type EnvironmentStats struct {
	TableStats []struct {
		Flows       int `json:"Flows"`
		Connections int `json:"Connections"`
		Apps        int `json:"Apps"`
		Customers   int `json:"Customers"`
		Constructs  int `json:"Constructs"`
		Users       int `json:"Users"`
		Events      int `json:"Events"`
		ID          struct {
			CompanyID string `json:"companyId"`
			Ts        int64  `json:"ts"`
		} `json:"_id"`
	} `json:"tableStats"`
	PopularFlows []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Name     string `json:"name,omitempty"`
	} `json:"popularFlows"`
	RunningFlowsCount []struct {
		KeyAsString time.Time `json:"key_as_string"`
		Key         int64     `json:"key"`
		DocCount    int       `json:"doc_count"`
	} `json:"runningFlowsCount"`
	EventOutcomesCount []interface{} `json:"eventOutcomesCount"`
	AllFlows           []string      `json:"allFlows"`
}

type Message struct {
	Message string `json:"message"`
}