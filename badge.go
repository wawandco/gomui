package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// BadgeVariant defines different badge style variants
type BadgeVariant string

const (
	BadgePrimary     BadgeVariant = "badge"
	BadgeSecondary   BadgeVariant = "badge-secondary"
	BadgeDestructive BadgeVariant = "badge-destructive"
	BadgeOutline     BadgeVariant = "badge-outline"
)

// Badge creates a colored status label
func Badge(variant BadgeVariant, children ...g.Node) g.Node {
	return h.Span(h.Class(string(variant)), g.Group(children))
}
