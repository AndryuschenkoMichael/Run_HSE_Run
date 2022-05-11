package handler

import (
	"Run_Hse_Run/pkg/service"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() chi.Router {
	router := chi.NewRouter()
	// fix if I will use websocket
	//router.Use(middleware.Logger, middleware.Timeout(time.Minute))
	router.Use(middleware.Logger)

	router.Route("/auth", func(router chi.Router) {
		router.Post("/send-email", h.sendEmail)
		router.Post("/check-auth", h.checkAuth)
		router.Post("/create-user", h.createUser)
	})

	router.Route("/api", func(router chi.Router) {
		router.Use(h.authorizationOnly)
		router.Route("/friends", func(router chi.Router) {
			router.Put("/add-friend", h.addFriend)
			router.Delete("/delete-friend", h.deleteFriend)
			router.Get("/get-friends", h.getFriends)
		})
		router.Route("/users", func(router chi.Router) {
			router.Get("/get-me", h.getMe)
			router.Get("/get-by-id", h.getUserById)
			router.Get("/get-by-nickname", h.getUserByNickname)
			router.Put("/change-nickname", h.changeNickname)
			router.Put("/change-profile-image", h.changeProfileImage)
		})

		router.Route("/game", func(router chi.Router) {
			router.Get("/get-rooms-by-code", h.getRoomByCode)
		})
	})

	router.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "%s", "pong")
		if err != nil {
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}

	})

	return router
}
