#!/bin/bash
SELINUX_STATUS=`getenforce`
SELINUX_ETC="/etc/selinux"
SELINUX_FS="/sys/fs/selinux"
SELINUX_LIB="/var/lib/selinux"
SELINUX_SHARE="/usr/share/selinux"
DEEPIN_SECURITY_DIR="/etc/deepin-security"
UNLABELED_FLAG_FILE=".unlabeled.tmp"
GRUB_CFG_DIR="/etc/default/grub.d"
GRUB_CMDLINE_SELINUX_FILE="01_deepin_selinux.cfg"
SEC_ADM="secadm"
ARCH_MIPS=0
ARCH_LOONGSON=0
NEED_EXECMEM=0

if [ -f ${DEEPIN_SECURITY_DIR}/${UNLABELED_FLAG_FILE} ]; then
    sed -i 's/ enforcing=0//g' ${GRUB_CFG_DIR}/${GRUB_CMDLINE_SELINUX_FILE} > /dev/null 2>&1
    update-grub > /dev/null 2>&1
    if [ $? = 0 ];then
        rm ${DEEPIN_SECURITY_DIR}/${UNLABELED_FLAG_FILE} > /dev/null 2>&1
    fi
    setenforce 1 > /dev/null 2>&1
fi
if [ ! $SELINUX_STATUS = "Disabled" ];then
    # For loongson
    CPU_ARCH=`uname -m`
    if [ $CPU_ARCH = "mips64" ];then
        NEED_EXECMEM=1
        ARCH_MIPS=1
    fi
    if [ $CPU_ARCH = "loongarch64" ];then
        NEED_EXECMEM=1
        ARCH_LOONGSON=1
    fi
    setsebool -P domain_can_execmem $NEED_EXECMEM > /dev/null 2>&1
    setsebool -P arch_is_mips $ARCH_MIPS > /dev/null 2>&1
    setsebool -P arch_is_loongson $ARCH_LOONGSON > /dev/null 2>&1
    if id -u $SEC_ADM > /dev/null 2>&1;then
        chown -R ${SEC_ADM}:${SEC_ADM} $SELINUX_ETC
        chown -R ${SEC_ADM}:${SEC_ADM} $SELINUX_FS
        chown -R ${SEC_ADM}:${SEC_ADM} $SELINUX_LIB
        chown -R ${SEC_ADM}:${SEC_ADM} $SELINUX_SHARE
    fi
    #setenforce 1 > /dev/null 2>&1
fi
