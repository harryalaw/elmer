package serialization

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strings"

	"github.com/harryalaw/elmer/internal/db"
)

func WriteDb(db *db.Db, filepath string) error {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)
	enc.Encode(db)

	dir := parsePath(filepath, "/")

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		fmt.Println("Mkdir failed: ", err)
		return err
	}

	file, err := os.Create(filepath)

	if err != nil {
		fmt.Println("Create failed: ", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(buf.Bytes())

	if err != nil {
		fmt.Println("Write failed: ", err)
		return err
	}

	return nil
}

func ImportDb(filepath string) (*db.Db, error) {
	info, err := os.Stat(filepath)

	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, fmt.Errorf("File is a directory")
	}

	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	dec := gob.NewDecoder(file)

	var db db.Db

	err = dec.Decode(&db)

	if err != nil {
		return nil, fmt.Errorf("Error decoding data: %+v", err)
	}

	return &db, nil
}

func parsePath(filepath string, sep string) string {
	parts := strings.Split(filepath, sep)

	dir := strings.Join(parts[:len(parts)-1], sep)

	return dir
}
