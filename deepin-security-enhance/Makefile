PROG_SECURITY_ENHANCE_DAEMON=deepin-security-enhance

PROG_SECURITY_ENHANCE_SUBCOMMAND_ALERTD=alertd
PROG_SECURITY_ENHANCE_SUBCOMMAND_LOGIN_COUNTS=login_counts

PROG_SECURITY_ENHANCE_SCRIPT=security_enhance.sh
PROG_SECADM_CHOWN_SCRIPT=secadm_chown.sh

PROG_FILE_CONF=com.deepin.daemon.SecurityEnhance.conf
PROG_FILE_DBUS_SERVICE=com.deepin.daemon.SecurityEnhance.service
PROG_FILE_SYSTEMD=deepin-security-enhance.service
PROG_FILE_SEPARATION_OF_POWERS=security-enhance.service

DIR_CONF=misc/conf
DIR_SYSTEMD=misc/systemd
DIR_SECURITY_ENHANCE=cmd/deepin-security-enhance-daemon
DIR_SUBCOMMAND=misc/subcommand
DIR_SCRIPT_FILE=misc/script
DIR_CONFIGURATION_FILE_COPY=misc/copy
DIR_ENHANCE_STATE=misc/state
DIR_LABEL=misc/label

PREFIX=/usr
PWD=$(shell pwd)
GOCODE=/usr/share/gocode
GOPATH_DIR=gopath
GOPKG_PREFIX=deepin-security-enhance
export GO111MODULE=off

all: build

prepare:
		@if [ ! -d ${GOPATH_DIR}/src/${GOPKG_PREFIX} ]; then \
				mkdir -p ${GOPATH_DIR}/src/${GOPKG_PREFIX}; \
				ln -sf ${PWD}/pkg ${GOPATH_DIR}/src/${GOPKG_PREFIX}; \
		fi
		chmod 0755 ${DIR_SCRIPT_FILE}/*
		chmod 0700 ${DIR_ENHANCE_STATE}/*
		chmod 0700 ${DIR_LABEL}/*

build: prepare
		env GOPATH=${PWD}/${GOPATH_DIR}:${GOCODE} go build -o ${PWD}/${PROG_SECURITY_ENHANCE_DAEMON} ./${DIR_SECURITY_ENHANCE}

install:
		mkdir -pv ${DESTDIR}${PREFIX}/bin/
		mkdir -pv ${DESTDIR}${PREFIX}/sbin/
		mkdir -pv ${DESTDIR}${PREFIX}/share/dbus-1/system.d/
		mkdir -pv ${DESTDIR}/lib/systemd/system/
		mkdir -pv ${DESTDIR}/usr/share/dbus-1/system-services/
		mkdir -pv ${DESTDIR}/etc/deepin-security/config/
		mkdir -pv ${DESTDIR}/etc/deepin-security/script/
		mkdir -pv ${DESTDIR}/etc/deepin-security/state/
		mkdir -pv ${DESTDIR}/etc/deepin-security/label/
		cp -rf ${PROG_SECURITY_ENHANCE_DAEMON} ${DESTDIR}${PREFIX}/sbin/
		cp -rf ${DIR_CONF}/${PROG_FILE_CONF} ${DESTDIR}${PREFIX}/share/dbus-1/system.d/
		cp -rf ${DIR_SYSTEMD}/${PROG_FILE_DBUS_SERVICE} ${DESTDIR}${PREFIX}/share/dbus-1/system-services/
		cp -rf ${DIR_SYSTEMD}/${PROG_FILE_SYSTEMD} ${DESTDIR}/lib/systemd/system/
		cp -rf ${DIR_SYSTEMD}/${PROG_FILE_SEPARATION_OF_POWERS} ${DESTDIR}/lib/systemd/system/
		cp -rf ${DIR_SUBCOMMAND}/* ${DESTDIR}${PREFIX}/bin/
		cp -rf ${DIR_CONFIGURATION_FILE_COPY}/* ${DESTDIR}/etc/deepin-security/config/
		cp -rf ${DIR_SCRIPT_FILE}/* ${DESTDIR}/etc/deepin-security/script/
		cp -rf ${DIR_ENHANCE_STATE}/* ${DESTDIR}/etc/deepin-security/state/
		cp -rf ${DIR_LABEL}/* ${DESTDIR}/etc/deepin-security/label/

uninstall:
		rm ${DESTDIR}${PREFIX}/sbin/${PROG_SECURITY_ENHANCE_DAEMON}
		rm ${DESTDIR}${PREFIX}/share/dbus-1/system.d/${PROG_FILE_CONF}
		rm ${DESTDIR}${PREFIX}/share/dbus-1/system-services/${PROG_FILE_DBUS_SERVICE}
		rm ${DESTDIR}/lib/systemd/system/${PROG_FILE_SYSTEMD}
		rm ${DESTDIR}/lib/systemd/system/${PROG_FILE_SEPARATION_OF_POWERS}
		rm ${DESTDIR}${PREFIX}/bin/${PROG_SECURITY_ENHANCE_SUBCOMMAND_ALERTD}
		rm ${DESTDIR}${PREFIX}/bin/${PROG_SECURITY_ENHANCE_SUBCOMMAND_LOGIN_COUNTS}
		rm ${DESTDIR}/etc/deepin-security/ -rf

clean:
		rm -rf ${GOPATH_DIR} ${PROG_SECURITY_ENHANCE_DAEMON}

rebuild: clean build
