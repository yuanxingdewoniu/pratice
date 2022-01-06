#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <dbus/dbus.h>
#include <pthread.h>
 
struct DBus
{
    char m_receiverBusName[64];
    char m_receiverPath[32];
    char m_receiverInterFace[32];
    char m_receiverMethod[16];
    char m_sendBusName[64];
};
 
DBusConnection * DbusInit(const char *name, DBusError *perr)
{
    int ret;
    DBusError err = *perr;
    DBusConnection *connection;
 
    pid_t pid;
    pid = getpid();
 
    dbus_error_init(&err);
    // 创建于session D-Bus的连接
    connection = dbus_bus_get(DBUS_BUS_SESSION, &err);
    if (!connection)
    {
        if (dbus_error_is_set(&err))
            printf("Connection Error %s\n", err.message);
        else
            printf("%s %d err\n", __func__, __LINE__);
        return NULL;
    }
 
    // 注册公共名
    printf("init PID = %d, name = %s\n", pid, name);
    ret = dbus_bus_request_name(connection, name,
                            DBUS_NAME_FLAG_REPLACE_EXISTING, &err);
    if (ret != DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER)
    {
        if (dbus_error_is_set(&err))
            printf("Name Error %s\n", err.message);
        else
            printf("%s %d err\n", __func__, __LINE__);
        return NULL;
    }
 
    //设置为当收到disconnect信号的时候不退出应用程序(_exit())
    dbus_connection_set_exit_on_disconnect(connection, FALSE);
 
    return connection;
}
 
void ReplyMethodCall(DBusMessage *msg, DBusConnection *conn)
{
    DBusMessageIter msg_arg;
    //从msg中读取参数，根据传入参数增加返回参数
    if (!dbus_message_iter_init(msg, &msg_arg))
    {
        printf("Message has NO Argument\n");
        return;
    }
 
    do
    {
        int ret = dbus_message_iter_get_arg_type(&msg_arg);
        if (DBUS_TYPE_STRING == ret)
        {
            char *pdata;
            dbus_message_iter_get_basic(&msg_arg, &pdata);
            printf("get Method Argument STRING: %s\n", pdata);
        }
        else
        {
            printf("Argument Type ERROR\n");
        }
 
    } while (dbus_message_iter_next(&msg_arg));
 
    dbus_connection_flush(conn);
}
 
/* 监听D-Bus消息，我们在上次的例子中进行修改 */
//void DbusReceive(DBusConnection *connection)
void * ReceiveThread(void *arg)
{
    struct DBus *dbus = (struct DBus *)arg;
    DBusError err;
    DBusConnection *connection = DbusInit(dbus->m_receiverBusName, &err);
    if (connection == NULL)
    {
        dbus_error_free(&err);
        return NULL;
    }
 
    printf("ReceiveThread %s, %s, %s, %s\n", dbus->m_receiverBusName, dbus->m_receiverPath, dbus->m_receiverInterFace, dbus->m_receiverMethod);
 
    DBusMessage *msg;
 
    while (1)
    {
        dbus_connection_read_write(connection, 0);
 
        msg = dbus_connection_pop_message(connection);
        if (msg == NULL)
        {
            sleep(1);
            continue;
        }
 
        const char *path = dbus_message_get_path(msg);
        if (path == NULL || strcmp(path, dbus->m_receiverPath))
        {
            printf("Wrong PATH: %s !=%s\n", (path==NULL ? "" : path), dbus->m_receiverPath);
            dbus_message_unref(msg);
            continue;
        }
 
        printf("Get a Message\n");
        if (dbus_message_is_method_call(msg, dbus->m_receiverInterFace, dbus->m_receiverMethod))
        {
            printf("Someone Call My Method\n");
            ReplyMethodCall(msg, connection);
        }
        else
        {
            printf("NOT a Signal OR a Method\n");
        }
 
        dbus_message_unref(msg);
    }
 
    dbus_error_free(&err);
}
 
static void DbusSend(struct DBus *dbus, DBusConnection *connection, const void *value)
{
    printf("DbusSend %s, %s, %s, %s\n", dbus->m_receiverBusName, dbus->m_receiverPath, dbus->m_receiverInterFace, dbus->m_receiverMethod);
    //针对目的地地址，创建一个method call消息。
    //Constructs a new message to invoke a method on a remote object.
    DBusMessage *msg = dbus_message_new_method_call(
        dbus->m_receiverBusName, dbus->m_receiverPath,
        dbus->m_receiverInterFace, dbus->m_receiverMethod);
    if (msg == NULL)
    {
        printf("Message NULL");
        return;
    }
 
    DBusMessageIter arg;
    dbus_message_iter_init_append(msg, &arg);
    if (!dbus_message_iter_append_basic(&arg, DBUS_TYPE_STRING, &value))
    {
        printf("Out of Memory!");
        dbus_message_unref(msg);
        return;
    }
 
    // 发送消息并获得reply的handle 。Queues a message to send, as with dbus_connection_send()
    // but also returns a DBusPendingCall used to receive a reply to the message.
    DBusPendingCall *pending;
    if (!dbus_connection_send_with_reply(connection, msg, &pending, -1))
    {
        printf("Out of Memory!");
        dbus_message_unref(msg);
        return;
    }
 
    if (pending == NULL)
    {
        printf("Pending Call NULL: connection is disconnected ");
        dbus_message_unref(msg);
        return;
    }
 
    dbus_connection_flush(connection);
    dbus_message_unref(msg);
}
 
int main(int argc, char *argv[])
{

	/*
	int i = 0 ;
     for (i = 0; i < argc; i++)
	{
		printf("%s\n", argv[i]);
		return 0;
	} */

	if (strncmp(argv[1], "1", 1) == 0)
    {
        struct DBus rdbus;
        strcpy(rdbus.m_receiverBusName, "com.bt.c2s");
        strcpy(rdbus.m_receiverPath, "/com/bt/c2s/object");
        strcpy(rdbus.m_receiverInterFace, "com.bt.c2s.interface");
        strcpy(rdbus.m_receiverMethod,"method");
        strcpy(rdbus.m_sendBusName, "");
 
        pthread_t thread;
        pthread_create(&thread, NULL, ReceiveThread, &rdbus);
 
        struct DBus wdbus;
        strcpy(wdbus.m_receiverBusName, "com.bt.s2c");
        strcpy(wdbus.m_receiverPath, "/com/bt/s2c/object");
        strcpy(wdbus.m_receiverInterFace, "com.bt.s2c.interface");
        strcpy(wdbus.m_receiverMethod,"method");
        strcpy(wdbus.m_sendBusName, "com.bt.s2c_send");
 
        DBusError err;
        DBusConnection *connection = DbusInit(wdbus.m_sendBusName, &err);
        if (connection == NULL)
        {
            printf("%s %d\n", __func__, __LINE__);
            dbus_error_free(&err);
            return -1;
        }
 
        while (1)
        {
            DbusSend(&wdbus, connection, "world");
            sleep(3);
        }
        dbus_error_free(&err);
    }
    else if (strncmp(argv[1], "2", 1) == 0)
    {
        struct DBus rdbus;
        strcpy(rdbus.m_receiverBusName, "com.bt.s2c");
        strcpy(rdbus.m_receiverPath, "/com/bt/s2c/object");
        strcpy(rdbus.m_receiverInterFace, "com.bt.s2c.interface");
        strcpy(rdbus.m_receiverMethod,"method");
        strcpy(rdbus.m_sendBusName, "");
 
        pthread_t thread;
        pthread_create(&thread, NULL, ReceiveThread, &rdbus);
 
        struct DBus wdbus;
        strcpy(wdbus.m_receiverBusName, "com.bt.c2s");
        strcpy(wdbus.m_receiverPath, "/com/bt/c2s/object");
        strcpy(wdbus.m_receiverInterFace, "com.bt.c2s.interface");
        strcpy(wdbus.m_receiverMethod,"method");
        strcpy(wdbus.m_sendBusName, "com.bt.c2s_send");
 
        DBusError err;
        DBusConnection *connection = DbusInit(wdbus.m_sendBusName, &err);
        if (connection == NULL)
        {
            printf("%s %d\n", __func__, __LINE__);
            dbus_error_free(&err);
            return -1;
        }
 
        while (1)
        {
            DbusSend(&wdbus, connection, "hello");
            sleep(3);
        }
 
        dbus_error_free(&err);
    }
 
    printf("%s %d\n", __func__, __LINE__);
 
    return 0;
}
