package route

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

type dbContextKey string

var dbKey = dbContextKey("databaseContextKey")

func InjectDBMiddleware(db *gorm.DB) func(http.HandlerFunc) http.HandlerFunc {
	return func (next http.HandlerFunc) http.HandlerFunc { 
		return func(w http.ResponseWriter, r *http.Request) {
				var ctx context.Context = context.WithValue(r.Context(), dbKey, db)

				r = r.WithContext(ctx)
				next(w, r)
			}
	}
}


func GetDBFromContext(ctx context.Context) *gorm.DB {
	var (
		db *gorm.DB
		ok bool
	)

	db, ok = ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		return nil
	}

	return db
}
