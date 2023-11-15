package internal

// PageNavEntry represents a page navigation entry
type PageNavEntry struct {
	Name string
	Key  string
}

// ContentBlock represents a content block entry
type ContentBlock struct {
	Title   string
	SubHead string
	Content string
	Link    ContentLink
}

// ContentLink represents a content link entry
type ContentLink struct {
	Label string
	URL   string
}
