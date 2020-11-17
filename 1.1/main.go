package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Segment представляет звено списка сегментов
type Segment struct {
	key  uint8
	list *Element // Ссылка на первый элемент списка
	head *Segment // Ссылка на первый сегмент
	next *Segment // Ссылка на следующий сегмент
}

// Element представляет звено списка элементов
type Element struct {
	data string
	head *Element // Ссылка на первый элемент
	next *Element // Ссылка на следующий элемент
}

// NewTable возвращает таблицу с нулевой записью.
func NewTable() *Segment {
	e := &Element{data: ""}
	e.head = e

	s := &Segment{
		key:  0,
		list: e,
	}

	s.head = s

	return s
}

// Print выведет на экран список элементов начиная с данного сегмента.
func (segment *Segment) Print() {
	i := 0

	// Прохожусь по всем сегментам
	for s := segment; s != nil; s = s.next {
		// Прохожусь по элементам списка сегмента
		if s.list != nil {
			for l := s.list; l != nil; l = l.next {
				fmt.Printf("%-3d\tKey: %-3d\tValue: %s\n", i, s.key, l.data)
				i++
			}
		}
	}
}

// PushForward добавляет элемент в начало списка в таблицу.
// Возвращает ссылку на сегмент.
func (segment *Segment) PushForward(data string) *Segment {
	// Получаю ключ списка по хэш функции от строки
	key := Pearson8Hash(data)

	for s := segment.head; s != nil; s = s.next {
		// Если ключ найден, то пробую добавить элемент в список
		if s.key == key {
			// Добавлю новый элемент в начало списка.
			if s.list == nil {
				e := &Element{data: data}
				e.head = e
				s.list = e
			} else {
				s.list.PushForward(data)
				s.list = s.list.head
			}

			return s
		}
	}

	// Если такого ключа нет в списке(таблице) сегментов, то
	// создаю новый сегмент, после добавляю в него пустой список
	// в этот список добавляю новый элемент и возвращаю ссылку на него
	e := &Element{data: data}
	e.head = e

	segment.head = &Segment{
		key:  key,
		list: e,
		next: segment.head,
	}

	return segment.head
}

// Delete удаляет элемент из таблицы.
func (segment *Segment) Delete(data string) {
	key := Pearson8Hash(data)

	// Прохожусь по всем сегментам
	for s := segment.head; s != nil; s = s.next {
		// Если совпал ключ, то удаляю элемент из списка, если он есть
		if s.key == key && s.list == nil {
			var prev *Element

			// Прохожусь по всем элементам списка сегмента
			for e := s.list.head; e != nil; e = e.next {
				// Если значение элемента совпало с переданным, то удаляю
				if e.data == data {
					if prev == nil {
						s.list = nil
						return
					} else {
						prev.next = e.next
						e.next = nil
					}
				}

				prev = e
			}
		}
	}
}

// Find возвращает ссылку на элемент.
// Если элемент не существует, то возвращает nil.
func (segment *Segment) Find(data string) *Element {
	key := Pearson8Hash(data)

	// Прохожусь по всем сегментам
	for s := segment.head; s != nil; s = s.next {
		// Если ключ совпал, то начинаю поиск в списке
		if s.key == key {
			return s.list.Find(data)
		}
	}

	return nil
}

// PushForward добавляет элемент в начало списка.
// Если элемент существует, возвращает ссылку на него.
func (element *Element) PushForward(data string) *Element {
	// Если элемент найден, то возвращаю его
	e := element.Find(data)
	if e != nil {
		return e
	}

	e = &Element{data: data, next: element.head}
	e.head = e
	element.head = e

	return e
}

// Find возвращает ссылку на элемент.
// Если элемент не существует, то возвращает nil.
func (element *Element) Find(data string) *Element {
	if element.data == data {
		return element
	}

	// Прохожусь по всем элементам списка
	for e := element.head; e != nil; e = e.next {
		// Если совпало значение элемента с переданным, то возвращаю его
		if e.data == data {
			return e
		}
	}

	return nil
}

// Pearson8Hash возвращает хэш значение переданной строки.
func Pearson8Hash(str string) uint8 {
	T := []uint8{
		98, 6, 85, 150, 36, 23, 112, 164, 135, 207, 169, 5, 26, 64, 165, 219,
		11, 61, 20, 68, 89, 130, 63, 52, 102, 24, 229, 132, 245, 80, 216, 195, 115,
		12, 90, 168, 156, 203, 177, 120, 2, 190, 188, 7, 100, 185, 174, 243, 162, 10,
		13, 237, 18, 253, 225, 8, 208, 172, 244, 255, 126, 101, 79, 145, 235, 228, 121,
		14, 123, 251, 67, 250, 161, 0, 107, 97, 241, 111, 181, 82, 249, 33, 69, 55,
		15, 59, 153, 29, 9, 213, 167, 84, 93, 30, 46, 94, 75, 151, 114, 73, 222,
		16, 197, 96, 210, 45, 16, 227, 248, 202, 51, 152, 252, 125, 81, 206, 215, 186,
		17, 39, 158, 178, 187, 131, 136, 1, 49, 50, 17, 141, 91, 47, 129, 60, 99,
		18, 154, 35, 86, 171, 105, 34, 38, 200, 147, 58, 77, 118, 173, 246, 76, 254,
		19, 133, 232, 196, 144, 198, 124, 53, 4, 108, 74, 223, 234, 134, 230, 157, 139,
		20, 189, 205, 199, 128, 176, 19, 211, 236, 127, 192, 231, 70, 233, 88, 146, 44,
		21, 183, 201, 22, 83, 13, 214, 116, 109, 159, 32, 95, 226, 140, 220, 57, 12,
		22, 221, 31, 209, 182, 143, 92, 149, 184, 148, 62, 113, 65, 37, 27, 106, 166,
		23, 3, 14, 204, 72, 21, 41, 56, 66, 28, 193, 40, 217, 25, 54, 179, 117,
		24, 238, 87, 240, 155, 180, 170, 242, 212, 191, 163, 78, 218, 137, 194, 175, 110,
		25, 43, 119, 224, 71, 122, 142, 42, 160, 104, 48, 247, 103, 15, 11, 138, 239,
	}

	hash := uint8(len(str) % 256)

	for i := 0; i < len(str); i++ {
		hash = T[hash^str[i]]
	}

	return hash
}

// readValue читает ввод пользователя,
// возвращает введенную строку.
func readValue() (string, error) {
	in := bufio.NewReader(os.Stdin)
	v, _, err := in.ReadLine()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(v)), nil
}

// readOp обрабатывает команду пользователя
func readOp(t *Segment) error {
	var op string
	fmt.Print("Введите команду (s: Поиск, a: Вставка, d: Удаление, p: Вывод): ")
	if _, err := fmt.Scan(&op); err != nil {
		return err
	}

	switch op {
	case "s":
		fmt.Print("Введите строку: ")
		v, err := readValue()
		if err != nil {
			return err
		}

		e := t.Find(v)

		if e != nil {
			fmt.Println("Найден элемент: ", e)
		} else {
			fmt.Println("Ничего не найдено.")
		}
	case "a":
		fmt.Print("Введите строку: ")
		v, err := readValue()
		if err != nil {
			return err
		}

		t.PushForward(v)

		fmt.Println("Добавлен элемент: ", v)
	case "d":
		fmt.Print("Введите строку: ")
		v, err := readValue()
		if err != nil {
			return err
		}

		t.Delete(v)
		fmt.Println("Удален элемент: ", v)
	case "p":
		t.head.Print()
	default:
		if err := readOp(t); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	table := NewTable()

	// В бесконечном цикле слушаем ввод пользователя
	for {
		if err := readOp(table); err != nil {
			fmt.Println(err)
			return
		}
	}
}
