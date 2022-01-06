ifneq ($(KERNELRELEASE),)
mymodule-objs := hello_module.o
obj-m := hello_module.o
 
else
PWD := $(shell pwd)
KVER ?= $(shell uname -r)
KDIR := /usr/src/linux-headers-4.19.0-amd64-desktop

all:
	(MAKE) -C $(KDIR) M=$(PWD)
clean:
	m -rf .*.cmd *.o *.mod.c *.ko .tmp_versions
endif
