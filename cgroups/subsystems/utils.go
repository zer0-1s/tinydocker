package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"github.com/pkg/errors"
)

// 传入cpu,memory等控制器的名称，返回对应的挂载点
func findCgroupMountPoint(subsystem string) string {
	// 通过 cat /proc/self/mountinfo 命令查看当前进程的挂载信息
	file, err := os.Open("/proc/self/mountinfo")
	if err!= nil {
		fmt.Println("Open mountinfo failed:", err)
		return ""
	}
	defer file.Close()
	// 读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        line := scanner.Text()
        
        // 检查每一行是否包含 "/sys/fs/cgroup"
        if strings.Contains(line, CgroupMountPoint) {
            // 提取挂载点信息
            fields := strings.Split(line, " ")
			subsystemPath := fields[MountInfoIndex]
			subsystems := strings.Split(fields[len(fields)-1], ",")
			for _, s := range subsystems {
				if s == subsystem {
					// 根据最后的字段判断是否是当前的控制器
					// 返回对应的挂载点根目录，例如 "/sys/fs/cgroup/memory"
					return subsystemPath
				}
			}
        }
    }
	if err := scanner.Err(); err != nil {
		fmt.Println("Scan mountinfo failed:", err)	
		return ""
	}
	return ""
}


func GetCgroupPath(subsystem string,cgroupPath string,autoCreate bool) (string, error) {
	subsystemPath := findCgroupMountPoint(subsystem)
	// 将 cgroup 挂载点和用户提供的 cgroup 路径拼接成绝对路径 absPath
	absPath :=path.Join(subsystemPath, cgroupPath)
	if !autoCreate {
		return absPath, nil
	}
	// 指定自动创建时才判断是否存在
	_, err := os.Stat(absPath)
	// 只有不存在才创建
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(absPath, RWX)
		return absPath, err
	}
	// 其他错误或者没有错误都直接返回，如果err=nil,那么errors.Wrap(err, "")也会是nil
	return absPath, errors.Wrap(err, "create cgroup")
}
