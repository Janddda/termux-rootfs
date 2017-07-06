CC=gcc
CFLAGS=-O2 -DSHOW_STATIC_MOTD -DSHOW_DYNAMIC_MOTD -DWARN_IF_PASSWD_NOT_SET
LDFLAGS=
LIBS=-lcrypto -lreadline

all: login passwd
	strip -s login
	strip -s passwd
	termux-elf-cleaner login
	termux-elf-cleaner passwd

login: login.c
	$(CC) $(CFLAGS) login.c -o login $(LDFLAGS) $(LIBS)

passwd: passwd.c
	$(CC) $(CFLAGS) passwd.c -o passwd $(LDFLAGS) $(LIBS)

.PHONY: clean
clean:
	rm -f login passwd
