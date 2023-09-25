RAW Sync
--------

For RAW+JPEG shooters. After culling JPEGs, delete every RAW file that does not have a corresponding JPEG file.

## Usage

The script expects a similar folder structure for RAW and JPEG files. Two common examples are:

- RAW files stored alongside JPEGs in the same folder.
- Separate RAW and JPEG folders, but following mirrored folder structures (for example, `${FORMAT}/${YEAR}/${DATE}/${FILE}`).

The script recursively inspects every RAW file, searches for a matching JPEG, and if it doesn't find it, moves the RAW file to a staging directory for inspection. The staging directory can then me manually inspected, and if everything went as expected, deleted.

Example run:

```bash
go run ./cmd/main.go -jpg-path Fotos/Originals/jpg -raw-path Raw/raf -destination-path /tmp/raw
```

The flag `-dry-run` can be used to display the list of RAW files to move instead of actually moving them.
