package remote

type ErrorResponse struct {
	Error string `json:"error"`
}

type RemoteResponse struct {
	RemoteName string `json:"remoteName"`
	Version    string `json:"version"`
	RemoteURL  string `json:"remoteURL"`
}
