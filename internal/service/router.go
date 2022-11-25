package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/config"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data/pg"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxContainerQ(pg.NewContainerQ(cfg.DB())),
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/KeyStorage", func(r chi.Router) {

		r.Get("/{id}", handlers.GetContainer)
		r.Post("/", handlers.CreateContainer)

	})

	return r
}
