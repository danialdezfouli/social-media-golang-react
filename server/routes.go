package server

import (
	"jupiter/app/auth"
	"jupiter/app/feeds"
)

func (r *Rest) routes() {
	auth.Routes(r.echo)
	feeds.Routes(r.echo)
}
