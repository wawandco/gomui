// Package gomui provides a comprehensive Go UI components library built on top of gomponents.
// It offers validated Basecoat UI components for creating beautiful, type-safe user interfaces
// in pure Go without any templating languages.
//
// All components are validated against Basecoat UI v0.3.10-beta.2 (https://basecoatui.com)
//
// Key features:
//   - Pure Go components rendering to semantic HTML
//   - Type-safe component API with compile-time checking
//   - Built-in dark mode with system preference detection and localStorage persistence
//   - Only includes components that actually exist in Basecoat UI
//   - Framework agnostic - works with any Go web framework
//
// Usage:
//
//	import (
//		g "maragu.dev/gomponents"
//		h "maragu.dev/gomponents/html"
//		gm "github.com/wawandco/gomui"
//	)
//
//	page := gm.Card(
//		gm.CardHeader(gm.CardTitle(g.Text("Hello World"))),
//		gm.CardContent(g.Text("Welcome to GomUI!")),
//	)
//
// Component Categories:
//   - Buttons: Button, ButtonGroup with multiple variants
//   - Forms: Input, Textarea, Form, Checkbox, Radio, Label, Field components
//   - Cards: Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter
//   - Overlays: Dialog, AlertDialog, Sheet, Popover, HoverCard with semantic HTML
//   - Feedback: Alert, Toast with proper ARIA attributes
//   - Navigation: Tabs, Command palette, DropdownMenu
//   - Data Display: Table components, Badge, Kbd
//   - Interactive: Select, Slider, Sidebar
//   - Theme: Dark mode utilities with system preference detection
//
// Note: This library only includes components officially supported by Basecoat UI.
// Components like Breadcrumb, Pagination, Avatar, Progress, Skeleton, and Tooltip
// are NOT included as Basecoat does not provide dedicated CSS classes for them.
// See README.md for alternative approaches using utility classes.
package gomui
