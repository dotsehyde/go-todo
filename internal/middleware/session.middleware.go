package middleware

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
)

var SessionManager *scs.SessionManager

func InitSessionManager(db *sql.DB) {
	SessionManager = scs.New()
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.Cookie.Persist = true
	SessionManager.Cookie.Secure = false //set true if on https
	SessionManager.Store = sqlite3store.New(db)
	// sqlite3store.NewWithCleanupInterval(1 * time.Hour)
	fmt.Println("SessionManager initialized")
}
