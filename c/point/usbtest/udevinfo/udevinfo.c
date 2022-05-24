#include <libudev.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h> 
/*
 * libudev api说明
 * https://mirrors.edge.kernel.org/pub/linux/utils/kernel/hotplug/libudev/ch01.html
 */
 
 
/**
 * 打印/dev/xxx设备节点的properties值
 * @devnode: 设备节点文件/dev/xxx,如:/dev/ttyUSB0
 */
#define false -1;
#define true 0;

typedef struct removable_storage_device_info{
    char* devNode;   // 设备节点  // DEVNAME
    char* name;      // 设备名称  ID_FS_LABEL_ENC
    char* time;      // 添加时间，格式为：“yyyy-MM-dd HH:mm:ss”add time
    char* serialNum; // 序列号 ID_SERIAL_SHORT
    char* vendorID;  //  产品id ID_VENDOR_ID
    char* productID; //厂商 id  ID_MODEL_ID
    long int size;       // 设备容量，单位byte
    int permMode;     // 访问控制规则，0 - 可读可写，1 - 只读，2 - 禁止
    char* extend;
} dev_info;

dev_info get_devnode_properties1(const char *devnode)
{
       //int	ret = false;
	dev_info info;
	struct udev *udev;
	struct udev_device *device;
	struct udev_enumerate *enumerate;
	struct udev_list_entry *first_entry;
	struct udev_list_entry *list_entry;
	const char *syspath;
 
	
	udev = udev_new();
	enumerate = udev_enumerate_new(udev);
	
	/* * 通过枚举器添加匹配条件,如果有多条匹配条件 */

	/*
	 * 根据枚举器设置的条件扫描所有的设备.
	 * 注意: 如果有多条匹配条件,符合其中一条就会被扫描到,
	 *       匹配条件越多,扫描越宽松
	 */
	udev_enumerate_add_match_property(enumerate, "ID_USB_DRIVER", "usb-storage"); //
	udev_enumerate_scan_devices(enumerate);
    
	list_entry = udev_enumerate_get_list_entry(enumerate);
 
	if (!list_entry) {
		//ret = false;
		goto ERROR1;
	}
 
	syspath = udev_list_entry_get_name(list_entry);
	device = udev_device_new_from_syspath(udev, syspath);
 
	if (!device) {
        //ret = false;		
		goto ERROR1;
	} 
 
	first_entry = udev_device_get_properties_list_entry(device);
 
	if (!first_entry) {
	//	ret = false;
		goto ERROR2;
	}

	udev_list_entry_foreach(list_entry, first_entry) {
      //printf("%s = %s\n", udev_list_entry_get_name(list_entry), udev_list_entry_get_value(list_entry));


    if( strcmp(udev_list_entry_get_name(list_entry) ,"DEVNAME")== 0)  {
  
    char *devNode =udev_list_entry_get_value(list_entry);

	
    info.devNode = (char * )malloc(128);
	memset  (info.devNode,0,128);	
	strncpy(info.devNode,devNode,strlen(devNode));
    printf("##info.devNode_address= %p,info.devNode = %s ", info.devNode, info.devNode); 
  }
    if(strcmp(udev_list_entry_get_name(list_entry),"ID_SERIAL_SHORT")==0) {
		   
	//	   info.serialNum=udev_list_entry_get_value(list_entry);
	char *serialNum =udev_list_entry_get_value(list_entry);
    info.serialNum = (char *)malloc(128);
	memset(info.serialNum,0,128);	
	strncpy(info.serialNum,serialNum,strlen(serialNum));
    printf("info_serialNum_address = %p info.serialNum   = %s ##\n", info.serialNum,info.serialNum );
	}
}
 
ERROR2:	
	udev_device_unref(device);
ERROR1:
	udev_enumerate_unref(enumerate);
	udev_unref(udev);
	return info;
} 
 
int  test_print_devnode_properties1(void)
{
	/* * 异常情况测试 */

	dev_info info= get_devnode_properties1("/dev/sda");
	printf("1111111111111RESULT=%p,%p  info.devNode = %s, info.serialNum= %s\n",info.devNode,info.serialNum,info.devNode,info.serialNum);
}

 int  get_devnode_properties2(const char *devnode,dev_info *info2)
{


    int	ret = false;
	struct udev *udev;
	struct udev_device *device;
	struct udev_enumerate *enumerate;
	struct udev_list_entry *first_entry;
	struct udev_list_entry *list_entry;
	const char *syspath;
 
	
	udev = udev_new();
	enumerate = udev_enumerate_new(udev);
	
	/* * 通过枚举器添加匹配条件,如果有多条匹配条件 */

/*  printf("%s  %d  match =  %d  \n",devnode,  udev_enumerate_add_match_property(enumerate,"DEVNAME",devnode),udev_enumerate_add_match_property(enumerate, "ID_USB_DRIVER", "usb-storage")); */
	/*
	 * 根据枚举器设置的条件扫描所有的设备.
	 * 注意: 如果有多条匹配条件,符合其中一条就会被扫描到,
	 *       匹配条件越多,扫描越宽松
	 */
	udev_enumerate_add_match_property(enumerate, "ID_USB_DRIVER", "usb-storage"); //
	udev_enumerate_scan_devices(enumerate);
    
	list_entry = udev_enumerate_get_list_entry(enumerate);
 
	if (!list_entry) {
		ret = false;
		goto ERROR1;
	}
 
	syspath = udev_list_entry_get_name(list_entry);
	device = udev_device_new_from_syspath(udev, syspath);
 
	if (!device) {
        ret = false;		
		goto ERROR1;
	} 
 
	first_entry = udev_device_get_properties_list_entry(device);
 
	if (!first_entry) {
		ret = false;
		goto ERROR2;
	}



	udev_list_entry_foreach(list_entry, first_entry) {
		 //printf("%s = %s\n", udev_list_entry_get_name(list_entry), udev_list_entry_get_value(list_entry));


		if( strcmp(udev_list_entry_get_name(list_entry) ,"DEVNAME")== 0)  {
		 //info2->devNode=udev_list_entry_get_value(list_entry);
		char *devNode=udev_list_entry_get_value(list_entry);
		strncpy(info2->devNode,devNode,strlen(devNode));
	    printf("##info.devNode_address= %p,info.devNode = %s ", info2->devNode, info2->devNode);
		}
        if(strcmp(udev_list_entry_get_name(list_entry),"ID_SERIAL_SHORT")==0) {
		   
		 //  info2->serialNum=udev_list_entry_get_value(list_entry);
		   char* serialnum = udev_list_entry_get_value(list_entry);
		   strncpy(info2->serialNum, serialnum, strlen(serialnum));
		   printf("info_serialNum_address = %p   info->devNode = %s ##\n", info2->serialNum,info2->serialNum );
       
		} 
      }
 
ERROR2:	
	udev_device_unref(device);
ERROR1:
	udev_enumerate_unref(enumerate);
	udev_unref(udev);
	return ret;
} 

 int  test_print_devnode_properties2(void)
{
     /* * 异常情况测试 */
 
    dev_info *info = (dev_info *)malloc(sizeof(dev_info));
	memset(info,0,sizeof(dev_info));

	info->devNode = (char *)malloc(64);
	memset(info->devNode,0,64);
	info->serialNum =(char *)malloc(128);
	memset(info->serialNum,0,128);
    printf("--------enter point address %p------- \n",info);
    get_devnode_properties2("/dev/sda1",info);
    printf("----RESULT=%p,%p, info.devNode = %s, info.serialNum= %s---------\n",info->devNode,info->serialNum,info->devNode,info->serialNum);

	dev_info info2;
	info2.devNode = (char *)malloc(64);
	memset(info2.devNode,0,64);
	info2.serialNum =(char *)malloc(128);
	memset(info2.serialNum,0,128);

   printf("--------enter point address %p ---------\n",&info2);
   get_devnode_properties2("/dev/sda1",&info2);
   printf(" !!!!    222     2RESULT=%p,%p, info.devNode = %s, info.serialNum= %s!!!!!!\n",info2.devNode,info2.serialNum,info2.devNode,info2.serialNum);

		if (info)
		free(info);
}


 int  get_devnode_properties3(const char *devnode)
{
    struct udev *udev;
	struct udev_device *device;
	struct udev_enumerate *enumerate;
	struct udev_list_entry *first_entry;
	struct udev_list_entry *list_entry;
	const char *syspath;
 
        
     dev_info *info = (char *)malloc(sizeof(dev_info));
     memset(info,0,sizeof(dev_info));
	
	udev = udev_new();
	enumerate = udev_enumerate_new(udev);
	
	/* * 通过枚举器添加匹配条件,如果有多条匹配条件 */

/*  printf("%s  %d  match =  %d  \n",devnode,  udev_enumerate_add_match_property(enumerate,"DEVNAME",devnode),udev_enumerate_add_match_property(enumerate, "ID_USB_DRIVER", "usb-storage")); */
	/*
	 * 根据枚举器设置的条件扫描所有的设备.
	 * 注意: 如果有多条匹配条件,符合其中一条就会被扫描到,
	 *       匹配条件越多,扫描越宽松
	 */
	udev_enumerate_add_match_property(enumerate, "ID_USB_DRIVER", "usb-storage"); //
	udev_enumerate_scan_devices(enumerate);
    
	list_entry = udev_enumerate_get_list_entry(enumerate);
 
	if (!list_entry) {
         info = NULL;
		goto ERROR1;
	}
 
	syspath = udev_list_entry_get_name(list_entry);
	device = udev_device_new_from_syspath(udev, syspath);
 
	if (!device) {
        info = NULL;		
		goto ERROR1;
	} 
 
	first_entry = udev_device_get_properties_list_entry(device);
 
	if (!first_entry) {
		info = NULL;
		goto ERROR2;
	}

	udev_list_entry_foreach(list_entry, first_entry) {
    // printf("%s = %s\n", udev_list_entry_get_name(list_entry), udev_list_entry_get_value(list_entry));

        
		if( strcmp(udev_list_entry_get_name(list_entry) ,"DEVNAME")== 0)  {

		// data = udev_list_entry_get_value(list_entry);
		// strncpy(info->devNode, data, strlen(data));
		

		info->devNode = (char*)malloc(64);
		memset(info->devNode,0,64);
	    char *devNode=udev_list_entry_get_value(list_entry);
		strncpy(info->devNode,devNode,strlen(devNode));
	    printf("##info.devNode_address= %p,info.devNode = %s ##", info->devNode, info->devNode);
		}
        if(strcmp(udev_list_entry_get_name(list_entry),"ID_SERIAL_SHORT")==0) {
		  info->serialNum = (char*)malloc(256);
          memset(info->serialNum,0,256);
		  char* serialnum = udev_list_entry_get_value(list_entry);
		   strncpy(info->serialNum, serialnum, strlen(serialnum));
		   printf(" info_serialNum_address = %p   info->devNode = %s  ##\n", info->serialNum,info->serialNum );
           
		}
      }
 
ERROR2:	
	udev_device_unref(device);
ERROR1:
	udev_enumerate_unref(enumerate);
	udev_unref(udev);
	return info;
} 

 int  test_print_devnode_properties3(void)
{
     /* * 异常情况测试 */
 
    dev_info *info ;
    info = get_devnode_properties3("/dev/sda");
 
   if (info == NULL)
       return -1;


    printf("333333333RESULT=%p,%p, info.devNode = %s, info.serialNum= %s\n",info->devNode,info->serialNum,info->devNode,info->serialNum);
       if (info->devNode)
		   free(info->devNode);
	   
	    if(info->serialNum)
	      free(info->serialNum);
	  
	   if (info)
        free(info);
 
 }


int  main(void)
{
 test_print_devnode_properties1();
  test_print_devnode_properties2();
  test_print_devnode_properties3();
  return 0;
}
