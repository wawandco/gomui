# Basecoat UI for Gomponents

A comprehensive Go wrapper library for [Basecoat UI](https://basecoatui.com/) components, built on top of [gomponents](https://maragu.dev/gomponents). Write beautiful, type-safe UI components in pure Go with no templating languages required.

## Features

✨ **Pure Go** - No templating languages, just Go functions  
🎨 **Component Library** - Pre-built Basecoat UI components  
🌓 **Dark Mode Support** - Built-in dark mode with system preference detection and localStorage persistence  
🔒 **Type-Safe** - Compile-time checking for your UI components  
⚡ **Minimal JavaScript** - Renders to pure HTML with minimal JavaScript  
🎯 **Framework Agnostic** - Works with any Go web framework  

## Installation

```bash
go get maragu.dev/gomponents
go get github.com/wawandco/gomui
```

Then include Basecoat UI CSS and JavaScript using our helper components in your HTML:

```html
<!-- Option 1: Include both CSS and JS (recommended) -->
<!-- In your Go code: gm.BasecoatAssets() -->

<!-- Option 2: Include CSS and JS separately -->
<!-- In your Go code: gm.BasecoatCSS() and gm.BasecoatJS() -->

<!-- Or manually include Basecoat assets -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.10-beta.2/dist/basecoat.cdn.min.css">
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.10-beta.2/dist/js/all.min.js" defer></script>
```

## Quick Start

```go
package main

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
    gm "github.com/wawandco/gomui"
)

func MyPage() g.Node {
    return Html(
        Head(
            TitleEl(g.Text("My App")),
            gm.DarkModeScript(), // Enable dark mode (place first)
            gm.BasecoatAssets(), // Include Basecoat CSS and JS
        ),
        Body(
            Div(Class("container mx-auto p-4"),
                gm.Card(
                    gm.CardHeader(
                        gm.CardTitle(g.Text("Welcome!")),
                        gm.CardDescription(g.Text("Get started with Basecoat UI")),
                    ),
                    gm.CardContent(
                        gm.Button(gm.ButtonPrimary, gm.ButtonDefault,
                            g.Text("Click me"),
                        ),
                    ),
                ),
            ),
        ),
    )
}
```

## Dark Mode

### Setup

Add the dark mode script in your `<head>` to initialize the theme:

```go
Head(
    gm.DarkModeScript(),
    // ... other head elements
)
```

### Theme Toggle Button

Add a theme toggle button anywhere in your UI:

```go
gm.ThemeToggle(
    // Add your SVG icons for light/dark mode here
    g.Raw(`<span class="hidden dark:block">☀️</span><span class="block dark:hidden">🌙</span>`),
)
```

The theme preference is automatically saved to `localStorage` and persists across sessions. The system also detects the user's OS preference on first load.

## Component Reference

### Buttons

```go
// Primary button
gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Primary"))

// Secondary button
gm.Button(gm.ButtonSecondary, gm.ButtonDefault, g.Text("Secondary"))

// Outline button
gm.Button(gm.ButtonOutline, gm.ButtonDefault, g.Text("Outline"))

// Ghost button
gm.Button(gm.ButtonGhost, gm.ButtonDefault, g.Text("Ghost"))

// Destructive button
gm.Button(gm.ButtonDestructive, gm.ButtonDefault, g.Text("Delete"))

// Different sizes
gm.Button(gm.ButtonPrimary, gm.ButtonSm, g.Text("Small"))
gm.Button(gm.ButtonPrimary, gm.ButtonLg, g.Text("Large"))
gm.Button(gm.ButtonPrimary, gm.ButtonIcon, g.Text("🔍"))
```

#### Link Buttons

Use `LinkButtonEl` when you need navigation that appears as a button:

```go
// Primary link button
gm.LinkButtonEl(gm.ButtonPrimary, gm.ButtonDefault, "/home", g.Text("Go Home"))

// Small secondary link button  
gm.LinkButtonEl(gm.ButtonSecondary, gm.ButtonSm, "/settings", g.Text("Settings"))

// Large outline link button
gm.LinkButtonEl(gm.ButtonOutline, gm.ButtonLg, "/docs", g.Text("Documentation"))
```

### Forms

```go
gm.Form(
    gm.Field(
        gm.Label(g.Attr("for", "email"), g.Text("Email")),
        gm.Input(h.Type("email"), h.ID("email"), h.Placeholder("you@example.com")),
        gm.FieldDescription(g.Text("We'll never share your email.")),
    ),
    gm.Field(
        gm.Label(g.Attr("for", "message"), g.Text("Message")),
        gm.Textarea(h.ID("message"), h.Placeholder("Your message...")),
    ),
    gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Submit")),
)
```

### Cards

```go
gm.Card(
    gm.CardHeader(
        gm.CardTitle(g.Text("Card Title")),
        gm.CardDescription(g.Text("Card description goes here")),
    ),
    gm.CardContent(
        g.Text("Your content here"),
    ),
    gm.CardFooter(
        gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Action")),
    ),
)
```

### Alerts

```go
// Default alert
gm.Alert(
    gm.AlertTitle(g.Text("Heads up!")),
    gm.AlertDescription(g.Text("This is an alert message.")),
)

// Destructive alert
gm.AlertDestructive(
    gm.AlertTitle(g.Text("Error")),
    gm.AlertDescription(g.Text("Something went wrong.")),
)
```

### Badges

```go
gm.Badge(gm.BadgePrimary, g.Text("Primary"))
gm.Badge(gm.BadgeSecondary, g.Text("Secondary"))
gm.Badge(gm.BadgeDestructive, g.Text("Destructive"))
gm.Badge(gm.BadgeOutline, g.Text("Outline"))
```

### Tables

```go
gm.Table(
    gm.TableHeader(
        gm.TableRow(
            gm.TableHead(g.Text("Name")),
            gm.TableHead(g.Text("Email")),
            gm.TableHead(g.Text("Role")),
        ),
    ),
    gm.TableBody(
        gm.TableRow(
            gm.TableCell(g.Text("John Doe")),
            gm.TableCell(g.Text("john@example.com")),
            gm.TableCell(g.Text("Admin")),
        ),
    ),
)
```

### Dialogs/Modals

```go
gm.Dialog(
    gm.DialogHeader(
        gm.DialogTitle(g.Text("Confirm Action")),
        gm.DialogDescription(g.Text("Are you sure?")),
    ),
    gm.DialogSection(
        // Dialog content here
    ),
    gm.DialogFooter(
        gm.Button(gm.ButtonOutline, gm.ButtonDefault, g.Text("Cancel")),
        gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Confirm")),
    ),
)
```

### Tabs

```go
gm.Tabs(
    gm.TabsList(
        gm.TabItem("account-tab", true, g.Text("Account")),
        gm.TabItem("password-tab", false, g.Text("Password")),
    ),
    gm.TabPanel("account-panel", "account-tab", true,
        gm.Card(
            gm.CardHeader(
                gm.CardTitle(g.Text("Account Settings")),
            ),
            gm.CardContent(
                g.Text("Account content here"),
            ),
        ),
    ),
    gm.TabPanel("password-panel", "password-tab", false,
        gm.Card(
            gm.CardHeader(
                gm.CardTitle(g.Text("Password Settings")),
            ),
            gm.CardContent(
                g.Text("Password content here"),
            ),
        ),
    ),
)
```

### Accordion

```go
gm.Accordion(
    gm.AccordionItem(
        gm.AccordionTrigger(g.Text("What is Basecoat UI?")),
        gm.AccordionContent(
            g.Text("Basecoat UI is a component library..."),
        ),
    ),
    gm.AccordionItem(
        gm.AccordionTrigger(g.Text("How do I use it?")),
        gm.AccordionContent(
            g.Text("Simply import the components..."),
        ),
    ),
)
```

### Dropdown Menu

```go
gm.DropdownMenu(
    gm.DropdownMenuTrigger(g.Text("Options")),
    gm.DropdownMenuContent(
        gm.DropdownMenuMenu(
            gm.DropdownMenuGroup(
                gm.DropdownMenuGroupHeading(g.Text("My Account")),
                gm.DropdownMenuItem(g.Text("Profile")),
                gm.DropdownMenuItem(g.Text("Settings")),
            ),
            gm.DropdownMenuSeparator(),
            gm.DropdownMenuItem(g.Text("Logout")),
        ),
    ),
)
```

## Complete Component List

All components are validated against [Basecoat UI v0.3.10-beta.2](https://basecoatui.com/) documentation.

### Layout & Structure
- `Card`, `CardWithClasses`, `CardHeader`, `CardTitle`, `CardDescription`, `CardContent`, `CardFooter`

### Navigation & Tabs
- `Tabs`, `TabsList`, `TabItem`, `TabPanel`
- `DropdownMenu`, `DropdownMenuTrigger`, `DropdownMenuContent`, `DropdownMenuMenu`, `DropdownMenuItem`, `DropdownMenuItemCheckbox`, `DropdownMenuItemRadio`, `DropdownMenuGroup`, `DropdownMenuGroupHeading`, `DropdownMenuSeparator`

### Forms & Inputs
- `Form`, `Field`, `FieldLabel`, `FieldDescription`, `FieldError`, `Fieldset`
- `Input`, `Textarea`
- `Label`
- `Checkbox`, `Radio`, `Switch`
- `Slider`, `SliderScript()`
- `Select` components (SelectContainer, SelectButton, SelectPopover, SelectHeader, SelectListbox, SelectOptionEl, SelectHiddenInput, SelectWithSearch, Combobox)

### Buttons & Actions
- `Button`, `ButtonEl`, `ButtonWithClasses` (with variants: Primary, Secondary, Outline, Ghost, Link, Destructive)
- `LinkButton`, `LinkButtonEl`, `LinkButtonWithClasses` - Anchor elements with button styling
- `ButtonGroup`, `ButtonGroupWithOrientation`, `ButtonGroupSeparator`

### Data Display
- `Table`, `TableEl`, `TableHeader`, `TableBody`, `TableFooter`, `TableRow`, `TableHead`, `TableData`, `TableCell`, `TableCaption`
- `Badge` (with variants: Primary, Secondary, Destructive, Outline)
- `Kbd`, `KbdWithClasses`, `KbdEl`

### Feedback
- `Alert`, `AlertDestructive`, `AlertTitle`, `AlertDescription`
- `Toast`, `ToastContent`, `ToastSection`, `ToastFooter`, `ToastTitle`, `ToastDescription`, `ToastAction`, `Toaster`

### Overlays
- `Dialog`, `DialogHeader`, `DialogTitle`, `DialogDescription`, `DialogSection`, `DialogFooter`
- `AlertDialog`, `AlertDialogHeader`, `AlertDialogTitle`, `AlertDialogDescription`, `AlertDialogSection`, `AlertDialogFooter`
- `Sheet`, `SheetHeader`, `SheetTitle`, `SheetDescription`, `SheetSection`, `SheetFooter`
- `Popover`, `PopoverContent`, `PopoverHeader`
- `HoverCard`, `HoverCardContent`, `HoverCardHeader`

### Command Palette
- `Command`, `CommandDialog`, `CommandHeader`, `CommandInput`, `CommandMenu`, `CommandGroup`, `CommandGroupHeading`, `CommandItem`, `CommandSeparator`

### Sidebar
- `Sidebar`, `SidebarHeader`, `SidebarNav`, `SidebarSection`, `SidebarGroup`, `SidebarHeading`, `SidebarList`, `SidebarItem`, `SidebarFooter`, `SidebarToggle`, `SidebarOpen`, `SidebarClose`, `SidebarAssets()`

### Theme
- `DarkModeScript()` - Initialize dark mode with system preference detection
- `ThemeToggle()` - Theme toggle button (dispatches `basecoat:theme` event)
- `BasecoatCSS()` - Include Basecoat CSS from CDN (v0.3.10-beta.2)
- `BasecoatJS()` - Include Basecoat JavaScript from CDN (v0.3.10-beta.2)
- `BasecoatAssets()` - Include both CSS and JavaScript

### Notes

**Components NOT included** (not in Basecoat or use Tailwind utilities instead):
- Breadcrumb, Pagination - Use Tailwind utility classes with button components
- Avatar, Skeleton, Progress - Use Tailwind utilities
- Separator - Use `<hr role="separator">` directly
- Tooltip - Use `data-tooltip` attribute on elements
- NavigationMenu, ContextMenu, Menubar - Not in Basecoat
- ScrollArea, AspectRatio, Calendar, Carousel - Not in Basecoat

## Framework Integration

### Standard Library (net/http)

```go
package main

import (
    "net/http"
    g "maragu.dev/gomponents"
    gm "github.com/wawandco/gomui"
)

func handler(w http.ResponseWriter, r *http.Request) {
    page := MyPage()
    page.Render(w)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Chi Router

```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    g "maragu.dev/gomponents"
    gm "github.com/wawandco/gomui"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        MyPage().Render(w)
    })
    http.ListenAndServe(":8080", r)
}
```

### Fiber

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    gm "github.com/wawandco/gomui"
)

func main() {
    app := fiber.New()
    
    app.Get("/", func(c *fiber.Ctx) error {
        c.Set("Content-Type", "text/html")
        return MyPage().Render(c)
    })
    
    app.Listen(":8080")
}
```

### Echo

```go
package main

import (
    "github.com/labstack/echo/v4"
    gm "github.com/wawandco/gomui"
)

func main() {
    e := echo.New()
    
    e.GET("/", func(c echo.Context) error {
        c.Response().Header().Set("Content-Type", "text/html")
        return MyPage().Render(c.Response())
    })
    
    e.Start(":8080")
}
```

## Best Practices

### 1. Organize Components

Create reusable component functions:

```go
func UserCard(name, email string) g.Node {
    return gm.Card(
        gm.CardHeader(
            gm.CardTitle(g.Text(name)),
            gm.CardDescription(g.Text(email)),
        ),
        gm.CardFooter(
            gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("View Profile")),
        ),
    )
}
```

### 2. Create Layout Components

```go
func PageLayout(title string, content g.Node) g.Node {
    return Html(
        Head(
            TitleEl(g.Text(title)),
            gm.DarkModeScript(), // Must be first
            gm.BasecoatAssets(),
        ),
        Body(
            Div(Class("container mx-auto p-4"),
                gm.ThemeToggle(
                    g.Raw(`<span class="hidden dark:block">☀️</span><span class="block dark:hidden">🌙</span>`),
                ),
                content,
            ),
        ),
    )
}
```

### 3. Use Conditional Rendering

```go
import "maragu.dev/gomponents"

func ConditionalAlert(hasError bool, message string) g.Node {
    return gomponents.If(hasError,
        gm.Alert(gm.AlertDestructive,
            gm.AlertTitle(g.Text("Error")),
            gm.AlertDescription(g.Text(message)),
        ),
    )
}
```

### 4. Map Over Data

```go
func UserList(users []User) g.Node {
    return Div(
        gomponents.Map(users, func(u User) g.Node {
            return UserCard(u.Name, u.Email)
        }),
    )
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Credits

- Built on top of [gomponents](https://maragu.dev/gomponents) by [@maragudk](https://github.com/maragudk)
- Uses [Basecoat UI](https://basecoatui.com/) components and styling (v0.3.10-beta.2)
- Inspired by [shadcn/ui](https://ui.shadcn.com/)

## Resources

- [Gomponents Documentation](https://www.gomponents.com/)
- [Basecoat UI Documentation](https://basecoatui.com/)
- [Go Documentation](https://golang.org/doc/)

---

Made with ❤️ for the Go community