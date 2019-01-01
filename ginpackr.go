package ginpackr


import (
	"github.com/gobuffalo/packr/v2"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type localFileSystem struct {
	// http.FileSystem
	box     *packr.Box
	indexes bool
}

func LocalFile(box *packr.Box, indexes bool) *localFileSystem {
	return &localFileSystem{
		// 	FileSystem: gin.Dir(root, indexes),
		box:     box,
		indexes: indexes,
	}
}

func (l *localFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if l.box.Has(p) {
			return true
		} else {
			return false
		}
	}
	return false
}

// Static returns a middleware handler that serves static files in the given directory.
func PackrServe(urlPrefix string, box *packr.Box) gin.HandlerFunc {
	fileserver := http.FileServer(box)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}

	return func(c *gin.Context) {
		if p := strings.TrimPrefix(c.Request.URL.Path, urlPrefix); len(p) < len(c.Request.URL.Path) {
			if box.Has(p) {
				fileserver.ServeHTTP(c.Writer, c.Request)
				c.Abort()
			}
		}
	}
}

