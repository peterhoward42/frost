package contract

type Comment struct {
	Type string
	Text string
}

func NewComment(text string) *Comment {
	return &Comment{
		Type: "Comment",
		Text: text,
	}
}
