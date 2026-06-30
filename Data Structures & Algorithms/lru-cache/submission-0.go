type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}

type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node // самый свежий
    tail     *Node // самый старый
}

func Constructor(capacity int) LRUCache {
    // Создаем фиктивные узлы для упрощения
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head
    
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*Node),
        head:     head,
        tail:     tail,
    }
}

// Удалить узел из списка
func (this *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

// Добавить узел после head (в начало)
func (this *LRUCache) addToFront(node *Node) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

// Переместить узел в начало (сделать самым свежим)
func (this *LRUCache) moveToFront(node *Node) {
    this.removeNode(node)
    this.addToFront(node)
}

// Удалить самый старый узел (перед tail)
func (this *LRUCache) removeLRU() {
    lru := this.tail.prev
    this.removeNode(lru)
    delete(this.cache, lru.key)
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        // Перемещаем узел в начало (он стал самым свежим)
        this.moveToFront(node)
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        // Ключ существует → обновляем значение и перемещаем в начало
        node.value = value
        this.moveToFront(node)
        return
    }
    
    // Новый ключ → создаем узел
    newNode := &Node{
        key:   key,
        value: value,
    }
    
    // Добавляем в map и в начало списка
    this.cache[key] = newNode
    this.addToFront(newNode)
    
    // Если превысили вместимость → удаляем самый старый
    if len(this.cache) > this.capacity {
        this.removeLRU()
    }
}