package unique

import "testing"

func Test_Unique_Number(t *testing.T) {
	var byt int
	var u08 int
	var u16 int
	var i08 int
	var i16 int
	{
		byt = maxOf[byte]()
		u08 = maxOf[uint8]()
		u16 = maxOf[uint16]()
		i08 = maxOf[int8]()
		i16 = maxOf[int16]()
	}

	if byt != 255 {
		t.Fatalf("expected %#v got %#v", 255, byt)
	}
	if u08 != 255 {
		t.Fatalf("expected %#v got %#v", 255, u08)
	}
	if u16 != 65535 {
		t.Fatalf("expected %#v got %#v", 65535, u16)
	}
	if i08 != 127 {
		t.Fatalf("expected %#v got %#v", 127, i08)
	}
	if i16 != 32767 {
		t.Fatalf("expected %#v got %#v", 32767, i16)
	}
}
