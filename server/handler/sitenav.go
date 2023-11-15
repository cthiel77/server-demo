package handler

import "github.com/cthiel77/server-demo/internal"

func getNavigationData() (nav []internal.PageNavEntry) {
	nav = append(nav,
		internal.PageNavEntry{Key: "", Name: "Home"},
		//internal.PageNavEntry{Key: "page2", Name: "Page 2"},
	)

	return
}
