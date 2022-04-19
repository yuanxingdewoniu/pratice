
一、breakpoint do_mount 
测试挂载

```
[Switching to Thread 1.2]

Thread 2 hit Breakpoint 11, do_mount (dev_name=0xffff8881f8053eb0 "/dev/loop0", dir_name=0x564264346600 "/home/minicoco/usbtest", type_page=0xffff8881f6078ca0 "vfat", flags=0, 
    data_page=0xffff8881f6146000) at fs/namespace.c:3063
3063    {
(gdb) l
3058     * Therefore, if this magic number is present, it carries no information
3059     * and must be discarded.
3060     */
3061    long do_mount(const char *dev_name, const char __user *dir_name,
3062                    const char *type_page, unsigned long flags, void *data_page)
3063    {
3064            struct path path;
3065            unsigned int mnt_flags = 0, sb_flags;
3066            int retval = 0;
3067
(gdb) 
```

```
(gdb) p (* (char *)data_page@128) 
$65 = "context=\"system_u:object_r:removable_t:s0\"", '\000' <repeats 14 times>, "A\000\000\000\000\000\000\000rw,context=\"system_u:object_r:removable_t:s0\"\000\060\"\000\000\000\000\000\000\000\000!\000\000\000\000\000\000"

(gdb) p (*(char *)data_page@1024)
$9 = "context=\"system_u:object_r:snappy_snap_t:s0\"",
```


```
	0xffffffff8135c707 in selinux_add_opt (token=0, s=0xffff8881ed22e240 "system_u:object_r:removable_t:s0", mnt_opts=0xffff8881f68a9f30) at security/selinux/hooks.c:999
999             if (token == Opt_seclabel)      /* eaten and completely ignored */
(gdb) p Opt_seclabel
$71 = Opt_seclabel
(gdb) ptype  Opt_seclabel
type = enum {Opt_error = -1, Opt_context, Opt_defcontext, Opt_fscontext, Opt_rootcontext, Opt_seclabel}
(gdb) p token
$72 = 0
```

```
      const struct cred *cred = current_cred();
2761
2762            if (flags & MS_REMOUNT)
2763                    return superblock_has_perm(cred, path->dentry->d_sb,
(gdb) l
2764                                               FILESYSTEM__REMOUNT, NULL);
2765            else
2766                    return path_has_perm(cred, path, FILE__MOUNTON);
```




```
2 /* Check whether a task can perform a filesystem operation. */
1963 static int superblock_has_perm(const struct cred *cred,
1964                    struct super_block *sb,
1965                    u32 perms,
1966                    struct common_audit_data *ad)
1967 {
1968     struct superblock_security_struct *sbsec;
1969     u32 sid = cred_sid(cred);
1970
1971     sbsec = sb->s_security;
1972     return avc_has_perm(&selinux_state,
1973                 sid, sbsec->sid, SECCLASS_FILESYSTEM, perms, ad);
1974 }
```
 小结：在do_mount 接口中 先调用security_sb_mount 进行一系列权限的检查，可以挂载的话。先进行挂载，
 security_sb_mount security_sb_remount 钩子 根据当前进程（mount）对路径（挂载点的权限）判断能否挂载。


有权限的基础上，在进行打标签。
目前是挂载的时候传入的参数，设置的标签。
修改标签的过程：
```
long do_mount(const char *dev_name, const char __user *dir_name,
		const char *type_page, unsigned long flags, void *data_page)

	else if (flags & MS_REMOUNT)
		retval = do_remount(&path, flags, sb_flags, mnt_flags,
				    data_page); 打标签
	else if (flags & MS_BIND)
		retval = do_loopback(&path, dev_name, flags & MS_REC);
	else if (flags & (MS_SHARED | MS_PRIVATE | MS_SLAVE | MS_UNBINDABLE))
		retval = do_change_type(&path, flags);
	else if (flags & MS_MOVE)
		retval = do_move_mount_old(&path, dev_name);
	else
		retval = do_new_mount(&path, type_page, sb_flags, mnt_flags,
				      dev_name, data_page); //打标签
dput_out:
	path_put(&path);
```

```
/*
 * change filesystem flags. dir should be a physical root of filesystem.
 * If you've mounted a non-root directory somewhere and want to do remount
 * on it - tough luck.
 */
 static int do_remount(struct path *path, int ms_flags, int sb_flags,
		      int mnt_flags, void *data)
 {
	int err;
	struct super_block *sb = path->mnt->mnt_sb;
	struct mount *mnt = real_mount(path->mnt);
	struct fs_context *fc;

	if (!check_mnt(mnt))
		return -EINVAL;

	if (path->dentry != path->mnt->mnt_root)
		return -EINVAL;

	if (!can_change_locked_flags(mnt, mnt_flags))
		return -EPERM;

	fc = fs_context_for_reconfigure(path->dentry, sb_flags, MS_RMT_MASK);
	if (IS_ERR(fc))
		return PTR_ERR(fc);

	err = parse_monolithic_mount_data(fc, data);
	if (!err) {
		down_write(&sb->s_umount);
		err = -EPERM;
		if (ns_capable(sb->s_user_ns, CAP_SYS_ADMIN)) {
			err = reconfigure_super(fc);
			if (!err)
				set_mount_attributes(mnt, mnt_flags);
		}
		up_write(&sb->s_umount);
	}

	mnt_warn_timestamp_expiry(path, &mnt->mnt);

	put_fs_context(fc);
	return err;
 }
```

```
int parse_monolithic_mount_data(struct fs_context *fc, void *data)
{
	int (*monolithic_mount_data)(struct fs_context *, void *);

	monolithic_mount_data = fc->ops->parse_monolithic;
	if (!monolithic_mount_data)
		monolithic_mount_data = generic_parse_monolithic;
	
	return monolithic_mount_data(fc, data);
}
```

```
int generic_parse_monolithic(struct fs_context *fc, void *data)
{
	char *options = data, *key;
	int ret = 0;

	if (!options)
		return 0;
	
	ret = security_sb_eat_lsm_opts(options, &fc->security);  //option 挂钩？？
	if (ret)
		return ret;
	
	while ((key = strsep(&options, ",")) != NULL) {
		if (*key) {
			size_t v_len = 0;
			char *value = strchr(key, '=');
	
			if (value) {
				if (value == key)
					continue;
				*value++ = 0;
				v_len = strlen(value);
			}
			ret = vfs_parse_fs_string(fc, key, value, v_len);
			if (ret < 0)
				break;
		}
	}
	
	return ret;
}
```

// options 标签
```
static int selinux_sb_eat_lsm_opts(char *options, void **mnt_opts)
{
	char *from = options;
	char *to = options;
	bool first = true;
	int rc;

	while (1) {
		int len = opt_len(from);
		int token;
		char *arg = NULL;
	
		token = match_opt_prefix(from, len, &arg);
	
		if (token != Opt_error) {
			char *p, *q;
	
			/* strip quotes */
			if (arg) {
				for (p = q = arg; p < from + len; p++) {
					char c = *p;
					if (c != '"')
						*q++ = c;
				}
				arg = kmemdup_nul(arg, q - arg, GFP_KERNEL);
				if (!arg) {
					rc = -ENOMEM;
					goto free_opt;
				}
			}
			rc = selinux_add_opt(token, arg, mnt_opts); // ???
			if (unlikely(rc)) {
				kfree(arg);
				goto free_opt;
			}
		} else {
			if (!first) {	// copy with preceding comma
				from--;
				len++;
			}
			if (to != from)
				memmove(to, from, len);
			to += len;
			first = false;
		}
		if (!from[len])
			break;
		from += len + 1;
	}
	*to = '\0';
	return 0;

free_opt:
	if (*mnt_opts) {
		selinux_free_mnt_opts(*mnt_opts);
		*mnt_opts = NULL;
	}
	return rc;
}
```


/*


后续进行USB 存储设备的管控，需要去修改标签。？？？ 限制修改标签的能力。 只有自己应用才有全力去调用修改！！！！！！！

不通过hook劫持挂载，直接读设备的udev信息，匹配管控的名单，修改对应挂载点的label





