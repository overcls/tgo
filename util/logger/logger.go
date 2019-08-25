package logger

import "fmt"

type Logger interface {
	Init()
}

type Log struct {
	Path string
	Name string
	Level int
}

func (logger *Log) Init() {

	fmt.Println("init logger")

}