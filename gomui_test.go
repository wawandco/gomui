package gomui_test

import (
	"testing"

	gm "github.com/wawandco/gomui"
	g "maragu.dev/gomponents"
)

func Test(t *testing.T) {
	// Test that components are accessible from main package
	page := gm.Card(
		gm.CardHeader(
			gm.CardTitle(g.Text("Card Title")),
			gm.CardDescription(g.Text("Card description")),
		),

		gm.CardContent(
			gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Click me")),
			gm.Badge(gm.BadgeSecondary, g.Text("Badge")),
		),
	)

	_ = page // Just verify compilation
}
