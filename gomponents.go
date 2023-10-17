package main

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	s "github.com/maragudk/gomponents/svg"
)

func NavbarDaisy(links []PageLink) g.Node {
	return h.Div(h.Class("navbar bg-base-100"),
		h.Div(h.Class("navbar-start"),
			h.Div(h.Class("dropdown"),
				h.Label(h.TabIndex("0"), h.Class("btn btn-ghost btn-circle"),
					h.SVG(h.Class("h-5 w-5"), s.Fill("none"), s.ViewBox("0 0 24 24"), s.Stroke("currentColor"),
						s.Path(s.StrokeWidth("2"), s.D("M4 6h16M4 18h7")),
					),
				),
				h.Ul(h.TabIndex("0"), h.Class("menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"),
					g.Group(g.Map(links, func(l PageLink) g.Node {
						return h.Li(h.A(h.Href(l.Path), g.Text(l.Name)))
					})),
				),
			),
		),
		h.Div(h.Class("navbar-center"),
			h.A(h.Href("/"), h.Class("btn btn-ghost normal-case text-xl"), g.Text("Bourgomponents")),
		),
		h.Div(h.Class("navbar-end"),
			h.Button(h.Class("btn btn-ghost btn-circle"),
				h.SVG(h.Class("h-5 w-5"), s.Fill("none"), s.ViewBox("0 0 24 24"), s.Stroke("currentColor"), s.Path(s.StrokeWidth("2"), s.D("M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"))),
			),
			h.Button(h.Class("btn btn-ghost btn-circle"),
				h.SVG(h.Class("h-5 w-5"), s.Fill("none"), s.ViewBox("0 0 24 24"), s.Stroke("currentColor"),
					s.Path(s.StrokeWidth("2"),
						s.D("M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"),
					),
				),
				h.Span(h.Class("badge badge-xs badge-primary indicator-item")),
			),
		),
	)
}
