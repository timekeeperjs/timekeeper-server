package remote

type Remote struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	RemoteName string `json:"remoteName"`
	Version    string `json:"version"`
	RemoteURL  string `json:"remoteURL"`
}
