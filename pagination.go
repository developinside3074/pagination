package library

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


func NewPaginator(currentPage int, totalItems int, perPage int, rangePage int) *Paginator{
    return &Paginator{
        CurrentPage: currentPage,
        TotalItems: totalItems,
        PerPage: perPage,
        RangePage: rangePage,
    }
}