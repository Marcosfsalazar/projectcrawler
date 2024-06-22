package models

type FileContent struct {
	Content string `json:"content"`
}

type Directory struct {
	SubDirs map[string]Directory `json:"sub_dirs,omitempty"`
	Files   map[string]FileContent `json:"files,omitempty"`
}
