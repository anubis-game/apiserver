package window

import "github.com/anubis-game/apiserver/pkg/object"

const (
	// Win describes half of the initial window size in pixels along X and Y. The
	// goal is to put the player into the middle of the screen, which means that
	// we have to define the edges of the visible view. E.g. a window size of
	// 320px implies the total window length along X and Y to be 640px.
	Win = 320
)

type Window struct {
	cbl object.Object
	ctr object.Object
	xfr map[int]int
	yfr map[int]int
}

func New() *Window {
	return &Window{
		cbl: object.Object{},
		ctr: object.Object{},
		xfr: map[int]int{},
		yfr: map[int]int{},
	}
}

func (w *Window) CBL() object.Object {
	return object.Object{X: w.cbl.X, Y: w.cbl.Y}
}

func (w *Window) CBR() object.Object {
	return object.Object{X: w.ctr.X, Y: w.cbl.Y}
}

func (w *Window) CTL() object.Object {
	return object.Object{X: w.cbl.X, Y: w.ctr.Y}
}

func (w *Window) CTR() object.Object {
	return object.Object{X: w.ctr.X, Y: w.ctr.Y}
}

// Exp allows this window to be expanded with a single coordinate object.  This
// is most relevant for player windows. See also player.New().
func (w *Window) Exp(obj object.Object, win int) {
	w.cbl = object.Object{
		X: obj.X - win,
		Y: obj.Y - win,
	}
	w.ctr = object.Object{
		X: obj.X + win,
		Y: obj.Y + win,
	}
}

// Has returns whether either of another window's 4 corners is inside of this
// window.
func (w *Window) Has(win *Window) bool {
	return w.has(win.CBL()) || w.has(win.CBR()) || w.has(win.CTL()) || w.has(win.CTR())
}

// has returns whether the given bucket resides inside the underlying Window. So
// if obj turns out to be outside of w, then Has returns false.
func (w *Window) has(obj object.Object) bool {
	return obj.X >= w.cbl.X && obj.X <= w.ctr.X && obj.Y >= w.cbl.Y && obj.Y <= w.ctr.Y
}

// Dec allows this window to shrink up to tai, if rem was part of this window's
// boundary.
func (w *Window) Dec(tai object.Object, rem object.Object) {
	{
		w.xfr[rem.X]--
		w.yfr[rem.Y]--
	}

	if rem.X == w.cbl.X && w.xfr[rem.X] == 0 {
		w.cbl.X = tai.X
	}
	if rem.X == w.ctr.X && w.xfr[rem.X] == 0 {
		w.ctr.X = tai.X
	}
	if rem.Y == w.cbl.Y && w.yfr[rem.Y] == 0 {
		w.cbl.Y = tai.Y
	}
	if rem.Y == w.ctr.Y && w.yfr[rem.Y] == 0 {
		w.ctr.Y = tai.Y
	}

	if w.xfr[rem.X] == 0 {
		delete(w.xfr, rem.X)
	}
	if w.yfr[rem.Y] == 0 {
		delete(w.yfr, rem.Y)
	}
}

// Inc allows this window to grow, if trg exceeds any boundary of this window.
func (w *Window) Inc(trg object.Object) {
	{
		w.xfr[trg.X]++
		w.yfr[trg.Y]++
	}

	if trg.X < w.cbl.X {
		w.cbl.X = trg.X
	}
	if trg.X > w.ctr.X {
		w.ctr.X = trg.X
	}
	if trg.Y < w.cbl.Y {
		w.cbl.Y = trg.Y
	}
	if trg.Y > w.ctr.Y {
		w.ctr.Y = trg.Y
	}
}

// Ini allows this window to be initialized with a single coordinate object.
// This is most relevant for vector windows. See also vector.New().
func (w *Window) Ini(obj object.Object) {
	{
		w.cbl = obj
		w.ctr = obj
	}

	{
		w.xfr = map[int]int{obj.X: 1}
		w.yfr = map[int]int{obj.Y: 1}
	}
}
