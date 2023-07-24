package minimax

import "net/http"

const (
	APIURLv1 = "https://api.minimax.chat/v1"
)

type ClientConfig struct {
	authToken  string
	BaseUrl    string
	OrgId      string
	HttpClient *http.Client
}

func DefaultConfig(authToken string, orgId string) ClientConfig {
	return ClientConfig{
		authToken:  authToken,
		OrgId:      orgId,
		BaseUrl:    APIURLv1,
		HttpClient: &http.Client{},
	}
}
