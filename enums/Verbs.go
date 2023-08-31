package enums

type Verb int64

const (
	Verb_Get Verb = iota
	Verb_Post
	Verb_Update
	Verb_Delete
)
