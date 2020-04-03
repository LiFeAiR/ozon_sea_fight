package dto

type Ship struct {
	Coordinates Coordinates
	Knock       int
}

type Coordinates struct {
	LeftTop     string
	RightBottom string
}
