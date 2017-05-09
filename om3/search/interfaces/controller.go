package interfaces

import (
	"flamingo/framework/web"
	"flamingo/framework/web/responder"
	"flamingo/om3/search/domain"
)

type (
	// ViewController demonstrates a search view controller
	ViewController struct {
		*responder.ErrorAware    `inject:""`
		*responder.RenderAware   `inject:""`
		domain.SearchService     `inject:""`
	}

	// ViewData is used for search rendering
	ViewData struct {
		SearchResult *domain.SearchResult
	}
)

// Get Response for search
func (vc *ViewController) Get(c web.Context) web.Response {

	searchResult, err := vc.SearchService.Search(c, c.MustQuery1("q"))

	// catch error
	if err != nil {
		return vc.Error(c, err)
	}

	// render page
	return vc.Render(c, "pages/search/view", ViewData{SearchResult: searchResult})
}