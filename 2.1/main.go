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
	if t.root == nil {
		t.root = NewNode(key)
	} else {
		t.root.insert(key)
	}
}

// DisplayPreOrder выводит на экран все значения
// в прямом порядке начиная с корня
func (t *Tree) DisplayPreOrder() {
	t.root.displayPreOrder()
}

// DisplayInOrder выводит на экран все значения
// в симметричном порядке начиная с корня
func (t *Tree) DisplayInOrder() {
	t.root.displayInOrder()
}

// DisplayInOrder выводит на экран все значения
// в обратном порядке начиная с корня
func (t *Tree) DisplayPostOrder() {
	t.root.displayPostOrder()
}

type Node struct {
	Key   int
	left  *Node
	right *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key}
}

// insert добавляет ключ в текущее дерево вершины (корень),
// после добавления балансирует дерево и возвращает указатель
// на новую вершину (корень) текущего дерева.
func (n *Node) insert(key int) {
	if n == nil {
		return
	}

	// Если значение искомого ключа, меньше,
	// чем значение ключа текущей вершины,
	// то добавляю значение в левое поддерево.
	if key < n.Key {
		if n.left == nil {
			n.left = NewNode(key)
		} else {
			n.left.insert(key)
		}
	}

	// Если значение искомого ключа, больше,
	// чем значение ключа текущей вершины,
	// то добавляю значение в правое поддерево.
	if key > n.Key {
		if n.right == nil {
			n.right = NewNode(key)
		} else {
			n.right.insert(key)
		}
	}
}

// displayInOrder выводит на экран значения в прямом порядке
func (n *Node) displayPreOrder() {
	if n == nil {
		return
	}

	fmt.Printf("%d ", n.Key)
	n.left.displayPreOrder()
	n.right.displayPreOrder()
}

// displayInOrder выводит на экран значения в симметричном порядке
func (n *Node) displayInOrder() {
	if n == nil {
		return
	}

	n.left.displayInOrder()
	fmt.Printf("%d ", n.Key)
	n.right.displayInOrder()

}

// displayPostOrder выводит на экран значения в обратном порядке
func (n *Node) displayPostOrder() {
	if n == nil {
		return
	}

	n.left.displayPostOrder()
	n.right.displayPostOrder()
	fmt.Printf("%d ", n.Key)
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
	fmt.Print("Введите команду (a: Вставка, p: Вывод): ")

	op, err := readValue()
	if err != nil {
		return err
	}

	switch op {
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
	case "p":
		fmt.Print("Прямой порядок: ")
		t.DisplayPreOrder()
		fmt.Println()

		fmt.Print("Симметричный порядок: ")
		t.DisplayInOrder()
		fmt.Println()

		fmt.Print("Обратный порядок: ")
		t.DisplayPostOrder()
		fmt.Println()

		fmt.Println()
	default:
		if err := readOp(t); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Онлайн: https://play.golang.org/p/JJF3bDwP4jx

	tree := NewTree()
	// В бесконечном цикле слушаю ввод пользователя
	for {
		if err := readOp(tree); err != nil {
			fmt.Println(err)
			return
		}
	}
}
