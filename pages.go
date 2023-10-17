package main

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func indexPage() (string, g.Node) {
	return "Welcome!", h.Article(
		h.H1(g.Text("Welcome to this example page")),
		h.P(g.Text("I hope it will make you happy. 😄 It's using TailwindCSS for styling.")),
	)
}

func contactPage() (string, g.Node) {
	return "Contact", h.Article(
		h.H1(g.Text("Contact us")),
		h.P(g.Text("Just do it.")),
	)
}

func aboutPage() (string, g.Node) {
	return "About", h.Article(
		h.H1(g.Text("About this site")),
		h.P(g.Text("This is a site showing off gomponents.")),
	)
}
