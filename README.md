# ventana

A small wrapper around [github.com/webview/webview_go](https://github.com/webview/webview_go) to reduce redundant code in small projects.

## Example

```go
package main

import (
	"embed"

	"github.com/ronelliott/ventana"
)

//go:embed index.html *.css *.js
var assets embed.FS

func main() {
	window, err := ventana.NewWindow(
		ventana.WithDebug(true),
		ventana.WithTitle("My Awesome Ventana"),
		ventana.WithSize(2400, 1400, ventana.HintFixed),
		ventana.WithServerEnabled(assets),
		ventana.WithInitialPath("index.html"),
	)
	if err != nil {
		panic(err)
		return
	}
	defer window.Close()
	window.Run()
}
```
