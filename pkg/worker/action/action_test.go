package action

import (
	"testing"
	"time"

	"github.com/anubis-game/apiserver/pkg/worker/record"
	"github.com/google/uuid"
)

func Test_Worker_Action_Hash(t *testing.T) {
	var act Interface
	{
		act = New(testConfig{})
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Sta().Get() != record.Unknown {
		t.Fatal("expected", record.Unknown, "got", act.Rec().Sta().Get())
	}

	{
		act.Rec().Sta().Set(record.Created)
	}

	if act.Rec().Sta().Get() != record.Created {
		t.Fatal("expected", record.Created, "got", act.Rec().Sta().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Sta().Get() != record.Unknown {
		t.Fatal("expected", record.Unknown, "got", act.Rec().Sta().Get())
	}

	{
		act.Rec().Sta().Set(record.Waiting)
	}

	if act.Rec().Sta().Get() != record.Waiting {
		t.Fatal("expected", record.Waiting, "got", act.Rec().Sta().Get())
	}

	if act.Rec().Get(0).Sta().Get() != record.Created {
		t.Fatal("expected", record.Created, "got", act.Rec().Get(0).Sta().Get())
	}

	if act.Rec().Get(1).Sta().Get() != record.Waiting {
		t.Fatal("expected", record.Waiting, "got", act.Rec().Get(1).Sta().Get())
	}
}

func Test_Worker_Action_Wait(t *testing.T) {
	var act Interface
	{
		act = New(testConfig{})
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 5*time.Second {
		t.Fatal("expected", 5*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 10*time.Second {
		t.Fatal("expected", 10*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 15*time.Second {
		t.Fatal("expected", 15*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 20*time.Second {
		t.Fatal("expected", 20*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 25*time.Second {
		t.Fatal("expected", 25*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 30*time.Second {
		t.Fatal("expected", 30*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 35*time.Second {
		t.Fatal("expected", 35*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 40*time.Second {
		t.Fatal("expected", 40*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 45*time.Second {
		t.Fatal("expected", 45*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 50*time.Second {
		t.Fatal("expected", 50*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 55*time.Second {
		t.Fatal("expected", 55*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 60*time.Second {
		t.Fatal("expected", 60*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 60*time.Second {
		t.Fatal("expected", 60*time.Second, "got", act.Rec().Wai().Get())
	}

	{
		act.Rec().Add()
	}

	if act.Rec().Wai().Get() != 60*time.Second {
		t.Fatal("expected", 60*time.Second, "got", act.Rec().Wai().Get())
	}
}

type testConfig struct{}

func (t testConfig) Arg() []byte {
	return nil
}

func (t testConfig) Rec() record.Interface {
	return record.NewSlicer(record.SlicerConfig{})
}

func (t testConfig) Typ() string {
	return ""
}

func (t testConfig) Uid() uuid.UUID {
	return uuid.New()
}
