package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ButtonVariant defines different button style variants
type ButtonVariant string

const (
	ButtonPrimary     ButtonVariant = "primary"
	ButtonSecondary   ButtonVariant = "secondary"
	ButtonOutline     ButtonVariant = "outline"
	ButtonGhost       ButtonVariant = "ghost"
	ButtonLink        ButtonVariant = "link"
	ButtonDestructive ButtonVariant = "destructive"
)

// ButtonSize defines different button sizes
type ButtonSize string

const (
	ButtonDefault ButtonSize = ""
	ButtonSm      ButtonSize = "sm"
	ButtonLg      ButtonSize = "lg"
	ButtonIcon    ButtonSize = "icon"
)

// Button creates a button with specified variant and size
// This is a convenience function that wraps ButtonEl without the isIcon parameter
func Button(variant ButtonVariant, size ButtonSize, children ...g.Node) g.Node {
	isIcon := size == ButtonIcon
	return ButtonEl(variant, size, isIcon, children...)
}

// ButtonEl creates a button element with specified variant, size, and icon flag
func ButtonEl(variant ButtonVariant, size ButtonSize, isIcon bool, children ...g.Node) g.Node {
	class := buildButtonClass(variant, size, isIcon)
	return h.Button(h.Class(class), g.Group(children))
}

// ButtonWithClasses creates a button element with additional custom classes
func ButtonWithClasses(classes string, variant ButtonVariant, size ButtonSize, isIcon bool, children ...g.Node) g.Node {
	class := buildButtonClass(variant, size, isIcon) + " " + classes
	return h.Button(h.Class(class), g.Group(children))
}

// LinkButtonEl creates an anchor element with button styling for navigation
// Use this when you need a link that visually appears as a button
func LinkButtonEl(variant ButtonVariant, size ButtonSize, isIcon bool, href string, children ...g.Node) g.Node {
	class := buildButtonClass(variant, size, isIcon)
	return h.A(h.Class(class), h.Href(href), g.Group(children))
}

// LinkButton creates a link with button styling (convenience function)
func LinkButton(variant ButtonVariant, size ButtonSize, href string, children ...g.Node) g.Node {
	isIcon := size == ButtonIcon
	return LinkButtonEl(variant, size, isIcon, href, children...)
}

// LinkButtonWithClasses creates an anchor element with button styling and additional custom classes
func LinkButtonWithClasses(classes string, variant ButtonVariant, size ButtonSize, isIcon bool, href string, children ...g.Node) g.Node {
	class := buildButtonClass(variant, size, isIcon) + " " + classes
	return h.A(h.Class(class), h.Href(href), g.Group(children))
}

// buildButtonClass constructs the appropriate button class based on variant, size, and icon flag
// Examples: "btn", "btn-primary", "btn-lg-destructive", "btn-sm-icon-outline"
func buildButtonClass(variant ButtonVariant, size ButtonSize, isIcon bool) string {
	class := "btn"

	// Build combined class: btn-{size}-{icon}-{variant}
	if size != "" {
		class += "-" + string(size)
	}
	if isIcon {
		class += "-icon"
	}
	if variant != "" && variant != ButtonPrimary {
		class += "-" + string(variant)
	}

	return class
}

// ButtonGroup components for grouping related buttons
// ButtonGroup creates a container that groups related buttons together
func ButtonGroup(children ...g.Node) g.Node {
	return h.Div(h.Role("group"), h.Class("button-group"), g.Group(children))
}

// ButtonGroupWithOrientation creates a button group with specified orientation
func ButtonGroupWithOrientation(orientation string, children ...g.Node) g.Node {
	return h.Div(
		h.Role("group"),
		h.Class("button-group"),
		g.Attr("data-orientation", orientation),
		g.Group(children),
	)
}

// ButtonGroupSeparator creates a visual separator between buttons in a group
func ButtonGroupSeparator() g.Node {
	return h.Hr(h.Role("separator"))
}
