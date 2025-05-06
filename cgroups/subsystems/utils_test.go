package subsystems

import (
	// "os"
	// "path"
	"testing"

)



func TestGetCgroupPath(t *testing.T) {
	// subsystem := "cpu" // 使用的子系统，可以根据实际情况修改
	// cgroupPath := "tinytest"
	// expectedPath := path.Join(findCgroupMountPoint(subsystem), cgroupPath)

	// // 测试：autoCreate 为 true，且路径不存在，期望成功创建
	// t.Run("AutoCreateTrueAndPathNotExist", func(t *testing.T) {
	// 	defer os.RemoveAll(expectedPath) // 清理测试创建的 Cgroup 路径
	// 	resultPath, err := GetCgroupPath(subsystem, cgroupPath, true)
	// 	if err != nil {
	// 		t.Fatalf("Expected no error, got %v", err)
	// 	}
	// 	if resultPath != expectedPath {
	// 		t.Fatalf("Expected path %s, got %s", expectedPath, resultPath)
	// 	}
	// 	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
	// 		t.Fatal("Expected created path to exist")
	// 	}
	// })

	// // 测试：autoCreate 为 true，且路径已存在，期望成功返回路径
	// t.Run("AutoCreateTrueAndPathExist", func(t *testing.T) {
	// 	if err := os.Mkdir(expectedPath, RWX); err != nil {
	// 		t.Fatalf("Failed to create test path: %v", err)
	// 	}
	// 	resultPath, err := GetCgroupPath(subsystem, cgroupPath, true)
	// 	if err != nil {
	// 		t.Fatalf("Expected no error, got %v", err)
	// 	}
	// 	if resultPath != expectedPath {
	// 		t.Fatalf("Expected path %s, got %s", expectedPath, resultPath)
	// 	}
	// })

	// // 测试：autoCreate 为 false，期望返回路径而不创建
	// t.Run("AutoCreateFalse", func(t *testing.T) {
	// 	resultPath, err := GetCgroupPath(subsystem, cgroupPath, false)
	// 	if err != nil {
	// 		t.Fatalf("Expected no error, got %v", err)
	// 	}
	// 	if resultPath != expectedPath {
	// 		t.Fatalf("Expected path %s, got %s", expectedPath, resultPath)
	// 	}
	// })
}