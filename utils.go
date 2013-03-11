package torgo

/*========================================================================
		Insion / email: insion@lihuashu.com
========================================================================*/
//v0.6

import (
	"strings"
	"time"
)

func webTime(t time.Time) string {
	ftime := t.Format(time.RFC1123)
	if strings.HasSuffix(ftime, "UTC") {
		ftime = ftime[0:len(ftime)-3] + "GMT"
	}
	return ftime
}
