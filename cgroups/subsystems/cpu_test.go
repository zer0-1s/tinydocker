package subsystems

import (
	"os"
	"path"
	"strconv"
	"testing"
	"fmt"
	"strings"
)

func TestCpuSubSystem(t *testing.T) {
	// 临时目录以隔离测试环境
	cpuSubSys := &CpuSubSystem{}
	cgroupPath := "test_cgroup"

	// 确保Cgroup的根目录存在
	if err := os.MkdirAll(path.Join(cpuSubSys.Name(), cgroupPath), 0755); err != nil {
		t.Fatalf("failed to create cgroup directory: %v", err)
	}

	
	res := &ResourceConfig{
		CpuCfsQuota: 10,  // 设定为 20% 的 CPU 配额
		CpuShare:     "512", // 设定 CPU 共享值
	}

	t.Run("SetMethod", func(t *testing.T) {
		if err := cpuSubSys.Set(cgroupPath, res); err != nil {
			t.Fatalf("Set method failed: %v", err)
		}

		// 检查文件的创建和内容
		quotaFile := path.Join(CgroupMountPoint,cpuSubSys.Name(),cgroupPath, "cpu.cfs_quota_us")
		shareFile := path.Join(CgroupMountPoint,cpuSubSys.Name(),cgroupPath, "cpu.shares")

		quotaContent, err := os.ReadFile(quotaFile)
		fmt.Println("[" + string(quotaContent) + "]") // 打印看看真实内容
        if err != nil {
            t.Fatalf("Failed to read cpu.cfs_quota_us: %v", err)
        }

        expectedQuota := (PeriodDefault / Percent * res.CpuCfsQuota)
		if strings.TrimSpace(string(quotaContent)) != strconv.Itoa(expectedQuota) {
			t.Errorf("Expected cpu.cfs_quota_us to be %d, got %s", expectedQuota, quotaContent)
		}
        shareContent, err := os.ReadFile(shareFile)
        if err != nil {
            t.Fatalf("Failed to read cpu.shares: %v", err)
        }
        
		if strings.TrimSpace(string(shareContent)) != res.CpuShare {
			t.Errorf("Expected cpu.shares to be %s, got %s", res.CpuShare, shareContent)
		}
    })
}