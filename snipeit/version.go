package snipeit

type Version struct {
	Version      string `json:"version"`
	BuildVersion string `json:"build_version"`
	HashVersion  string `json:"hash_version"`
	FullVersion  string `json:"full_version"`
}
