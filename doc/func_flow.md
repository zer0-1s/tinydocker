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
                 |-- 挂载 /proc
                 |-- syscall.Exec("/bin/sh")
                     |
                     |-- [容器环境的 PID 1：/bin/sh]

```

