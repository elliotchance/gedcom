package main

import "github.com/elliotchance/gedcom/html"

// navTabs is a group of tabs.
type navTabs struct {
	items []*navItem
}

func newNavTabs(items []*navItem) *navTabs {
	return &navTabs{
		items: items,
	}
}

func (c *navTabs) String() string {
	tabs := ""
	for _, tab := range c.items {
		tabs += tab.String()
	}

	return html.Sprintf(`
    <div class="row">
        <div class="col">
            <ul class="nav nav-tabs">
                %s
            </ul>
        </div>
    </div>`, tabs)
}
