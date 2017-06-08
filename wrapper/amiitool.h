#include <stdio.h>
#include "nfc3d/amiibo.h"

#define NTAG215_SIZE 540

#define AMIIBO_UNPACK_CANNOT_OPEN_DUMP_FILE        -1001
#define AMIIBO_UNPACK_FORMAT_INVALID               -1002
#define AMIIBO_UNPACK_ERROR                        -1003
#define AMIIBO_UNPACK_CANNOT_OPEN_UNPACK_FILE      -1004
#define AMIIBO_UNPACK_CANNOT_SAVE_UNPACK_FILE      -1005
#define AMIIBO_UNPACK_CANNOT_SAVE_MORE_UNPACK_FILE -1006

#define AMIIBO_PACK_CANNOT_OPEN_DUMP_FILE        -1031
#define AMIIBO_PACK_FORMAT_INVALID               -1032
#define AMIIBO_PACK_CANNOT_OPEN_UNPACK_FILE      -1034
#define AMIIBO_PACK_CANNOT_SAVE_UNPACK_FILE      -1035
#define AMIIBO_PACK_CANNOT_SAVE_MORE_UNPACK_FILE -1036

bool amiibo_load_keys(nfc3d_amiibo_keys * amiibo_keys, const char * amiibo_keys_file_path);
int amiibo_unpack_dump_file(const nfc3d_amiibo_keys * amiibo_keys, char * amiibo_dump_file_path, char * amiibo_unpack_file_path);
int amiibo_pack_dump_file(const nfc3d_amiibo_keys * amiibo_keys, char * amiibo_unpack_file_path, char * amiibo_pack_file_path);
