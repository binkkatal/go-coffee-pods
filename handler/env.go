package handler

import (
	"github.com/binkkatal/go-coffee-pods/dao"
)

// Env will be the receiver for our handlers. It can pass various
// environment variables we may need in our handlers, like the
// database pool as implement here, along with loggers, templates etc.
type Env struct {
	DS        dao.Datastore
	AssetsDir *string
}
