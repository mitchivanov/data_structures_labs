package main

import "fmt"

// Определяем структуру Node, которая представляет элемент связного списка
type Node struct {
	data int   // Данные, хранящиеся в узле
	next *Node // Указатель на следующий узел в списке
}

// Определяем структуру List, которая представляет связный список
type List struct {
	head *Node // Указатель на головной узел списка
}

// Создаем метод append для добавления нового узла в конец списка
func (l *List) append(data int) {
	// Создаем новый узел с данными
	newNode := &Node{data: data}
	// Если список пуст, то делаем новый узел головным
	if l.head == nil {
		l.head = newNode
		return
	}
	// Иначе ищем последний узел в списке
	last := l.head
	for last.next != nil {
		last = last.next
	}
	// Добавляем новый узел в конец списка
	last.next = newNode
}

// Создаем метод remove для удаления головного узла списка и возвращения его данных
func (l *List) remove() int {
	// Проверяем, что список не пуст
	if l.head == nil {
		fmt.Println("Список пуст")
		return 0
	}
	// Запоминаем данные головного узла
	data := l.head.data
	// Удаляем головной узел и делаем следующий узел головным
	l.head = l.head.next
	// Возвращаем данные удаленного узла
	return data
}

// Создаем метод length для подсчета длины списка
func (l *List) length() int {
	// Начинаем с головного узла и счетчика, равного нулю
	current := l.head
	count := 0
	// Пока не достигнем конца списка, увеличиваем счетчик на единицу и переходим к следующему узлу
	for current != nil {
		count++
		current = current.next
	}
	// Возвращаем счетчик в качестве длины списка
	return count
}

// Создаем метод print для вывода элементов списка на экран
func (l *List) print() {
	// Начинаем с головного узла
	current := l.head
	// Пока не достигнем конца списка, выводим данные каждого узла
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// Определяем структуру Queue, которая представляет очередь на основе связного списка
type Queue struct {
	list *List // Указатель на связный список, который хранит элементы очереди
}

// Создаем метод enqueue для добавления элемента в конец очереди
func (q *Queue) enqueue(data int) {
	// Добавляем элемент в конец связного списка, который представляет очередь
	q.list.append(data)
}

// Создаем метод dequeue для удаления элемента из начала очереди и возвращения его данных
func (q *Queue) dequeue() int {
	// Удаляем элемент из начала связного списка, который представляет очередь, и возвращаем его данные
	return q.list.remove()
}

// Создаем метод size для определения размера очереди
func (q *Queue) size() int {
	// Возвращаем длину связного списка, который представляет очередь
	return q.list.length()
}

// Создаем метод print для вывода элементов очереди на экран
func (q *Queue) print() {
	// Выводим элементы связного списка, который представляет очередь
	q.list.print()
}

func main() {
	// Создаем пустую очередь на основе пустого связного списка
	queue := &Queue{list: &List{}}
	// Добавляем несколько элементов в очередь
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(4)
	queue.enqueue(5)
	// Выводим очередь на экран
	fmt.Println("Очередь после добавления элементов:")
	queue.print()
	// Удаляем элемент из очереди и выводим его данные
	fmt.Println("Удаленный элемент из очереди:")
	fmt.Println(queue.dequeue())
	// Выводим очередь на экран
	fmt.Println("Очередь после удаления элемента:")
	queue.print()
}