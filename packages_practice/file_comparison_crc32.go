package main

import (
	"errors"
	"fmt"
	"hash/crc32"
	"os"
)

func getHash(filename string) (uint32, error) {
	//открываем файл
	file, err := os.Open(filename)
	if err != nil {
		err := errors.New("getHash: ошибка открытия файла")
		fmt.Println(err)
	}
	//получаем информацию о файле
	stat, err := file.Stat()
	if err != nil {
		err := errors.New("getHash: ошибка получение Stat файла")
		fmt.Println(err)
	}
	//создаем срез для хранения файла(указываем размер среза = размеру файла)
	bs := make([]byte, stat.Size())
	//записываем содержимое файла в переменную
	_, err = file.Read(bs)
	if err != nil {
		err := errors.New("getHash: ошибка чтения файла")
		fmt.Println(err)
	}
	//создаем объект crc32
	h := crc32.NewIEEE()
	//запись данных файла в crc32
	h.Write(bs)
	return h.Sum32(), err
}

func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
