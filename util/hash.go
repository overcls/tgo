package util

type IHash interface {
	Encode() string
	Decode() string
}

type Hash struct {

	Config string
}

type Video struct {
	Hash
}

type Comment struct {
	Hash
}

func (hash *Video) Encode(keyName string) string  {

	return "video encode"

}

func (hash *Video) Decode(keyName string) string  {

	return "video decode"

}

func (hash *Comment) Encode(keyName string) string  {

	return "comment encode"

}

func (hash *Comment) Decode(keyName string) string  {

	return "comment decode"

}