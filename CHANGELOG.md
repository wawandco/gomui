# Changelog

All notable changes to the GomUI project are documented in this file.

## [Unreleased] - Basecoat UI v0.3.10-beta.2 Validation

### Breaking Changes

This release includes a comprehensive validation of all components against Basecoat UI v0.3.10-beta.2 documentation. Components that do not exist in Basecoat or use incorrect CSS classes have been removed or fixed.

#### Removed Components (Not in Basecoat UI)

The following components have been **completely removed** because Basecoat UI does not provide dedicated CSS classes or patterns for them:

**Navigation Components:**
- `Breadcrumb`, `BreadcrumbList`, `BreadcrumbItem`, `BreadcrumbLink`, `BreadcrumbSeparator` - Basecoat states "There is NO dedicated Breadcrumb component"
- `Pagination`, `PaginationContent`, `PaginationItem`, `PaginationLink`, `PaginationPrevious`, `PaginationNext` - Basecoat states "There is NO dedicated pagination component"
- `ContextMenu` - Not provided by Basecoat
- `Menubar` - Not provided by Basecoat
- `NavigationMenu` - Not provided by Basecoat

**Layout Components:**
- `Separator` - No dedicated separator component
- `Avatar`, `AvatarImage`, `AvatarFallback` - Basecoat states "There is no dedicated Avatar component"
- `ScrollArea` - Not provided by Basecoat
- `AspectRatio` - Not provided by Basecoat
- `Calendar` - Not provided by Basecoat
- `Carousel` and all sub-components (`CarouselContent`, `CarouselItem`, `CarouselPrevious`, `CarouselNext`, `CarouselIndicators`) - Not provided by Basecoat

**Feedback Components:**
- `Progress`, `ProgressEl` - Basecoat states "There is no dedicated Progress component"
- `Tooltip`, `TooltipTrigger`, `TooltipContent` - Use `data-tooltip` attribute directly on elements instead
- `Skeleton` - Basecoat states "There is no dedicated skeleton component"

**Form Components:**
- `RadioGroup` - Use standard `<fieldset>` or `Fieldset` component instead

**Migration Guide:**
- For **Breadcrumbs**: Use Tailwind utility classes with `Button` components. See [Basecoat Breadcrumb docs](https://basecoatui.com/components/breadcrumb)
- For **Pagination**: Use Tailwind utility classes with `Button` components. See [Basecoat Pagination docs](https://basecoatui.com/components/pagination)
- For **Tooltips**: Use `data-tooltip` attribute directly: `g.Attr("data-tooltip", "Your tooltip text")`
- For **Progress**: Build custom progress bars using Tailwind utilities
- For **RadioGroup**: Use `<fieldset>` elements or the `Fieldset` component

#### Component Renamings

**Forms:**
- `FormField` → `Field`
- `FormLabel` → `FieldLabel` (note: `Label` still exists as an alias)
- `FormDescription` → `FieldDescription`
- `FormMessage` → `FieldError`

#### Badge Variant Changes

**Removed invalid Badge variants:**
- `BadgeSuccess` - Not provided by Basecoat
- `BadgeWarning` - Not provided by Basecoat

**Valid Badge variants remain:**
- `BadgePrimary`
- `BadgeSecondary`
- `BadgeDestructive`
- `BadgeOutline`

#### Alert Variant Changes

**Removed Alert variants:**
- `AlertSuccess` - Not provided by Basecoat
- `AlertWarning` - Not provided by Basecoat
- `AlertInfo` - Not provided by Basecoat

**Valid Alert variants:**
- `AlertDefault` (just `"alert"`)
- `AlertDestructive` (`"alert-destructive"`)

#### Fixed Components

**Overlays:**
- **Dialog, AlertDialog, Sheet**: Now use semantic HTML (`<header>`, `<section>`, `<footer>`) without custom CSS classes
  - Added: `DialogSection`, `AlertDialogSection`, `SheetSection`
- **Popover, HoverCard**: Fixed to use `data-popover` attribute instead of CSS classes
  - Added: `PopoverHeader`, `HoverCardHeader`

**Navigation:**
- **Tabs**: 
  - Added `aria-orientation="horizontal"` to `TabsList`
  - Removed `TabsTrigger` and `TabsContent`
  - Added proper `TabPanel` with ARIA attributes
- **Command**: Complete restructure from CSS classes to ARIA roles
  - Added: `CommandDialog`, `CommandHeader`, `CommandInput`, `CommandMenu`, `CommandGroup`, `CommandGroupHeading`
  - Removed: `CommandList`, `CommandEmpty`
  - Fixed: `CommandSeparator` now uses `<hr role="separator">`
- **DropdownMenu**: Fixed to use `data-popover` and ARIA roles

**Forms:**
- **Checkbox**: Changed from `class="checkbox"` to `class="input"`
- **Radio**: Changed from `class="radio"` to `class="input"`
- Removed `SwitchLabel` (use standard `Label` instead)

**Feedback:**
- **Alert**: Removed Success/Warning/Info variants (only `alert` and `alert-destructive` exist)
- **AlertTitle**: Now plain `<h2>` without class
- **AlertDescription**: Now plain `<section>` without class
- **Toast**: Added required ARIA attributes, `ToastContent`, `ToastSection`, `ToastFooter`
- **ToastTitle**: Now plain `<h2>`
- **ToastDescription**: Now plain `<p>`
- **ToastAction**: Changed to `class="btn"`

**Layout:**
- **CardTitle**: Now plain `<h2>` without class
- **CardDescription**: Now plain `<p>` without class

**Table:**
- Removed ALL custom CSS classes from table sub-elements (Basecoat doesn't use them):
  - `table-header` → plain `<thead>`
  - `table-body` → plain `<tbody>`
  - `table-footer` → plain `<tfoot>`
  - `table-row` → plain `<tr>`
  - `table-head` → plain `<th>`
  - `table-cell` → plain `<td>`
  - `table-caption` → plain `<caption>`
- Only the main `<table class="table">` has a CSS class

**Theme:**
- **DarkModeScript()**: Completely rewritten to match Basecoat's approach
  - Now uses `basecoat:theme` event listener
  - Proper localStorage handling with fallback to system preference (`prefers-color-scheme`)
  - Uses `themeMode` localStorage key
- **ThemeToggle()**: 
  - Removed `id` parameter (not needed)
  - Now dispatches `basecoat:theme` event
  - Changed to use `h.Aria("label", ...)` instead of inline JS

### Updated

- **CDN URLs**: Updated from v0.3.9 to v0.3.10-beta.2
  - `BasecoatCSS()`
  - `BasecoatJS()`
  - `SelectAssets()`
  - `SidebarAssets()`

- **README.md**: Completely updated
  - Fixed component count claims
  - Updated all code examples to reflect actual API
  - Added "Components NOT included" section with explanations
  - Fixed Basecoat URL from `www.basecoat.dev` to `basecoatui.com`
  - Updated credits section with version number

- **Package Documentation** (gomui.go):
  - Removed inaccurate "40+" component claim
  - Added validation notice against Basecoat UI v0.3.10-beta.2
  - Updated feature list
  - Added comprehensive component categories list
  - Added note about excluded components

### Added

- **File-level documentation comments** explaining:
  - Why certain components were removed (forms.go, feedback.go)
  - Component naming changes (forms.go)
  - Basecoat validation status

- **AGENTS.md**: Development guide for AI agents with:
  - Build, test, and lint commands
  - Code style guidelines
  - Basecoat integration principles
  - Testing guidelines
  - Common patterns

### Compatibility Notes

This is a **major breaking change release**. If you are upgrading from a previous version:

1. **Review removed components** - Check if your code uses any of the removed components listed above
2. **Update import names** - Change `FormField` → `Field`, `FormLabel` → `FieldLabel`, etc.
3. **Update Badge/Alert variants** - Remove usage of Success/Warning variants
4. **Update ThemeToggle calls** - Remove the `id` parameter
5. **Update CDN references** - If you're using custom CDN links, update to v0.3.10-beta.2
6. **Test thoroughly** - Many components have structural changes (semantic HTML, ARIA attributes)

### Verification

All components have been:
- ✅ Validated against official Basecoat UI v0.3.10-beta.2 documentation
- ✅ Tested for compilation
- ✅ Formatted with `go fmt`
- ✅ Documented with updated examples

---

## Previous Versions

No changelog was maintained for versions prior to the Basecoat UI v0.3.10-beta.2 validation.
