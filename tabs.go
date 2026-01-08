package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Tabs components for content switching
// Tabs creates a tab container
func Tabs(children ...g.Node) g.Node {
	return h.Div(h.Class("tabs w-full"), g.Group(children))
}

// TabsList creates a list of tab triggers
func TabsList(children ...g.Node) g.Node {
	return h.Nav(h.Class("w-full"), h.Role("tablist"), h.Aria("orientation", "horizontal"), g.Group(children))
}

// TabItem creates a single tab button
func TabItem(id string, selected bool, children ...g.Node) g.Node {
	return h.Button(
		g.If(selected, g.Attr("aria-selected", "true")),
		h.Type("button"), h.Role("tab"), h.TabIndex("0"), h.ID(id), g.Group(children))
}

// TabPanel creates a tab panel with proper ARIA attributes
func TabPanel(id string, labelledBy string, selected bool, children ...g.Node) g.Node {
	attrs := []g.Node{
		h.Role("tabpanel"),
		h.ID(id),
		h.Aria("labelledby", labelledBy),
		h.TabIndex("-1"),
	}
	if selected {
		attrs = append(attrs, h.Aria("selected", "true"))
	} else {
		attrs = append(attrs, h.Aria("selected", "false"), g.Attr("hidden", ""))
	}
	attrs = append(attrs, g.Group(children))
	return h.Div(attrs...)
}
