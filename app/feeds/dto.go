package feeds

type profileDTO struct {
	ID uint `param:"id"`
}

type timelineDTO struct {
	Offset uint `query:"offset"`
}
