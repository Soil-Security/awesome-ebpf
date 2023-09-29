# eBPF Code Snippets

``` c
#define LAST_32_BITS(x) x & 0xFFFFFFFF
#define FIRST_32_BITS(x) x >> 32

SEC("tracepoint/syscalls/sys_enter_execve")
int enter_execve(struct execve_entry_args_t *args) {
    u64 pid_tgid = bpf_get_current_pid_tgid();
    u32 pid = LAST_32_BITS(pid_tgid);

    return 0;
}
```

```
git rebase -i $(git merge-base origin/main $(git rev-parse --abbrev-ref HEAD))
```
