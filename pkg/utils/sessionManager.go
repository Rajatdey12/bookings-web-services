package utils

import (
	"net/http"
	"time"

	"github.com/Rajatdey12/bookings-web-services/pkg/config"
	"github.com/alexedwards/scs/v2"
)

// CreateSession for the app..
func CreateSession(aConfig config.AppConfig) *scs.SessionManager {

	session := scs.New()
	session.Lifetime = 30 * time.Minute
	session.Cookie.HttpOnly = true
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = aConfig.InProduction

	return session
}
