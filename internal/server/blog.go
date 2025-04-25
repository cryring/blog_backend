package server

import (
	"strconv"

	"github.com/cryring/blog_backend/internal/blog"
	"github.com/cryring/blog_backend/internal/log"
	"github.com/gin-gonic/gin"
)

const (
	defaultPageSize = 10
)

type UriBlogs struct {
	Category blog.Category `uri:"category" binding:"required"`
}

func (srv *Server) handleBlogs(c *gin.Context) {
	var uri UriBlogs
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var (
		paramPage     = c.Params.ByName("page")
		paramPageSize = c.Params.ByName("page_size")

		page     int
		pageSize = defaultPageSize
	)

	num, err := strconv.Atoi(paramPageSize)
	if err != nil {
		log.Errorf("invalid param page_size: [%s]", paramPageSize)
	}
	if num > 0 && num < 100 {
		pageSize = num
	}

	num, err = strconv.Atoi(paramPage)
	if err != nil {
		log.Errorf("invalid param page: [%s]", paramPage)
	}
	if num >= 0 {
		page = num
	}

	log.Debugf("handleBlogs params: category[%v], page[%v], pageSize[%v]", uri.Category, page, pageSize)

	// TODO[erik]: support tag query
	blogs, err := srv.db.GetBlogs(uri.Category, page, pageSize)
	if err != nil {
		log.Errorf("db get blogs[%v:%v:%v] failed: %v", uri.Category, page, pageSize, err)
		ResponseOK(c, gin.H{})
		return
	}

	ResponseOK(c, gin.H{"blogs": blogs})
}

type UriBlogContent struct {
	Category blog.Category `uri:"category" binding:"required"`
	FileName string        `uri:"filename" binding:"required"`
}

func (srv *Server) handlBlogContent(c *gin.Context) {
	var uri UriBlogContent
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	log.Debugf("handlBlogContent params: category[%v], filename[%v]", uri.Category, uri.FileName)

	b := blog.New(uri.Category, uri.FileName)
	if err := b.Load(); err != nil {
		ResponseOK(c, gin.H{})
		return
	}
	ResponseOK(c, gin.H{"content": b.Content()})
}
