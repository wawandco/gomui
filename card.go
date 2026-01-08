package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Card components for creating content containers
// Card creates a main card container
func Card(children ...g.Node) g.Node {
	return h.Div(h.Class("card"), g.Group(children))
}

// CardWithClasses creates a card container with additional classes
func CardWithClasses(classes string, children ...g.Node) g.Node {
	return h.Div(h.Class("card "+classes), g.Group(children))
}

// CardHeader creates the top section of a card
func CardHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// CardTitle creates a title heading for card header (uses h2 without class)
func CardTitle(children ...g.Node) g.Node {
	return h.H2(g.Group(children))
}

// CardDescription creates descriptive text for card header (uses p without class)
func CardDescription(children ...g.Node) g.Node {
	return h.P(g.Group(children))
}

// CardContent creates the main content area of a card
func CardContent(children ...g.Node) g.Node {
	return h.Section(g.Group(children))
}

// CardFooter creates the bottom section of a card
func CardFooter(children ...g.Node) g.Node {
	return h.Footer(g.Group(children))
}
