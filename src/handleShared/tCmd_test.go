package handleShared

import (
	"testing"
)

func TestHandleSharedCmdOk(t *testing.T) {
	//初始化参数
	length := 22
	data := []int{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	serNum := 0xFF

	if HandleSharedCmdOk(length, data, serNum) == nil {
		t.Log("---> testHandleSharedCmdOk err.")
	}
}

func TestHandleSharedCmdError(t *testing.T) {
	//初始化参数
	length := 22
	data := []int{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	serNum := 0xFF

	if HandleSharedCmdError(length, data, serNum) == nil {
		t.Log("---> testHandleSharedCmdOk err.")
	}
}

