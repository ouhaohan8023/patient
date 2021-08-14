package lib

import (
	"errors"
	"github.com/kataras/iris/v12"
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}

func ValidateUsPhoneNumber(phone string) bool {
	// todo
	return true
	//regular := "/^(1?|(1\\-)?)\\d{10,12}$/"
	//re := regexp.MustCompile(regular)
	//return re.MatchString(phone)
}

func CommonResponse(code int, msg string, content iris.Map) iris.Map {
	ctx := iris.Map{
		"status":  code,
		"msg":     msg,
		"content": content,
	}
	return ctx
}
