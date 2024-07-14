package ventana

import webview "github.com/webview/webview_go"

// Hint represents a hint for the size of a window.
type Hint webview.Hint

const (
	// HintNone is the default size.
	HintNone Hint = webview.HintNone
	// HintFixed is the size that cannot be changed by a user.
	HintFixed Hint = webview.HintFixed
	// HintMin is the minimum size.
	HintMin Hint = webview.HintMin
	// HintMax is the maximum size.
	HintMax Hint = webview.HintMax
)
