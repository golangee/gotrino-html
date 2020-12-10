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

package html

import (
	"strings"

	"github.com/golangee/dom"
	"github.com/golangee/wui"
)

func Div(e ...wui.Renderable) wui.Node {
	return wui.Element("div", e...)
}

func IFrame(e ...wui.Renderable) wui.Node {
	return wui.Element("iframe", e...)
}

func Hr(e ...wui.Renderable) wui.Node {
	return wui.Element("hr", e...)
}

func Button(e ...wui.Renderable) wui.Node {
	return wui.Element("button", e...)
}

func Nav(e ...wui.Renderable) wui.Node {
	return wui.Element("nav", e...)
}

// Class sets and replaces all existing class definitions.
func Class(classes ...string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		if len(classes) == 0 {
			return
		}

		if len(classes) == 1 {
			e.SetClassName(classes[0])

			return
		}

		e.SetClassName(strings.Join(classes, " "))
	})
}

// AddClass appends each given class.
func AddClass(classes ...string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		if len(classes) == 0 {
			return
		}

		for _, class := range classes {
			if strings.ContainsRune(class, ' ') {
				// separate it
				for _, s := range strings.Split(class, " ") {
					s = strings.TrimSpace(s)
					if s != "" {
						e.AddClass(s)
					}
				}
			} else {
				e.AddClass(class)
			}
		}
	})
}

// RemoveClass remove each given class.
func RemoveClass(classes ...string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		if len(classes) == 0 {
			return
		}

		for _, class := range classes {
			if strings.ContainsRune(class, ' ') {
				// separate it
				for _, s := range strings.Split(class, " ") {
					e.RemoveClass(s)
				}
			} else {
				e.RemoveClass(class)
			}
		}
	})
}

func SetAttribute(attr, value string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetAttribute(attr, value)
	})
}

func RemoveAttribute(attr string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.RemoveAttribute(attr)
	})
}

// Style sets a single CSS property.
func Style(property, value string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.Style().SetProperty(property, value)
	})
}

func Text(t string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.AppendTextNode(t)
	})
}

func InnerHTML(t string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetInnerHTML(t)
	})
}

func Src(src string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.Set("src", src)
	})
}

func TabIndex(t string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("tabindex", t)
	})
}

func I(mods ...wui.Renderable) wui.Node {
	return wui.Element("i", mods...)
}

func A(mods ...wui.Renderable) wui.Node {
	return wui.Element("a", mods...)
}

func Em(mods ...wui.Renderable) wui.Node {
	return wui.Element("em", mods...)
}

func Alt(a string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.Set("alt", a)
	})
}

func Title(a string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.Set("title", a)
	})
}

func Href(href string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.Set("href", href)
	})
}

func Width(w string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.Set("width", w)
	})
}

func Height(h string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.Set("height", h)
	})
}

func Figure(mods ...wui.Renderable) wui.Node {
	return wui.Element("figure", mods...)
}

func Ul(mods ...wui.Renderable) wui.Node {
	return wui.Element("ul", mods...)
}

func Li(mods ...wui.Renderable) wui.Node {
	return wui.Element("li", mods...)
}

func Ol(mods ...wui.Renderable) wui.Node {
	return wui.Element("ol", mods...)
}

func Img(mods ...wui.Renderable) wui.Node {
	return wui.Element("img", mods...)
}

func P(mods ...wui.Renderable) wui.Node {
	return wui.Element("p", mods...)
}

func Pre(mods ...wui.Renderable) wui.Node {
	return wui.Element("pre", mods...)
}

func Code(mods ...wui.Renderable) wui.Node {
	return wui.Element("code", mods...)
}

func Aside(mods ...wui.Renderable) wui.Node {
	return wui.Element("aside", mods...)
}

func Blockquote(mods ...wui.Renderable) wui.Node {
	return wui.Element("blockquote", mods...)
}

func Figcaption(mods ...wui.Renderable) wui.Node {
	return wui.Element("figcaption", mods...)
}

func Span(mods ...wui.Renderable) wui.Node {
	return wui.Element("span", mods...)
}

func ID(id string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetID(id)
	})
}

func AddEventListener(eventType string, f func()) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, false, f)
	})
}

func AddClickListener(f func()) wui.Modifier {
	return AddEventListener("click", f)
}

func AddKeyDownListener(f func(keyCode int)) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.AddKeyListener("keyup", f)
	})
}

func AddEventListenerOnce(eventType string, f func()) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, true, f)
	})
}

func AriaLabel(label string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-label", label)
	})
}

func AriaOrientation(orientation string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-orientation", orientation)
	})
}

func AriaLabelledby(label string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-labelledby", label)
	})
}

func Role(role string) wui.Modifier {
	return wui.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("role", role)
	})
}
