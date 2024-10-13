package middleware

import (
	"SimpleForum/internal/transport/customHttp"
	"net/http"
)

func HomePageRoleDispatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		role := r.Context().Value("Role").(string)

		switch role {
		case "Admin":
			// AdminHomePage(w, r)
			//Admin homepage handler
		case "Moderator":
			// Moderator homepage handler
		case "User":
			// User homepage handler
		default:
			customHttp.HomePage(w, r)
			// Guest homepage handler

		}
	})
}
