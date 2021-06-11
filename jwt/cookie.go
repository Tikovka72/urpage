package jwt

import (
	"net/http"
	"strconv"
	"time"
)

func AddJWTCookie(token string, tokenExpireDate time.Time, writer http.ResponseWriter) {
	cookieToken := http.Cookie{
		Name:     "JWT",
		Value:    token,
		Expires:  tokenExpireDate,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(writer, &cookieToken)
}

func AddRefreshTokenCookie(refreshToken string, tokenId int64, userId int, refreshExpireDate time.Time, writer http.ResponseWriter) {
	cookieRefresh := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshToken,
		Expires:  refreshExpireDate,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	cookieRefreshId := http.Cookie{
		Name:     "RefreshTokenId",
		Value:    strconv.FormatInt(tokenId, 10),
		Expires:  refreshExpireDate,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	cookieRefreshUserId := http.Cookie{
		Name:     "RefreshTokenUserId",
		Value:    strconv.Itoa(userId),
		Expires:  refreshExpireDate,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(writer, &cookieRefresh)
	http.SetCookie(writer, &cookieRefreshId)
	http.SetCookie(writer, &cookieRefreshUserId)
}

func DeleteJWTCookie(writer http.ResponseWriter) {
	http.SetCookie(writer,
		&http.Cookie{
			Name:    "JWT",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		})
}

func DeleteRefreshTokenCookie(writer http.ResponseWriter) {
	http.SetCookie(writer,
		&http.Cookie{
			Name:    "RefreshToken",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		})

	http.SetCookie(writer,
		&http.Cookie{
			Name:    "RefreshTokenId",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		})

	http.SetCookie(writer,
		&http.Cookie{
			Name:    "RefreshTokenUserId",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		})
}
