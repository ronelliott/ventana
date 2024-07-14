package ventana

import (
	"fmt"
	"io/fs"
	"math/rand"
	"net/http"

	webview "github.com/webview/webview_go"
)

// WindowOption represents an option for a Window.
type WindowOption func(*windowImpl) error

// WithBoundFunction binds a function to the window.
func WithBoundFunction(name string, fn interface{}) WindowOption {
	return func(window *windowImpl) error {
		return window.Bind(name, fn)
	}
}

// WithDebug enables debug mode for the window. Note that this must be the first
// option in an option list, and must always be present.
func WithDebug(debug bool) WindowOption {
	return func(window *windowImpl) error {
		window.WebView = webview.New(debug)
		return nil
	}
}

// WithInitialPath sets the initial path of the window.
func WithInitialPath(path string) WindowOption {
	return func(window *windowImpl) error {
		return WithInitialURL(fmt.Sprintf("http://%s/%s", window.port, path))(window)
	}
}

// WithInitialURL sets the initial URL of the window.
func WithInitialURL(url string) WindowOption {
	return func(window *windowImpl) error {
		window.Navigate(url)
		return nil
	}
}

// WithSize sets the size of the window.
func WithSize(width, height int, hint Hint) WindowOption {
	return func(window *windowImpl) error {
		window.SetSize(width, height, webview.Hint(hint))
		return nil
	}
}

// WithTitle sets the title of the window.
func WithTitle(title string) WindowOption {
	return func(window *windowImpl) error {
		window.SetTitle(title)
		return nil
	}
}

// WithUIEventHandlerName sets the name of the UI event handler.
func WithUIEventHandlerName(name string) WindowOption {
	return func(window *windowImpl) error {
		window.uiEventHandlerName = name
		return nil
	}
}

// WithServerEnabled enables the HTTP server for the window. Note that
// this option is required if you want to serve UI assets and must be used in
// conjunction and after WithPort.
func WithServerEnabled(filesystem fs.FS) WindowOption {
	return func(window *windowImpl) error {
		if len(window.port) == 0 {
			window.port = fmt.Sprintf("127.0.0.1:%d", rand.Intn(54000)+4000)
		}

		fileHandler := http.FileServer(http.FS(filesystem))
		window.server = &http.Server{
			Addr:    window.port,
			Handler: fileHandler,
		}

		return nil
	}
}

// WithPort sets the port used by the HTTP server created by the window.
func WithPort(port string) WindowOption {
	return func(window *windowImpl) error {
		window.port = port
		return nil
	}
}
