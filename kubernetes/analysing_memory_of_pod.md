When you check the running container and navigate to the folder path `/sys/fs/cgroup/memory`, 
you obtain all the memory details of the container. In this directory, 
you can find memory metrics such as usage, limits, cache, and so on.

You can exec into the Running pod and then check the following

```
/ $ cd /sys/fs/cgroup/memory/
/sys/fs/cgroup/memory $ cat memory.stat
cache 49627136
rss 59940864
rss_huge 2097152
mapped_file 1196032
swap 0
pgpgin 49075
pgpgout 23347
pgfault 51653
pgmajfault 54
inactive_anon 0
active_anon 59895808
inactive_file 35459072
active_file 14168064
unevictable 0
hierarchical_memory_limit 134217728
hierarchical_memsw_limit 134217728
total_cache 49627136
total_rss 59940864
total_rss_huge 2097152
total_mapped_file 1196032
total_swap 0
total_pgpgin 0
total_pgpgout 0
total_pgfault 0
total_pgmajfault 0
total_inactive_anon 0
total_active_anon 59895808
total_inactive_file 35459072
total_active_file 14168064
total_unevictable 0
/sys/fs/cgroup/memory $
```
