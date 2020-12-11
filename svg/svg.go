// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package svg

import (
	"github.com/golangee/dom"
	. "github.com/golangee/gotrino"
)

func Svg(e ...Renderable) Node {
	return ElementNS("http://www.w3.org/2000/svg", "svg", e...)
}

func Fill(fill string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("fill", fill)
	})
}

func Stroke(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke", stroke)
	})
}

func StrokeLinecap(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke-linecap", stroke)
	})
}

func Cx(cx string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("cx", cx)
	})
}

func Cy(cy string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("cy", cy)
	})
}

func R(r string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("r", r)
	})
}

func StrokeLinejoin(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke-linejoin", stroke)
	})
}

func StrokeWidth(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke-width", stroke)
	})
}

func StrokeMiterlimit(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke-miterlimit", stroke)
	})
}

func ViewBox(viewBox string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("viewBox", viewBox)
	})
}

func Path(e ...Renderable) Node {
	return ElementNS("http://www.w3.org/2000/svg", "path", e...)
}

func Circle(e ...Renderable) Node {
	return ElementNS("http://www.w3.org/2000/svg", "circle", e...)
}

func D(d string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("d", d)
	})
}
