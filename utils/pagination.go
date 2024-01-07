package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func Pagination(c *gin.Context, nextpage, page, perpage, prevpage, total, totalpages string) {
func Pagination(c *gin.Context, totalProjects, projectsPerPage int) {
	page := "1"

	if c.Query("page") != "" {
		page = c.Query("page")
	}

	currentPage, err := strconv.Atoi(page)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}

	pageCount := totalProjects / projectsPerPage

	c.Writer.Header().Set("x-next-page", "")
	c.Writer.Header().Set("x-prev-page", "")

	c.Header("x-page", fmt.Sprint(currentPage))
	c.Header("x-per-page", fmt.Sprint(projectsPerPage))

	// if the current page is less than the total pages
	// then we can set the next page header
	if currentPage < pageCount {
		c.Header("x-next-page", fmt.Sprint(currentPage+1))
	}

	// if the current page is less than or equal to the total pages
	// then we can set the prev page header
	if currentPage <= pageCount {
		// the pages start at 1, so we can't have a prev page of 0
		if currentPage-1 != 0 {
			c.Header("x-prev-page", fmt.Sprint(currentPage-1))
		}
	}

	c.Header("x-total", fmt.Sprintf("%d", totalProjects))
	c.Header("x-total-pages", fmt.Sprintf("%d", pageCount))
}
