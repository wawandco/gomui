package gomui

// Feedback components validated against Basecoat UI v0.3.10-beta.2
//
// Note: The following components are NOT included as Basecoat UI does not
// provide dedicated CSS classes or patterns for them:
//   - Progress/ProgressBar - no dedicated component in Basecoat
//   - Tooltip (container) - use data-tooltip attribute directly on elements
//   - Skeleton - no dedicated skeleton component in Basecoat
//
// Alert variants have been limited to only those officially supported:
//   - alert (default)
//   - alert-destructive
// Previous Success, Warning, and Info variants have been removed.

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// AlertVariant defines different alert style variants
type AlertVariant string

const (
	AlertDefault     AlertVariant = "alert"
	AlertDestructive AlertVariant = "alert-destructive"
)

// Alert creates a message container with variant styling
func Alert(variant AlertVariant, children ...g.Node) g.Node {
	return h.Div(h.Class(string(variant)), g.Group(children))
}

// AlertTitle creates a heading for alert content (uses plain h2 without class)
func AlertTitle(children ...g.Node) g.Node {
	return h.H2(g.Group(children))
}

// AlertDescription creates descriptive text for alerts (uses section without class)
func AlertDescription(children ...g.Node) g.Node {
	return h.Section(g.Group(children))
}

// Toast creates a notification toast container
func Toast(children ...g.Node) g.Node {
	return h.Div(
		h.Class("toast"),
		h.Role("status"),
		g.Attr("aria-atomic", "true"),
		g.Attr("aria-hidden", "false"),
		g.Group(children),
	)
}

// ToastContent creates the content wrapper for toast (required by Basecoat)
func ToastContent(children ...g.Node) g.Node {
	return h.Div(h.Class("toast-content"), g.Group(children))
}

// ToastSection creates a section wrapper for toast title and description
func ToastSection(children ...g.Node) g.Node {
	return h.Section(g.Group(children))
}

// ToastTitle creates a heading for toast content (uses h2 without class)
func ToastTitle(children ...g.Node) g.Node {
	return h.H2(g.Group(children))
}

// ToastDescription creates descriptive text for toasts (uses p without class)
func ToastDescription(children ...g.Node) g.Node {
	return h.P(g.Group(children))
}

// ToastFooter creates a footer wrapper for toast buttons
func ToastFooter(children ...g.Node) g.Node {
	return h.Footer(g.Group(children))
}

// ToastAction creates a button for toast interaction
func ToastAction(children ...g.Node) g.Node {
	return h.Button(h.Class("btn"), g.Group(children))
}

// Toaster creates the container for displaying toasts
func Toaster(children ...g.Node) g.Node {
	return h.Div(h.ID("toaster"), h.Class("toaster"), g.Group(children))
}
