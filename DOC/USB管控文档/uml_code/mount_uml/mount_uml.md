```
sequenceDiagram
mount -> do_mount
do_mount -> path_mount
path_mount <-> sb_mount: seucurity_sb_mount的回调
path_mount -> path_mount : do_remount
path_mount -> path_mount : do_new_mount
UOS_sb_mount->sb_mount: hook_mananager

```