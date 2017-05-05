package changer

/*
#cgo CFLAGS: -I${SRCDIR}/amiitool/include
#cgo LDFLAGS: -L${SRCDIR}/amiitool -L${SRCDIR}/amiitool/mbedtls/library -lamiibo -lkeygen -ldrbg -lmbedtls -lmbedx509 -lmbedcrypto
#include "changer.h"
 */
import "C"

import (
    "fmt"
    "errors"
)

const (
    CHANGER_UNPACK_CANNOT_OPEN_DUMP_FILE        = C.AMIIBO_UNPACK_CANNOT_OPEN_DUMP_FILE
    CHANGER_UNPACK_FORMAT_INVALID               = C.AMIIBO_UNPACK_FORMAT_INVALID
    CHANGER_UNPACK_ERROR                        = C.AMIIBO_UNPACK_ERROR
    CHANGER_UNPACK_CANNOT_OPEN_UNPACK_FILE      = C.AMIIBO_UNPACK_CANNOT_OPEN_UNPACK_FILE
    CHANGER_UNPACK_CANNOT_SAVE_UNPACK_FILE      = C.AMIIBO_UNPACK_CANNOT_SAVE_UNPACK_FILE
    CHANGER_UNPACK_CANNOT_SAVE_MORE_UNPACK_FILE = C.AMIIBO_UNPACK_CANNOT_SAVE_MORE_UNPACK_FILE

    CHANGER_PACK_CANNOT_OPEN_DUMP_FILE          = C.AMIIBO_PACK_CANNOT_OPEN_DUMP_FILE
    CHANGER_PACK_FORMAT_INVALID                 = C.AMIIBO_PACK_FORMAT_INVALID
    CHANGER_PACK_CANNOT_OPEN_UNPACK_FILE        = C.AMIIBO_PACK_CANNOT_OPEN_UNPACK_FILE
    CHANGER_PACK_CANNOT_SAVE_UNPACK_FILE        = C.AMIIBO_PACK_CANNOT_SAVE_UNPACK_FILE
    CHANGER_PACK_CANNOT_SAVE_MORE_UNPACK_FILE   = C.AMIIBO_PACK_CANNOT_SAVE_MORE_UNPACK_FILE
)

func LoadAmiiboKeys(path string) (C.nfc3d_amiibo_keys, error) {
    var nfc3dAmiiboKeys C.nfc3d_amiibo_keys

    keyPath := C.CString(path)

    if C.amiibo_load_keys(&nfc3dAmiiboKeys, keyPath) == false {
        return nfc3dAmiiboKeys, errors.New(fmt.Sprintf("Could not load keys from %s", path))
    }

    return nfc3dAmiiboKeys, nil
}
