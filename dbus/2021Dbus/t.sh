#!/bin/bash
gcc send.c  -I /usr/include/dbus-1.0/ -I /usr/include/glib-2.0/ -I /usr/lib/x86_64-linux-gnu/dbus-1.0/include  `pkg-config --libs --cflags dbus-1`  -o send
gcc rec.c  -I /usr/include/dbus-1.0/ -I /usr/include/glib-2.0/ -I /usr/lib/x86_64-linux-gnu/dbus-1.0/include  `pkg-config --libs --cflags dbus-1`  -o rec
