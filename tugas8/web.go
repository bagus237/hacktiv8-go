package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var path = "./data.json"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	data := map[string]int{"water": 11, "wild":5}
	dataJ, _ := json.Marshal(data)
	_, err = file.Write(dataJ)
	if isError(err) {
		return
	}
	defer file.Close()

	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

func main() {
	// createFile()
	writeFile()
}