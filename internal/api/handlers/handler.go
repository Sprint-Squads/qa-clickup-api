package handlers

import (
	"github.com/Sprint-Squads/qa-clickup-api/internal/manager"
	"github.com/Sprint-Squads/qa-clickup-api/pkg/application"
)

// Handler model
type Handler struct {
	App application.Application
	Manager *manager.Manager
}

// Get - Handler initializer
func Get(app application.Application) *Handler {
	var handler Handler
	handler.App = app
	manager, _ := manager.Get(&app)

	return &Handler{
		App:     app,
		Manager: manager,
	}
}
