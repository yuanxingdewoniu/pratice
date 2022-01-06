#!/bin/bash
gcc -g meth_send.c  -I /usr/include/dbus-1.0/ -I /usr/include/glib-2.0/ -I /usr/lib/x86_64-linux-gnu/dbus-1.0/include  `pkg-config --libs --cflags dbus-1` `pkg-config --libs glib-2.0`   -o meth_send
gcc -g  meth_rec.c  -I /usr/include/dbus-1.0/ -I /usr/include/glib-2.0/ -I /usr/lib/x86_64-linux-gnu/dbus-1.0/include  `pkg-config --libs --cflags dbus-1`  `pkg-config --libs glib-2.0`   -o meth_rec
