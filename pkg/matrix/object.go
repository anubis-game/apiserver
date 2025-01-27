package matrix

import "github.com/google/uuid"

type Object struct {
	Bck Bucket
	Pxl Pixel
	Pro Profile
	Uid uuid.UUID
}
