package models

type Callback struct {
	Object string  `json:"object,omitempty"`
	Entry  []Entry `json:"entry,omitempty"`
}
