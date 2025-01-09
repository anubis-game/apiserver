package cache

import (
	"testing"
)

func Test_Cache_Data_Create_And_Escape(t *testing.T) {
	createAndEscape(t, NewData[int]())
}

func Test_Cache_Data_Lifecycle(t *testing.T) {
	lifecycle(t, NewData[int]())
}

func Test_Cache_Data_Ranger(t *testing.T) {
	ranger(t, NewData[int]())
}
