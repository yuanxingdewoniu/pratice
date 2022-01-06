#!/bin/bash

###########################################3
#
#  此脚本可根据日志文件自动生成.te
#  参数1：日志文件
#  参数2：模块名

if [ -f $1 ];then
    cat $1 | audit2allow -m $2 > $2.te
    checkmodule -M -m -o $2.mod $2.te 
    semodule_package -o $2.pp -m $2.mod 
fi


