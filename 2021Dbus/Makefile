.PHONY:rec send
CC = gcc
sec:
	gcc rec.c  -I /usr/include/dbus-1.0/ -I /usr/include/glib-2.0/ -I /usr/lib/x86_64-linux-gnu/dbus-1.0/include  `pkg-config --libs --cflags dbus-1`  
send:
	gcc send.c  -I /usr/include/dbus-1.0/ -I /usr/include/glib-2.0/ -I /usr/lib/x86_64-linux-gnu/dbus-1.0/include  `pkg-config --libs --cflags dbus-1`  

.PHONY:clean
clean:
	rm sec send
