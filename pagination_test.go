package pagination

import (
	"testing"
)

//Test if setTotalPages set the correct value to TotalPages property
func TestSetTotalPages(t *testing.T) {

	currentPage := 1
	totalItems := 50
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)

	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.setTotalPages()

	if pagination.TotalPages != 10 {
		t.Errorf("SetTotalPages(): expected pagination.TotalPages equals 10 but obtained %v", pagination.TotalPages)
	}

}

//Test if setTotalPages set the correct value to TotalPages property if TotalItems is zero
func TestSetTotalPagesWithZero(t *testing.T) {

	currentPage := 1
	totalItems := 0
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.setTotalPages()

	if pagination.TotalPages != 0 {
		t.Errorf("SetTotalPages(): expected pagination.TotalPages equals 0 but obtained %v", pagination.TotalPages)
	}

}

//Test if checkCurrentPage set the correct value to CurrentPage property if CurrentPage is zero
func TestCheckCurrentPage(t *testing.T) {

	currentPage := 0
	totalItems := 0
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.checkCurrentPage()

	if pagination.CurrentPage <= 0 {
		t.Errorf("TestCheckCurrentPage(): expected pagination.CurrrentPage equals 1 but obtained %v", pagination.CurrentPage)
	}

}

//Test if setPreviousPage set the correct value to PreviousPage if CurrentPage is zero
func TestSetPreviousPage(t *testing.T) {

	currentPage := 0
	totalItems := 50
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.setPreviousPage()

	if pagination.PreviousPage != 0 {
		t.Errorf("TestPreviousPage(): expected pagination.PreviousPage equals 0 but obtained %v", pagination.PreviousPage)
	}

}

//Test if setNextPage set the correct value to NextPage property
func TestSetNextPage(t *testing.T) {

	currentPage := 1
	totalItems := 50
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.setNextPage()

	if pagination.NextPage != 2 {
		t.Errorf("TestNextPage(): expected pagination.NextPage equals 2 but obtained %v", pagination.NextPage)
	}

}

//Test if setBottomItem set the correct vaule to BottomItem property
func TestSetBottomItem(t *testing.T) {

	currentPage := 1
	totalItems := 50
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.setBottomItem()

	if pagination.BottomItem != 0 {
		t.Errorf("TestBottomItem(): expected pagination.BottomItem equals 0 but obtained %v", pagination.BottomItem)
	}

}

//Test if setBottomItem set the correct vaule to BottomItem property when current page is bigger than 1
func TestSetBottomItemWithCurrentPageBiggerThanOne(t *testing.T) {

	currentPage := 5
	totalItems := 50
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.setBottomItem()

	if pagination.BottomItem != 20 {
		t.Errorf("TestBottomItem(): expected pagination.BottomItem equals 200 but obtained %v", pagination.BottomItem)
	}

}

//Test if setNextPage set the correct value to NextPage property
//Even more if totalPages is equals one
func TestSetNextPageIfDontHaveMorePages(t *testing.T) {

	currentPage := 1
	totalItems := 1
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.setNextPage()

	if pagination.NextPage != 0 {
		t.Errorf("Expected equals 0, but got %v", pagination.NextPage)
	}

}

//Test if SetValues set the correct values to Paginator object
func TestSetValues(t *testing.T) {

	currentPage := 1
	totalItems := 50
	perPage := 5
	rangePage := 5

	pagination := new(Pagination)
	pagination.CurrentPage = currentPage
	pagination.TotalItems = totalItems
	pagination.PerPage = perPage
	pagination.RangePage = rangePage

	pagination.SetValues()

	if pagination.NextPage != 2 {
		t.Errorf("SetValues(): expected pagination.NextPage equals 2 but obtained %v", pagination.NextPage)
	}

	if pagination.PreviousPage != 0 {
		t.Errorf("SetValues(): expected pagination.PreviousPage equals 0 but obtained %v", pagination.PreviousPage)
	}

	if pagination.TotalPages != 10 {
		t.Errorf("SetValues(): expected pagination.TotalPages equals 10 but obtained %v", pagination.TotalPages)
	}

}
