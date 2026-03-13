package bookmark

type Bookmark struct {
	Name string
	URL  string
}

type BookmarkService struct {
	bookmarks map[string]*Bookmark
}

func NewService() *BookmarkService {
	return &BookmarkService{
		bookmarks: make(map[string]*Bookmark),
	}

}
func (b *BookmarkService) AddBookmark(name string, url string) {
	b.bookmarks[name] = &Bookmark{
		Name: name,
		URL:  url,
	}

}

func (b *BookmarkService) DeleteBookmark(name string) (string, bool) {
	bookmark, ok := b.bookmarks[name]
	if !ok {
		return "", false
	}
	delete(b.bookmarks, name)
	return bookmark.URL, true
}
func (b *BookmarkService) ListBookmarks() map[string]*Bookmark {
	result := make(map[string]*Bookmark, len(b.bookmarks))
	for name, bookmark := range b.bookmarks {
		result[name] = bookmark
	}
	return result
}
