//go:test ignoretest

package models

type Remote struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	RemoteName string `json:"remoteName"`
	Version    string `json:"version"`
	RemoteURL  string `json:"remoteURL"`
}

type SuccessResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type RemoteResponse struct {
	RemoteName string `json:"remoteName"`
	Version    string `json:"version"`
	RemoteURL  string `json:"remoteURL"`
}
