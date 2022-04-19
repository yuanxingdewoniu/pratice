## 1.文件防护需求清单

【功能介绍】该功能主要是保护文件被恶意程序删除或篡改，基于系统安全提供的基于应用的访问控制接口，提供给安全中心及dde组策略进行相关的图形化配置，本接口兼容的版本为桌面专业版、服务器企业版。具体功能如下：
#### 1 可将某文件/目录加入到保护清单，/usr、/opt、/boot、/dev、/proc、/run、/sys、/tmp目录及其子目录下的文件不可加入（2月24日修改）。加入文件保护后默认不可被任意用户/程序删除、修改、重命名和移动路径，包括root用户。将目录加入到保护清单后，即该目录及其下所有文件都继承此设定，不能被任意用户/程序删除、修改、重命名和移动路径。系统下所有文件和目录均能加入保护，包括系统文件和用户文件。
#### 2  系统更新不受影响，可以正常读写加入保护清单的文件/目录。(直接放过)
#### 3  可对加入保护清单的文件设置例外原则，允许该文件/目录能被指定应用进行读写，允许指定多个应用，除该应用外其他所有应用和用户均不能删除、修改、重命名和移动路径。（可以进行更新和升级，其他操作不允许，*允许被指定应用读写*，需要给应用设置新的标签）
###  4 加入保护的文件，被不允许读写的用户或应用进行修改、删除、重命名、移动路径时，统一进行提示“该文件已被加入保护，无权限修改，请联系系统管理员”，包括在文管或命令行进行修改均需提示。
####  5 默认保护清单为空。
####  6 该接口可被安全中心、DDE组策略、域管进行调用。  

* 

  方法名称：设置保护文件目录

  接口定义：setFileProtect(In: protect_file_path, Out: err_code)
  
  传入参数：要保护文件的绝对路径
  
  返回参数：错误码信息
  
  
  
  方法名称：取消文件保护目录
  
  接口定义：setFileProtect(In: protect_file_path, Out: err_code)
  
  传入参数：取消保护文件的绝对路径
  
  返回参数：错误码信息
  
  
  
  
  
  方法名称：设置例外进程
  
  接口定义：SetProcessException(In: exec_file_path, Out: err_code) 
  
  传入参数：需要设置例外策略的可执行文件路径
  
  传回参数：err_code错误码， 0成功， 其他失败
  
  
  
  方法名称：取消例外访问进程
  
  接口定义：SetProcessException(In: exec_file_path, Out: err_code) 
  
  传入参数：取消的例外策略的可执行文件路径
  
  传回参数：err_code错误码， 0成功， 其他失败

####  7 错误码需要在内核、libc库里增加，统一给到上层应用，且需要支持国际化。【使用角色】未开启等保下，仅限系统管理员使用；已开启等保，仅限安全管理员使用【前置条件】有管理员权限才可调用

####  *【流程说明】暂无【相关方】安全中心---调用此接口，提供图形化管控界面文件管理器---涉及权限修改和对用户的提示

## 2.文档防护的实现逻辑

* 开启等保,同时使用原有的标签.
* 给需要管控的文档打只读标签，同时针对更新进程和应用商店，赋予写的权限
* 给文件夹授权，相应的域赋予目录的对应权限
* 对于其他的文件,访问不控制.

## 3. 测试用例

用例的主要部分
```
policy_module(file_protect, 1.0.0)

########################################
#
# Declarations
#
type only_read_t;
type read_write_t;
type file_protect_t;

require {
        attribute domain;
        type deepin_domain_exec_t;
        type apt_exec_t;
}


#################read file ###########################
allow  domain only_read_t:file  {read_file_perms link};
allow {apt_exec_t deepin_domain_exec_t } only_read_t:file {manage_file_perms};

allow  domain read_write_t:file { manage_file_perms};
allow {apt_exec_t deepin_domain_exec_t } read_write_t:file {manage_file_perms};


####################read dir #################################
allow  domain only_read_t:dir {link list_dir_perms}; ##只读权限，可以创建软硬链接
allow {apt_exec_t deepin_domain_exec_t } only_read_t:dir {all_dir_perms}; 

allow  domain read_write_t:dir {list_dir_perms};
allow {apt_exec_t deepin_domain_exec_t } read_write_t:dir {all_dir_perms};


```

## 4. 文件管控DEMO测试
### 4.1 测试过程

* a.测试对文件标签：
```
secadm@uos-PC: setenforce 0 关闭 selinux
```
```
uos@uos-PC:~$ ls -all |grep testonly ##查看文件权限
drwxr-xr-x.  2 uos  uos      4096 4月   6 10:40 testonlyRead
uos@uos-PC:~$ echo 22222 > testonlyRead/2.c 
uos@uos-PC:~$ chcon -t only_read_t  testonlyRead 
uos@uos-PC:~$ ls -lZ |grep testonlyRead/ 修改文件标签
drwxr-xr-x.  2 uos uos user_u:object_r:only_read_t:s0         4096 4月   6 10:41 testonlyRead

```

* b.开启等保
```
secadm@uos-PC: setenforce 1 开启selinux 
```

* c.测试读写删除文件

```
uos@uos-PC:~$ echo 22222222222222222 > testonlyRead/3.c 
-bash: testonlyRead/3.c: 权限不够
uos@uos-PC:~$ cat testonlyRead/1.c
11111111111111111

uos@uos-PC:~$ rm testonlyRead/1.c 
rm: 无法删除'testonlyRead/1.c': 权限不够

```

* d.测试恢复文档标签。
```
uos@uos-PC:~$ restorecon -R testonlyRead/1.c 
uos@uos-PC:~$ ls -lZ testonlyRead/1.c 
-rw-r--r--. 1 uos uos user_u:object_r:user_home_t:s0 18 4月   6 10:40 testonlyRead/1.c # 恢复文件标签为原始标签，父目录被管控，还是无法删除
uos@uos-PC:~$ echo 11111111111111111111 testonlyRead/1.c 
11111111111111111111 testonlyRead/1.c
uos@uos-PC:~$ echo 11111111111111111111 > testonlyRead/1.c  
uos@uos-PC:~$ rm testonlyRead/1.c 
rm: 无法删除'testonlyRead/1.c': 权限不够
uos@uos-PC:~$ rm testonlyRead/1.c -rf 
rm: 无法删除'testonlyRead/1.c': 权限不够
```


```
uos@uos-PC:~$ restorecon -R testonlyRead 
restorecon: Could not set context for /home/uos/testonlyRead:  Permission denied
uos@uos-PC:~$ restorecon -R testonlyRead
uos@uos-PC:~$ rm testonlyRead/1.c -rf 
```

```
uos@uos-PC:~/testonlyRead$ ls
2.c
uos@uos-PC:~/testonlyRead$ ls -LZ 
user_u:object_r:user_home_t:s0 2.c
uos@uos-PC:~/testonlyRead$ echo testwrite > 2.c 
uos@uos-PC:~/testonlyRead$ cd ../
uos@uos-PC:~$ ls -lZ |grep testonlyRead
drwxr-xr-x.  2 uos uos user_u:object_r:user_home_t:s0         4096 4月   6 10:51 testonlyRead
uos@uos-PC:~$ chcon -t only_read_t testonlyRead 
uos@uos-PC:~$ ls -lZ |grep testonlyRead
drwxr-xr-x.  2 uos uos user_u:object_r:only_read_t:s0         4096 4月   6 10:51 testonlyRead
uos@uos-PC:~$ ls -lZ testonlyRead/* 
-rw-r--r--. 1 uos uos user_u:object_r:user_home_t:s0 10 4月   6 11:02 testonlyRead/2.c
uos@uos-PC:~$ echo writedata >  testonlyRead/2.c  


// 关闭selinux 在操作
uos@uos-PC:~$ chcon -t only_read_t testonlyRead/2.c ## setenforce 0 先关闭selinux
// 开启selinux 在操作
uos@uos-PC:~$ echo writedata >  testonlyRead/2.c  
-bash: testonlyRead/2.c: 权限不够
uos@uos-PC:~$ echo 11111111 > testonlyRead/1.c 
-bash: testonlyRead/1.c: 权限不够

//关闭selinux 创建 新的文件，测试可以自动集成父目录的标签
uos@uos-PC:~$ echo 3333333333333 > testonlyRead/3.c 
uos@uos-PC:~$ ls -lZ testonlyRead/*
-rw-r--r--. 1 uos uos user_u:object_r:only_read_t:s0 10 4月   6 11:04 testonlyRead/2.c
-rw-r--r--. 1 uos uos user_u:object_r:only_read_t:s0 14 4月   6 11:05 testonlyRead/3.c
```



* e. 软硬链接读写测试

```
uos@uos-PC:~$ ln testonlyRead/2.c ln2.c 
uos@uos-PC:~$ cat ln2.c 
writedata

// 限制通过硬连接写
uos@uos-PC:~$ echo  writedatatotestonlyreadfold > ln2.c 
-bash: ln2.c: 权限不够
uos@uos-PC:~$ ln  -s testonlyRead/2.c lns2.c  
uos@uos-PC:~$ cat lns2.c 
writedata
// 限制通过软连接写
2016-08-01
uos@uos-PC:~$ echo  writedatatotestonlyreadfold > lns2.c 
-bash: lns2.c: 权限不够
uos@uos-PC:~$ 
```


* f.  测试apt  source拉取项目源码 

```
* 测试apt 拉去项目源码 
提示git 无权限，增加git 的标签，进行测试 注意，很多应用直接bin_t 标签，所以这些应用也会获取权限，所以实际应用，需要定义不同标签，避免给其他应用赋权
```


https://salsa.debian.org/glibc-team/glibc.git
请使用：
git clone https://salsa.debian.org/glibc-team/glibc.git
获得该软件包的最近更新(可能尚未正式发布)。
需要下载 18.2 MB 的源代码包。
错误:1 http://pools.uniontech.com/desktop-professional eagle/main glibc 2.28.19-1+dde (diff)
  无法打开文件 glibc_2.28.19-1+dde.debian.tar.xz - open (13: 权限不够) [IP: 10.0.32.42 80]
错误:2 http://pools.uniontech.com/desktop-professional eagle/main glibc 2.28.19-1+dde (dsc)
  无法打开文件 glibc_2.28.19-1+dde.dsc - open (13: 权限不够) [IP: 10.0.32.42 80]
错误:3 http://pools.uniontech.com/desktop-professional eagle/main glibc 2.28.19-1+dde (tar)
  无法打开文件 glibc_2.28.19.orig.tar.xz - open (13: 权限不够) [IP: 10.0.32.42 80]
W: 用 unlink 删除文件 glibc_2.28.19-1+dde.debian.tar.xz 出错 - PrepareFiles (13: 权限不够)
W: 用 unlink 删除文件 glibc_2.28.19-1+dde.dsc 出错 - PrepareFiles (13: 权限不够)
W: 用 unlink 删除文件


*修改te文件。直接放开，可以更新
```
policy_module(file_protect, 1.0.0)

type only_read_t;
type read_write_t;
type file_protect_t;

require {
        attribute domain;
        type deepin_domain_exec_t;
        type apt_exec_t;
        type bin_t 
        type shell_exec_t;
}

#################read file ###########################
allow  domain only_read_t:file  {read_file_perms link};
allow {apt_exec_t deepin_domain_exec_t bin_t shell_exec_t } only_read_t:file {manage_file_perms};

allow  domain read_write_t:file { manage_file_perms};
allow {apt_exec_t deepin_domain_exec_t  bin_t shell_exec_t} read_write_t:file {manage_file_perms};

####################read dir #################################
allow  domain only_read_t:dir {link list_dir_perms}; ##只读权限，可以创建软硬链接
allow {apt_exec_t deepin_domain_exec_t  bin_t shell_exec_t  } only_read_t:dir {all_dir_perms}; 

allow  domain read_write_t:dir {list_dir_perms};
allow {apt_exec_t deepin_domain_exec_t bin_t shell_exec_t } read_write_t:dir {all_dir_perms};

```



* 测试apt 拉取项目源码 
提示git 无权限，增加git的相关标签(bin_t shell_exec_t)，进行测试 注意，很多应用直接bin_t 标签，所以这些应用也会获取权，所以实际应用，需要定义不同标签，避免给其他应用赋权。

* 修改te文件。直接放开，可以更新

###  4.2 测试结果:


|测试方法|预期结果|实际结果|
|---|---|----|
|测试only_read_t 标签| 可以打标签，只能读，不能删除 |可以打标签，只能读，无法删除，删除报权限错误|
|测试创建硬链接 |正常创建，可以访问，无法修改|正常创建，可以访问，无法修改|
|测试创建软链接正常|正常创建，可以访问，无法修改|正常创建，可以访问，无法修改|
|设置bin_tapt_t 标签的应用对打only_read_t标签文件写操作|可以写|可以写|


##  5. U盘管控问题实现
* man mount 在挂载时候进行处理，进行打标签

* 对U盘打三类标签， 禁止读写 ，只读，读写。

* 根据三类标签进行相应的授权控制。