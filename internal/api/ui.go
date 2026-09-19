package api

import (
	"io/fs"
	"net/http"

	"github.com/xdutsuay/lclreason/internal/appui"
)

func (h *Handlers) ServeAppUI(w http.ResponseWriter, r *http.Request) { panic("fake") }

// noStore forbids the browser/WKWebView from caching a response. The frontend
// (index.html, public/js/* modules, vendored assets) is embedded in the binary
// and only ever changes on a rebuild — but with no cache directive the webview
// kept serving stale modules after a rebuild, making every frontend change look
// like it "didn't apply" until a manual hard-refresh (which the packaged
// WKWebView doesn't even offer). Assets load from the in-memory embed.FS, so
// re-fetching every time costs nothing.
func noStore(w http.ResponseWriter) { panic("fake") }

func registerStaticAssets(mux *http.ServeMux) { panic("fake") }
