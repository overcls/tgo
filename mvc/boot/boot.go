package boot

import "github.com/gin-gonic/gin"

type IBoot interface {
	Engine() *gin.Engine
	Context() *gin.Context
	NameSpaces()
}

type Boot struct {
	Router *gin.Engine
}

var App Boot
func init()  {

	App = Boot{
		Router:new(gin.Engine),
	}


}

func (boot *Boot) NameSpaces(ns string)  {

	boot.Router.Group(ns)

}