package library

import(
    "testing"
)

func TestSetTotalPages(t *testing.T) {
    currentPage := 1
    totalItems := 50
    perPage := 5
    rangePage := 5

    paginator := NewPaginator(currentPage, totalItems, perPage, rangePage)

    paginator.setTotalPages()

    if paginator.TotalPages != 10{
        t.Errorf("SetTotalPages(): expected paginator.TotalPages equals 10 but but obtained %v", paginator.TotalPages)
    }

}