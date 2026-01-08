package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Kbd components for displaying keyboard keys and shortcuts

// Kbd creates a keyboard key display element
func Kbd(children ...g.Node) g.Node {
	return h.Kbd(h.Class("kbd"), g.Group(children))
}

// KbdWithClasses creates a keyboard key display element with additional classes
func KbdWithClasses(classes string, children ...g.Node) g.Node {
	return h.Kbd(h.Class("kbd "+classes), g.Group(children))
}

// KbdEl creates a keyboard key display element with custom attributes
func KbdEl(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{h.Class("kbd")}
	allAttrs = append(allAttrs, attrs...)
	return h.Kbd(allAttrs...)
}
