package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DarkModeScript initializes theme detection and applies saved theme preference
// This script should be placed in the <head> of your page to avoid flash of unstyled content
func DarkModeScript() g.Node {
	return h.Script(g.Raw(`
		(() => {
			try {
				const stored = localStorage.getItem('themeMode');
				if (stored ? stored === 'dark'
						  : matchMedia('(prefers-color-scheme: dark)').matches) {
					document.documentElement.classList.add('dark');
				}
			} catch (_) {}

			const apply = dark => {
				document.documentElement.classList.toggle('dark', dark);
				try { localStorage.setItem('themeMode', dark ? 'dark' : 'light'); } catch (_) {}
			};

			document.addEventListener('basecoat:theme', (event) => {
				const mode = event.detail?.mode;
				apply(mode === 'dark' ? true
					: mode === 'light' ? false
					: !document.documentElement.classList.contains('dark'));
			});
		})();
	`))
}

// ThemeToggle creates a theme switcher button for toggling between light/dark modes
func ThemeToggle(children ...g.Node) g.Node {
	return h.Button(
		h.Type("button"),
		h.Class("btn-icon-outline size-8"),
		h.Aria("label", "Toggle dark mode"),
		g.Attr("onclick", `document.dispatchEvent(new CustomEvent('basecoat:theme'))`),
		g.Group(children),
	)
}

// BasecoatCSS includes Basecoat CSS from CDN
func BasecoatCSS() g.Node {
	return h.Link(
		h.Rel("stylesheet"),
		h.Href("https://cdn.jsdelivr.net/npm/basecoat-css@0.3.10-beta.2/dist/basecoat.cdn.min.css"),
	)
}

// BasecoatJS includes Basecoat JavaScript from CDN
func BasecoatJS() g.Node {
	return h.Script(
		h.Src("https://cdn.jsdelivr.net/npm/basecoat-css@0.3.10-beta.2/dist/js/all.min.js"),
		g.Attr("defer"),
	)
}

// BasecoatAssets includes both Basecoat CSS and JavaScript for convenience
func BasecoatAssets() g.Node {
	return g.Group([]g.Node{
		BasecoatCSS(),
		BasecoatJS(),
	})
}
