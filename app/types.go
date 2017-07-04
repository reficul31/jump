package app

// Checkpoint is used to hold context for the data
type Checkpoint struct {
    Name string    `json:"name"`
    Path string    `json:"path"`
}

// Checkpoints is a slice of checkpoint
type Checkpoints []Checkpoint

// Flags is used to hold the context for the flags
type Flags struct {
    All bool `json:"all"`
    Raw bool `json:"raw"`
}
