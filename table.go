package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// TableEl creates a styled table container
func TableEl(children ...g.Node) g.Node {
	return h.Table(h.Class("table"), g.Group(children))
}

// Table creates a styled table container (convenience alias for TableEl)
func Table(children ...g.Node) g.Node {
	return TableEl(children...)
}

// TableHeader creates the header section of a table
func TableHeader(children ...g.Node) g.Node {
	return h.THead(g.Group(children))
}

// TableBody creates the main content section of a table
func TableBody(children ...g.Node) g.Node {
	return h.TBody(g.Group(children))
}

// TableFooter creates the footer section of a table
func TableFooter(children ...g.Node) g.Node {
	return h.TFoot(g.Group(children))
}

// TableRow creates a single row in a table
func TableRow(children ...g.Node) g.Node {
	return h.Tr(g.Group(children))
}

// TableHead creates a header cell in a table
func TableHead(children ...g.Node) g.Node {
	return h.Th(g.Group(children))
}

// TableData creates a table data cell in a table
func TableData(children ...g.Node) g.Node {
	return h.Td(g.Group(children))
}

// TableCell creates a table data cell (convenience alias for TableData)
func TableCell(children ...g.Node) g.Node {
	return TableData(children...)
}

// TableCaption creates a caption for the table
func TableCaption(children ...g.Node) g.Node {
	return h.Caption(g.Group(children))
}
