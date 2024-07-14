package ventana

import (
	"context"
	"net/http"

	webview "github.com/webview/webview_go"
)

// Window represents a window.
type Window interface {
	webview.WebView
	// Close closes the window and cleans up. The window will be unusable after a call to this method.
	Close()
	// SendEvent sends an event to the window.
	SendEvent(*Event) error
}

// windowImpl is the default implementation of Window.
type windowImpl struct {
	webview.WebView
	// port is the port used by the HTTP server created by the window.
	port string
	// server is the HTTP server used to serve the UI assets.
	server *http.Server
	// uiEventHandlerName is the name of the event handler used by the window.
	uiEventHandlerName string
}

// NewWindow creates a new window with the given options.
func NewWindow(opts ...WindowOption) (Window, error) {
	window := &windowImpl{
		uiEventHandlerName: "onEventReceived",
	}

	for _, opt := range opts {
		if err := opt(window); err != nil {
			return nil, err
		}
	}

	return window, nil
}

// Close closes the window and cleans up. The window will be unusable after a call to this method.
func (window *windowImpl) Close() {
	if window.server != nil {
		window.server.Shutdown(context.TODO())
	}

	window.Destroy()
}

// Run runs the window.
func (window *windowImpl) Run() {
	if window.server != nil {
		go window.server.ListenAndServe()
	}

	window.WebView.Run()
}
