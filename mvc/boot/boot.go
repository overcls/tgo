package boot

import "github.com/gin-gonic/gin"

type IBoot interface {
	Engine() *gin.Engine
	Context() *gin.Context
	NameSpaces()
}

type Boot struct {
	App *gin.Engine
}


func init()  {




}

func (boot *Boot) Init()  {
	boot.App=new(gin.Engine)
}

func (boot *Boot) NameSpaces(ns string)  {

	boot.App.Group(ns)

}