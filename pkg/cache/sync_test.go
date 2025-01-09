package cache

import (
	"testing"
)

func Test_Cache_Sync_Create_And_Escape(t *testing.T) {
	createAndEscape(t, NewSync[int]())
}

func Test_Cache_Sync_Lifecycle(t *testing.T) {
	lifecycle(t, NewSync[int]())
}

func Test_Cache_Sync_Ranger(t *testing.T) {
	ranger(t, NewSync[int]())
}
