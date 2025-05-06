```bash
#  cat /proc/self/mountinfo | grep "memory"
 cat /proc/self/mountinfo | grep "memory"
1492 1482 0:40 / /sys/fs/cgroup/memory rw,nosuid,nodev,noexec,relatime - cgroup cgroup rw,memory
```
这里的`/sys/fs/cgroup/memory`缺乏写权限，所以硬编码是无法写入的
```bash
(base) hids@ubuntu:~/tinydocker$ ls -ld /sys/fs/cgroup/memory
dr-xr-xr-x 5 root root 0 May  2 19:54 /sys/fs/cgroup/memory
```


```bash
(base) hids@ubuntu:~/tinydocker/cgroups/subsystems$ cat /proc/self/mountinfo | grep "/sys/fs/cgroup"
33 24 0:28 / /sys/fs/cgroup ro,nosuid,nodev,noexec shared:9 - tmpfs tmpfs ro,mode=755,inode64
34 33 0:29 / /sys/fs/cgroup/unified rw,nosuid,nodev,noexec,relatime shared:10 - cgroup2 cgroup2 rw,nsdelegate
35 33 0:30 / /sys/fs/cgroup/systemd rw,nosuid,nodev,noexec,relatime shared:11 - cgroup cgroup rw,xattr,name=systemd
38 33 0:33 / /sys/fs/cgroup/cpuset rw,nosuid,nodev,noexec,relatime shared:15 - cgroup cgroup rw,cpuset
39 33 0:34 / /sys/fs/cgroup/misc rw,nosuid,nodev,noexec,relatime shared:16 - cgroup cgroup rw,misc
40 33 0:35 / /sys/fs/cgroup/cpu,cpuacct rw,nosuid,nodev,noexec,relatime shared:17 - cgroup cgroup rw,cpu,cpuacct
41 33 0:36 / /sys/fs/cgroup/freezer rw,nosuid,nodev,noexec,relatime shared:18 - cgroup cgroup rw,freezer
42 33 0:37 / /sys/fs/cgroup/rdma rw,nosuid,nodev,noexec,relatime shared:19 - cgroup cgroup rw,rdma
43 33 0:38 / /sys/fs/cgroup/pids rw,nosuid,nodev,noexec,relatime shared:20 - cgroup cgroup rw,pids
44 33 0:39 / /sys/fs/cgroup/net_cls,net_prio rw,nosuid,nodev,noexec,relatime shared:21 - cgroup cgroup rw,net_cls,net_prio
45 33 0:40 / /sys/fs/cgroup/memory rw,nosuid,nodev,noexec,relatime shared:22 - cgroup cgroup rw,memory
46 33 0:41 / /sys/fs/cgroup/blkio rw,nosuid,nodev,noexec,relatime shared:23 - cgroup cgroup rw,blkio
47 33 0:42 / /sys/fs/cgroup/hugetlb rw,nosuid,nodev,noexec,relatime shared:24 - cgroup cgroup rw,hugetlb
48 33 0:43 / /sys/fs/cgroup/devices rw,nosuid,nodev,noexec,relatime shared:25 - cgroup cgroup rw,devices
49 33 0:44 / /sys/fs/cgroup/perf_event rw,nosuid,nodev,noexec,relatime shared:26 - cgroup cgroup rw,perf_event
```
上述输出的规律，按空格划分，第5列是挂载点的根目录,最后一列是挂载点的挂载选项对应子系统的名称（rw,subsystem）