package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func main() {
	http.Handle("/", createHandler(indexPage()))
	http.Handle("/contact", createHandler(contactPage()))
	http.Handle("/about", createHandler(aboutPage()))
	log.Println("Running at port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Could not listen on port 8081 %v\n", err)
	}
}

func createHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = Page(title, r.URL.Path, body).Render(w)
	}
}

func Page(title, path string, body g.Node) g.Node {
	links := []PageLink{
		{Path: "/contact", Name: "Contact"},
		{Path: "/about", Name: "About"},
	}

	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			h.Link(h.Rel("stylesheet"), h.Type("text/css"), h.Href("https://cdn.jsdelivr.net/npm/daisyui@3.9.2/dist/full.css")),
			h.Script(h.Src("https://cdn.tailwindcss.com?plugins=typography")),
		},
		Body: []g.Node{
			NavbarDaisy(links), Container(Prose(body), PageFooter()),
		},
	})
}

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return h.Nav(h.Class("bg-gray-700 mb-4"),
		Container(
			h.Div(h.Class("flex justify-between items-center space-x-4 h-16"),
				h.H1(
					h.Class("text-3xl font-semibold"),
					h.A(h.Href("/"), g.Text("Bourgomponents")),
				),
				h.Div(
					h.Class("flex items-center space-x-4"),
					g.Group(g.Map(links, func(l PageLink) g.Node {
						return NavbarLink(l.Path, l.Name, currentPath == l.Path)
					})),
				),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) g.Node {
	return h.A(h.Href(path), g.Text(text),
		// Apply CSS classes conditionally
		c.Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
			"text-white bg-gray-900":                           active,
			"text-gray-300 hover:text-white hover:bg-gray-700": !active,
		},
	)
}

func Container(children ...g.Node) g.Node {
	return h.Main(h.Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}

func Prose(children ...g.Node) g.Node {
	return h.Section(h.Class("prose"), g.Group(children))
}

func PageFooter() g.Node {
	return h.Footer(h.Class("prose prose-sm prose-indigo"),
		h.P(
			// We can use string interpolation directly, like fmt.Sprintf.
			g.Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			g.If(time.Now().Second()%2 == 0, g.Text("It's an even second.")),
			g.If(time.Now().Second()%2 == 1, g.Text("It's an odd second.")),
		),

		h.P(h.A(h.Href("https://www.gomponents.com"), g.Text("gomponents"))),
	)
}
