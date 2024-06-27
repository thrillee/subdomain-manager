package httpsvr

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type authHandler func(http.ResponseWriter, *http.Request)

func middlewareWhitelistedIP(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-FORWARDED-FOR")
		whitelistedIPs := os.Getenv("WHITELISTED_IPS")

		if strings.Contains(whitelistedIPs, ip) {
			responseWithError(w, 401, fmt.Sprintf("Invalid Caller: %v", ip))
			return
		}

		handler(w, r)
	}
}
