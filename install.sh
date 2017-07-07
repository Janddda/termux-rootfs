#!/data/data/com.termux/files/usr/bin/sh
##
##    Termux RootFS Installer
##

VERSION="2.1"
SHA256="483bd4b84f0475c72e5c3b8f75c9755610cceebc97b116795b0fbec798c0dae8"

###############################################################################
##
##    Setting up environment
##
###############################################################################

BASEDIR="/data/data/com.termux/files"

ARCHIVE_NAME="termux-rootfs-v${VERSION}.tar.bz2"
ARCHIVE_PATH="${BASEDIR}/${ARCHIVE_NAME}"
ARCHIVE_URL="https://github.com/xeffyr/termux-rootfs/releases/download/v${VERSION}/${ARCHIVE_NAME}"

ROOTFS_DIR="termux-rootfs-v${VERSION}"

## use command busybox's version of command 'echo'
alias echo='busybox echo'

###############################################################################
##
##    Common functions
##
###############################################################################

is_binary_installed()
{
    if [ -x "${PREFIX}/bin/${1}" ]; then
        return 0
    else
        return 1
    fi
}

###############################################################################
##
##    Check packages needed for installation
##
###############################################################################

NEEDED_PACKAGES=""

for bin in bash bzip2 coreutils tar wget; do
    if ! is_binary_installed ${bin}; then
        NEEDED_PACKAGES="${NEEDED_PACKAGES} ${bin}"
    fi
done

if [ ! -z "${NEEDED_PACKAGES}" ]; then
    if is_binary_installed apt; then
        echo "[*] installing packages..."
        apt install ${NEEDED_PACKAGES}

        for bin in bash bzip2 coreutils tar wget; do
            if ! is_binary_installed ${bin}; then
                echo "[!] cannot find binary '${bin}'"\
                exit 1
            fi
        done
    else
        echo "[!] the following packages is not available:"

        echo
        for bin in ${NEEDED_PACKAGES}; do
            echo "    + ${bin}"
        done
        echo
    fi
fi

###############################################################################
##
##    Installation
##
###############################################################################

if [ ! -e "${ARCHIVE_PATH}" ]; then
    echo "[*] downloading archive..."
    if ! wget -O ${ARCHIVE_PATH} ${ARCHIVE_URL}; then
        echo "[!] failed to download archive"
        exit 1
    fi
fi

echo -n "[*] checking SHA-256... "
if [ "$(sha256sum ${ARCHIVE_PATH} | awk '{ print $1 }')" != "${SHA256}" ]; then
    echo "FAIL"
    exit 1
else
    echo "OK"
fi

if [ ! -e "${BASEDIR}/${ROOTFS_DIR}" ]; then
    echo -n "[*] extracting data... "
    if ! tar jxf ${ARCHIVE_PATH} -C ${BASEDIR} > /dev/null 2>&1; then
        echo "FAIL"
        exit 1
    else
        echo "OK"
    fi
else
    echo "[!] remove '${BASEDIR}/${ROOTFS_DIR}'"
    exit 1
fi

if [ ! -e "${BASEDIR}/usr.old" ]; then
    echo -n "[*] replacing rootfs... "
    if /system/bin/mv ${BASEDIR}/usr ${BASEDIR}/usr.old; then
        if /system/bin/mv ${BASEDIR}/termux-rootfs-v${VERSION} ${BASEDIR}/usr; then
            echo "OK"
        else
            echo "FAIL"
            exit 1
        fi
    else
        echo "FAIL"
        exit 1
    fi
else
    echo "[!] remove '${BASEDIR}/usr.old'"
    exit 1
fi

rm -f ${ARCHIVE_PATH}

exec ${PREFIX}/bin/bash ${PREFIX}/bin/termux-setup-rootfs
