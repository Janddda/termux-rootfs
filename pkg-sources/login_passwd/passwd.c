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
//  Set a new password for 'login'
//

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <termios.h>
#include <unistd.h>

#include <readline/readline.h>
#include <openssl/evp.h>

char* getpass(const char *prompt) {
    struct termios term_old, term_new;
    int nread;

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

int main(int argc, char *argv[]) {
    const char *password;
    const char *password_confirmation;
    const char *salt = "Termux";
    const int iterations = 16384;
    const int length = 256;
    unsigned char *output;

    FILE *passwd_file = NULL;
    const char *passwd_file_path = "/data/data/com.termux/files/usr/etc/login.pwd";

    if ((output = (unsigned char *)malloc(length * sizeof(char))) == NULL) {
        puts("Failed to allocate memory.");
        return 1;
    }

    password = getpass("New password: ");
    puts("");

    if(password == NULL) {
        puts("Failed to read password input.");
        return 1;
    }

    if(strlen(password) < 8) {
        puts("Minimum password size is 8 characters.");
        return 1;
    }

    password_confirmation = getpass("Retype new password: ");
    puts("");

    if(password_confirmation == NULL) {
        puts("Failed to read password input.");
        return 1;
    }

    if(strcmp(password, password_confirmation) != 0) {
        puts("Sorry, passwords do not match.");
        return 1;
    }

    PKCS5_PBKDF2_HMAC_SHA1(password, strlen(password), salt, strlen(salt),
                           iterations, length, output);

    if((passwd_file = fopen(passwd_file_path, "w")) != NULL) {
        int w = fwrite(output, length, 1, passwd_file);
        fclose(passwd_file);

        if (w == 0) {
            puts("Failed to write output to login.pwd.");
            return 1;
        }

        // Set permissions.
        // 384 = -rw------- (600)
        chmod(passwd_file_path, 384);

        puts("New password was successfully set.");
        return 0;
    }
    else {
        puts("Failed to open login.pwd for writing.");
        return 1;
    }

    return 0;
}
