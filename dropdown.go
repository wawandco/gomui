package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DropdownMenu creates a dropdown menu container
func DropdownMenu(children ...g.Node) g.Node {
	return h.Div(h.Class("dropdown-menu"), g.Group(children))
}

// DropdownMenuTrigger creates a button to open dropdown
func DropdownMenuTrigger(children ...g.Node) g.Node {
	return h.Button(h.Type("button"), g.Group(children))
}

// DropdownMenuContent creates the popover content area
func DropdownMenuContent(children ...g.Node) g.Node {
	return h.Div(g.Attr("data-popover", ""), h.Aria("hidden", "true"), g.Group(children))
}

// DropdownMenuMenu creates the menu wrapper with role="menu"
func DropdownMenuMenu(children ...g.Node) g.Node {
	return h.Div(h.Role("menu"), g.Group(children))
}

// DropdownMenuItem creates a clickable menu item with role="menuitem"
func DropdownMenuItem(children ...g.Node) g.Node {
	return h.Div(h.Role("menuitem"), g.Group(children))
}

// DropdownMenuItemCheckbox creates a checkbox menu item with role="menuitemcheckbox"
func DropdownMenuItemCheckbox(children ...g.Node) g.Node {
	return h.Div(h.Role("menuitemcheckbox"), g.Group(children))
}

// DropdownMenuItemRadio creates a radio menu item with role="menuitemradio"
func DropdownMenuItemRadio(children ...g.Node) g.Node {
	return h.Div(h.Role("menuitemradio"), g.Group(children))
}

// DropdownMenuGroup creates a grouped section with role="group"
func DropdownMenuGroup(children ...g.Node) g.Node {
	return h.Div(h.Role("group"), g.Group(children))
}

// DropdownMenuGroupHeading creates a group heading with role="heading"
func DropdownMenuGroupHeading(children ...g.Node) g.Node {
	return h.Div(h.Role("heading"), g.Group(children))
}

// DropdownMenuSeparator creates a visual divider between items
func DropdownMenuSeparator() g.Node {
	return h.Hr(h.Role("separator"))
}
