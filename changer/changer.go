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

type ChangerEngine struct {
    nfc3dAmiiboKeys C.nfc3d_amiibo_keys
}

func NewChangerEngine() *ChangerEngine {
    return &ChangerEngine{}
}

func (this *ChangerEngine) LoadAmiiboKeys(path string) error {
    keyPath := C.CString(path)

    if C.amiibo_load_keys(&this.nfc3dAmiiboKeys, keyPath) == false {
        return errors.New(fmt.Sprintf("Could not load keys from %s", path))
    }

    return nil
}
