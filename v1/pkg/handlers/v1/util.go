package v1

import (
	ua "github.com/mileusna/useragent"
)

// === request util === //

// parseUserAgent parsea userAgentStr y retorna la informacion
func parseUserAgent(userAgentStr string) *userAgent {
	d := ua.Parse(userAgentStr)
	return &userAgent{
		Name:   d.Name,
		OS:     d.OS,
		Device: d.Device,
		String: d.String,
	}
}

type userAgent struct {
	Name   string
	OS     string
	Device string
	String string
}
