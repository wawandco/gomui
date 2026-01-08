package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DialogEl creates a modal dialog overlay
func DialogEl(open bool, children ...g.Node) g.Node {
	if open {
		return h.Dialog(h.Class("dialog"), g.Attr("open"), g.Group(children))
	}
	return h.Dialog(h.Class("dialog"), g.Group(children))
}

// Dialog creates a modal dialog overlay (convenience alias for DialogEl)
func Dialog(open bool, children ...g.Node) g.Node {
	return DialogEl(open, children...)
}

// DialogContent creates a wrapper div for dialog content (no class)
func DialogContent(children ...g.Node) g.Node {
	return h.Div(g.Group(children))
}

// DialogHeader creates the header section of a dialog
func DialogHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// DialogTitle creates a heading for dialog content (typically h2)
func DialogTitle(children ...g.Node) g.Node {
	return h.H2(g.Group(children))
}

// DialogDescription creates supporting text for dialogs
func DialogDescription(children ...g.Node) g.Node {
	return h.P(g.Group(children))
}

// DialogSection creates the content section of a dialog
func DialogSection(children ...g.Node) g.Node {
	return h.Section(g.Group(children))
}

// DialogFooter creates the footer section of a dialog
func DialogFooter(children ...g.Node) g.Node {
	return h.Footer(g.Group(children))
}

// Sheet (Drawer/Sidebar) components for slide-out panels
// Sheet creates a slide-out drawer component
func Sheet(open bool, children ...g.Node) g.Node {
	if open {
		return h.Div(h.Class("sheet"), g.Attr("open"), g.Group(children))
	}
	return h.Div(h.Class("sheet"), g.Group(children))
}

// SheetContent creates a wrapper div for sheet content (no class)
func SheetContent(children ...g.Node) g.Node {
	return h.Div(g.Group(children))
}

// SheetHeader creates the header section of a sheet
func SheetHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// SheetTitle creates a heading for sheet content
func SheetTitle(children ...g.Node) g.Node {
	return h.H2(g.Group(children))
}

// SheetDescription creates supporting text for sheets
func SheetDescription(children ...g.Node) g.Node {
	return h.P(g.Group(children))
}

// SheetSection creates the content section of a sheet
func SheetSection(children ...g.Node) g.Node {
	return h.Section(g.Group(children))
}

// SheetFooter creates the footer section of a sheet
func SheetFooter(children ...g.Node) g.Node {
	return h.Footer(g.Group(children))
}

// Popover components for floating content panels
// Popover creates a popover container
func Popover(children ...g.Node) g.Node {
	return h.Div(h.Class("popover"), g.Group(children))
}

// PopoverTrigger creates a button to toggle popover
func PopoverTrigger(children ...g.Node) g.Node {
	return h.Button(h.Type("button"), g.Group(children))
}

// PopoverContent creates the content area of popover (uses data-popover attribute, not class)
func PopoverContent(children ...g.Node) g.Node {
	return h.Div(g.Attr("data-popover"), h.Aria("hidden", "true"), g.Group(children))
}

// PopoverHeader creates a header section inside popover content
func PopoverHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// HoverCard components for hover-triggered content
// HoverCard creates a hover card container
func HoverCard(children ...g.Node) g.Node {
	return h.Div(h.Class("hover-card"), g.Group(children))
}

// HoverCardTrigger creates an element to trigger hover card
func HoverCardTrigger(children ...g.Node) g.Node {
	return h.Div(g.Group(children))
}

// HoverCardContent creates the content area of hover card
func HoverCardContent(children ...g.Node) g.Node {
	return h.Div(g.Attr("data-popover"), h.Aria("hidden", "true"), g.Group(children))
}

// HoverCardHeader creates a header section inside hover card content
func HoverCardHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// AlertDialog components for modal dialogs that require user response
// AlertDialog creates an alert dialog that interrupts the user
// Unlike regular Dialog, it has no close button and cannot be dismissed by clicking the backdrop
func AlertDialog(open bool, children ...g.Node) g.Node {
	if open {
		return h.Dialog(h.Class("dialog"), g.Attr("open"), g.Group(children))
	}
	return h.Dialog(h.Class("dialog"), g.Group(children))
}

// AlertDialogContent creates a wrapper div for alert dialog content (no class)
func AlertDialogContent(children ...g.Node) g.Node {
	return h.Div(g.Group(children))
}

// AlertDialogHeader creates the header section of an alert dialog
func AlertDialogHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// AlertDialogTitle creates a heading for alert dialog content
func AlertDialogTitle(children ...g.Node) g.Node {
	return h.H2(g.Group(children))
}

// AlertDialogDescription creates supporting text for alert dialogs
func AlertDialogDescription(children ...g.Node) g.Node {
	return h.P(g.Group(children))
}

// AlertDialogSection creates the content section of an alert dialog
func AlertDialogSection(children ...g.Node) g.Node {
	return h.Section(g.Group(children))
}

// AlertDialogFooter creates the footer section of an alert dialog with action buttons
func AlertDialogFooter(children ...g.Node) g.Node {
	return h.Footer(g.Group(children))
}
