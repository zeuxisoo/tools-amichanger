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
    ERROR_UNPACK_CANNOT_OPEN_DUMP_FILE        = C.AMIIBO_UNPACK_CANNOT_OPEN_DUMP_FILE
    ERROR_UNPACK_FORMAT_INVALID               = C.AMIIBO_UNPACK_FORMAT_INVALID
    ERROR_UNPACK_ERROR                        = C.AMIIBO_UNPACK_ERROR
    ERROR_UNPACK_CANNOT_OPEN_UNPACK_FILE      = C.AMIIBO_UNPACK_CANNOT_OPEN_UNPACK_FILE
    ERROR_UNPACK_CANNOT_SAVE_UNPACK_FILE      = C.AMIIBO_UNPACK_CANNOT_SAVE_UNPACK_FILE
    ERROR_UNPACK_CANNOT_SAVE_MORE_UNPACK_FILE = C.AMIIBO_UNPACK_CANNOT_SAVE_MORE_UNPACK_FILE

    ERROR_PACK_CANNOT_OPEN_DUMP_FILE          = C.AMIIBO_PACK_CANNOT_OPEN_DUMP_FILE
    ERROR_PACK_FORMAT_INVALID                 = C.AMIIBO_PACK_FORMAT_INVALID
    ERROR_PACK_CANNOT_OPEN_UNPACK_FILE        = C.AMIIBO_PACK_CANNOT_OPEN_UNPACK_FILE
    ERROR_PACK_CANNOT_SAVE_UNPACK_FILE        = C.AMIIBO_PACK_CANNOT_SAVE_UNPACK_FILE
    ERROR_PACK_CANNOT_SAVE_MORE_UNPACK_FILE   = C.AMIIBO_PACK_CANNOT_SAVE_MORE_UNPACK_FILE
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

func (this *ChangerEngine) UnpackAmiibo(originalFilePath string, unpackFilePath string) error {
    fromFilePath    := C.CString(originalFilePath)
    toFilePath      := C.CString(unpackFilePath)
    upackFileStatus := C.amiibo_unpack_dump_file(&this.nfc3dAmiiboKeys, fromFilePath, toFilePath)

    if upackFileStatus == ERROR_UNPACK_CANNOT_OPEN_DUMP_FILE {
        return errors.New(fmt.Sprintf("[UNPACK] Could not open input file %s", originalFilePath))
    }

    if upackFileStatus == ERROR_UNPACK_FORMAT_INVALID {
        return errors.New(fmt.Sprintf("[UNPACK] Could not read from input %s", originalFilePath))
    }

    if upackFileStatus == ERROR_UNPACK_ERROR {
        return errors.New(fmt.Sprintf("[UNPACK] Tag signature was NOT valid"))
    }

    if upackFileStatus == ERROR_UNPACK_CANNOT_OPEN_UNPACK_FILE {
        return errors.New(fmt.Sprintf("[UNPACK] Could not open output file %s", unpackFilePath))
    }

    if upackFileStatus == ERROR_UNPACK_CANNOT_SAVE_UNPACK_FILE {
        return errors.New(fmt.Sprintf("[UNPACK] Could not write to output %s", unpackFilePath))
    }

    if upackFileStatus == ERROR_UNPACK_CANNOT_SAVE_MORE_UNPACK_FILE {
        return errors.New(fmt.Sprintf("[UNPACK] Could not write more to output %s", unpackFilePath))
    }

    return nil
}

