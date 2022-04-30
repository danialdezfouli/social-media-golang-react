package dto

type ProfileDTO struct {
	Username string `param:"id"`
}

type TimelineDTO struct {
	Offset uint `query:"offset"`
}

type SearchDTO struct {
	Query string `query:"q"`
}

type PostDTO struct {
	ID uint `param:"id"`
}
