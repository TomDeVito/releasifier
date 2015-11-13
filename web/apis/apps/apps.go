package app

import (
	"net/http"

	"github.com/alinz/releasifier/lib/utils"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

func getAllApps(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, 200, "all apps")
}

func createApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, 200, "create an app")
}

func removeApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	appID := chi.URLParams(ctx)["appID"]
	utils.Respond(w, 200, "delete an app with id of "+appID)
}