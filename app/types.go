package app

type Checkpoint struct {
	Name string    `json:"name"`
	Path string    `json:"path"`
}

type Checkpoints []Checkpoint

type Flags struct {
	All bool `json:"all"`
	Raw bool `json:"raw"`
}