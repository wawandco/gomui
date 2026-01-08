package gomui

// Form components validated against Basecoat UI v0.3.10-beta.2
//
// Note: This package uses Field, FieldLabel, FieldDescription, and FieldError
// components following Basecoat's official Field pattern. Previous versions may
// have used FormField, FormLabel, FormDescription, and FormMessage - these have
// been renamed to align with Basecoat's naming conventions.
//
// RadioGroup component has been removed - use standard <fieldset> elements or
// the Fieldset component instead, as Basecoat does not provide a RadioGroup class.

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// InputEl creates a styled text input field
func InputEl(attrs ...g.Node) g.Node {
	return h.Input(h.Class("input"), g.Group(attrs))
}

// Input creates a styled text input field (convenience alias for InputEl)
func Input(attrs ...g.Node) g.Node {
	return InputEl(attrs...)
}

// InputWithClasses creates a styled text input field with additional classes
func InputWithClasses(classes string, attrs ...g.Node) g.Node {
	return h.Input(h.Class("input "+classes), g.Group(attrs))
}

// TextareaEl creates a multi-line text input field
func TextareaEl(attrs ...g.Node) g.Node {
	return h.Textarea(h.Class("textarea"), g.Group(attrs))
}

// Textarea creates a multi-line text input field (convenience alias for TextareaEl)
func Textarea(attrs ...g.Node) g.Node {
	return TextareaEl(attrs...)
}

// FormComponent creates a form wrapper element
func FormComponent(attrs ...g.Node) g.Node {
	return h.Form(h.Class("form"), g.Group(attrs))
}

// Form creates a form wrapper element (convenience alias for FormComponent)
func Form(children ...g.Node) g.Node {
	return FormComponent(children...)
}

// Checkbox component for toggle selection
// Checkbox creates a checkbox input element
func Checkbox(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		h.Type("checkbox"),
		h.Class("input"),
	}
	allAttrs = append(allAttrs, attrs...)
	return h.Input(allAttrs...)
}

// Radio component for single selection
// Radio creates a radio input element
func Radio(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		h.Type("radio"),
		h.Class("input"),
	}
	allAttrs = append(allAttrs, attrs...)
	return h.Input(allAttrs...)
}

// Switch component for binary toggles
// Switch creates a toggle switch input
func Switch(checked bool, attrs ...g.Node) g.Node {
	return h.Input(
		g.If(checked, h.Checked()),
		g.Group(attrs),
		h.Type("checkbox"),
		h.Class("input"),
		h.Role("switch"),
	)
}

// InputLabel creates a label for form inputs
func InputLabel(forID string, children ...g.Node) g.Node {
	return h.Label(
		h.Class("label"),
		h.For(forID),
		g.Group(children),
	)
}

// Label creates a label for form inputs (convenience alias for InputLabel)
func Label(forID string, children ...g.Node) g.Node {
	return InputLabel(forID, children...)
}

// Field creates a container for individual form fields (label, input, helper text)
// This is the Basecoat Field pattern - uses role="group" and class="field"
func Field(children ...g.Node) g.Node {
	return h.Div(h.Role("group"), h.Class("field"), g.Group(children))
}

// FieldWithOrientation creates a field with specific orientation
// Use "horizontal" to align controls horizontally
func FieldWithOrientation(orientation string, children ...g.Node) g.Node {
	return h.Div(
		h.Role("group"),
		h.Class("field"),
		g.Attr("data-orientation", orientation),
		g.Group(children),
	)
}

// FieldLabel creates a label for a field (uses standard label styling)
func FieldLabel(forID string, children ...g.Node) g.Node {
	return h.Label(h.For(forID), g.Group(children))
}

// FieldDescription creates descriptive helper text for fields
// Typically paired with aria-describedby on the input
func FieldDescription(id string, children ...g.Node) g.Node {
	return h.P(h.ID(id), g.Group(children))
}

// FieldError creates an error message for fields
// Should be paired with aria-describedby and aria-invalid="true" on the input
func FieldError(id string, children ...g.Node) g.Node {
	return h.P(h.ID(id), h.Role("alert"), g.Group(children))
}

// FieldSection creates a section wrapper for label and description
// Used when label sits beside input (e.g. with data-orientation="horizontal")
func FieldSection(children ...g.Node) g.Node {
	return h.Section(g.Group(children))
}

// Fieldset creates a container for grouping related fields
func Fieldset(children ...g.Node) g.Node {
	return g.El("fieldset", h.Class("fieldset"), g.Group(children))
}

// FieldsetLegend creates a heading for a fieldset
func FieldsetLegend(children ...g.Node) g.Node {
	return g.El("legend", g.Group(children))
}
