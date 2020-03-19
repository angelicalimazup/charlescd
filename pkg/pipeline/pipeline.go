package pipeline

type Version struct {
	Version    string `json:"version"`
	VersionURL string `json:"versionUrl"`
}

type GithubAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Istio struct {
	VirtualService   map[string]interface{} `json:"virtualService"`
	DestinationRules map[string]interface{} `json:"destinationRules"`
}

type Pipeline struct {
	Name           string        `json:"appName"`
	Namespace      string        `json:"appNamespace""`
	Versions       []*Version    `json:"versions"`
	UnusedVersions []*Version    `json:"unusedVersions"`
	Webhook        string        `json:"webhookUrl"`
	HelmRepository string        `json:"helmUrl"`
	GithubAccount  GithubAccount `json:"github"`
	Istio          Istio         `json:"istio"`
}
