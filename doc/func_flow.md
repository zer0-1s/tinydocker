```
宿主机
|
|-- tinydocker run /bin/sh
    |
    |-- fork 出来新的进程
         |
         |-- 执行自己 tinydocker init /bin/sh
             |
             |-- initCommand:
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
```