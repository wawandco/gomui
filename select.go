package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// SelectWithSearchProps defines the properties for the SelectWithSearch component
type SelectWithSearchProps struct {
	ID    string
	Label string
	Name  string
	Value string

	Options SelectOptions

	Extras           g.Group
	SearchInputAttrs g.Group
}

// SelectWithSearch creates a select component with search functionality
// and a hidden input to store the selected value
func SelectWithSearch(props SelectWithSearchProps) g.Node {
	selectedLabel := props.Options.SelectedLabel()
	return SelectContainer(props.ID,
		h.Button(
			h.Type("button"),
			h.Class("btn-outline"),
			h.ID(props.ID+"-trigger"),
			h.Aria("expanded", "false"),
			h.Aria("controls", props.ID+"-listbox"),
			h.Span(
				h.Class("truncate"),
				g.If(selectedLabel != "", g.Text(selectedLabel)),
				g.If(selectedLabel == "", g.Text(props.Label)),
			),

			g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-chevrons-up-down-icon lucide-chevrons-up-down text-muted-foreground opacity-50 shrink-0">
				<path d="m7 15 5 5 5-5" />
				<path d="m7 9 5-5 5 5" />
			</svg>`),
		),
		SelectPopover(props.ID+"-popover",
			SelectHeader(
				g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-search-icon lucide-search">
						<circle cx="11" cy="11" r="8" />
						<path d="m21 21-4.3-4.3" />
					</svg>
				`),
				h.Input(
					h.Type("text"),
					h.Placeholder("Search..."),
					g.Attr("autocomplete", "off"),
					g.Attr("autocorrect", "off"),
					g.Attr("spellcheck", "false"),
					h.Aria("autocomplete", "list"),
					h.Role("combobox"),
					h.Aria("expanded", "false"),
					h.Aria("controls", props.ID+"-listbox"),
					h.Aria("labelledby", props.ID+"-trigger"),
					g.Group(props.SearchInputAttrs),
				),
			),

			SelectListbox(props.ID+"-listbox", props.ID+"-trigger",
				g.Map(props.Options, func(opt SelectOption) g.Node {
					return SelectOptionEl(
						props.ID+"-option-"+opt.Value,
						opt.Value,
						opt.Selected,
						g.Text(opt.Label),
					)
				}),
			),
		),
		SelectHiddenInput(props.Name, props.Value, props.Extras...),
	)
}

// SelectContainer creates the main container for the select component
func SelectContainer(id string, children ...g.Node) g.Node {
	return h.Div(h.ID(id), h.Class("select"), g.Group(children))
}

// SelectButton creates the button to toggle the select
func SelectButton(id, listboxId string, children ...g.Node) g.Node {
	return h.Button(
		h.Type("button"),
		h.ID(id),
		h.Aria("haspopup", listboxId),
		h.Aria("expanded", "false"),
		h.Aria("controls", listboxId),
		g.Group(children),
	)
}

// SelectPopover creates the popover container
func SelectPopover(id string, children ...g.Node) g.Node {
	return h.Div(
		h.ID(id),
		h.Data("popover", ""),
		h.Aria("hidden", "true"),
		g.Group(children),
	)
}

// SelectHeader creates an optional header section (for search)
func SelectHeader(children ...g.Node) g.Node {
	return h.Header(g.Group(children))
}

// SelectListbox creates the listbox container
func SelectListbox(id, labelledBy string, children ...g.Node) g.Node {
	return h.Div(
		h.ID(id),
		h.Role("listbox"),
		h.Aria("orientation", "vertical"),
		h.Aria("labelledby", labelledBy),
		g.Group(children),
	)
}

// SelectOption creates an option element
func SelectOptionEl(id, value string, selected bool, children ...g.Node) g.Node {
	return h.Div(
		h.ID(id),
		h.Role("option"),
		h.Data("value", value),
		g.If(selected, h.Aria("selected", "true")),
		g.Group(children),
	)
}

// SelectHiddenInput creates a hidden input to store the value
func SelectHiddenInput(name, value string, children ...g.Node) g.Node {
	return h.Input(
		h.Type("hidden"),
		h.Name(name),
		h.Value(value),
		g.Group(children),
	)
}

// SelectOption represents an option in the Select component
type SelectOption struct {
	Value    string
	Label    string
	Selected bool
}

type SelectOptions []SelectOption

func (so SelectOptions) SelectedLabel() string {
	for _, opt := range so {
		if opt.Selected {
			return opt.Label
		}
	}
	return ""
}

// Select component for basic dropdowns
func Select(opts []SelectOption, children ...g.Node) g.Node {
	return h.Select(
		h.Class("select"),
		g.Group(children),
		g.Map(opts, func(opt SelectOption) g.Node {

			return h.Option(
				h.Value(opt.Value),
				g.Text(opt.Label),
				g.If(opt.Selected, h.Selected()),
			)
		}),
	)
}

// SelectAssets includes the necessary JS for Select components
func SelectAssets() g.Node {
	return h.Script(h.Src("https://cdn.jsdelivr.net/npm/basecoat-css@0.3.10-beta.2/dist/js/select.min.js"), h.Defer())
}

// Combobox is an alias for SelectWithSearch - they use the same structure
// Combobox provides autocomplete input with a list of suggestions
func Combobox(props SelectWithSearchProps) g.Node {
	return SelectWithSearch(props)
}

// ComboboxProps is an alias for SelectWithSearchProps
type ComboboxProps = SelectWithSearchProps
