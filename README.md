# Termux-RootFS
This is a static root filesystem for [Termux](https://github.com/termux/termux-app) (AArch64). It is based on packages from [official termux repository](http://termux.net) and some
advanced software. It's working on Samsung Galaxy S7, but no guarantee that it will work on other devices.

## Features
* __Development tools:__ clang, gcc, go, perl, python 2 & 3, ruby, rust
* __Emulators:__ qemu-i386-static, qemu-x86_64-static, riscvemu
* __Games:__ 2048, curseofwar, moon-buggy, nsnake, pacman
* __Pentesting:__ aircrack-ng, reaver, mdk3, john, fcrackzip
* __Servers:__ nginx, openssh, polipo, privoxy, tor, transfer.sh
* __Databases:__ mariadb, postgresql, redis, sqlite3
* __VCS:__ mercurial, git
* __[More](#available-software)__

## Requirements
### System properties:
* __AArch64 architecture__
* __Android 6.x.x__ (security features of Android 7.0+ may break some apps)
* __SELinux permissive__ (if enforced, some apps may not work)
* __Root__ (a few apps won't work if your device not rooted)
* __3.5+ GB of free space in /data__

### Software needed for installation:
* __bash__
* __bzip2__
* __coreutils__
* __git__
* __tar__
* __wget__
* __xz-utils__

For optimal experience, you need to use the following modified Termux apps:
* [Termux](https://github.com/xeffyr/termux-app) (Termux:Boot, Termux:Styling, Termux:Widget are integrated)
* [Termux:API](https://github.com/xeffyr/termux-api)

## How to install
Please, __backup your home dir and current rootfs__ and __move to the safe place__ before installation of the termux-rootfs.

__Installation:__
```
  $ git clone https://github.com/xeffyr/termux-rootfs.git
  $ ./termux-rootfs/install.sh
```

Then stop Termux app and reopen it. If shell working, probably installation was ok.
If bad things happened, then use your backup (if you did it).

## How to upgrade
Incremental updates can be installed with command ```termux-upgrade-rootfs```. If the
upgrading process fails, you can try to apply patch manually with command:
```
  $ termux-apply-patch termux-rootfs-patch-inc-v${VERSION}.bin
```
File 'termux-rootfs-patch-inc-v${VERSION}.bin' is a binary patch file that should
be downloaded from the [releases page](https://github.com/xeffyr/termux-rootfs/releases). Label '${VERSION}' should be replaced with
next version, for example 3.3.

__Warning:__ big updates that patch several thousands of files may require a lot of RAM (~1.5GB).

## Password login protection
You can prevent using of termux by unwanted users by setting password with command '__passwd__' or '__termux-setup-rootfs__'.
__If you want to use a such feature, you must use a [patched Termux app](https://github.com/xeffyr/termux-app) to prevent
login bypassing with a 'failsafe' shell.__
To remove password login, delete file '__$PREFIX/etc/login.pwd__'.

## Available software
__admin tools:__
```
  apt, bmon, cpulimit, debootstrap, dnsutils, dpkg, fsmon, geoip, htop, httping, hping3,
  iproute2, iperf3, iw, macchanger, ngrep, nmap, proot, pwgen, ranpwd, sensors, sslscan,
  sipcalc, tracepath, whois, wireless-tools
```
__android:__
```
  adb, apktool, apk-utils, bootimg-tools, fastboot, resetprop, sdat2img, sparse-image-tools,
  termux-api
```
__archivers/compressors:__
```
  bsdtar, bzip2, cpio, gzip, lhasa, lzip, lzop, par2, p7zip, tar, unrar, unzip,
  xz-utils, zip
```
__binary file processors:__
```
  bvi, hexcurse, hexedit, hte, radare2
```
__console utils:__
```
  abduco, dialog, dvtm, screen, tmux, ttyrec
```
__databases:__
```
  mariadb, postgresql, redis, sqlite3
```
__data processors:__
```
  bc, datamash, docx2txt, dos2unix, ed, hunspell, micro, nano, pcapfix, poppler,
  stag, txt2man, vim, xmlstarlet, xsltproc
```
__data rescue:__
```
  ddrescue, extundelete, photorec, testdisk
```
__development:__
```
  astyle, autoconf, automake, bash-bats, binutils, bison, cargo, cccc, cfr,
  cgdb, cmake, clang, cppi, cpplint, cproto, cscope, ctags, diff2html,
  diffstat, ecj, elfutils, expect, flex, indent, jack, gcc, gdb, go, gperf,
  llvm, ltrace, lua, m4, make, nodejs, openjdk7-jre, patchelf, perl, python2,
  python3, ruby, rust, tcl, texinfo, unifdef, valac, yasm
```
__encryption:__
```
  aespipe, codecrypt, cryptsetup, encfs, gnupg, gnutls, openssl, scrypt, steghide
```
__filesystem:__
```
  btrfs-progs, duff, e2fsprogs, exfat-utils, lvm2, squashfs-tools, zerofree
```
__games:__
```
  2048, bs, curseofwar, hangman, moon-buggy, nsnake, nudoku, pacman,
  typespeed, vitetris
```
__generic utilities:__
```
  ag, bash, busybox, coreutils, dash, diffutils, file, findutils, fzf, gawk,
  gettext, global, grep, inetutils, info, less, man, mktorrent, patch, procps,
  rhash, rsync, tree, tasksh, taskwarrior, timewarrior, units, util-linux,
  xdelta3, zsh
```
__libraries:__
```
  apr, apr-util, boost, cairo, c-ares, db, expat, fftw, flac, fontconfig,
  freetype, gdbm, glib, gnutls, harfbuzz, harfbuzz-icu, icu, imlib2,
  libandroid-glob, libandroid-shmem, libandroid-support, libcaca, libconfig,
  libconfuse, libclang, libcroco, libcrypt, libcryptopp, libcurl, libedit,
  libevent, libffi, libgcrypt, libgd, libgit2, libgrpc, libid3tag, libidn,
  libisl, libjansson, libjasper, libjpeg-turbo, libmad, libmp3lame, libmpc,
  libmpfr, libnet, libnl, libnpth, libogg, libpcap, libpcre, libpipeline,
  libpng, libpopt, libprotobuf, libqrencode, librsync, libsodium, libssh,
  libssh2, libtalloc, libtiff, libunistring, libutil, libuuid, libvorbis,
  libx264, libx265, libxml2, libxslt, libyaml, libzmq, libzopfli, ldns,
  leptonica, littlecms, miniupnpc, ncurses, nettle, nghttp2, openblas,
  openjpeg, openssl, opus, pango, poppler, readline, c-toxcore, zlib
```
__media:__
```
  dcraw, ffmpeg, figlet, graphicsmagick, optipng, play-audio, sox, tesseract, toilet
```
__misc:__
```
  crunch, cmatrix, ent, eschalot, kona, lolcat, mathomatic, pick, sc,
  vanitygen-plus
```
__networking:__
```
  aria2, cryptcat, curl, elinks, ftp, irssi, lftp, lynx, megatools, netcat,
  socat, syncthing, tcpdump, telnet, torsocks, toxic, transmission, upnpc,
  wget, wput, zsync
```
__pentesting & cracking:__
```
  aircrack-ng, bettercap, fcrackzip, hydra, john, mdk3, metasploit-framework,
  pkcrack, reaver
```
__python 3 modules:__
```
  Automat, Cython, Django, Jinja2, Markdown, MarkupSafe, Pillow, PyBrain, PyDispatcher,
  PyMySQL, PyYAML, Pygments, SQLAlchemy, Scrapy, Twisted, Unidecode, WebOb, WebTest,
  Werkzeug, asn1crypto, astroid, attrs, autopep8, bash-kernel, beautifulsoup4, bleach,
  certifi, cffi, chardet, click, constantly, coverage, cryptography, cssselect, cycler,
  decorator, dismagic, dj-database-url, dj-static, django-bootstrap3, django-ckeditor,
  django-js-asset, django-orm-magic, django-profiler-middleware, django-pygments,
  djangorestframework, entrypoints, et-xmlfile, gevent, greenlet, guess-language-spirit,
  gunicorn, html5lib, httpie, httplib2, hyperlink, idna, incremental, iotop, ipdb, ipykernel,
  ipynose, ipyparallel, ipytest, ipython, ipython-genutils, ipywidgets, isort, jdcal, jedi,
  jsonschema, jupyter, jupyter-c-kernel, jupyter-client, jupyter-console, jupyter-core,
  jupyter-fortran-kernel, lazy-object-proxy, line-profiler, lxml, markdown2, matplotlib,
  mccabe, memory-profiler, mistune, mpmath, nbconvert, nbextensions, nbformat, nose,
  notebook, numpy, odfpy, olefile, openpyxl, pandas, pandocfilters, parsel, path.py, pbr,
  pep257, pep8, pexpect, pickleshare, pip, prompt-toolkit, psutil, psycopg2, ptyprocess,
  pyOpenSSL, pyasn1, pyasn1-modules, pycodestyle, pycparser, pydocstyle, pyflakes, pylama,
  pylint, pyparsing, python-dateutil, pytz, pyzmq, qrcode, qtconsole, queuelib, redis,
  redis-kernel, requests, scapy-python3, scikit-learn, scipy, service-identity, setuptools,
  sh, simplegeneric, six, snowballstemmer, static3, sympy, tablib, terminado, testpath,
  texttable, tornado, traitlets, unicodecsv, urllib3, virtualenv, virtualenv-clone, w3lib,
  waitress, wcwidth, webencodings, wheel, widgetsnbextension, wrapt, xlrd, xlwt, zope.interface
```
__servers:__
```
  nginx, openssh, polipo, privoxy, stunnel, tor, transfer.sh
```
__special/custom:__
```
  create-android-app, batteryinfo, buildapk, fake-chroot, login, linkchk,
  mount-sdext, myip, passwd, runlinux, secret-space-encryptor, service-manager,
  sudo, umount-sdext, update-config-guess, wifi-dump, wifi-jam, wttr.in
```
__vcs:__
```
  git, mercurial, tig
```
