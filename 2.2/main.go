package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

// Insert добавляет ключ в дерево начиная с корня.
func (t *Tree) Insert(key int) {
	t.root = t.root.insert(key)
}

// Remove удаляет ключ из дерево начиная с корня.
func (t *Tree) Remove(key int) {
	t.root.remove(key)
}

// Display выводит дерево начиная с корня на экран.
func (t *Tree) Display() {
	t.root.display(0, 0)
}

// Search ищет ключ в девере начиная с корня.
// Возвращает ссылку на вершину с заданным ключом.
// Если такой вершины нет, то вернет nil.
func (t *Tree) Search(key int) *Node {
	return t.root.search(key)
}

type Node struct {
	Key    int
	left   *Node
	right  *Node
	height int
}

func NewNode(key int) *Node {
	return &Node{Key: key, height: 1}
}

// Height возвращает высоту
// переданной вершины (корень).
func (n *Node) Height() int {
	// Если переданная вершина не существует,
	// то возвращаю 0
	if n == nil {
		return 0
	}

	return n.height
}

// balanceFactor возвращает разницу
// между высотами левого и правого
// поддеревьев переданной вершины (корень).
func (n *Node) balanceFactor() int {
	// Если переданная вершина не существует,
	// то возвращаю 0
	if n == nil {
		return 0
	}

	return n.left.Height() - n.right.Height()
}

// fixHeight пересчитывает разницы
// между высотами двух поддеревьев
// переданной вершины (корень).
func (n *Node) fixHeight() {
	hl := n.left.Height()
	hr := n.right.Height()

	if hl > hr {
		n.height = hl + 1
	} else {
		n.height = hr + 1
	}
}

// rotateLeft возвращает результат правого поворота
// вокруг переданной вершины (корень).
func (n *Node) rotateRight() *Node {
	// Новый корень - указатель
	// на левое поддерево текущей вершины
	newRoot := n.left

	// Левое поддерево - указатель
	// на правое поддерево нового корня
	// (правое поддерево левого поддерева текущей вершины)
	n.left = newRoot.right

	// Правое поддерево нового корня -
	// указатель на левое поддерево текущей вершины
	newRoot.right = n.left

	// Исправляю высоты после изменений
	n.fixHeight()
	newRoot.fixHeight()

	// Возвращаю новый корень
	return newRoot
}

// rotateLeft возвращает результат левого поворота
// вокруг переданной вершины (корень).
func (n *Node) rotateLeft() *Node {
	// Новый корень - указатель
	// на правое поддерево текущей вершины
	newRoot := n.right

	// Правое поддерево - указатель
	// на левое поддерево нового корня
	// (левое поддерево правого поддерева текущей вершины)
	n.right = newRoot.left

	// Левое поддерево нового корня -
	// указатель на текущую вершину
	newRoot.left = n

	// Исправляю высоты после изменений
	n.fixHeight()
	newRoot.fixHeight()

	// Возвращаю новый корень
	return newRoot
}

// findMin возвращает указатель на вершину с минимальным
// значением из дерева переданной вершины (корень).
func (n *Node) findMin() *Node {
	// Если нет левого поддерева, то текущая вершина
	// и есть минимальная, возвращаю ее
	if n.left == nil {
		return n
	}

	// Иначе, возвращаю результат поиска
	// минимального значения из левого поддерева
	return n.left.findMin()
}

// removeMin удаляет вершину с минимальным
// значением из дерева переданной вершины (корень).
func (n *Node) removeMin() *Node {
	// Если нет левого поддерева,
	// то возвращаю указатель на правое поддерево
	if n.left == nil {
		return n.right
	}

	// Иначе удаляю минимальное значение из левого поддерева
	n.left = n.left.removeMin()

	// Балансирую дерево после изменений
	return n.balance()
}

// balance балансирует дерево переданной вершины.
// Возвращает указатель на новую вершину (корень)
// текущего дерева.
func (n *Node) balance() *Node {
	// Если переданная вершина не существует,
	// то возвращаю пустую ссылку
	if n == nil {
		return n
	}

	// Пересчитываю высоту текущей вершины
	n.fixHeight()

	balanceFactor := n.balanceFactor()

	// Если разница высот левого и правого поддеревьев
	// текущей вершины равна 2 или -2, то выполняю повороты
	if balanceFactor == 2 {
		// Если высота правого поддерева больше высоты левого поддерева,
		// то делаю левый поворот вокруг вершины левого поддерева
		if n.left.balanceFactor() < 0 {
			n.left = n.left.rotateLeft()
		}

		// Возвращаю результат правого поворота вокруг текущей вершины
		return n.rotateRight()
	}

	if balanceFactor == -2 {
		// Если высота левого поддерева больше высоты правого поддерева,
		// то делаю правый поворот вокруг вершины правого поддерева
		if n.right.balanceFactor() > 0 {
			n.right = n.right.rotateRight()
		}

		// Возвращаю результат левого поворота вокруг текущей вершины
		return n.rotateLeft()
	}

	// Если дерево текущей вершины сбалансировано,
	// то возвращаю текущую вершину
	return n
}

// insert добавляет ключ в текущее дерево вершины (корень),
// после добавления балансирует дерево и возвращает указатель
// на новую вершину (корень) текущего дерева.
func (n *Node) insert(key int) *Node {
	// Если переданная вершина не существует,
	// то возвращаю новую вершину
	if n == nil {
		return NewNode(key)
	}

	// Если значение искомого ключа, меньше,
	// чем значение ключа текущей вершины,
	// то добавляю значение в левое поддерево.
	if key < n.Key {
		n.left = n.left.insert(key)
	}

	// Если значение искомого ключа, больше,
	// чем значение ключа текущей вершины,
	// то добавляю значение в правое поддерево.
	if key > n.Key {
		n.right = n.right.insert(key)
	}

	// Балансирую дерево после изменений
	return n.balance()
}

// remove удаляет ключ из текущего дерева вершины,
// после удаления балансирует дерево и возвращает указатель
// на новую вершину (корень) текущего дерева.
func (n *Node) remove(key int) *Node {
	// Если переданная вершина не существует,
	// то возвращаю пустую ссылку
	if n == nil {
		return nil
	}

	// Если значение искомого ключа, меньше,
	// чем значение ключа текущей вершины,
	// то удаляю значение из левого поддерева.
	//
	// Если значение искомого ключа, больше,
	// чем значение ключа текущей вершины,
	// то удаляю значение из правого поддерева.
	//
	// Если значение искомого ключа, равно
	// значению текущей вершины, то произвожу удаление.
	if key < n.Key {
		n.left = n.left.remove(key)
	} else if key > n.Key {
		n.right = n.right.remove(key)
	} else {
		// Если текущая вершина имеет левое и правое поддерево,
		// то нахожу наименьшее значение из правого поддерева,
		// меняю местами значение текущей вершины на минимальное
		// из правого поддерева, и удаляю из правого поддерева
		// минимальное значение.
		//
		// Если отсутствует правое поддерево,
		// то заменяю указатель текущей вершины
		// на указатель правого поддерева.
		//
		// Если отсутствует левое поддерево,
		// то заменяю указатель текущей вершины
		// на указатель левого поддерева.
		//
		// Если текущая вершина - лист, то удаляю указатель.
		// Сборщик мусора сам удалит значение из памяти.
		if n.left != nil && n.right != nil {
			min := n.right.findMin()
			n.Key = min.Key
			n.right = n.right.remove(min.Key)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n
		}
	}

	// Балансирую дерево после изменений
	return n.balance()
}

// search возвращает указатель на вершину,
// начиная с заданной вершины.
func (n *Node) search(key int) *Node {
	// Если ключи совпали, то вернуть текущую вершину
	if key == n.Key {
		return n
	}

	// Если значение ключа больше, чем значение ключа текущей
	// вершины, то начинаю поиск с правого поддерева, если оно существует
	if key > n.Key && n.right != nil {
		return n.right.search(key)
	}

	// Если значение ключа меньше, чем значение ключа текущей
	// вершины, то начинаю поиск с левого поддерева, если оно существует
	if key < n.Key && n.left != nil {
		return n.left.search(key)
	}

	// Во всех остальных значениях возвращаю nil,
	// что означает - ключ не был найдет
	return nil
}

// display выводит дерево на экран
func (n *Node) display(level int, direction int) {
	pre := ""
	post := ""

	if n.right != nil && n.left != nil {
		post = "─┤"
	} else if n.right != nil {
		post = "─┘"
	} else if n.left != nil {
		post = "─┐"
	}

	if level > 0 {
		if direction == 1 {
			pre = "┌─"
		} else {
			pre = "└─"
		}

	}

	if n.right != nil {
		n.right.display(level+1, 1)
	}

	fmt.Printf("%*s%*d %s\n", level*8, pre, 4, n.Key, post)

	if n.left != nil {
		n.left.display(level+1, -1)
	}
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
func readOp(t *Tree) error {
	fmt.Print("Введите команду (s: Поиск, a: Вставка, p: Вывод): ")

	op, err := readValue()
	if err != nil {
		return err
	}

	switch op {
	case "s":
		fmt.Print("Введите строку: ")
		v, err := readValue()
		if err != nil {
			return err
		}

		key, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		e := t.Search(key)

		if e != nil {
			fmt.Println("Найден элемент: ", e)
		} else {
			fmt.Println("Ничего не найдено.")
		}

		break
	case "a":
		fmt.Print("Введите числа, разделенные пробелом: ")
		v, err := readValue()
		if err != nil {
			return err
		}

		keys := strings.Split(v, " ")

		for i := range keys {
			key, err := strconv.Atoi(keys[i])
			if err != nil {
				return err
			}

			t.Insert(key)
		}

		fmt.Println("Добавлены элементы:", keys)
		break
	case "p":
		t.Display()
		break
	default:
		if err := readOp(t); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Онлайн: https://play.golang.org/p/llvfb7nk2oB

	tree := NewTree()
	// В бесконечном цикле слушаю ввод пользователя
	for {
		if err := readOp(tree); err != nil {
			fmt.Println(err)
			return
		}
	}
}
