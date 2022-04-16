package server

import (
	"jupiter/app/auth"
	"jupiter/app/feeds"
	"jupiter/app/relationship"
)

func (r *Rest) routes() {
	auth.Routes(r.echo)
	feeds.Routes(r.echo)
	relationship.Routes(r.echo)
}
