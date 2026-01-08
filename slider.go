package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Slider component for range selection
// Slider creates a range input slider
func Slider(attrs ...g.Node) g.Node {
	return h.Input(
		g.Group(attrs),
		h.Type("range"),
		h.Class("input"),
		h.Style("width: 100%;"),
	)
}

// SliderScript includes the necessary JS to enhance slider functionality
func SliderScript() g.Node {
	return h.Script(g.Raw(`
		(() => {
		const sliders = document.querySelectorAll('input[type="range"].input');
		if (!sliders) return;

		const updateSlider = (el) => {
			const min = parseFloat(el.min || 0);
			const max = parseFloat(el.max || 100);
			const value = parseFloat(el.value);
			const percent = (max === min) ? 0 : ((value - min) / (max - min)) * 100;
			el.style.setProperty('--slider-value', percent + '%');
		};

		sliders.forEach(slider => {
			updateSlider(slider);
			slider.addEventListener('input', (event) => updateSlider(event.target));
		});
	})();
	`))
}
