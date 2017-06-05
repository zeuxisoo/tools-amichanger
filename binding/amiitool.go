package binding

/*
#cgo CFLAGS: -I${SRCDIR}/amiitool/include
#cgo LDFLAGS: -L${SRCDIR}/amiitool -L${SRCDIR}/amiitool/mbedtls/library -lamiibo -lkeygen -ldrbg -lmbedtls -lmbedx509 -lmbedcrypto
#include "amiitool.h"
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


type AmiiToolEngine struct {
    nfc3dAmiiboKeys C.nfc3d_amiibo_keys
}

func NewAmiiToolEngine() *AmiiToolEngine {
    return &AmiiToolEngine{}
}

func (this *AmiiToolEngine) LoadAmiiboKeys(path string) error {
    keyPath := C.CString(path)

    if C.amiibo_load_keys(&this.nfc3dAmiiboKeys, keyPath) == false {
        return errors.New(fmt.Sprintf("Could not load keys from %s", path))
    }

    return nil
}

func (this *AmiiToolEngine) UnpackAmiibo(originalFilePath string, unpackFilePath string) error {
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

func (this *AmiiToolEngine) PackAmiibo(unpackFilePath string, packFilePath string) error {
    fromFilePath   := C.CString(unpackFilePath)
    toFilePath     := C.CString(packFilePath)
    packFileStatus := C.amiibo_pack_dump_file(&this.nfc3dAmiiboKeys, fromFilePath, toFilePath)

    if packFileStatus == ERROR_PACK_CANNOT_OPEN_DUMP_FILE {
        return errors.New(fmt.Sprintf("[PACK] Could not open input file %s", unpackFilePath))
    }

    if packFileStatus == ERROR_PACK_FORMAT_INVALID {
        return errors.New(fmt.Sprintf("[PACK] Could not read from input %s", unpackFilePath))
    }

    if packFileStatus == ERROR_PACK_CANNOT_OPEN_UNPACK_FILE {
        return errors.New(fmt.Sprintf("[PACK] Could not open output file %s", packFilePath))
    }

    if packFileStatus == ERROR_PACK_CANNOT_SAVE_UNPACK_FILE {
        return errors.New(fmt.Sprintf("[PACK] Could not write to output %s", packFilePath))
    }

    if packFileStatus == ERROR_PACK_CANNOT_SAVE_MORE_UNPACK_FILE {
        return errors.New(fmt.Sprintf("[PACK] Could not write more to output %s", packFilePath))
    }

    return nil
}
