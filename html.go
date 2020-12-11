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
	gt "github.com/golangee/gotrino"
)

func Div(e ...gt.Renderable) gt.Node {
	return gt.Element("div", e...)
}

func IFrame(e ...gt.Renderable) gt.Node {
	return gt.Element("iframe", e...)
}

func Hr(e ...gt.Renderable) gt.Node {
	return gt.Element("hr", e...)
}

func Button(e ...gt.Renderable) gt.Node {
	return gt.Element("button", e...)
}

func Nav(e ...gt.Renderable) gt.Node {
	return gt.Element("nav", e...)
}

// Class sets and replaces all existing class definitions.
func Class(classes ...string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
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
func AddClass(classes ...string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
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
func RemoveClass(classes ...string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
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

func SetAttribute(attr, value string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetAttribute(attr, value)
	})
}

func RemoveAttribute(attr string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.RemoveAttribute(attr)
	})
}

// Style sets a single CSS property.
func Style(property, value string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.Style().SetProperty(property, value)
	})
}

func Text(t string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.AppendTextNode(t)
	})
}

func InnerHTML(t string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetInnerHTML(t)
	})
}

func Src(src string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.Set("src", src)
	})
}

func TabIndex(t string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("tabindex", t)
	})
}

func I(mods ...gt.Renderable) gt.Node {
	return gt.Element("i", mods...)
}

func A(mods ...gt.Renderable) gt.Node {
	return gt.Element("a", mods...)
}

func Em(mods ...gt.Renderable) gt.Node {
	return gt.Element("em", mods...)
}

func Alt(a string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.Set("alt", a)
	})
}

func Title(a string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.Set("title", a)
	})
}

func Href(href string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.Set("href", href)
	})
}

func Width(w string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.Set("width", w)
	})
}

func Height(h string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.Set("height", h)
	})
}

func Figure(mods ...gt.Renderable) gt.Node {
	return gt.Element("figure", mods...)
}

func Ul(mods ...gt.Renderable) gt.Node {
	return gt.Element("ul", mods...)
}

func Li(mods ...gt.Renderable) gt.Node {
	return gt.Element("li", mods...)
}

func Ol(mods ...gt.Renderable) gt.Node {
	return gt.Element("ol", mods...)
}

func Img(mods ...gt.Renderable) gt.Node {
	return gt.Element("img", mods...)
}

func P(mods ...gt.Renderable) gt.Node {
	return gt.Element("p", mods...)
}

func Pre(mods ...gt.Renderable) gt.Node {
	return gt.Element("pre", mods...)
}

func Code(mods ...gt.Renderable) gt.Node {
	return gt.Element("code", mods...)
}

func Aside(mods ...gt.Renderable) gt.Node {
	return gt.Element("aside", mods...)
}

func Blockquote(mods ...gt.Renderable) gt.Node {
	return gt.Element("blockquote", mods...)
}

func Figcaption(mods ...gt.Renderable) gt.Node {
	return gt.Element("figcaption", mods...)
}

func Span(mods ...gt.Renderable) gt.Node {
	return gt.Element("span", mods...)
}

func ID(id string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetID(id)
	})
}

func AddEventListener(eventType string, f func()) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, false, f)
	})
}

func AddClickListener(f func()) gt.Modifier {
	return AddEventListener("click", f)
}

func AddKeyDownListener(f func(keyCode int)) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.AddKeyListener("keyup", f)
	})
}

func AddEventListenerOnce(eventType string, f func()) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, true, f)
	})
}

func AriaLabel(label string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-label", label)
	})
}

func AriaOrientation(orientation string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-orientation", orientation)
	})
}

func AriaLabelledby(label string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-labelledby", label)
	})
}

func Role(role string) gt.Modifier {
	return gt.ModifierFunc(func(e dom.Element) {
		e.SetAttribute("role", role)
	})
}
