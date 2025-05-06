```
宿主机
|
|-- tinydocker run -it /bin/sh (父进程)
    |
    |-- fork/exec /proc/self/exe init /bin/sh (子进程)
         |
         |-- 执行自己 tinydocker init /bin/sh
             |
             |-- init参数执行initCommand:
                 |-- mount /proc
                 |-- syscall.Exec("/bin/sh")
                     |
                     |-- [容器环境的 PID 1：/bin/sh]

```

```
root@ubuntu:/home/hids/tinydocker# ./tinydocker run -it /bin/sh
Init container
# ps -ef
ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 04:50 pts/7    00:00:00 /bin/sh
root           6       1  0 04:50 pts/7    00:00:00 ps -ef

由于没有适用chroot或者pivot_root，所以子进程直接继承了父进程的文件系统
hids@ubuntu:~/tinydocker$ sudo ./tinydocker run -it /bin/ls
[sudo] password for hids: 
Init container
cmd  doc  go.mod  go.sum  main.go  Makefile  tinydocker
hids@ubuntu:~/tinydocker$ ls
cmd  doc  go.mod  go.sum  main.go  Makefile  tinydocker
```

虽然做了 "进程隔离"（用的是 CLONE_NEWPID 等 namespace）和 "挂载隔离"（重新挂了 /proc），
但还没有做的是：

文件系统（rootfs）隔离 👉 即还没有做 chroot / pivot_root 切换根目录

控制组（cgroup）限制 👉 比如 CPU、内存资源的限制

网络隔离的真正配置 👉 只是新建了 netns，但还没配置 veth pair、桥接之类的网络设备
### key points
- /proc/self/exe：调用自身 init 命令，初始化容器环境,因为/proc/self/exe（是个符号链接）实际上是当前进程可执行文件的路径。
- tty：实现交互
- Namespace 隔离：通过在 fork 时指定对应 Cloneflags+ Unshareflags来实现创建新 Namespace
- proc 隔离：通过重新 mount /proc 文件系统来实现进程信息隔离
- execve 系统调用：使用指定进程覆盖 init 进程



```bash
sudo ./tinydocker run --memory 256m --cpu-cfs-quota 20000 --cpu-share 512 --tty --interactive /bin/bash
```

```
root@ubuntu:/sys/fs/cgroup/cpu,cpuacct# ls
cgroup.clone_children  cpuacct.stat       cpuacct.usage_percpu       cpuacct.usage_sys   cpu.cfs_period_us  cpu.shares  notify_on_release  tasks
cgroup.procs           cpuacct.usage      cpuacct.usage_percpu_sys   cpuacct.usage_user  cpu.cfs_quota_us   cpu.stat    release_agent      tinydocker
cgroup.sane_behavior   cpuacct.usage_all  cpuacct.usage_percpu_user  cpu.cfs_burst_us    cpu.idle           init.scope  system.slice       user.slice
```