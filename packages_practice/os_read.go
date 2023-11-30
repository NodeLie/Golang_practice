package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		//перехват ошибки
		return
	}
	//используем отложенный вызов, для закрытия чтения файла после выполнения программы
	defer file.Close()

	//получаем информацию о файле, в данном случае нам необходимо знать только размер файла
	stat, err := file.Stat()
	if err != nil {
		return
	}
	//инициализируем переменную типа []byte для хранения данных из файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	//преобразуем []byte в string для вывода в терминал
	str := string(bs)
	fmt.Println(str)
}
