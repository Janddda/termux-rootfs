#!/data/data/com.termux/files/usr/bin/sh
##
##    Termux RootFS incremental updater
##

###############################################################################
##
##    Setting up environment
##
###############################################################################

SCRIPT_PATH=$(realpath "${0}")
SCRIPT_DIR=$(dirname "${SCRIPT_PATH}")

if [ ! -f "${SCRIPT_DIR}/VERSION_INFO" ]; then
    echo "[!] Cannot obtain version info"
    exit 1
fi

## Set variables: CURRENT_VERSION, ENABLE_INC_UPDATES, PREVIOUS_VERSION,
##                ARCHIVE_SHA256, PATCH_SHA256
. "${SCRIPT_DIR}"/VERSION_INFO

if [ "$(echo ${ENABLE_INC_UPDATES} | tr '[[:upper:]]' '[[:lower:]]')" = "no" ]; then
    echo "[!] Incremental updates disabled."
    exit 1
fi

PATCH_NAME="termux-rootfs-patch-inc-v${CURRENT_VERSION}.bin"
PATCH_FILE="${TMPDIR}/${PATCH_NAME}"
PATCH_URL="https://github.com/xeffyr/termux-rootfs/releases/download/v${CURRENT_VERSION}/${PATCH_NAME}"

## use command busybox's version of command 'echo'
alias echo='busybox echo'

## Check temporary directory
if [ -z "${TMPDIR}" ]; then
    export TMPDIR="${PREFIX}/tmp"
fi

if [ ! -e "${TMPDIR}" ]; then
    echo -n "[*] Temporary directory is not exist. Creating... "
    if mkdir -p "${TMPDIR}" > /dev/null 2>&1; then
        echo "OK"
    else
        echo "FAIL"
        exit 1
    fi
fi

###############################################################################
##
##    Check version of current Termux RootFS
##
###############################################################################

if [ -f "${PREFIX}/etc/os-release" ]; then
    if ! (
        . "${PREFIX}"/etc/os-release

        if [ ! -z "${VERSION}" ]; then
            if [ "${VERSION}" = "${PREVIOUS_VERSION}" ]; then
                echo "[*] Updating from ${VERSION} to ${CURRENT_VERSION}"
                exit 0
            elif [ "${VERSION}" = "${CURRENT_VERSION}" ]; then
                echo "[*] Your Termux RootFS is already patched."
                exit 1
            else
                echo "[*] Cannot upgrade from ${VERSION} to ${CURRENT_VERSION}"
                exit 1
            fi
        else
            echo "[!] Cannot check current version of Termux-RootFS"
            exit 1
        fi
    ); then
        exit 1
    fi
else
    echo "[!] This version of Termux RootFS does not support"
    echo "    incremental updates but you can try this on your"
    echo "    own risk."
    echo -n "    Do you still want to apply patch ? (y/n): "

    read -r CHOICE

    if [ "${CHOICE}" = "Y" ] || [ "${CHOICE}" = "y" ]; then
        echo "[*] Accepting user's choice..."
    else
        echo "[*] Exiting."
        exit 0
    fi
fi

###############################################################################
##
##    Applying patch
##
###############################################################################

if [ ! -f "${PATCH_FILE}" ]; then
    echo "[*] Downloading patch file..."
    if ! wget -O "${PATCH_FILE}" "${PATCH_URL}"; then
        echo "[!] Failed to download patch file"
        exit 1
    fi
fi

echo -n "[*] Checking SHA-256... "
if [ "$(sha256sum ${PATCH_FILE} | awk '{ print $1 }')" != "${PATCH_SHA256}" ]; then
    echo "FAIL"
    echo "[!] Deleting corrupted patch file."
    rm -f "${PATCH_FILE}"
    exit 1
else
    echo "OK"
fi

if cd "${PREFIX}" > /dev/null 2>&1; then
    echo "[*] Patching started..."
    if git apply --whitespace=nowarn --verbose "${PATCH_FILE}"; then
        cd - > /dev/null
        echo "[*] Patching successfully done."
        rm -f "${PATCH_FILE}"
        exit 0
    else
        cd - > /dev/null
        echo "[!] Patching done with errors."
        echo "    Patch file is saved to '\${TMPDIR}/${PATCH_NAME}'."
        exit 1
    fi
else
    echo "[!] Cannot cd to \${PREFIX}. Aborting."
    echo "    Patch file is saved to '\${TMPDIR}/${PATCH_NAME}'."
    exit 1
fi
