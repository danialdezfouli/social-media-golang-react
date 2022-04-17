package dto

type ProfileDTO struct {
	ID uint `param:"id"`
}

type TimelineDTO struct {
	Offset uint `query:"offset"`
}
