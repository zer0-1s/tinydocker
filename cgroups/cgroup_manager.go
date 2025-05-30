package cgroups

import (
	"tinydocker/cgroups/subsystems"
	"fmt"
)

type CgroupManager struct {
	// cgroup在hierarchy中的路径 相当于创建的cgroup目录相对于root cgroup目录的路径
	Path string
	// 资源配置
	Resource *subsystems.ResourceConfig
}

// Set 设置cgroup资源限制
func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		err := subSysIns.Set(c.Path, res)
		if err != nil {
			fmt.Printf("apply subsystem:%s err:%s\n", subSysIns.Name(), err)
		}
	}
	return nil
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

// 具体实现interface方法
// Apply 将进程pid加入到这个cgroup中
func (c *CgroupManager) Apply(pid int, res *subsystems.ResourceConfig) error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		err := subSysIns.Apply(c.Path, pid, res)
		if err != nil {
			fmt.Printf("apply subsystem:%s err:%s\n", subSysIns.Name(), err)
		}
	}
	return nil
}



// Destroy 释放cgroup
func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			fmt.Printf("remove cgroup subsystem:%s err:%s\n", subSysIns.Name(), err)
		}
	}
	return nil
}