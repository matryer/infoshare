package comments

import (
	"net/url"
	"strings"
)

// SimplifyURL simplifies a URL.
func SimplifyURL(url *url.URL) string {
	return strings.ToLower(url.Host + "/" + strings.TrimPrefix(url.Path, "/"))
}
