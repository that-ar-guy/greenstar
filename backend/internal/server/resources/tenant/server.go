// Code generated by greenstar scripts; DO NOT EDIT.

package tenant

import (
	"github.com/arikkfir/greenstar/backend/internal/server/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type Server struct {
	h    Handler
	pool *pgxpool.Pool
}

func (s *Server) Register(mux *http.ServeMux) {
	mux.Handle("/tenants", &middleware.Handlers{
		GET:    middleware.PostgresMiddleware(s.pool, pgx.ReadOnly, http.HandlerFunc(s.List)),
		POST:   middleware.PostgresMiddleware(s.pool, pgx.ReadWrite, http.HandlerFunc(s.Create)),
		DELETE: middleware.PostgresMiddleware(s.pool, pgx.ReadWrite, http.HandlerFunc(s.DeleteAll)),
	})
	mux.Handle("/tenants/{id}", &middleware.Handlers{
		GET:    middleware.PostgresMiddleware(s.pool, pgx.ReadOnly, http.HandlerFunc(s.Get)),
		PATCH:  middleware.PostgresMiddleware(s.pool, pgx.ReadWrite, http.HandlerFunc(s.Patch)),
		PUT:    middleware.PostgresMiddleware(s.pool, pgx.ReadWrite, http.HandlerFunc(s.Update)),
		DELETE: middleware.PostgresMiddleware(s.pool, pgx.ReadWrite, http.HandlerFunc(s.Delete)),
	})
}

func NewServer(pool *pgxpool.Pool, handler Handler) *Server {
	return &Server{
		h:    handler,
		pool: pool,
	}
}
