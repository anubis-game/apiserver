package matrix

import "github.com/google/uuid"

type Object struct {
	Bck Bucket
	Pxl Pixel
	Pro Profile
	Uid uuid.UUID
}

func (o Object) Inside(win Window) bool {
	// TODO
	return false
}
