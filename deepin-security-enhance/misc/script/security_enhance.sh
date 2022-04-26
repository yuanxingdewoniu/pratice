#!/bin/bash

#set -e
###################################################
###
### README：运行此脚本之前，先确保做以下操作已完成：
### 1、创建sysadm、secadm、audadm三大管理员用户
### 2、创建/etc/deepin-security、/etc/deepin-security/config和/etc/deepin-security/script目录
### 3、将配置文件都拷贝到/etc/deepin-security/config目录，将secadm_chown.sh（权限改为0755）拷贝到/etc/deepin-security/script目录
### 4、将login_counts和alertd文件拷贝到/usr/bin目录
### 5、将security-enhance.service文件拷贝到/usr/lib/systemd/system目录
### 6、本脚本自己（权限改为0755）也拷贝到拷贝到/etc/deepin-security/script目录下

###################################################
### dbus开关命令：
### 开启：dbus-send --system --print-reply --dest=com.deepin.daemon.SecurityEnhance  /com/deepin/daemon/SecurityEnhance com.deepin.daemon.SecurityEnhance.Enable boolean:true boolean:true/false string:sysadmpasswd string:secadmpasswd string:audadmpasswd
### 关闭：dbus-send --system --print-reply --dest=com.deepin.daemon.SecurityEnhance  /com/deepin/daemon/SecurityEnhance com.deepin.daemon.SecurityEnhance.Enable boolean:false boolean:true/false string:  string:  string: 
### 查询：dbus-send --system --print-reply --dest=com.deepin.daemon.SecurityEnhance  /com/deepin/daemon/SecurityEnhance com.deepin.daemon.SecurityEnhance.Status

###################################################
### 执行此脚本参数定义：
### $1:enable/disable 开启/关闭等保三级
### $2:true/false 是否删除已有管理员账号，关闭等保三级时有效，开启等保三级时忽略此参数
### $3:系统管理员密码
### $4:安全管理员密码
### $5:审计管理员密码
### $6:是否开启等保四级,false:不开启,true:开启    (预留功能，暂不支持)

ROOT_USER_ID=`id | grep "uid=0(root)"`
DEEPIN_SECURITY_DIR="/etc/deepin-security"
DS_BACKUP_DIR=${DEEPIN_SECURITY_DIR}"/back_up"
DS_CONFIG_DIR=${DEEPIN_SECURITY_DIR}"/config"
DS_SCRIPT_DIR=${DEEPIN_SECURITY_DIR}"/script"
SELINUX_DEFAULT_POLICY_FILE="/etc/selinux/default/policy/policy.*"
SELINUX_DEFAULT_MLS_FILE="/etc/selinux/mls/policy/policy.*"
GRUB_CFG_DIR="/etc/default/grub.d"
GRUB_CMDLINE_SELINUX_FILE="01_deepin_selinux.cfg"
UNLABELED_FLAG_FILE=".unlabeled.tmp"
SYSTEMD_SERVICE_DIR="/lib/systemd/system"
SECURITY_ENHANCE_SERVICE="security-enhance.service"
AUDITD_SERVICE="auditd.service"
SECURITY_ENHANCE_SERVICE_PATH=$SYSTEMD_SERVICE_DIR"/"$SECURITY_ENHANCE_SERVICE
AUDITD_SERVICE_PATH=$SYSTEMD_SERVICE_DIR"/"$AUDITD_SERVICE
SYS_ADM="sysadm"
SEC_ADM="secadm"
AUD_ADM="audadm"

# 三大管理员的加密密码，只会在enable的时候用到
SYS_PASSWD=$3
SEC_PASSWD=$4
AUD_PASSWD=$5

# 是否开启等保四级 (预留功能，暂不支持)
#SECURITY_LEVEL_FOUR=$6
SECURITY_LEVEL_FOUR="false"

function back_up_config()
{
	if [ -f /etc/profile ];then
		cp /etc/profile ${DS_BACKUP_DIR}/etc_profile > /dev/null 2>&1
	fi

	if [ -f /etc/security/pwquality.conf ];then
		cp /etc/security/pwquality.conf ${DS_BACKUP_DIR}/pam_pwquality.conf > /dev/null 2>&1
	fi

	if [ -f /etc/pam.d/common-password ];then
		cp /etc/pam.d/common-password ${DS_BACKUP_DIR}/pam_common-password > /dev/null 2>&1
	fi

	# 先注释掉，定制服务器版本的时候可能没有上层应用来设置，需要在此脚本里设置
#	if [ -f /etc/deepin/dde.conf ];then
#		cp /etc/deepin/dde.conf ${DS_BACKUP_DIR}/pam_dde.conf > /dev/null 2>&1
#	fi

	if [ -f /etc/pam.d/sshd ];then
		cp /etc/pam.d/sshd ${DS_BACKUP_DIR}/pam_sshd > /dev/null 2>&1
	fi

	if [ -f /etc/pam.d/login ];then
		cp /etc/pam.d/login ${DS_BACKUP_DIR}/pam_login > /dev/null 2>&1
	fi

	if [ -f /etc/pam.d/lightdm ];then
		cp /etc/pam.d/lightdm ${DS_BACKUP_DIR}/pam_lightdm > /dev/null 2>&1
	fi

	if [ -f /etc/systemd/logind.conf ];then
		cp /etc/systemd/logind.conf ${DS_BACKUP_DIR}/system_logind.conf > /dev/null 2>&1
	fi

	if [ -f /etc/selinux/config ];then
		cp /etc/selinux/config ${DS_BACKUP_DIR}/selinux_config > /dev/null 2>&1
	fi

	if [ -f /etc/bash.bashrc ];then
		cp /etc/bash.bashrc ${DS_BACKUP_DIR}/bash.bashrc > /dev/null 2>&1
	fi 

	if [ $SECURITY_LEVEL_FOUR = "true" ];then
		#wkhtmltopdf(等保四级要求)
		if [ -f /usr/bin/wkhtmltopdf ];then
			cp /usr/bin/wkhtmltopdf ${DS_BACKUP_DIR}/wkhtmltopdf > /dev/null 2>&1
		fi

		#sysrq(等保四级要求)
		if [ -f /etc/sysctl.conf ];then
			cp /etc/sysctl.conf ${DS_BACKUP_DIR}/sysctl.conf > /dev/null 2>&1
		fi
	fi
}

function restore_config()
{
	if [ -f ${DS_BACKUP_DIR}/etc_profile ];then
		cp ${DS_BACKUP_DIR}/etc_profile /etc/profile > /dev/null 2>&1
	fi

	if [ -f ${DS_BACKUP_DIR}/pam_pwquality.conf ];then
		cp ${DS_BACKUP_DIR}/pam_pwquality.conf /etc/security/pwquality.conf > /dev/null 2>&1
	fi

	if [ -f ${DS_BACKUP_DIR}/pam_common-password ];then
		cp ${DS_BACKUP_DIR}/pam_common-password /etc/pam.d/common-password > /dev/null 2>&1
	fi

	# 先注释掉，定制服务器版本的时候可能没有上层应用来设置，需要在此脚本里设置
#	if [ -f ${DS_BACKUP_DIR}/pam_dde.conf ];then
#		cp ${DS_BACKUP_DIR}/pam_dde.conf /etc/deepin/dde.conf > /dev/null 2>&1
#	fi

	if [ -f ${DS_BACKUP_DIR}/pam_sshd ];then
		cp ${DS_BACKUP_DIR}/pam_sshd /etc/pam.d/sshd > /dev/null 2>&1
	fi

	if [ -f ${DS_BACKUP_DIR}/pam_login ];then
		cp ${DS_BACKUP_DIR}/pam_login /etc/pam.d/login > /dev/null 2>&1
	fi

	if [ -f ${DS_BACKUP_DIR}/pam_lightdm ];then
		cp ${DS_BACKUP_DIR}/pam_lightdm /etc/pam.d/lightdm > /dev/null 2>&1
	fi

	if [ -f ${DS_BACKUP_DIR}/system_logind.conf ];then
		cp ${DS_BACKUP_DIR}/system_logind.conf /etc/systemd/logind.conf > /dev/null 2>&1
	fi

	if [ -f ${DS_BACKUP_DIR}/selinux_config ];then
		cp ${DS_BACKUP_DIR}/selinux_config /etc/selinux/config > /dev/null 2>&1
	fi

	if [ -f ${DS_BACKUP_DIR}/bash.bashrc ];then
		cp ${DS_BACKUP_DIR}/bash.bashrc /etc/bash.bashrc > /dev/null 2>&1
	fi

	if [ $SECURITY_LEVEL_FOUR = "true" ];then
		#wkhtmltopdf(等保四级要求)
		if [ -f ${DS_BACKUP_DIR}/wkhtmltopdf ];then
			cp ${DS_BACKUP_DIR}/wkhtmltopdf /usr/bin/wkhtmltopdf > /dev/null 2>&1
		fi

		#sysrq(等保四级要求)
		if [ -f ${DS_BACKUP_DIR}/sysctl.conf ];then
			cp ${DS_BACKUP_DIR}/sysctl.conf /etc/sysctl.conf > /dev/null 2>&1
		fi
	fi

	if [ -d ${DS_BACKUP_DIR} ];then
		rm -rf ${DS_BACKUP_DIR} > /dev/null 2>&1
	fi
}

function restore_root()
{
	# 去除SELinux自动打标签以及还未打标签标志
	if [ -f /.autorelabel ];then
		rm /.autorelabel > /dev/null 2>&1
	fi
	if [ -f ${DEEPIN_SECURITY_DIR}/${UNLABELED_FLAG_FILE} ];then
		rm ${DEEPIN_SECURITY_DIR}/${UNLABELED_FLAG_FILE} > /dev/null 2>&1
	fi

	# 确保是在default模式下
	sed -i 's/^SELINUXTYPE=.*/SELINUXTYPE=default/g' /etc/selinux/config

	semanage login -d $SYS_ADM > /dev/null 2>&1
	semanage login -d $SEC_ADM > /dev/null 2>&1
	semanage login -d $AUD_ADM > /dev/null 2>&1
	semanage login -m -s unconfined_u root > /dev/null 2>&1

	if [ $SECURITY_LEVEL_FOUR = "true" ];then
		# 确保是在mls模式下
		sed -i 's/^SELINUXTYPE=.*/SELINUXTYPE=mls/g' /etc/selinux/config
		semanage login -d $SYS_ADM > /dev/null 2>&1
		semanage login -d $SEC_ADM > /dev/null 2>&1
		semanage login -d $AUD_ADM > /dev/null 2>&1
		semanage login -m -s unconfined_u root > /dev/null 2>&1
	fi

	usermod -U root > /dev/null 2>&1
	chown -R 0:0 /var/lib/selinux > /dev/null 2>&1
	chown -R 0:0 /var/lib/sepolgen > /dev/null 2>&1
	chown -R 0:0 /etc/selinux > /dev/null 2>&1
	chown -R 0:0 /usr/share/selinux > /dev/null 2>&1
	chmod -s /sbin/setfiles > /dev/null 2>&1
	chmod -s /sbin/restorecon > /dev/null 2>&1
	chmod -s /sbin/semodule > /dev/null 2>&1

	chown -R 0:0 /etc/audit > /dev/null 2>&1
	chown -R 0:0 /etc/audisp > /dev/null 2>&1

	if [ -f $SECURITY_ENHANCE_SERVICE_PATH ];then
		systemctl disable $SECURITY_ENHANCE_SERVICE > /dev/null 2>&1
	fi

	# 关闭审计服务
	if [ -f $AUDITD_SERVICE_PATH ];then
		systemctl disable $AUDITD_SERVICE > /dev/null 2>&1
	fi
}

function delete_account()
{
	# 检查管理员是否存在
	if  id -u $SEC_ADM > /dev/null 2>&1;then
		usermod -s /usr/sbin/nologin -L $SEC_ADM > /dev/null 2>&1
	fi
	if  id -u $AUD_ADM > /dev/null 2>&1;then
		usermod -G $AUD_ADM -s /usr/sbin/nologin -L $AUD_ADM > /dev/null 2>&1
	fi

	if [ $1 = "true" ];then
		# 删除管理员
		if id -u $SEC_ADM > /dev/null 2>&1;then
			userdel -r $SEC_ADM > /dev/null 2>&1
		fi
		if id -u $AUD_ADM > /dev/null 2>&1;then
			userdel -r $AUD_ADM > /dev/null 2>&1
		fi
	fi
}

function exit_and_restore()
{
	restore_root
	delete_account "true"

	# 配置还原必须放最后，保证配置还原后不会再误修改
	restore_config
	exit 0
}

# 是否使用了特权运行此脚本
if [ -z "$ROOT_USER_ID" ];then
	echo -n '-1'
	exit 0
fi 

# 检查安全目录
if [ ! -d ${DEEPIN_SECURITY_DIR} ];then
	echo -n '-2'
	exit 0
fi

if [ ! -d ${DS_BACKUP_DIR} ];then
	mkdir ${DS_BACKUP_DIR}
fi

if [ ! -d ${DS_CONFIG_DIR} ];then
	echo -n '-3'
	exit 0
fi

if [ ! -d ${DS_SCRIPT_DIR} ];then
	echo -n '-4'
	exit 0
fi

case "$1" in
	enable)
		# 将需要修改的的配置文件备份
		back_up_config

		################ 身份鉴别 ################
		## 强化口令
		DEEPIN_PW_CHECK=`dpkg -l | grep deepin-pw-check`
		PAM_PWQUALITY_PKG_CHECK_SUCCESS=`dpkg -l | grep "ii  libpam-pwquality"`
		PAM_CRACKLIB_PKG_CHECK_SUCCESS=`dpkg -l | grep "ii  libpam-cracklib"`

		#先查看是否有pam_deepin_pw_check模块，如果有，则由上层调用pam模块接口配置，如果没有，则使用pam_pwquality.so，在此脚本里手动配置
		if [ -z "$DEEPIN_PW_CHECK" ];then
			if [ -z "$PAM_PWQUALITY_PKG_CHECK_SUCCESS" ];then
				if [ -z "$PAM_CRACKLIB_PKG_CHECK_SUCCESS" ];then
					echo -n '-101'
					exit_and_restore
				fi

				if [ -f ${DS_CONFIG_DIR}/pam_common-password_cracklib ];then
					cp ${DS_CONFIG_DIR}/pam_common-password_cracklib /etc/pam.d/common-password > /dev/null 2>&1
					if [ $? != 0 ];then
						echo -n '-102'
						exit_and_restore
					fi
				else
					#sed -i '/pam_cracklib.so/d' /etc/pam.d/common-password
					#sed -i '25ipassword required pam_cracklib.so retry=3 minlen=8' /etc/pam.d/common-password
					echo -n '-103'
					exit_and_restore
				fi
			else
				if [ -f ${DS_CONFIG_DIR}/pam_common-password_pwquality ];then
					cp ${DS_CONFIG_DIR}/pam_common-password_pwquality /etc/pam.d/common-password > /dev/null 2>&1
					if [ $? != 0 ];then
						echo -n '-104'
						exit_and_restore
					fi
				else
					#sed -i '/pam_pwquality.so/d' /etc/pam.d/common-password
					#sed -i '25ipassword     requisite       pam_pwquality.so retry=3' /etc/pam.d/common-password
					echo -n '-105'
					exit_and_restore
				fi

				if [ -f ${DS_CONFIG_DIR}/pam_pwquality.conf ];then
					cp ${DS_CONFIG_DIR}/pam_pwquality.conf /etc/security/pwquality.conf > /dev/null 2>&1
					if [ $? != 0 ];then
						echo -n '-106'
						exit_and_restore
					fi
				else
					#sed -i 's/.*minlen.*=.*/minlen = 8/g' /etc/security/pwquality.conf
					#sed -i 's/.*minclass.*/minclass = 3/g' /etc/security/pwquality.conf
					#sed -i 's/.*maxsequence.*/maxsequence = 3/g' /etc/security/pwquality.conf
					#sed -i 's/.*maxrepeat.*/maxrepeat = 3/g' /etc/security/pwquality.conf
					echo -n '-107'
					exit_and_restore
				fi
			fi
#(先注释掉，定制服务器版本的时候可能没有上层应用来设置，需要在此脚本里设置)
#		else
#			if [ -f ${DS_CONFIG_DIR}/pam_dde.conf ];then
#				cp ${DS_CONFIG_DIR}/pam_dde.conf /etc/deepin/dde.conf > /dev/null 2>&1
#				if [ $? != 0 ];then
#					echo -n '-108'
#					exit_and_restore
#				fi
#			else
#				#sed -i 's/^PASSWORD_MIN_LENGTH =.*/PASSWORD_MIN_LENGTH=8/g' /etc/deepin/dde.conf
#				#sed -i 's/^CONSECUTIVE_SAME_CHARACTER_NUM =.*/CONSECUTIVE_SAME_CHARACTER_NUM=4/g' /etc/deepin/dde.conf
#				#sed -i 's/^VALIDATE_REQUIRED =.*/VALIDATE_REQUIRED=3/g' /etc/deepin/dde.conf
#				#sed -i 's/^WORD_CHECK =.*/WORD_CHECK = 1/g' /etc/deepin/dde.conf
#				echo -n '-109'
#				exit_and_restore
#			fi
		fi

		## 连续登录一定次失败后锁定20分钟
		DEEPIN_AUTHENTICATE=`dpkg -l | grep deepin-authenticate`

		#先查看是否有pam_deepin_authentication.so模块，如果有，则由上层调用pam_deepin_authentication接口自行配置，如果没有，则使用pam_tally2.so，在此脚本里手动配置
		if [ -z "$DEEPIN_AUTHENTICATE" ];then
			if [ -f ${DS_CONFIG_DIR}/pam_sshd ];then
				cp ${DS_CONFIG_DIR}/pam_sshd /etc/pam.d/sshd > /dev/null 2>&1
				if [ $? != 0 ];then
					echo -n '-112'
					exit_and_restore
				fi
			else
				#sed -i '/pam_tally2.so/d' /etc/pam.d/sshd
				#sed -i '2iauth     required       pam_tally2.so deny=3 unlock_time=1200' /etc/pam.d/sshd
				echo -n '-113'
				exit_and_restore
			fi

			if [ -f ${DS_CONFIG_DIR}/pam_login ];then
				cp ${DS_CONFIG_DIR}/pam_login /etc/pam.d/login > /dev/null 2>&1
				if [ $? != 0 ];then
					echo -n '-114'
					exit_and_restore
				fi
			else
				#sed -i '/pam_tally2.so/d' /etc/pam.d/login
				#sed -i '2iauth     required       pam_tally2.so deny=3 unlock_time=1200' /etc/pam.d/login
				echo -n '-115'
				exit_and_restore
			fi

			if [ -f ${DS_CONFIG_DIR}/pam_lightdm ];then
				cp ${DS_CONFIG_DIR}/pam_lightdm /etc/pam.d/lightdm > /dev/null 2>&1
				if [ $? != 0 ];then
					echo -n '-116'
					exit_and_restore
				fi
			else
				#sed -i '/pam_tally2.so/d' /etc/pam.d/lightdm
				#sed -i '2iauth     required       pam_tally2.so deny=3 unlock_time=1200' /etc/pam.d/lightdm
				echo -n '-117'
				exit_and_restore
			fi
		#(先注释掉，定制服务器版本的时候可能没有上层应用来设置，需要在此脚本里设置)
#		else
		fi

		################ 自主访问控制 ################
		# linux本身就满足，不需要加固

		################ 标记和强制访问控制 ################

		# 创建系统管理员并设置密码,传下来的密码是经过加密的，故要加-p参数
		if ! id -u $SYS_ADM > /dev/null 2>&1;then
			useradd -m $SYS_ADM -s /bin/bash -p $SYS_PASSWD > /dev/null 2>&1
		else
			usermod $SYS_ADM -s /bin/bash -p $SYS_PASSWD > /dev/null 2>&1
		fi
		if [ $? != 0 ];then
			echo -n '-301'
			exit_and_restore
		fi

		# 创建安全管理员并设置密码,传下来的密码是经过加密的，故要加-p参数
		if ! id -u $SEC_ADM > /dev/null 2>&1;then
			useradd -m $SEC_ADM -s /bin/bash -p $SEC_PASSWD > /dev/null 2>&1
		else
			usermod $SEC_ADM -s /bin/bash -p $SEC_PASSWD > /dev/null 2>&1
		fi
		if [ $? != 0 ];then
			echo -n '-302'
			exit_and_restore
		fi

		# 创建审计管理员并设置密码,传下来的密码是经过加密的，故要加-p参数
		if ! id -u $AUD_ADM > /dev/null 2>&1;then
			useradd -m $AUD_ADM -s /bin/bash -p $AUD_PASSWD > /dev/null 2>&1
		else
			usermod $AUD_ADM -s /bin/bash -p $AUD_PASSWD > /dev/null 2>&1
		fi
		if [ $? != 0 ];then
			echo -n '-303'
			exit_and_restore
		fi

		# 检查是否安装了selinux policy包
		if [ ! -f ${SELINUX_DEFAULT_POLICY_FILE} ];then
			echo -n '-304'
			exit_and_restore
		fi

		#设置打印水印 (等保四级要求)
		if [ $SECURITY_LEVEL_FOUR = "true" ];then
			if [ ! -d /usr/lib/cups/filter/ ];then
		    	mkdir -p /usr/lib/cups/filter > /dev/null 2>&1
			fi 
	    
			if [ -f ${DS_CONFIG_DIR}/watermark.py ];then
				cp ${DS_CONFIG_DIR}/watermark.py /usr/lib/cups/filter/watermarkpdf > /dev/null 2>&1
				if [ $? != 0 ];then
					echo -n '-324'
					exit_and_restore
				fi
			else
				echo -n '-325'
				exit_and_restore
			fi
			chmod 755 /usr/lib/cups/filter/watermarkpdf > /dev/null 2>&1

			#cp -R /usr/share/deepin-security/watermark/home/* /home/

			if [ -f /usr/bin/wkhtmltopdf ];then
				mv /usr/bin/wkhtmltopdf /usr/bin/wkhtmltopdf.old
			fi

			if [ -f ${DS_CONFIG_DIR}/wkhtmltopdf ];then
				cp ${DS_CONFIG_DIR}/wkhtmltopdf  /usr/bin/wkhtmltopdf > /dev/null 2>&1
				if [ $? != 0 ];then
					echo -n '-326'
					exit_and_restore
				fi
			else
				echo -n '-327'
				exit_and_restore
			fi
			chmod +x /usr/bin/wkhtmltopdf > /dev/null 2>&1
		fi

		#保证管理员目录下文件的属主
		if [ -d /home/${SYS_ADM} ];then
			chown -R $SYS_ADM /home/${SYS_ADM} > /dev/null 2>&1
			chmod o+rx /home/${SYS_ADM} > /dev/null 2>&1
		fi
		if [ -d /home/${SEC_ADM} ];then
			chown -R $SEC_ADM /home/${SEC_ADM} > /dev/null 2>&1
			chmod o+rx /home/${SEC_ADM} > /dev/null 2>&1
		fi
		if [ -d /home/${AUD_ADM} ];then
			chown -R $AUD_ADM /home/${AUD_ADM} > /dev/null 2>&1
			chmod o+rx /home/${AUD_ADM} > /dev/null 2>&1
		fi

		################ 安全审计 ################
		# 检查是否安装了auditd
		AUDITD_PKG_CHECK_SUCCESS=`dpkg -l | grep "ii  auditd"`
		if [ -z "$AUDITD_PKG_CHECK_SUCCESS" ];then
			echo -n '-401'
			exit_and_restore
		fi

		# 设置应用权限
		if [ -f /sbin/auditctl ];then
			setcap "cap_audit_write,cap_audit_read,cap_audit_control,cap_sys_nice,cap_dac_override+ep" /sbin/auditctl > /dev/null 2>&1
		fi

		if [ -f /sbin/auditd ];then
			setcap "cap_audit_write,cap_audit_read,cap_audit_control,cap_sys_nice,cap_dac_override+ep" /sbin/auditd > /dev/null 2>&1
		fi

		# 开启审计服务
		if [ -f $AUDITD_SERVICE_PATH ];then
			systemctl enable $AUDITD_SERVICE > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-402'
				exit_and_restore
			fi
		else
			echo -n '-403'
			exit_and_restore
		fi

		################ 数据完整性 ################
		# 本章节均为功能要求，且依赖硬件，无需也无法本加固脚本配置
		#


		################ 数据保密性 ################
		# 本章节均为功能要求，且依赖硬件以及内核配置，无需也无法通过本加固脚本配置
		#


		################ 网络安全保护 ################
		# 本章节均为功能要求，只要提供功能就行，具体配置没有要求，故不需要也无法给出默认配置
		# 需要由系统管理员根据实际环境配置防火墙等
		#


		################ 运行安全保护 ################
		# 启用sysrq,这是等保四级要求
		if [ $SECURITY_LEVEL_FOUR = "true" ];then
			if [ -f /etc/sysctl.conf ];then
				sed -i '/kernel.sysrq=/d' /etc/sysctl.conf
				echo "kernel.sysrq=1" >> /etc/sysctl.conf

				sed -i '/vm.panic_on_oom=/d' /etc/sysctl.conf
				echo "vm.panic_on_oom=1" >> /etc/sysctl.conf

				sed -i '/kernel.panic=/d' /etc/sysctl.conf
				echo "kernel.panic=10" >> /etc/sysctl.conf  
			else
				echo -n '-801'
				exit_and_restore
			fi
		fi

		# 限制控制台使用（仅提供机制，需要用户自己根据用户环境的实际情况和需求来配置）

		################ 资源利用 ################
		#本章节均为机制要求，只要提供机制就行，具体配置没有要求（也没法要求，因为这完全是根据用户自己的实际需要来设置的），故不需要也无法给出默认配置
		#比如：/etc/security/limits.conf可以给每个用户限定分配的资源，但是因系统硬件资源以及用户实际需要的资源无法预知，故不可能给出合理的默认配置

		################ 用户登陆访问控制 ################
		# 限制系统只开启6个tty
		if [ -f ${DS_CONFIG_DIR}/system_logind.conf ];then
			cp ${DS_CONFIG_DIR}/system_logind.conf /etc/systemd/logind.conf > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-1001'
				exit_and_restore
			fi
		else
			#sed -i 's/^#NAutoVTs=6/NAutoVTs=6/g' /etc/systemd/logind.conf
			echo -n '-1002'
			exit_and_restore
		fi

		# 打印用户上次登录的信息以及登陆成功/失败次数统计(login、sshd文件已在已在身份鉴别阶段一起配置好)
		#sed -i 's/pam_lastlog.so/pam_lastlog.so showfailed/g' /etc/pam.d/login
		#sed -i 'N;42isession    optional     pam_lastlog.so showfailed' /etc/pam.d/sshd

		if [ ! -f /etc/.lasts ];then
			touch /etc/.lasts > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-1003'
				exit_and_restore
			fi
		fi
		if [ ! -f /etc/.lastsb ];then
			touch /etc/.lastsb > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-1004'
				exit_and_restore
			fi
		fi
		chmod 0666 /etc/.lasts > /dev/null 2>&1
		chmod 0666 /etc/.lastsb > /dev/null 2>&1
		if [ -f /var/log/auth.log ];then
			chmod 0644 /var/log/auth.log > /dev/null 2>&1
		fi

		if [ -f /etc/bash.bashrc ];then
			sed -i '/login_counts/d' /etc/bash.bashrc
			echo "login_counts" >> /etc/bash.bashrc
		fi 

		# 超时自动退出
		if [ -f /etc/profile ];then
			sed -i '/TMOUT=/d' /etc/profile
			echo "TMOUT=120" >> /etc/profile
		else
			echo -n '-1005'
			exit_and_restore
		fi

		################ 可信度量 ################
		#需要硬件支持，无法配置

		################ 开启SELinux ################
		# 默认启用selinux
		if [ -f ${DS_CONFIG_DIR}/selinux_config ];then
			cp ${DS_CONFIG_DIR}/selinux_config /etc/selinux/config > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-315'
				exit_and_restore
			fi
		else
			#sed -i 's/^SELINUX=permissive/SELINUX=enforcing/g' /etc/selinux/config
			echo -n '-316'
			exit_and_restore
		fi

		# selinux三权分立配置
		str_user="$SYS_ADM $SEC_ADM $AUD_ADM"
		for i in ${str_user}
		do
			semanage login -a -s ${i}_u ${i} > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-305'
				exit_and_restore
			fi
		done
		semanage login -m -s user_u __default__ > /dev/null 2>&1
		semanage login -m -s system_u root > /dev/null 2>&1

		#若开启的是等保四级，则进入mls模式
		if [ $SECURITY_LEVEL_FOUR = "true" ];then
			#若开启的是等保四级，则进入mls模式
			sed -i 's/SELINUX=enforcing/SELINUX=permissive/g' /etc/selinux/config
			sed -i 's/SELINUXTYPE=default/SELINUXTYPE=mls/g' /etc/selinux/config

			###selinux三权分立配置
			mls_user="$SYS_ADM $SEC_ADM"
			for i in ${mls_user}
			do
				semanage login -a -s ${i}_u ${i} > /dev/null 2>&1
				if [ $? != 0 ];then
					echo -n '-328'
					exit_and_restore
				fi
			done
			semanage login -a -s ${AUD_ADM}_u -r s15:c0.c1023-s15:c0.c1023 $AUD_ADM > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-329'
				exit_and_restore
			fi
			semanage login -m -s user_u __default__ > /dev/null 2>&1
			semanage login -m -s system_u root > /dev/null 2>&1
		fi

		# 系统管理员dac权限配置
		usermod $SYS_ADM -aG sudo > /dev/null 2>&1
		usermod $SYS_ADM -aG 0 > /dev/null 2>&1
		if [ $? != 0 ];then
			echo -n '-306'
			exit_and_restore
		fi

		# 安全管理员dac权限配置
		if [ -d /var/lib/selinux ];then
			chown -R ${SEC_ADM}:${SEC_ADM} /var/lib/selinux > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-307'
				exit_and_restore
			fi
		fi
		if [ -d /var/lib/sepolgen ];then
			chown -R ${SEC_ADM}:${SEC_ADM} /var/lib/sepolgen > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-308'
				exit_and_restore
			fi
		fi
		if [ -d /etc/selinux ];then
			chown -R ${SEC_ADM}:${SEC_ADM} /etc/selinux > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-309'
				exit_and_restore
			fi
		fi
		if [ -d /usr/share/selinux ];then
			chown -R ${SEC_ADM}:${SEC_ADM} /usr/share/selinux > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-310'
				exit_and_restore
			fi
		fi

		if [ -f /sbin/setfiles ];then
			chmod +s /sbin/setfiles > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-311'
				exit_and_restore
			fi
		fi
		if [ -f /sbin/restorecon ];then
			chmod +s /sbin/restorecon > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-330'
				exit_and_restore
			fi
		fi
		if [ -f /sbin/semodule ];then
			chmod +s /sbin/semodule > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-332'
				exit_and_restore
			fi
		fi

		# 审计管理员dac权限配置
		usermod $AUD_ADM -aG sudo,adm > /dev/null 2>&1
		if [ $? != 0 ];then
			echo -n '-312'
			exit_and_restore
		fi
		if [ -d /etc/audit ];then
			chown -R 0:adm /etc/audit > /dev/null 2>&1
			chmod -R g+w /etc/audit > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-313'
				exit_and_restore
			fi
		fi
		if [ -d /etc/audisp ];then
			chown -R 0:adm  /etc/audisp > /dev/null 2>&1
			chmod -R g+w /etc/audisp > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-314'
				exit_and_restore
			fi
		fi
		#chmod 0770 /var/log/audit > /dev/null 2>&1
		#chmod 0660 /var/log/audit/* > /dev/null 2>&1

		# 锁定原root用户
		usermod -L root > /dev/null 2>&1

		# 确保安全管理员和审计管理员每次开机都拥有正确的dac权限
		if [ -f $SECURITY_ENHANCE_SERVICE_PATH ];then
			systemctl enable $SECURITY_ENHANCE_SERVICE > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-317'
				exit_and_restore
			fi
		else
			echo -n '-318'
			exit_and_restore
		fi

		# SELinux自动打标签
		fixfiles -F onboot > /dev/null 2>&1
		if [ $? != 0 ];then
			echo -n '-319'
			exit_and_restore
		fi

		if [ ! -f ${DEEPIN_SECURITY_DIR}/${UNLABELED_FLAG_FILE} ];then
			touch ${DEEPIN_SECURITY_DIR}/${UNLABELED_FLAG_FILE} > /dev/null 2>&1
		fi

		# 修改grub
		if [ -d ${GRUB_CFG_DIR} ];then
			echo "GRUB_CMDLINE_LINUX_DEFAULT=\"\$GRUB_CMDLINE_LINUX_DEFAULT security=selinux checkreqprot=1 enforcing=0\"" | tee ${GRUB_CFG_DIR}/${GRUB_CMDLINE_SELINUX_FILE} > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-320'
				exit_and_restore
			fi
			update-grub > /dev/null 2>&1
			if [ $? != 0 ];then
				rm ${GRUB_CFG_DIR}/${GRUB_CMDLINE_SELINUX_FILE} > /dev/null 2>&1
				echo -n '-321'
				exit_and_restore
			fi
		fi

		# 去除开发者模式开关的"i"属性，以便打标签能成功
		chattr -R -i /var/lib/deepin > /dev/null 2>&1
	;;

	disable)
		################ 关闭SELinux ################
		SELINUX_STATUS=`getenforce`
		if [ $SELINUX_STATUS = "Enforcing" ];then
			setenforce 0 > /dev/null 2>&1
			if [ $? != 0 ];then
				echo -n '-322'
				exit 0
			fi
		fi

		# 修改grub
		rm ${GRUB_CFG_DIR}/${GRUB_CMDLINE_SELINUX_FILE} > /dev/null 2>&1
		update-grub > /dev/null 2>&1
		if [ $? != 0 ];then
			echo "GRUB_CMDLINE_LINUX_DEFAULT=\"\$GRUB_CMDLINE_LINUX_DEFAULT security=selinux checkreqprot=1 enforcing=0\"" | tee ${GRUB_CFG_DIR}/${GRUB_CMDLINE_SELINUX_FILE} > /dev/null 2>&1
			if [ $SELINUX_STATUS = "Enforcing" ];then
				setenforce 1 > /dev/null 2>&1
			fi
			echo -n '-323'
			exit 0
		fi

		# 此时,SELinux已经关闭了，后续只是还原一些配置，没必要再返错了
		restore_root

		delete_account $2

		# 恢复开发者模式开关的"i"属性
		chattr +i /var/lib/deepin/deepin_security_verify.whitelist > /dev/null 2>&1
		chattr +i /var/lib/deepin/developer-mode/enabled > /dev/null 2>&1

		# 配置还原必须放最后，保证配置还原后不会再误修改（除了下一步的特例）
		restore_config

		# 强制disable SELinux功能，防止用户原来的配置里是开启的
		if [ -f /etc/selinux/config ];then
			sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config
			sed -i 's/SELINUX=permissive/SELINUX=disabled/g' /etc/selinux/config
		fi
	;;
esac

echo -n 0
