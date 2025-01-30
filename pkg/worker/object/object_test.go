package object

import (
	"testing"

	"github.com/anubis-game/apiserver/pkg/transaction"
	"github.com/ethereum/go-ethereum/common"
)

func Test_Worker_Object(t *testing.T) {
	var foo testStruct
	{
		foo = testStruct{
			hsh: New[common.Hash](),
		}
	}

	if !foo.Object().Emp() {
		t.Fatal("expected", true, "got", false)
	}

	var exp common.Hash
	{
		exp = common.HexToHash("0x1897455d07ff874248ecef1cfe4f3802e6ca69fc1fe9d0644c65910ace3a570e")
	}

	{
		foo.Object().Set(exp)
	}

	if foo.Object().Emp() {
		t.Fatal("expected", false, "got", true)
	}

	var act common.Hash
	{
		act = foo.Object().Get()
	}

	if !transaction.Equal(act, exp) {
		t.Fatal("expected", exp, "got", act)
	}
}

type testStruct struct {
	hsh *Object[common.Hash]
}

func (t testStruct) Object() *Object[common.Hash] {
	return t.hsh
}
