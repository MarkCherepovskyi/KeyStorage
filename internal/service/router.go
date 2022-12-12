package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/config"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data/pg"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/handlers"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxContainerQ(pg.NewContainerQ(cfg.DB())),
			helpers.CtxLog(s.log),
		),
	)
	r.Route("/KeyStorage", func(r chi.Router) {

		r.Post("/get", handlers.GetContainer)
		r.Post("/create", handlers.CreateContainer)

	})

	return r
}
