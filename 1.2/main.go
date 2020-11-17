package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Stack представляет список элементов,
// организованных по принципу LIFO
type Stack struct {
	top *Node
}

// Node представляет элемента стека
type Node struct {
	value rune
	next  *Node // top(*Node) -> *Node -> *Node -> nil
}

// NewStack возвращает новый стек.
func NewStack() *Stack {
	return &Stack{nil}
}

// Peek возвращает значение, при этом не удаляет его.
func (s *Stack) Peek() rune {
	if s.top == nil {
		return 0
	}

	return s.top.value
}

// Pop удаляет и возвращает значение.
func (s *Stack) Pop() rune {
	if s.top == nil {
		return 0
	}

	n := s.top
	s.top = n.next
	return n.value
}

// Push добавляет значение на верх стека.
func (s *Stack) Push(value rune) {
	n := &Node{value: value, next: s.top}
	s.top = n
}

// Reverse вовзращает обращенную строку.
func Reverse(in string) (out string) {
	for _, r := range in {
		if r == '(' {
			out = ")" + out
		} else if r == ')' {
			out = "(" + out
		} else {
			out = string(r) + out
		}
	}

	return out
}

// Precedence возвращает приоритет функции.
func Precedence(char int32) int {
	switch char {
	case '(', ')':
		return 0
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	}

	return -1
}

// ToPostfix возвращает постфиксную форму выражения.
func ToPostfix(expression string) string {
	var postfix string
	s := NewStack()

	// Прохожусь по всем символам переданной строки
	for _, char := range expression {
		// Если пробел, то пропускаю
		if unicode.IsSpace(char) {
			continue
		}

		// Если буква или число,
		// то добавляю в конец постфиксной строки
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			postfix += string(char)
			continue
		}

		// Если скобка открывающая или возведение в степень,
		// то добавляю символ в стек
		if char == '(' || char == '^' {
			s.Push(char)
			continue
		}

		// Если скобка закрывающая, то забираю из
		// стека все символы до открывающей скобки
		// и добавляю их в конец постфиксной строки,
		// удаляю открывающую скобку из стека
		if char == ')' {
			for s.top != nil && s.Peek() != '(' {
				r := s.Pop()
				postfix += string(r)
			}

			// Удаляю открывающую скобку
			s.Pop()
			continue
		}

		// Всех остальные символы, чей приоритет больше или равен
		// приоритету символа, забираю из стека
		// и добавляю в конец постфиксной строки
		for s.top != nil && Precedence(char) <= Precedence(s.Peek()) {
			r := s.Pop()
			postfix += string(r)
		}

		// Добавляю символ в стек
		s.Push(char)
	}

	// Если стек не пуст, то забираю все символы из стека
	// и добавляю их в конец постфиксной строку
	for s.top != nil {
		r := s.Pop()
		postfix += string(r)
	}

	return postfix
}

// ToPrefix возвращает префиксную форму выражения.
func ToPrefix(expression string) string {
	var prefix string

	postfix := ToPostfix(Reverse(expression))

	// Воспользуюсь алгоритмом постфиксной трансляции.
	// Полученную строку записываю справа налево
	// в префиксную строку
	for _, char := range postfix {
		prefix = string(char) + prefix
	}

	return prefix
}

func main() {
	fmt.Print("Введите выражение в инфиксной форме: ")

	in := bufio.NewReader(os.Stdin)
	expression, _, err := in.ReadLine()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if len(expression) == 0 {
		fmt.Println("Введена пустая строка")
		return
	}

	prefix := ToPrefix(string(expression))
	postfix := ToPostfix(string(expression))

	fmt.Println("Выражение в префиксной форме:", prefix)
	fmt.Println("Выражение в постфиксной форме:", postfix)
}
