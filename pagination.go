package library

import(
    "math"
)

//Define paginator type
type Paginator struct{
    CurrentPage int
    PreviousPage int
    NextPage int
    PerPage int
    TotalPages float64
    TotalItems int
    RangePage int
    HasMore bool
}

//Initialize a new construct and return this instance
func NewPaginator(currentPage int, totalItems int, perPage int, rangePage int) *Paginator{
    return &Paginator{
        CurrentPage: currentPage,
        TotalItems: totalItems,
        PerPage: perPage,
        RangePage: rangePage,
    }
}

//Set or update total pages value of object
func (p *Paginator) setTotalPages() {
    if p.TotalItems == 0 || p.PerPage == 0{
        p.TotalPages = 0
    } else{
        p.TotalPages = math.Ceil(float64(p.TotalItems)/float64(p.PerPage))
    }    
}