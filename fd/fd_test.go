package fd

import "testing"

func TestIsSocket(t *testing.T) {
	var file = File{
		Fd:   "7",
		Link: "'socket:[869857]'",
	}
	if !file.IsSocket() {
		t.Fail()
	}
}
func TestInode(t *testing.T) {
	var file = File{
		Fd:   "7",
		Link: "'socket:[869857]'",
	}
	if !file.IsSocket() {
		t.Fail()
	}
	inode, err := file.Inode()
	println(inode)
	if err != nil {
		t.Fail()
	}
	if inode != "869857" {
		t.Fail()
	}
}
