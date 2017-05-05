#include "changer.h"

uint8_t original[NTAG215_SIZE];
uint8_t modified[NFC3D_AMIIBO_SIZE];

bool amiibo_load_keys(nfc3d_amiibo_keys * amiibo_keys, const char * amiibo_keys_file_path) {
    return nfc3d_amiibo_load_keys(amiibo_keys, amiibo_keys_file_path);
}

int amiibo_unpack_dump_file(const nfc3d_amiibo_keys * amiibo_keys, char * amiibo_dump_file_path, char * amiibo_unpack_file_path) {
    // Open and read dump file
    FILE * dump_file = fopen(amiibo_dump_file_path, "rb");

    if (!dump_file) {
        return AMIIBO_UNPACK_CANNOT_OPEN_DUMP_FILE;
    }

    size_t readPages = fread(original, 4, NTAG215_SIZE / 4, dump_file);

    if (readPages < NFC3D_AMIIBO_SIZE / 4) {
        return AMIIBO_UNPACK_FORMAT_INVALID;
    }

    fclose(dump_file);

    // Try to unpack
    if (!nfc3d_amiibo_unpack(amiibo_keys, original, modified)) {
        return AMIIBO_UNPACK_ERROR;
    }

    // Save unpacked file
    FILE * unpack_file = fopen(amiibo_unpack_file_path, "wb");

    if (!unpack_file) {
        return AMIIBO_UNPACK_CANNOT_OPEN_UNPACK_FILE;
    }

    if (fwrite(modified, NFC3D_AMIIBO_SIZE, 1, unpack_file) != 1) {
        return AMIIBO_UNPACK_CANNOT_SAVE_UNPACK_FILE;
    }

    if (readPages > NFC3D_AMIIBO_SIZE / 4) {
        if (fwrite(original + NFC3D_AMIIBO_SIZE, readPages * 4 - NFC3D_AMIIBO_SIZE, 1, unpack_file) != 1) {
            return AMIIBO_UNPACK_CANNOT_SAVE_MORE_UNPACK_FILE;
        }
    }

    fclose(unpack_file);

    return true;
}

int amiibo_pack_dump_file(const nfc3d_amiibo_keys * amiibo_keys, char * amiibo_unpack_file_path, char * amiibo_pack_file_path) {
    // Open and read dump file
    FILE * unpack_file = fopen(amiibo_unpack_file_path, "rb");

    if (!unpack_file) {
        return AMIIBO_PACK_CANNOT_OPEN_DUMP_FILE;
    }

    size_t readPages = fread(original, 4, NTAG215_SIZE / 4, unpack_file);

    if (readPages < NFC3D_AMIIBO_SIZE / 4) {
        return AMIIBO_PACK_FORMAT_INVALID;
    }

    fclose(unpack_file);

    // Try to unpack
    nfc3d_amiibo_pack(amiibo_keys, original, modified);

    // Save unpacked file
    FILE * pack_file = fopen(amiibo_pack_file_path, "wb");

    if (!pack_file) {
        return AMIIBO_PACK_CANNOT_OPEN_UNPACK_FILE;
    }

    if (fwrite(modified, NFC3D_AMIIBO_SIZE, 1, pack_file) != 1) {
        return AMIIBO_PACK_CANNOT_SAVE_UNPACK_FILE;
    }

    if (readPages > NFC3D_AMIIBO_SIZE / 4) {
        if (fwrite(original + NFC3D_AMIIBO_SIZE, readPages * 4 - NFC3D_AMIIBO_SIZE, 1, pack_file) != 1) {
            return AMIIBO_PACK_CANNOT_SAVE_MORE_UNPACK_FILE;
        }
    }

    fclose(pack_file);

    return true;
}
