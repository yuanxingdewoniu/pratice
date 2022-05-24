#include <libudev.h>
#include <stdio.h>
#include <string.h>
 
/*
 * libudev api说明
 * https://mirrors.edge.kernel.org/pub/linux/utils/kernel/hotplug/libudev/ch01.html
 */
 
 
/**
 * 打印/dev/xxx设备节点的properties值
 * @devnode: 设备节点文件/dev/xxx,如:/dev/ttyUSB0
 */

struct removable_storage_device_info{
    char* devNode;   // 设备节点  // DEVNAME
    char* name;      // 设备名称  ID_FS_LABEL_ENC
    char* time;      // 添加时间，格式为：“yyyy-MM-dd HH:mm:ss”add time
    char* serialNum; // 序列号 ID_SERIAL_SHORT
    char* vendorID;  //  产品id ID_VENDOR_ID
    char* productID; //厂商 id  ID_MODEL_ID
    long int size;       // 设备容量，单位byte
    int permMode;     // 访问控制规则，0 - 可读可写，1 - 只读，2 - 禁止
    char* extend;
};


int get_devnode_properties(const char *devnode,struct removable_storage_device_info *info)
{
	int ret = 0;
	struct udev *udev;
	struct udev_device *device;
	struct udev_enumerate *enumerate;
	struct udev_list_entry *first_entry;
	struct udev_list_entry *list_entry;
	const char *syspath;
 
	udev = udev_new();
	
	enumerate = udev_enumerate_new(udev);
 
	/*
	 * 通过枚举器添加匹配条件,如果有多条匹配条件
	 */
	//udev_enumerate_add_match_property(enumerate, "DEVNAME", devnode);

  printf("%s  %d  match =  %d  \n",devnode,  udev_enumerate_add_match_property(enumerate,"DEVNAME",devnode),udev_enumerate_add_match_property(enumerate, "ID_USB_DRIVER", "usb-storage"));
/*
   if (udev_enumerate_add_match_property(enumerate, "DEVNAME", devnode)  &&   udev_enumerate_add_match_property(enumerate, "ID_USB_DRIVER", "usb-storage")) */

  
/*
	 * 根据枚举器设置的条件扫描所有的设备.
	 * 注意: 如果有多条匹配条件,符合其中一条就会被扫描到,
	 *       匹配条件越多,扫描越宽松
	 */
	udev_enumerate_scan_devices(enumerate);
 
	list_entry = udev_enumerate_get_list_entry(enumerate);
 
	if (!list_entry) {
		ret = -1;
		goto ERROR1;
	}
 
	syspath = udev_list_entry_get_name(list_entry);
	device = udev_device_new_from_syspath(udev, syspath);
 
	if (!device) {
		ret = -1;
		goto ERROR1;
	} 
 
	first_entry = udev_device_get_properties_list_entry(device);
 
	if (!first_entry) {
		ret = -1;
		goto ERROR2;
	}
 
	udev_list_entry_foreach(list_entry, first_entry) {
		printf("%s = %s\n", udev_list_entry_get_name(list_entry), udev_list_entry_get_value(list_entry));
		 if( strcmp(udev_list_entry_get_name(list_entry) ,"DEVNAME")== 0)  {
		 printf("############%s = %s#################\n", udev_list_entry_get_name(list_entry), udev_list_entry_get_value(list_entry));
		 info->devNode=udev_list_entry_get_value(list_entry);
         }


          if(strcmp(udev_list_entry_get_name(list_entry),"ID_SERIAL_SHORT")==0) {
		    printf("^^^^^^^^^^^^%s = %s^^^^^^^^^^^^^\n", udev_list_entry_get_name(list_entry), udev_list_entry_get_value(list_entry));
 			info->serialNum=udev_list_entry_get_value(list_entry);
         }

      }
 
ERROR2:	
	udev_device_unref(device);
ERROR1:
	udev_enumerate_unref(enumerate);
	udev_unref(udev);
	return ret;
}
 
 
void test_print_devnode_properties(void)
{



/*
	 * 异常情况测试
	 */
     struct removable_storage_device_info device_info;
     

     get_devnode_properties("/dev/sda",&device_info);

     get_devnode_properties("/dev/wsda",&device_info);

     //printf("device_info->devNode = %s, device_name= %s\n",device_info.devNode,device_info.name);


     get_devnode_properties("/dev/sdb",&device_info);
     printf("device_info->devNode = %s, device_name= %s\n",device_info.devNode,device_info.name);
}

int  main(void)
{
    test_print_devnode_properties();
    return 0;
}
