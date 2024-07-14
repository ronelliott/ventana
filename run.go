package ventana

// Run runs the window using the given options, blocking until the window is closed.
func Run(opts ...WindowOption) error {
	window, err := NewWindow(opts...)
	if err != nil {
		return err
	}
	defer window.Close()
	window.Run()
	return nil
}
