package main

import "fmt"

type navPills struct {
	links []fmt.Stringer
}

func newNavPills(links []fmt.Stringer) *navPills {
	return &navPills{
		links: links,
	}
}

func (c *navPills) String() string {
	s := `<ul class="nav nav-pills nav-fill">`

	for _, link := range c.links {
		s += link.String()
	}

	s += `</ul>`

	return s
}
