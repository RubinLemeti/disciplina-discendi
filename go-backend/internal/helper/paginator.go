package helper

import "strconv"

func GetPageUrl(path string, limit int, pageNumber int) string {
	var offset int

	if pageNumber > 1 {
		offset = limit & (pageNumber - 1)
		pathToPage := path + "?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)
		return pathToPage
	}

	pathToPage := path + "?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)
	return pathToPage
}

func Paginate(total int, limit int, offset int, pageRecords int, path string) *Pagination{
	firstPage := 1
	currentPage := (offset / limit) + 1
	lastPage := total / limit

	var previousPage int
	if currentPage > 1 {
		previousPage = currentPage - 1
	}

	var nextPage int
	if lastPage > 1 && currentPage < lastPage {
		nextPage = currentPage + 1
	}

	var currentPageFirstRecord int
	var currentPageLastRecord int

	if pageRecords > 0 && currentPage > 0 {
		currentPageFirstRecord = (currentPage-1)*limit + 1
		currentPageLastRecord = (currentPage-1)*limit + pageRecords
	}

	records := Records{
		First:  currentPageFirstRecord,
		Last:   currentPageLastRecord,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}

	links := Links{
		First:    Href{Href: GetPageUrl(path, limit, firstPage)},
		Previous: Href{Href: GetPageUrl(path, limit, previousPage)},
		Current:  Href{Href: GetPageUrl(path, limit, currentPage)},
		Next:     Href{Href: GetPageUrl(path, limit, nextPage)},
		Last:     Href{Href: GetPageUrl(path, limit, lastPage)},
	}

	paginated := Pagination{
		First:    firstPage,
		Previous: previousPage,
		Current:  currentPage,
		Next:     nextPage,
		Last:     lastPage,
		Links:    links,
		Records:  records,
	}

	return &paginated
}
