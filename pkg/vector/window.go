package vector

const (
	// Win describes half of the initial window size along x and y axis. The goal
	// is to put the player into the middle of this window, which means that we
	// have to define the edges and the center of the window. E.g. a size of 5
	// implies the total window length along x and y axis of 11 inner buckets,
	// which puts the player into the middle of the window at the relative
	// coordinates x=5 y=5. The player has then 5 inner buckets all around the
	// inner bucket that the player is put into.
	Win byte = 5
)

type Window struct {
	OBL Object
	OTR Object
}

// Has returns whether the given bucket resides inside the underlying Window. So
// if obj turns out to be outside of w, then Has returns false.
func (w Window) Has(obj Object) bool {
	return obj.X >= w.OBL.X && obj.X <= w.OTR.X && obj.Y >= w.OBL.Y && obj.Y <= w.OTR.Y
}

// TODO change Window.Has to check if either of another window's 4 corners is inside of itself.
//      We need to check whether a vector window is inside a player window.
