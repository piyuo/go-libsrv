package util

import (
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"

	useragent "github.com/mileusna/useragent"
)

// GetUserAgentID return short id from user agent
//
//	txt := GetUserAgentID(request)
//
func GetUserAgentID(r *http.Request) string {
	u := useragent.Parse(r.UserAgent())
	return u.Device + ", " + u.OS + " " + u.OSVersion + ", " + u.Name + " " + u.Version
}

// ParseUserAgent return browser name,browser version,os name,os version,device from user agent
//
//	browserName,browserVer,osName,osVer,device := ParseUserAgent(ua)
//
func ParseUserAgent(ua string) (string, string, string, string, string) {
	u := useragent.Parse(ua)
	return u.Name, u.Version, u.OS, u.OSVersion, u.Device
}

// GetUserAgent return user agent
//
//	ua := GetUserAgent(request)
//
func GetUserAgent(r *http.Request) string {
	return r.UserAgent()
}

// GetIP return ip from request
//
//	ip := GetIP(request)
//
func GetIP(r *http.Request) string {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	if ips != "" {
		splitIps := strings.Split(ips, ",")
		for _, ip := range splitIps {
			netIP := net.ParseIP(ip)
			if netIP != nil {
				return ip
			}
		}
	}

	//Get IP from RemoteAddr
	if r.RemoteAddr != "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return ""
		}
		netIP = net.ParseIP(ip)
		if netIP != nil {
			return ip
		}
	}
	return ""
}

// GetLocale parse http header Accept-Language field and return first one as default language
//
//	defaultLocale := GetLocale(request)
//
func GetLocale(r *http.Request) string {
	languages := parseAcceptLanguage(r.Header.Get("Accept-Language"))
	return languages[0]
}

// GetAcceptLanguage parse http header Accept-Language field to sorted string list
//
//	list := GetAcceptLanguage(request)
//
func GetAcceptLanguage(r *http.Request) []string {
	return parseAcceptLanguage(r.Header.Get("Accept-Language"))
}

// parseAcceptLanguage parse http header Accept-Language field to sorted string list
//
//	list := parseAcceptLanguage("da, en-gb;q=0.8, en;q=0.7") // []string{"da","en-gb","en"}
//
func parseAcceptLanguage(acptLang string) []string {
	if acptLang == "" {
		return []string{"en-us"}
	}

	type langQ struct {
		Lang string
		Q    float64
	}

	langQS := []*langQ{}
	accepts := strings.Split(acptLang, ",")
	for _, accept := range accepts {
		accept = strings.Trim(accept, " ")
		args := strings.Split(accept, ";")
		if len(args) == 1 {
			langQS = append(langQS, &langQ{
				Lang: args[0],
				Q:    1,
			})
		} else {
			qp := strings.Split(args[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err == nil {
				langQS = append(langQS, &langQ{
					Lang: args[0],
					Q:    q,
				})
			}
		}
	}

	sort.SliceStable(langQS, func(i, j int) bool {
		return langQS[i].Q > langQS[j].Q
	})

	result := []string{}
	for _, lq := range langQS {
		result = append(result, strings.ToLower(lq.Lang))
	}

	return result
}
