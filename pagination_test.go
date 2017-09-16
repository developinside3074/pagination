package library

import(
    "testing"
)

//
func TestSetTotalPages(t *testing.T) {
    
    currentPage := 1
    totalItems := 50
    perPage := 5
    rangePage := 5

    pagination := NewPagination(currentPage, totalItems, perPage, rangePage)

    pagination.setTotalPages()

    if pagination.TotalPages != 10{
        t.Errorf("SetTotalPages(): expected pagination.TotalPages equals 10 but obtained %v", pagination.TotalPages)
    }

}

func TestSetTotalPagesWithZero(t *testing.T) {
    
    currentPage := 1
    totalItems := 0
    perPage := 5
    rangePage := 5

    pagination := NewPagination(currentPage, totalItems, perPage, rangePage)

    pagination.setTotalPages()

    if pagination.TotalPages != 0{
        t.Errorf("SetTotalPages(): expected pagination.TotalPages equals 0 but obtained %v", pagination.TotalPages)
    }

}

func TestCheckCurrentPage(t *testing.T) {
    
    currentPage := 0
    totalItems := 0
    perPage := 5
    rangePage := 5

    pagination := NewPagination(currentPage, totalItems, perPage, rangePage)

    pagination.checkCurrentPage()

    if pagination.CurrentPage <= 0 {
        t.Errorf("TestCheckCurrentPage(): expected pagination.CurrrentPage equals 1 but obtained %v", pagination.CurrentPage)
    }

}


func TestSetPreviousPage(t *testing.T) {

    currentPage := 1
    totalItems := 50
    perPage := 5
    rangePage := 5

    pagination := NewPagination(currentPage, totalItems, perPage, rangePage)

    pagination.setPreviousPage()

    if pagination.PreviousPage != 0 {
        t.Errorf("TestPreviousPage(): expected pagination.PreviousPage equals 0 but obtained %v", pagination.PreviousPage)
    }

}