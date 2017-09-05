/*
** This file is part of login system for termux-rootfs.
**
** Copyright (C) 2017 Leonid Plyushch <leonid.plyushch@gmail.com>
**
** This program is free software: you can redistribute it and/or modify
** it under the terms of the GNU General Public License as published by
** the Free Software Foundation, either version 3 of the License, or
** (at your option) any later version.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//
//  Prompt password when opening new session in Termux app
//
//  Warning: Normally Termux has 'Failsafe' session mode for
//  using '/system/bin/sh' as login shell instead of '$PREFIX/bin/login',
//  so you will need to modify app to prevent login bypassing via 'Failsafe'.
//

#include <libgen.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <termios.h>
#include <unistd.h>

#include <readline/readline.h>
#include <openssl/evp.h>

#define LOGIN_BANNER_PATH "/data/data/com.termux/files/usr/etc/login-banner.txt"
#define MOTD_PATH "/data/data/com.termux/files/usr/etc/motd"
#define DYN_MOTD_PATH "/data/data/com.termux/files/usr/etc/motd.dyn"

void show_banner(const char *banner_path) {
    FILE *file;

    if ((file = fopen(banner_path, "r")) != NULL) {
        char *buffer = NULL;
        size_t len = 0;
        ssize_t chars_read = 0;

        while ((chars_read = getline(&buffer, &len, file)) != -1) {
            printf("%s", buffer);
        }

        if (buffer) {
            free(buffer);
        }

        fclose(file);
    }
}

char* getpass(const char *prompt) {
    struct termios term_old, term_new;

    /* Turn echoing off and fail if we can't. */
    if (tcgetattr (0, &term_old) != 0) {
        return NULL;
    }

    term_new = term_old;
    term_new.c_lflag &= ~ECHO;

    if (tcsetattr (0, TCSAFLUSH, &term_new) != 0) {
        return NULL;
    }

    /* Read the password. */
    char *password = readline(prompt);

    /* Restore terminal. */
    (void) tcsetattr (0, TCSAFLUSH, &term_old);

    return password;
}

void userlogin() {
    char link_data[1024];
    char *shell;

#ifdef SHOW_STATIC_MOTD
    if (access(MOTD_PATH, R_OK) == 0) {
        show_banner(MOTD_PATH);
    }
#endif

#ifdef SHOW_DYNAMIC_MOTD
    if (access(DYN_MOTD_PATH, R_OK | X_OK) == 0) {
        system(DYN_MOTD_PATH);
    }
#endif

    if(access("/data/data/com.termux/files/home/.termux/shell", F_OK) == 0) {
        ssize_t len = readlink("/data/data/com.termux/files/home/.termux/shell",
                               link_data, sizeof(link_data)-1);

        if(len != -1) {
            link_data[len] = '\0';
            shell = link_data;
        }
        else {
            shell = "/data/data/com.termux/files/usr/bin/bash";
        }
    }
    else {
        shell = "/data/data/com.termux/files/usr/bin/bash";
    }

    execl(shell, "-", (char *) 0);
}

int main(int argc, char *argv[]) {
    const char *progname = basename(argv[0]);
    const char *salt = "Termux";
    const int iterations = 16384;
    const int length = 256;
    char *login_password;
    unsigned char *input_pbkdf_data;
    unsigned char *passwd_pbkdf_data;

    FILE *passwd_file = NULL;
    const char *passwd_file_path = "/data/data/com.termux/files/usr/etc/login.pwd";

#ifdef SHOW_LOGIN_BANNER
    if (access(LOGIN_BANNER_PATH, R_OK) == 0) {
        show_banner(LOGIN_BANNER_PATH);
    }
#endif

    if (access(passwd_file_path, F_OK) != 0) {
#ifdef WARN_IF_PASSWD_NOT_SET
        puts("Looks like the login password is not set.");
        puts("Please, set it with command 'passwd'.\n");
#endif
        userlogin(progname);
    }
    else {
        // make sure that we can read /etc/login.pwd
        // 384 = -rw------- (600)
        chmod(passwd_file_path, 384);
    }

    while (1) {
        if ((input_pbkdf_data = (unsigned char *)malloc(length * sizeof(char))) == NULL) {
            puts("Failed to allocate memory.");
            return 1;
        }

        if ((passwd_pbkdf_data = (unsigned char *)malloc(length * sizeof(char))) == NULL) {
            puts("Failed to allocate memory.");
            return 1;
        }

        login_password = getpass("Enter password: ");
        puts("");

        if (login_password == NULL) {
            puts("Failed to read password input.");
            return 1;
        }

        PKCS5_PBKDF2_HMAC_SHA1(login_password, strlen(login_password), salt,
                               strlen(salt), iterations, length, input_pbkdf_data);
        free(login_password);

        if ((passwd_file = fopen(passwd_file_path, "r")) != NULL) {
            int r = fread(passwd_pbkdf_data, length, 1, passwd_file);
            fclose(passwd_file);

            if (r == 0) {
                puts("Failed to read data from login.pwd.");
                return 1;
            }
        }
        else {
            puts("Failed to open login.pwd for reading.");
            return 1;
        }

        int check = memcmp(input_pbkdf_data, passwd_pbkdf_data, length);
        free(input_pbkdf_data);
        free(passwd_pbkdf_data);

        if (check == 0) {
            puts("");
            userlogin(progname);
        }
        else {
            puts("Invalid password.");
        }
    }

    return 0;
}
