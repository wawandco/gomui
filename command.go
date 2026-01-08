package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Command components for command palette/search functionality
// Command creates a command palette container
func Command(children ...g.Node) g.Node {
	return h.Div(h.Class("command"), g.Group(children))
}

// CommandDialog creates a dialog wrapper for command palette
func CommandDialog(children ...g.Node) g.Node {
	return h.Dialog(h.Class("command-dialog"), g.Group(children))
}

// CommandHeader creates a header with search input for command palette
func CommandHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// CommandInput creates a search input for command palette with proper ARIA attributes
func CommandInput(id string, menuID string, attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		h.Type("text"),
		h.ID(id),
		h.Role("combobox"),
		h.Aria("expanded", "true"),
		h.Aria("controls", menuID),
		h.Aria("autocomplete", "list"),
		g.Attr("autocomplete", "off"),
		g.Attr("autocorrect", "off"),
		g.Attr("spellcheck", "false"),
	}
	allAttrs = append(allAttrs, attrs...)
	return h.Input(allAttrs...)
}

// CommandMenu creates a menu container for command items with role="menu"
func CommandMenu(id string, children ...g.Node) g.Node {
	return h.Div(
		h.Role("menu"),
		h.ID(id),
		h.Aria("orientation", "vertical"),
		g.Group(children),
	)
}

// CommandGroup creates a grouped section with role="group"
func CommandGroup(labelledBy string, children ...g.Node) g.Node {
	return h.Div(h.Role("group"), h.Aria("labelledby", labelledBy), g.Group(children))
}

// CommandGroupHeading creates a group heading with role="heading"
func CommandGroupHeading(id string, children ...g.Node) g.Node {
	return h.Span(h.Role("heading"), h.ID(id), g.Group(children))
}

// CommandItem creates a single selectable command item with role="menuitem"
func CommandItem(children ...g.Node) g.Node {
	return h.Div(h.Role("menuitem"), g.Group(children))
}

// CommandSeparator creates a visual separator between command groups
func CommandSeparator() g.Node {
	return h.Hr(h.Role("separator"))
}
