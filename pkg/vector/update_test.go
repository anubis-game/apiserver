package vector

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Test_Vector_Update(t *testing.T) {
	testCases := []struct {
		qdr byte
		agl byte
		hid int
		len int
		nod []matrix.Coordinate
		ocd Change
		upd func(*Vector)
	}{
		// Case 000, siz: 5
		{
			qdr: 0x1,
			agl: 0x0,
			hid: 0,
			len: 2,
			nod: []matrix.Coordinate{
				{X: 1000, Y: 1005}, // H
				{X: 1000, Y: 1000}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1000, Y: 1005},
			},
			upd: func(vec *Vector) {
				for range 1 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
			},
		},
		// Case 001, siz: 10
		{
			qdr: 0x1,
			agl: 0x0,
			hid: 0,
			len: 2,
			nod: []matrix.Coordinate{
				{X: 1000, Y: 1010}, // H
				{X: 1000, Y: 1005}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1000, Y: 1010},
				Tai: matrix.Coordinate{X: 1000, Y: 1005},
				Rem: []matrix.Coordinate{
					{X: 1000, Y: 1000},
				},
			},
			upd: func(vec *Vector) {
				for range 2 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
			},
		},
		// Case 002, siz: 15
		{
			qdr: 0x1,
			agl: 0x0,
			hid: 1,
			len: 3,
			nod: []matrix.Coordinate{
				{X: 1000, Y: 1015}, // H
				// x=1000 y=1010
				{X: 1000, Y: 1005}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1000, Y: 1015},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
			},
		},
		// Case 003, siz: 20
		{
			qdr: 0x1,
			agl: 0x3a,
			hid: 2,
			len: 4,
			nod: []matrix.Coordinate{
				{X: 1002, Y: 1020}, // H
				// x=1000 y=1015
				// x=1000 y=1010
				{X: 1000, Y: 1005}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1002, Y: 1020},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 1 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 004, siz: 25
		{
			qdr: 0x1,
			agl: 0x74,
			hid: 3,
			len: 5,
			nod: []matrix.Coordinate{
				{X: 1005, Y: 1024}, // H
				// x=1002 y=1020
				// x=1000 y=1015
				// x=1000 y=1010
				{X: 1000, Y: 1005}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1005, Y: 1024},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 2 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 005, siz: 30
		{
			qdr: 0x1,
			agl: 0xae,
			hid: 3,
			len: 6,
			nod: []matrix.Coordinate{
				{X: 1009, Y: 1026}, // H
				{X: 1005, Y: 1024},
				// x=1002 y=1020
				// x=1000 y=1015
				// x=1000 y=1010
				{X: 1000, Y: 1005}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1009, Y: 1026},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 3 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 006, siz: 35
		{
			qdr: 0x1,
			agl: 0xe8,
			hid: 3,
			len: 6,
			nod: []matrix.Coordinate{
				{X: 1014, Y: 1027}, // H
				// x=1009 y=1026
				{X: 1005, Y: 1024},
				// x=1002 y=1020
				// x=1000 y=1015
				{X: 1001, Y: 1010}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1014, Y: 1027},
				Tai: matrix.Coordinate{X: 1001, Y: 1010},
				Rem: []matrix.Coordinate{
					{X: 1000, Y: 1005},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 4 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 007, siz: 40
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 4,
			len: 7,
			nod: []matrix.Coordinate{
				{X: 1019, Y: 1027}, // H
				// x=1014 y=1027
				// x=1009 y=1026
				{X: 1005, Y: 1024},
				// x=1002 y=1020
				// x=1000 y=1015
				{X: 1001, Y: 1010}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1019, Y: 1027},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 5 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 008, siz: 45
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 5,
			len: 8,
			nod: []matrix.Coordinate{
				{X: 1024, Y: 1027}, // H
				// x=1019 y=1027
				// x=1014 y=1027
				// x=1009 y=1026
				{X: 1005, Y: 1024},
				// x=1002 y=1020
				// x=1000 y=1015
				{X: 1001, Y: 1010}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1024, Y: 1027},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 6 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 009, siz: 50
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 4,
			len: 8,
			nod: []matrix.Coordinate{
				{X: 1029, Y: 1027}, // H
				{X: 1024, Y: 1027},
				// x=1019 y=1027
				// x=1014 y=1027
				// x=1009 y=1026
				{X: 1005, Y: 1024},
				// x=1002 y=1020
				{X: 1002, Y: 1015}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1029, Y: 1027},
				Tai: matrix.Coordinate{X: 1002, Y: 1015},
				Rem: []matrix.Coordinate{
					{X: 1001, Y: 1010},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 010, siz: 45
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 4,
			len: 8,
			nod: []matrix.Coordinate{
				{X: 1034, Y: 1027}, // H
				// x=1029 y=1027
				{X: 1024, Y: 1027},
				// x=1019 y=1027
				// x=1014 y=1027
				// x=1009 y=1026
				{X: 1005, Y: 1024},
				{X: 1004, Y: 1020}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1034, Y: 1027},
				Tai: matrix.Coordinate{X: 1004, Y: 1020},
				Rem: []matrix.Coordinate{
					{X: 1002, Y: 1015},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 1 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 011, siz: 40
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 4,
			len: 7,
			nod: []matrix.Coordinate{
				{X: 1039, Y: 1027}, // H
				// x=1034 y=1027
				// x=1029 y=1027
				{X: 1024, Y: 1027},
				// x=1019 y=1027
				// x=1014 y=1027
				{X: 1010, Y: 1025}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1039, Y: 1027},
				Tai: matrix.Coordinate{X: 1010, Y: 1025},
				Rem: []matrix.Coordinate{
					{X: 1005, Y: 1024},
					{X: 1004, Y: 1020},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 2 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 012, siz: 35
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 3,
			len: 6,
			nod: []matrix.Coordinate{
				{X: 1044, Y: 1027}, // H
				// x=1039 y=1027
				// x=1034 y=1027
				// x=1029 y=1027
				{X: 1024, Y: 1027},
				{X: 1020, Y: 1027}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1044, Y: 1027},
				Tai: matrix.Coordinate{X: 1020, Y: 1027},
				Rem: []matrix.Coordinate{
					{X: 1015, Y: 1026},
					{X: 1010, Y: 1025},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 3 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 013, siz: 30
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 3,
			len: 6,
			nod: []matrix.Coordinate{
				{X: 1049, Y: 1027}, // H
				{X: 1044, Y: 1027},
				// x=1039 y=1027
				// x=1034 y=1027
				// x=1029 y=1027
				{X: 1024, Y: 1027}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1049, Y: 1027},
				Tai: matrix.Coordinate{X: 1024, Y: 1027},
				Rem: []matrix.Coordinate{
					{X: 1020, Y: 1027},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 4 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 014, siz: 25
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 2,
			len: 5,
			nod: []matrix.Coordinate{
				{X: 1054, Y: 1027}, // H
				// x=1049 y=1027
				{X: 1044, Y: 1027},
				// x=1039 y=1027
				{X: 1034, Y: 1027}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1054, Y: 1027},
				Tai: matrix.Coordinate{X: 1034, Y: 1027},
				Rem: []matrix.Coordinate{
					{X: 1029, Y: 1027},
					{X: 1024, Y: 1027},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 5 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 015, siz: 20
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 2,
			len: 4,
			nod: []matrix.Coordinate{
				{X: 1059, Y: 1027}, // H
				// x=1054 y=1027
				// x=1049 y=1027
				{X: 1044, Y: 1027}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1059, Y: 1027},
				Tai: matrix.Coordinate{X: 1044, Y: 1027},
				Rem: []matrix.Coordinate{
					{X: 1039, Y: 1027},
					{X: 1034, Y: 1027},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 6 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 016, siz: 15
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 1,
			len: 3,
			nod: []matrix.Coordinate{
				{X: 1064, Y: 1027}, // H
				// x=1059 y=1027
				{X: 1054, Y: 1027}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1064, Y: 1027},
				Tai: matrix.Coordinate{X: 1054, Y: 1027},
				Rem: []matrix.Coordinate{
					{X: 1049, Y: 1027},
					{X: 1044, Y: 1027},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 017, siz: 10
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 0,
			len: 2,
			nod: []matrix.Coordinate{
				{X: 1069, Y: 1027}, // H
				{X: 1064, Y: 1027}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1069, Y: 1027},
				Tai: matrix.Coordinate{X: 1064, Y: 1027},
				Rem: []matrix.Coordinate{
					{X: 1059, Y: 1027},
					{X: 1054, Y: 1027},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 8 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 018, siz: 5
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 0,
			len: 2,
			nod: []matrix.Coordinate{
				{X: 1074, Y: 1027}, // H
				{X: 1069, Y: 1027}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1074, Y: 1027},
				Tai: matrix.Coordinate{X: 1069, Y: 1027},
				Rem: []matrix.Coordinate{
					{X: 1064, Y: 1027},
				},
			},
			upd: func(vec *Vector) {
				for range 3 {
					vec.Update(int(Si/Li), 0x1, 0x0, Nrm)
				}
				for range 7 {
					vec.Update(int(Si/Li), 0x2, 0x0, Nrm)
				}
				for range 9 {
					vec.Update(-int(Si/Li), 0x2, 0x0, Nrm)
				}
			},
		},
		// Case 019, len: 78, siz: 785
		{
			qdr: 0x2,
			agl: 0x0,
			hid: 57,
			len: 78,
			nod: []matrix.Coordinate{
				{X: 1304, Y: 1303}, // H
				{X: 1284, Y: 1303},
				{X: 1264, Y: 1303},
				{X: 1244, Y: 1303},
				{X: 1224, Y: 1303},
				{X: 1204, Y: 1303},
				{X: 1184, Y: 1303},
				{X: 1164, Y: 1303},
				{X: 1144, Y: 1303},
				{X: 1124, Y: 1303},
				{X: 1104, Y: 1303},
				{X: 1084, Y: 1303},
				{X: 1064, Y: 1303},
				{X: 1044, Y: 1303},
				{X: 1024, Y: 1303},
				{X: 1005, Y: 1299},
				{X: 1000, Y: 1280},
				{X: 1000, Y: 1260},
				{X: 1000, Y: 1240},
				{X: 1000, Y: 1220},
				{X: 1000, Y: 1215}, // T
			},
			ocd: Change{
				Hea: matrix.Coordinate{X: 1304, Y: 1303},
			},
			upd: tesUpd,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec *Vector
			{
				vec = tesVec()
			}

			{
				tc.upd(vec)
			}

			var hea matrix.Coordinate
			var tai matrix.Coordinate
			{
				hea = tc.nod[0]
				tai = tc.nod[len(tc.nod)-1]
			}

			var nod []matrix.Coordinate
			vec.Ranger(func(o matrix.Coordinate) {
				nod = append(nod, o)
			})

			{
				sort.Sort(matrix.Coordinates(tc.nod))
				sort.Sort(matrix.Coordinates(nod))
			}

			if vec.hea.crd != hea {
				t.Fatalf("expected %#v got %#v", hea, vec.hea.crd)
			}
			if vec.tai.crd != tai {
				t.Fatalf("expected %#v got %#v", tai, vec.tai.crd)
			}
			if !slices.Equal(nod, tc.nod) {
				t.Fatalf("expected %#v got %#v", tc.nod, nod)
			}
			if !reflect.DeepEqual(vec.ocd, tc.ocd) {
				t.Fatalf("expected %#v got %#v", tc.ocd, vec.ocd)
			}
			if vec.hidden() != tc.hid {
				t.Fatalf("expected %#v got %#v", tc.hid, vec.hidden())
			}
			if vec.len != tc.len {
				t.Fatalf("expected %#v got %#v", tc.len, vec.len)
			}
			if vec.len != len(tc.nod)+tc.hid {
				t.Fatalf("expected %#v got %#v", len(tc.nod)+tc.hid, vec.len)
			}
			if vec.mot.Qdr != tc.qdr {
				t.Fatalf("expected %#v got %#v", tc.qdr, vec.mot.Qdr)
			}
			if vec.mot.Agl != tc.agl {
				t.Fatalf("expected %#v got %#v", tc.agl, vec.mot.Agl)
			}
		})
	}
}

func Test_Vector_Update_angle(t *testing.T) {
	testCases := []struct {
		siz float64
		agl byte
	}{
		// Case 000
		{
			siz: 10,
			agl: 30,
		},
		// Case 001
		{
			siz: 50,
			agl: 29,
		},
		// Case 002
		{
			siz: 100,
			agl: 28,
		},
		// Case 003
		{
			siz: 250,
			agl: 27,
		},
		// Case 004
		{
			siz: 500,
			agl: 26,
		},
		// Case 005
		{
			siz: 1_000,
			agl: 24,
		},
		// Case 006
		{
			siz: 2_500,
			agl: 21,
		},
		// Case 007
		{
			siz: 5_000,
			agl: 19,
		},
		// Case 008
		{
			siz: 10_000,
			agl: 15,
		},
		// Case 009
		{
			siz: 20_000,
			agl: 10,
		},
		// Case 010
		{
			siz: 30_000,
			agl: 6,
		},
		// Case 011
		{
			siz: 40_000,
			agl: 3,
		},
		// Case 012
		{
			siz: 50_000,
			agl: 1,
		},
		// Case 013
		{
			siz: 60_000,
			agl: 1,
		},
		// Case 014
		{
			siz: 70_000,
			agl: 1,
		},
		// Case 015
		{
			siz: 80_000,
			agl: 1,
		},
		// Case 016
		{
			siz: 90_000,
			agl: 1,
		},
		// Case 017
		{
			siz: 100_000,
			agl: 1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			agl := angle(tc.siz)

			if agl != tc.agl {
				t.Fatalf("expected %d got %d", tc.agl, agl)
			}
		})
	}
}

func Test_Vector_Update_length(t *testing.T) {
	testCases := []struct {
		siz float64
		len int
	}{
		// Case 000
		{
			siz: 10,
			len: 2,
		},
		// Case 001
		{
			siz: 50,
			len: 8,
		},
		// Case 002
		{
			siz: 100,
			len: 14,
		},
		// Case 003
		{
			siz: 250,
			len: 30,
		},
		// Case 004
		{
			siz: 500,
			len: 53,
		},
		// Case 005
		{
			siz: 1_000,
			len: 95,
		},
		// Case 006
		{
			siz: 2_500,
			len: 204,
		},
		// Case 007
		{
			siz: 5_000,
			len: 364,
		},
		// Case 008
		{
			siz: 10_000,
			len: 650,
		},
		// Case 009
		{
			siz: 20_000,
			len: 1_161,
		},
		// Case 010
		{
			siz: 30_000,
			len: 1_630,
		},
		// Case 011
		{
			siz: 40_000,
			len: 2_074,
		},
		// Case 012
		{
			siz: 50_000,
			len: 2_500,
		},
		// Case 013
		{
			siz: 60_000,
			len: 2_500,
		},
		// Case 014
		{
			siz: 70_000,
			len: 2_500,
		},
		// Case 015
		{
			siz: 80_000,
			len: 2_500,
		},
		// Case 016
		{
			siz: 90_000,
			len: 2_500,
		},
		// Case 017
		{
			siz: 100_000,
			len: 2_500,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			len := length(tc.siz)

			if len != tc.len {
				t.Fatalf("expected %#v got %#v", tc.len, len)
			}
		})
	}
}

func Test_Vector_Update_radius(t *testing.T) {
	testCases := []struct {
		siz float64
		rad byte
	}{
		// Case 000
		{
			siz: 10,
			rad: 10,
		},
		// Case 001
		{
			siz: 50,
			rad: 10,
		},
		// Case 002
		{
			siz: 100,
			rad: 12,
		},
		// Case 003
		{
			siz: 250,
			rad: 22,
		},
		// Case 004
		{
			siz: 500,
			rad: 33,
		},
		// Case 005
		{
			siz: 1_000,
			rad: 47,
		},
		// Case 006
		{
			siz: 2_500,
			rad: 73,
		},
		// Case 007
		{
			siz: 5_000,
			rad: 98,
		},
		// Case 008
		{
			siz: 10_000,
			rad: 128,
		},
		// Case 009
		{
			siz: 20_000,
			rad: 166,
		},
		// Case 010
		{
			siz: 30_000,
			rad: 190,
		},
		// Case 011
		{
			siz: 40_000,
			rad: 210,
		},
		// Case 012
		{
			siz: 50_000,
			rad: 225,
		},
		// Case 013
		{
			siz: 60_000,
			rad: 225,
		},
		// Case 014
		{
			siz: 70_000,
			rad: 225,
		},
		// Case 015
		{
			siz: 80_000,
			rad: 225,
		},
		// Case 016
		{
			siz: 90_000,
			rad: 225,
		},
		// Case 017
		{
			siz: 100_000,
			rad: 225,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			rad := radius(tc.siz)

			if rad != tc.rad {
				t.Fatalf("expected %d got %d", tc.rad, rad)
			}
		})
	}
}

func Test_Vector_Update_sight(t *testing.T) {
	testCases := []struct {
		siz float64
		fos byte
	}{
		// Case 000
		{
			siz: 10,
			fos: 2,
		},
		// Case 001
		{
			siz: 50,
			fos: 2,
		},
		// Case 002
		{
			siz: 100,
			fos: 2,
		},
		// Case 003
		{
			siz: 250,
			fos: 3,
		},
		// Case 004
		{
			siz: 500,
			fos: 3,
		},
		// Case 005
		{
			siz: 1_000,
			fos: 4,
		},
		// Case 006
		{
			siz: 2_500,
			fos: 4,
		},
		// Case 007
		{
			siz: 5_000,
			fos: 5,
		},
		// Case 008
		{
			siz: 10_000,
			fos: 6,
		},
		// Case 009
		{
			siz: 20_000,
			fos: 6,
		},
		// Case 010
		{
			siz: 30_000,
			fos: 7,
		},
		// Case 011
		{
			siz: 40_000,
			fos: 7,
		},
		// Case 012
		{
			siz: 50_000,
			fos: 7,
		},
		// Case 013
		{
			siz: 60_000,
			fos: 7,
		},
		// Case 014
		{
			siz: 70_000,
			fos: 7,
		},
		// Case 015
		{
			siz: 80_000,
			fos: 7,
		},
		// Case 016
		{
			siz: 90_000,
			fos: 7,
		},
		// Case 017
		{
			siz: 100_000,
			fos: 7,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			fos := sight(tc.siz)

			if fos != tc.fos {
				t.Fatalf("expected %#v got %#v", tc.fos, fos)
			}
		})
	}
}

func Test_Vector_Update_trgAgl(t *testing.T) {
	testCases := []struct {
		pqd byte
		pag byte
		nqd byte
		nag byte
		lim byte
		qdr byte
		agl byte
	}{
		// Case 000, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to remain inside the 1st
		// quadrant, towards the 2nd quadrant.
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(2),
			nag: byte(87),
			lim: byte(100),
			qdr: byte(1),
			agl: byte(207),
		},
		// Case 001, the desired range of motion is granted, because the maximum
		// angle deviation of 100 bytes is respected, which then does not force the
		// allowed range of motion to be restricted.
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(1),
			nag: byte(87),
			lim: byte(100),
			qdr: byte(1),
			agl: byte(87),
		},
		// Case 002, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to remain inside the 1st
		// quadrant, towards the 2nd quadrant.
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(4), // under 180°, move clockwise
			lim: byte(100),
			qdr: byte(1),
			agl: byte(105),
		},
		// Case 003, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to remain inside the 1st
		// quadrant, towards the 2nd quadrant.
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(5), // exactly 180°, move clockwise
			lim: byte(100),
			qdr: byte(1),
			agl: byte(105),
		},
		// Case 004, the desired range of motion cannot be granted, because the
		// desired range of motion exceeds the maximum angle deviation of 100 bytes,
		// which forces the allowed range of motion to underflow into the 4th
		// quadrant.
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(6), // over 180°, move counter clockwise
			lim: byte(100),
			qdr: byte(4),
			agl: byte(161),
		},
		// Case 005
		{
			pqd: byte(4),
			pag: byte(161),
			nqd: byte(3),
			nag: byte(1), // under 180°, move counter clockwise
			lim: byte(100),
			qdr: byte(4),
			agl: byte(61),
		},
		// Case 006
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(77), // under 180°, move clockwise
			lim: byte(75),
			qdr: byte(2),
			agl: byte(163),
		},
		// Case 007
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(99), // above 180°, move counter clockwise
			lim: byte(75),
			qdr: byte(2),
			agl: byte(13),
		},
		// Case 008
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(90), // under 180°, move clockwise
			lim: byte(175),
			qdr: byte(2),
			agl: byte(90),
		},
		// Case 009
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(85), // under 180°, move clockwise
			lim: byte(175),
			qdr: byte(2),
			agl: byte(85),
		},
		// Case 010
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(1),
			nag: byte(2), // under 180°, move counter clockwise
			lim: byte(175),
			qdr: byte(1),
			agl: byte(169),
		},
		// Case 011
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(88), // exactly 180°, move clockwise
			lim: byte(175),
			qdr: byte(3),
			agl: byte(7),
		},
		// Case 012
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(84), // under 180°, move clockwise
			lim: byte(175),
			qdr: byte(3),
			agl: byte(7),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var qdr byte
			var agl byte
			{
				qdr, agl = trgAgl(tc.pqd, tc.pag, tc.nqd, tc.nag, tc.lim)
			}

			if qdr != tc.qdr {
				t.Fatalf("expected %#v got %#v", tc.qdr, qdr)
			}
			if agl != tc.agl {
				t.Fatalf("expected %#v got %#v", tc.agl, agl)
			}
		})
	}
}

// ~212 ns/op 1 allocs/op
func Benchmark_Vector_Update(b *testing.B) {
	var vec *Vector
	{
		vec = tesVec()
	}

	{
		tesUpd(vec)
	}

	for b.Loop() {
		vec.Update(0, 0x2, 0x0, Nrm)
	}
}

func Benchmark_Vector_Update_trgAgl(b *testing.B) {
	testCases := []struct {
		pqd byte
		pag byte
		nqd byte
		nag byte
		lim byte
	}{
		// Case 000, ~2 ns/op
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(2),
			nag: byte(87),
			lim: byte(100),
		},
		// Case 001, ~2 ns/op
		{
			pqd: byte(1),
			pag: byte(107),
			nqd: byte(1),
			nag: byte(87),
			lim: byte(100),
		},
		// Case 002, ~2 ns/op
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(4), // under 180°, move clockwise
			lim: byte(100),
		},
		// Case 003, ~2 ns/op
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(5), // exactly 180°, move clockwise
			lim: byte(100),
		},
		// Case 004, ~2 ns/op
		{
			pqd: byte(1),
			pag: byte(5),
			nqd: byte(3),
			nag: byte(6), // over 180°, move counter clockwise
			lim: byte(100),
		},
		// Case 005, ~2 ns/op
		{
			pqd: byte(4),
			pag: byte(161),
			nqd: byte(3),
			nag: byte(1), // under 180°, move counter clockwise
			lim: byte(100),
		},
		// Case 006, ~2 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(77), // under 180°, move clockwise
			lim: byte(75),
		},
		// Case 007, ~2 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(99), // above 180°, move counter clockwise
			lim: byte(75),
		},
		// Case 008, ~2 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(90), // under 180°, move clockwise
			lim: byte(175),
		},
		// Case 009, ~2 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(2),
			nag: byte(85), // under 180°, move clockwise
			lim: byte(175),
		},
		// Case 010, ~2 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(1),
			nag: byte(2), // under 180°, move counter clockwise
			lim: byte(175),
		},
		// Case 011, ~2 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(88), // exactly 180°, move clockwise
			lim: byte(175),
		},
		// Case 012, ~2 ns/op
		{
			pqd: byte(2),
			pag: byte(88),
			nqd: byte(4),
			nag: byte(84), // under 180°, move clockwise
			lim: byte(175),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				trgAgl(tc.pqd, tc.pag, tc.nqd, tc.nag, tc.lim)
			}
		})
	}
}
