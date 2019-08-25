package tgo

import (
	"github.com/gin-gonic/gin"
	"github.com/overcls/tgo/pool/orm"
	"github.com/overcls/tgo/util"
)

var (
 DB *orm.DB

 Web *gin.Engine

 VideoID *util.Video
 CommentID *util.Comment

)
