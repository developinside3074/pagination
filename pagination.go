//Package pagination implements a simple and general porpouse pagination to a "list" of items
package pagination

import(
    "math"
)

//Define paginator type
type Pagination struct{
    CurrentPage int
    PreviousPage int
    NextPage int
    PerPage int
    TotalPages float64
    TotalItems int
    RangePage int
    BottomItem int
}

//Initialize a new construct and return this instance
func NewPagination(currentPage int, totalItems int, perPage int, rangePage int) *Pagination{
    return &Pagination{
        CurrentPage: currentPage,
        TotalItems: totalItems,
        PerPage: perPage,
        RangePage: rangePage,
        NextPage: 1,
    }
}

//Set or update total pages value of object
func (p *Pagination) setTotalPages() {
    if p.TotalItems <= 0 || p.PerPage <= 0{
        p.TotalPages = 0
    } else{
        p.TotalPages = math.Ceil(float64(p.TotalItems)/float64(p.PerPage))
    }    
}

//Check the value of current page and set it to one if necessary
func (p *Pagination) checkCurrentPage() {
    if p.CurrentPage <= 0 {
        p.CurrentPage = 1
    }
}

//Set the value of previous page
func (p *Pagination) setPreviousPage() {
    if p.TotalItems > 0 {
        p.checkCurrentPage()
        p.PreviousPage = p.CurrentPage - 1;
    }
}

//Set the value of next page
func (p *Pagination) setNextPage() {
    if p.TotalItems > 0 {

        if p.TotalPages == 0 {
            p.setTotalPages()
        }

        if p.TotalPages == float64(p.NextPage){
            p.NextPage = 0
        } else {
            p.NextPage = p.CurrentPage + 1
        }

    }
}

//Set bottom value of pagination
func (p *Pagination) setBottomItem(){
    p.checkCurrentPage()
    if p.CurrentPage <= 1 {
        p.BottomItem = 0
    } else {
        p.BottomItem = (p.CurrentPage -1) * p.PerPage
    }
}

//Set Pagination Values
func (p *Pagination) SetValues() {
    p.setTotalPages()
    p.checkCurrentPage()
    p.setPreviousPage()
    p.setBottomItem()
    p.setNextPage()
}