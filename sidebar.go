package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type SidebarPosition string

const (
	SidebarLeft  SidebarPosition = "left"
	SidebarRight SidebarPosition = "right"
)

// Sidebar component for side navigation panels
// Sidebar creates a sidebar element with specified side and visibility
func Sidebar(side SidebarPosition, children ...g.Node) g.Node {
	return h.Aside(
		h.Class("sidebar"),
		g.Attr("aria-hidden", "true"),
		g.Attr("data-side", string(side)),
		g.Group(children),
	)
}

// SidebarHeader creates the header section of the sidebar
func SidebarHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// SidebarNav creates the navigation wrapper
func SidebarNav(label string, children ...g.Node) g.Node {
	return h.Nav(
		h.Aria("label", label),
		h.Aria("expanded", "false"),
		h.Aria("haspopup", "menu"),
		g.Group(children),
	)
}

// SidebarSection creates a scrollable section
func SidebarSection(children ...g.Node) g.Node {
	return h.Section(h.Class("scrollbar"), g.Group(children))
}

// SidebarGroup creates a grouped section
func SidebarGroup(children ...g.Node) g.Node {
	return h.Div(h.Role("group"), g.Group(children))
}

// SidebarHeading creates a group heading
func SidebarHeading(id string, children ...g.Node) g.Node {
	return h.H3(h.ID(id), g.Group(children))
}

// SidebarList creates a list of items
func SidebarList(children ...g.Node) g.Node {
	return h.Ul(g.Group(children))
}

// SidebarItem creates a list item
func SidebarItem(children ...g.Node) g.Node {
	return h.Li(g.Group(children))
}

// SidebarFooter creates an optional footer section
func SidebarFooter(children ...g.Node) g.Node {
	return h.Footer(g.Group(children))
}

// SidebarToggle creates a button to toggle the sidebar visibility
// The button dispatches a custom event 'basecoat:sidebar' when clicked
func SidebarToggle(children ...g.Node) g.Node {
	return h.Button(
		h.Type("button"),
		g.Group(children),
		g.Attr("onclick", `document.dispatchEvent(new CustomEvent('basecoat:sidebar'))`),
	)
}

// SidebarOpen creates a button to open the sidebar
func SidebarOpen(sidebarID string, children ...g.Node) g.Node {
	return h.Button(
		h.Type("button"),
		g.Group(children),
		g.Attr("onclick", "document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail: { id: '"+sidebarID+"', action: 'open' } }));"),
	)
}

// SidebarClose creates a button to close the sidebar
// The button dispatches a custom event 'basecoat:sidebar' when clicked
func SidebarClose(sidebarID string, children ...g.Node) g.Node {
	return h.Button(
		h.Type("button"),
		g.Group(children),
		g.Attr("onclick", "document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail: { id: '"+sidebarID+"', action: 'close' } }));"),
	)
}

// SidebarAssets includes the necessary JS for sidebar functionality
// Make sure to include this in your HTML head or before the closing body tag
func SidebarAssets() g.Node {
	return h.Script(h.Src("https://cdn.jsdelivr.net/npm/basecoat-css@0.3.10-beta.2/dist/js/sidebar.min.js"), h.Defer())
}
