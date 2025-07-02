# Data Structures Comprehensive Analysis

This document provides detailed analysis of all data structure implementations in this Go project, including their properties, operations, complexity analysis, and real-world applications.

## Table of Contents
1. [Linear Data Structures](#linear-data-structures)
   - [Array-based Stack](#array-based-stack)
   - [Array-based Queue](#array-based-queue)
   - [Circular Queue](#circular-queue)
   - [Double-ended Queue (Deque)](#double-ended-queue-deque)
   - [Simply Linked List](#simply-linked-list)
   - [Doubly Linked List](#doubly-linked-list)
   - [Circular Linked List](#circular-linked-list)
2. [Tree Data Structures](#tree-data-structures)
   - [Binary Tree](#binary-tree)
   - [Binary Search Tree](#binary-search-tree)
   - [Trie (Prefix Tree)](#trie-prefix-tree)
3. [Heap Data Structures](#heap-data-structures)
   - [Min Heap](#min-heap)
   - [Max Heap](#max-heap)

---

## Linear Data Structures

### Array-based Stack

**File:** `/datastructures/stack/Stack.go`

#### Overview and Properties
A stack is a Last-In-First-Out (LIFO) data structure implemented using Go's dynamic array (slice). Elements are added and removed from the same end (top of the stack).

#### Memory Layout
```
Stack: [10, 20, 30, 40]
Index:  0   1   2   3
              ^top (index 3)
```

#### Implementation Analysis
- **Storage**: Uses Go slice `[]int` for dynamic resizing
- **Top tracking**: Implicitly tracked using slice length
- **Operations**: Push, Pop, Peek, isEmpty, size

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Push      | O(1) amortized | O(1)            | May trigger slice reallocation |
| Pop       | O(1)          | O(1)            | Returns value and boolean |
| Peek      | O(1)          | O(1)            | Non-destructive top access |
| isEmpty   | O(1)          | O(1)            | Length check |
| size      | O(1)          | O(1)            | Length access |

#### Detailed Operation Traces

**Push Operation Trace:**
```go
stack := &Stack{}
stack.Push(10)  // items = [10]
stack.Push(20)  // items = [10, 20]
stack.Push(30)  // items = [10, 20, 30]
```

**Pop Operation Trace:**
```go
// Starting state: items = [10, 20, 30]
val, ok := stack.Pop()  // val=30, ok=true, items=[10, 20]
val, ok = stack.Pop()   // val=20, ok=true, items=[10]
val, ok = stack.Pop()   // val=10, ok=true, items=[]
val, ok = stack.Pop()   // val=0, ok=false, items=[]
```

#### Stack Applications Implemented

**1. String Reversal** (`ReverseString.go`)
- Converts string to rune array via stack
- Handles Unicode characters correctly
- Time: O(n), Space: O(n)

**2. Balanced Parentheses** (`stack-balanced-parentheses.go`)
- Validates matching brackets: (), [], {}
- Uses stack to track opening brackets
- Time: O(n), Space: O(n)

#### Real-world Applications
- Function call management (call stack)
- Undo operations in text editors
- Expression evaluation and syntax parsing
- Browser history (back button)
- Memory management in recursive algorithms

#### Comparison with Related Structures
- **vs Queue**: LIFO vs FIFO ordering
- **vs Deque**: Single-ended vs double-ended access
- **vs Linked List**: O(1) top access vs O(n) arbitrary access

---

### Array-based Queue

**File:** `/datastructures/queue/queue.go`

#### Overview and Properties
A queue is a First-In-First-Out (FIFO) data structure where elements are added at the rear and removed from the front. This implementation uses a simple array-based approach.

#### Memory Layout
```
Queue: [10, 20, 30, 40]
Index:  0   1   2   3
       ^front      ^rear
```

#### Implementation Analysis
- **Storage**: Go slice `[]int`
- **Enqueue**: Append to end of slice
- **Dequeue**: Remove from front (shifts all elements)
- **Inefficiency**: Dequeue operation is O(n) due to element shifting

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Enqueue   | O(1) amortized | O(1)            | Append operation |
| Dequeue   | O(n)          | O(1)            | Requires shifting elements |
| Front     | O(1)          | O(1)            | Access first element |
| isEmpty   | O(1)          | O(1)            | Length check |
| size      | O(1)          | O(1)            | Length access |

#### Detailed Operation Traces

**Enqueue Operation:**
```go
q := &Queue{}
q.Enqueue(10)  // items = [10]
q.Enqueue(20)  // items = [10, 20]
q.Enqueue(30)  // items = [10, 20, 30]
```

**Dequeue Operation (Inefficient):**
```go
// items = [10, 20, 30]
val, ok := q.Dequeue()  // val=10, ok=true, items=[20, 30] (shift occurred)
val, ok = q.Dequeue()   // val=20, ok=true, items=[30] (shift occurred)
```

#### Variants and Optimizations
1. **Circular Queue**: More efficient implementation (see below)
2. **Linked List Queue**: O(1) dequeue without shifting
3. **Double-ended Queue**: Supports both ends

#### Real-world Applications
- Process scheduling in operating systems
- Breadth-first search algorithms
- Print job management
- Buffer for data streams
- Level-order tree traversal

---

### Circular Queue

**File:** `/datastructures/circularQueue/circularQueue.go`

#### Overview and Properties
An optimized queue implementation using a fixed-size circular buffer. It eliminates the need for element shifting by using front and rear pointers that wrap around.

#### Memory Layout
```
CircularQueue (capacity=5):
items: [_, 20, 30, 40, _]
index:  0  1   2   3   4
       ^rear      ^front

front=3, rear=0, size=3
```

#### Implementation Analysis
- **Storage**: Fixed-size array `[]int`
- **Pointers**: `front`, `rear` indices with modular arithmetic
- **Size tracking**: Explicit `size` counter
- **Wrap-around**: Uses modulo operation for circular behavior

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Enqueue   | O(1)          | O(1)            | No shifting required |
| Dequeue   | O(1)          | O(1)            | No shifting required |
| Front     | O(1)          | O(1)            | Direct access |
| IsFull    | O(1)          | O(1)            | Size comparison |
| IsEmpty   | O(1)          | O(1)            | Size check |

#### Detailed Operation Traces

**Circular Wrap-around Example:**
```go
cq := NewCircularQueue(3)
// Initial: front=0, rear=-1, size=0, items=[_, _, _]

cq.Enqueue(1)  // front=0, rear=0, size=1, items=[1, _, _]
cq.Enqueue(2)  // front=0, rear=1, size=2, items=[1, 2, _]
cq.Enqueue(3)  // front=0, rear=2, size=3, items=[1, 2, 3] (full)

val, _ := cq.Dequeue()  // val=1, front=1, rear=2, size=2, items=[1, 2, 3]
cq.Enqueue(4)          // front=1, rear=0, size=3, items=[4, 2, 3] (wrapped)
```

#### Memory Efficiency
- **Fixed Memory**: No dynamic allocation after initialization
- **No Fragmentation**: Reuses same memory locations
- **Cache Friendly**: Contiguous memory access pattern

#### Real-world Applications
- Ring buffers in embedded systems
- Keyboard input buffering
- Network packet queues
- Producer-consumer problems with bounded buffers
- CPU scheduling (round-robin)

#### Comparison with Simple Queue
- **Space**: Fixed vs dynamic
- **Dequeue**: O(1) vs O(n)
- **Memory**: Efficient reuse vs potential waste

---

### Double-ended Queue (Deque)

**File:** `/datastructures/deque/deque.go`

#### Overview and Properties
A deque (double-ended queue) allows insertion and deletion at both ends. This implementation uses a simple array-based approach with operations at both front and back.

#### Memory Layout
```
Deque: [10, 20, 30, 40]
Index:  0   1   2   3
       ^front      ^back
```

#### Implementation Analysis
- **Storage**: Go slice `[]int`
- **Front operations**: Use array prepending/slicing
- **Back operations**: Use array appending/slicing
- **Limitation**: Front operations are O(n) due to shifting

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| PushFront | O(n)          | O(1)            | Requires shifting all elements |
| PushBack  | O(1) amortized| O(1)            | Simple append |
| PopFront  | O(n)          | O(1)            | Requires shifting all elements |
| PopBack   | O(1)          | O(1)            | Simple slice operation |
| Front     | O(1)          | O(1)            | Direct access |
| Back      | O(1)          | O(1)            | Direct access |

#### Detailed Operation Traces

**Mixed Operations:**
```go
d := &Deque{}
d.PushBack(20)   // items = [20]
d.PushBack(30)   // items = [20, 30]
d.PushFront(10)  // items = [10, 20, 30] (shift occurred)

// Front: 10, Back: 30
val, _ := d.PopFront()  // val=10, items=[20, 30] (shift occurred)
val, _ = d.PopBack()    // val=30, items=[20]
```

#### Variants and Optimizations
1. **Circular Deque**: Use circular buffer for O(1) operations
2. **Linked List Deque**: Doubly linked list implementation
3. **Segmented Deque**: Combine arrays and linked lists

#### Real-world Applications
- Palindrome checking
- Sliding window algorithms
- Undo/redo operations in applications
- Job scheduling with priority changes
- Browser navigation (forward/back)

---

### Simply Linked List

**File:** `/datastructures/simplylinkedlist/simplylinkedlist.go`

#### Overview and Properties
A singly linked list is a linear data structure where elements (nodes) are stored in sequence, with each node containing data and a pointer to the next node.

#### Memory Layout
```
head -> [10|next] -> [20|next] -> [30|null]
        Node1        Node2        Node3
```

#### Implementation Analysis
- **Node Structure**: Contains `data int` and `next *Node`
- **Head Pointer**: Points to first node
- **Insertion**: Only at beginning (O(1))
- **Traversal**: Sequential access only

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| InsertAtBeginning | O(1) | O(1)       | Direct head manipulation |
| Traverse          | O(n) | O(1)       | Visit each node once |
| Search            | O(n) | O(1)       | Linear search required |
| Delete            | O(n) | O(1)       | Need to find predecessor |

#### Detailed Operation Traces

**Insertion Sequence:**
```go
list := SimplyLinkedList{}
// head = nil

list.InsertAtBeginning(10)
// head -> [10|nil]

list.InsertAtBeginning(20)
// head -> [20|next] -> [10|nil]

list.InsertAtBeginning(30)
// head -> [30|next] -> [20|next] -> [10|nil]
```

#### Memory Characteristics
- **Dynamic Size**: Grows/shrinks as needed
- **Non-contiguous**: Nodes scattered in memory
- **Memory Overhead**: Extra pointer per node
- **Cache Performance**: Poor due to pointer chasing

#### Variants and Extensions
1. **Singly Linked List with Tail**: O(1) insertion at end
2. **Sorted Linked List**: Maintains order
3. **Circular Singly Linked List**: Last node points to first

#### Real-world Applications
- Implementation of other data structures (stacks, queues)
- Memory management (free lists)
- Music playlist (next song)
- Web browser history
- Undo functionality in text editors

---

### Doubly Linked List

**File:** `/datastructures/doublylinkedlist/doublylinkedlist.go`

#### Overview and Properties
A doubly linked list allows traversal in both directions by maintaining both next and previous pointers in each node, along with head and tail references.

#### Memory Layout
```
head -> [prev|10|next] <-> [prev|20|next] <-> [prev|30|next] <- tail
        null              ^               ^              null
                      Node1           Node2           Node3
```

#### Implementation Analysis
- **Node Structure**: `data int`, `next *DoubleNode`, `prev *DoubleNode`
- **List Structure**: `head *DoubleNode`, `tail *DoubleNode`
- **Bidirectional**: Forward and backward traversal
- **Limited Operations**: Only insertion at end implemented

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| InsertAtEnd    | O(1)        | O(1)            | Direct tail manipulation |
| ReverseTraverse| O(n)        | O(1)            | Backward traversal |
| Forward Traverse| O(n)       | O(1)            | Standard traversal |
| Search         | O(n)        | O(1)            | Can search from either end |

#### Detailed Operation Traces

**Insertion at End:**
```go
list := &DoublyLinkedList{}
// head = nil, tail = nil

list.InsertAtEnd(10)
// head -> [nil|10|nil] <- tail

list.InsertAtEnd(20)
// head -> [nil|10|next] <-> [prev|20|nil] <- tail

list.InsertAtEnd(30)
// head -> [nil|10|next] <-> [prev|20|next] <-> [prev|30|nil] <- tail
```

#### Memory Characteristics
- **Extra Memory**: Additional pointer per node (prev)
- **Faster Operations**: O(1) insertion/deletion when node reference known
- **Bidirectional Access**: Can traverse both ways efficiently

#### Variants and Optimizations
1. **Circular Doubly Linked List**: Head.prev = tail, tail.next = head
2. **XOR Linked List**: Memory-efficient using XOR for both pointers
3. **Skip List**: Multi-level doubly linked structure

#### Real-world Applications
- Browser navigation (forward/back buttons)
- Music player (previous/next track)
- Text editors (cursor movement)
- LRU cache implementation
- Undo/redo systems

#### Comparison with Singly Linked List
- **Memory**: ~2x memory usage
- **Traversal**: Bidirectional vs unidirectional
- **Deletion**: O(1) vs O(n) when node is known
- **Complexity**: More complex pointer management

---

### Circular Linked List

**File:** `/datastructures/circularlinkedlist/circularlinkedlist.go`

#### Overview and Properties
A circular linked list is a variation where the last node points back to the first node, creating a circular structure. Only a head pointer is maintained.

#### Memory Layout
```
head -> [1|next] -> [2|next] -> [3|next] -> [4|next]
         ^                                     |
         |_____________________________________|
```

#### Implementation Analysis
- **Circular Property**: Last node's next points to head
- **Head Management**: Single head pointer for entire list
- **Insertion**: At end (requires traversal to find last node)
- **Deletion**: Complex due to circular nature

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Insert    | O(n)          | O(1)            | Must find last node |
| Delete    | O(n)          | O(1)            | Must handle circular links |
| Search    | O(n)          | O(1)            | Circular traversal |
| Traverse  | O(n)          | O(1)            | Must detect full circle |

#### Detailed Operation Traces

**Insertion Process:**
```go
list := &CircularLinkedList{}
// head = nil

list.Insert(1)
// head -> [1|next] where next points to itself
//         ^____________|

list.Insert(2)
// head -> [1|next] -> [2|next]
//         ^__________________|

list.Insert(3)
// head -> [1|next] -> [2|next] -> [3|next]
//         ^___________________________|
```

**Deletion Complexity:**
```go
// Deleting head node (data = 1):
// Before: head -> [1] -> [2] -> [3] -> [1]...
// Process: 
// 1. Find tail node (traverse until node.next == head)
// 2. Update head = head.next
// 3. Update tail.next = new head
// After: head -> [2] -> [3] -> [2]...
```

#### Edge Cases and Challenges
1. **Single Node**: Node points to itself
2. **Head Deletion**: Requires updating tail's next pointer
3. **Empty List**: Proper nil handling
4. **Infinite Loops**: Traversal termination conditions

#### Real-world Applications
- Round-robin scheduling in operating systems
- Circular buffers in embedded systems
- Multi-player games (turn rotation)
- Resource allocation in networks
- Josephus problem solution

#### Comparison with Linear Linked Lists
- **Memory**: Same per-node overhead
- **Traversal**: No natural end point
- **Operations**: More complex due to circularity
- **Use Cases**: Cyclic processes vs linear sequences

---

## Tree Data Structures

### Binary Tree

**File:** `/datastructures/tree/Tree.go`

#### Overview and Properties
A binary tree is a hierarchical data structure where each node has at most two children (left and right). This implementation focuses on tree traversal algorithms.

#### Memory Layout
```
        4
       / \
      2   6
     / \ / \
    1  3 5  7

Nodes in memory (not necessarily contiguous):
Node{Value: 4, Left: *Node2, Right: *Node6}
Node{Value: 2, Left: *Node1, Right: *Node3}
Node{Value: 6, Left: *Node5, Right: *Node7}
Node{Value: 1, Left: nil, Right: nil}
...
```

#### Implementation Analysis
- **Node Structure**: `Value int`, `Left *Node`, `Right *Node`
- **No Tree Wrapper**: Direct node manipulation
- **Traversal Methods**: In-order, pre-order, post-order
- **Recursive Approach**: All traversals use recursion

#### Traversal Algorithms

**1. In-Order Traversal (Left-Root-Right):**
```
Algorithm: inOrderTraversal(node)
1. if node == nil: return
2. inOrderTraversal(node.Left)
3. process(node.Value)
4. inOrderTraversal(node.Right)

Result: 1 2 3 4 5 6 7 (sorted order for BST)
```

**2. Pre-Order Traversal (Root-Left-Right):**
```
Algorithm: preOrderTraversal(node)
1. if node == nil: return
2. process(node.Value)
3. preOrderTraversal(node.Left)
4. preOrderTraversal(node.Right)

Result: 4 2 1 3 6 5 7 (root first)
```

**3. Post-Order Traversal (Left-Right-Root):**
```
Algorithm: postOrderTraversal(node)
1. if node == nil: return
2. postOrderTraversal(node.Left)
3. postOrderTraversal(node.Right)
4. process(node.Value)

Result: 1 3 2 5 7 6 4 (children before parent)
```

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| In-Order    | O(n)          | O(h)            | h = height of tree |
| Pre-Order   | O(n)          | O(h)            | Recursive stack space |
| Post-Order  | O(n)          | O(h)            | Deep recursion possible |
| Construction| O(n)          | O(n)            | Manual tree building |

#### Detailed Operation Traces

**In-Order Traversal Execution:**
```
Tree:    4
        / \
       2   6
      / \ / \
     1  3 5  7

Call Stack Trace:
inOrder(4)
├── inOrder(2)
│   ├── inOrder(1)
│   │   ├── inOrder(nil) → return
│   │   ├── print(1)
│   │   └── inOrder(nil) → return
│   ├── print(2)
│   └── inOrder(3)
│       ├── inOrder(nil) → return
│       ├── print(3)
│       └── inOrder(nil) → return
├── print(4)
└── inOrder(6)
    ├── inOrder(5)
    │   ├── inOrder(nil) → return
    │   ├── print(5)
    │   └── inOrder(nil) → return
    ├── print(6)
    └── inOrder(7)
        ├── inOrder(nil) → return
        ├── print(7)
        └── inOrder(nil) → return

Output: 1 2 3 4 5 6 7
```

#### Tree Properties Analysis
- **Height**: Maximum depth from root to leaf
- **Balance**: Difference in heights of subtrees
- **Complete**: All levels filled except possibly last
- **Full**: Every node has 0 or 2 children

#### Real-world Applications
- Expression trees in compilers
- Decision trees in machine learning
- File system hierarchies
- XML/HTML document structure
- Huffman coding trees

---

### Binary Search Tree

**File:** `/datastructures/binarysearchtree/binarysearchtree.go`

#### Overview and Properties
A Binary Search Tree (BST) is a binary tree with the ordering property: for each node, all values in the left subtree are less than the node's value, and all values in the right subtree are greater.

#### Memory Layout
```
BST Structure (values: 8,3,10,1,6,14,4,7,13):
        8
       / \
      3   10
     / \    \
    1   6    14
       / \   /
      4   7 13

Property: Left < Root < Right (for each subtree)
```

#### Implementation Analysis
- **BST Wrapper**: Contains root pointer
- **Ordering Property**: Maintained during insertion
- **Recursive Operations**: Both insert and search
- **No Deletion**: Not implemented in current version

#### Operation Complexity Analysis
| Operation | Average Case | Worst Case | Best Case | Notes |
|-----------|--------------|------------|-----------|-------|
| Insert    | O(log n)     | O(n)       | O(1)      | Depends on tree balance |
| Search    | O(log n)     | O(n)       | O(1)      | Skewed tree worst case |
| Space     | O(n)         | O(n)       | O(n)      | One node per element |

#### Detailed Operation Traces

**Insertion Sequence:**
```go
bst := &BST{}
values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}

Insert(8):  [8]

Insert(3):      8
               /
              3

Insert(10):     8
               / \
              3   10

Insert(1):      8
               / \
              3   10
             /
            1

Insert(6):      8
               / \
              3   10
             / \
            1   6

Insert(14):     8
               / \
              3   10
             / \   \
            1   6   14

... (complete tree shown above)
```

**Search Operation Trace:**
```go
Search(6) in the BST:
1. Start at root (8): 6 < 8, go left
2. At node (3): 6 > 3, go right  
3. At node (6): 6 == 6, found! return true

Search(5) in the BST:
1. Start at root (8): 5 < 8, go left
2. At node (3): 5 > 3, go right
3. At node (6): 5 < 6, go left
4. At nil: not found, return false
```

#### BST Properties and Invariants
1. **Ordering Property**: Left < Root < Right
2. **In-Order Traversal**: Produces sorted sequence
3. **Unique Values**: No duplicates (as implemented)
4. **Dynamic Structure**: Can become unbalanced

#### Performance Analysis

**Balanced vs Unbalanced Trees:**
```
Balanced BST (O(log n)):     Skewed BST (O(n)):
        8                           1
       / \                           \
      4   12                          2
     / \ / \                           \
    2  6 10 14                          3
                                         \
                                          4
                                           \
                                            5
```

#### Variants and Optimizations
1. **Self-Balancing Trees**: AVL, Red-Black trees
2. **B-Trees**: For disk-based storage
3. **Splay Trees**: Recently accessed nodes move to root
4. **Treap**: Randomized BST

#### Real-world Applications
- Database indexing
- Expression parsing
- File system organization
- Priority queues (with balancing)
- Symbol tables in compilers

#### Comparison with Other Search Structures
- **vs Hash Table**: Ordered traversal, range queries
- **vs Array**: Dynamic size, O(log n) vs O(n) insertion
- **vs Trie**: General data vs string-specific

---

### Trie (Prefix Tree)

**File:** `/datastructures/Trie/Trie.go`

#### Overview and Properties
A Trie is a tree-like data structure used for storing strings efficiently. Each node represents a character, and paths from root to nodes represent prefixes. It's optimized for prefix-based operations.

#### Memory Layout
```
Trie containing ["apple", "app", "bat", "bath", "batman", "cat"]:

            root
           / | \
          a  b  c
         /   |   \
        p    a    a
       /     |     \
      p      t      t(end)
     /       |
    l        h(end)
   /         |
  e(end)     m
            /
           a
          /
         n(end)

Each node: {children: map[rune]*TrieNode, isEndOfWord: bool}
```

#### Implementation Analysis
- **Node Structure**: Map of children + end-of-word flag
- **Character Mapping**: Uses `map[rune]*TrieNode` for Unicode support
- **Path Representation**: Root-to-node path forms string prefix
- **Space Optimization**: Shared prefixes stored once

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Insert    | O(m)          | O(m×ALPHABET)   | m = word length |
| Search    | O(m)          | O(1)            | Exact word lookup |
| StartsWith| O(m)          | O(1)            | Prefix existence |
| Space     | O(TOTAL_CHARS)| O(TOTAL_CHARS)  | Shared prefixes |

#### Detailed Operation Traces

**Insertion Process:**
```go
trie := NewTrie()

Insert("app"):
root -> 'a' -> 'p' -> 'p'(end)

Insert("apple"):
root -> 'a' -> 'p' -> 'p'(end) -> 'l' -> 'e'(end)
                   (shared prefix)

Insert("bat"):
root -> 'a' -> 'p' -> 'p'(end) -> 'l' -> 'e'(end)
     -> 'b' -> 'a' -> 't'(end)

Complete structure after all insertions:
root
├── 'a' -> 'p' -> 'p'(end) -> 'l' -> 'e'(end)
├── 'b' -> 'a' -> 't'(end) -> 'h'(end) -> 'm' -> 'a' -> 'n'(end)
└── 'c' -> 'a' -> 't'(end)
```

**Search Operation:**
```go
Search("app"):
1. root -> children['a'] (exists)
2. 'a' -> children['p'] (exists)  
3. 'p' -> children['p'] (exists)
4. 'p' -> isEndOfWord == true → return true

Search("ap"):
1. root -> children['a'] (exists)
2. 'a' -> children['p'] (exists)
3. 'p' -> isEndOfWord == false → return false
```

**StartsWith Operation:**
```go
StartsWith("bat"):
1. root -> children['b'] (exists)
2. 'b' -> children['a'] (exists)
3. 'a' -> children['t'] (exists)
4. All characters found → return true

StartsWith("xyz"):
1. root -> children['x'] (not exists) → return false
```

#### Memory Efficiency Analysis

**Space Comparison:**
```
Traditional Storage:        Trie Storage:
["app", "apple"]           root
                          /
Words: 8 characters       a(1)
Storage: 8 chars          |
                         p(1)  
                         |
                         p(1) - shared!
                         |
                         l(1)
                         |
                         e(1)
                         
Total: 5 unique chars vs 8 individual chars
```

#### Trie Variants and Optimizations
1. **Compressed Trie (Radix Tree)**: Merge chains of single-child nodes
2. **Suffix Trie**: For substring searches
3. **Ternary Search Tree**: Space-efficient alternative
4. **AC Automaton**: Multiple pattern matching

#### Real-world Applications
- Autocomplete/suggestions in search engines
- Spell checkers and word validation
- IP routing tables (longest prefix matching)
- Dictionary implementations
- DNA sequence analysis
- Phone number prefix matching

#### Performance Characteristics
- **Prefix Operations**: Exceptionally fast
- **Memory Usage**: Can be high for sparse data
- **Cache Performance**: Good for common prefixes
- **Unicode Support**: Handles international characters

#### Comparison with Other String Structures
- **vs Hash Table**: Prefix queries, ordered traversal
- **vs BST**: Specialized for strings, prefix operations
- **vs Suffix Array**: Different use cases, space trade-offs

---

## Heap Data Structures

### Min Heap

**File:** `/datastructures/minheap/MinHeap.go`

#### Overview and Properties
A Min Heap is a complete binary tree where each parent node is smaller than or equal to its children. The smallest element is always at the root. This implementation uses an array-based representation.

#### Memory Layout
```
Heap Array: [1, 2, 5, 3, 7, 8]
Index:       0  1  2  3  4  5

Tree Representation:
        1
       / \
      2   5
     / \ /
    3  7 8

Array Index Relationships:
- Parent of index i: (i-1)/2
- Left child of i: 2*i + 1  
- Right child of i: 2*i + 2
```

#### Implementation Analysis
- **Storage**: Dynamic array `[]int` for complete binary tree
- **Heap Property**: Parent ≤ children (min-heap)
- **Operations**: Insert with heapify-up, extract with heapify-down
- **Complete Tree**: Always fills left to right, level by level

#### Operation Complexity Analysis
| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Insert    | O(log n)      | O(1)            | Bubble up operation |
| Extract   | O(log n)      | O(1)            | Bubble down operation |
| Peek      | O(1)          | O(1)            | Access root element |
| Heapify   | O(log n)      | O(1)            | Up or down operation |
| Build     | O(n)          | O(n)            | Bottom-up construction |

#### Detailed Operation Traces

**Insert Operation (heapifyUp):**
```go
h := &MinHeap{}
Insert sequence: [5, 3, 8, 1, 2, 7]

Insert(5): [5]
           5

Insert(3): [3, 5]  (3 < 5, swap)
           3
          /
         5

Insert(8): [3, 5, 8]
           3
          / \
         5   8

Insert(1): [1, 3, 8, 5]  (1 bubbles up to root)
Step 1: [3, 5, 8, 1]  (insert at end)
Step 2: [3, 1, 8, 5]  (1 < 5, swap with parent)
Step 3: [1, 3, 8, 5]  (1 < 3, swap with parent)
           1
          / \
         3   8
        /
       5

Insert(2): [1, 2, 8, 5, 3]
Step 1: [1, 3, 8, 5, 2]  (insert at end)
Step 2: [1, 2, 8, 5, 3]  (2 < 3, swap with parent)
           1
          / \
         2   8
        / \
       5   3

Insert(7): [1, 2, 7, 5, 3, 8]
           1
          / \
         2   7
        / \ /
       5  3 8
```

**Extract Operation (heapifyDown):**
```go
// Starting heap: [1, 2, 7, 5, 3, 8]
Extract(): returns 1

Step 1: Move last element to root: [8, 2, 7, 5, 3]
Step 2: Heapify down from root:
   8 > min(2, 7) = 2, swap with left child: [2, 8, 7, 5, 3]
   8 > min(5, 3) = 3, swap with right child: [2, 3, 7, 5, 8]

Final heap: [2, 3, 7, 5, 8]
           2
          / \
         3   7
        / \
       5   8
```

#### Heap Property Maintenance

**HeapifyUp Algorithm:**
```
heapifyUp(index):
1. while index > 0 and parent(index) > current(index):
2.    swap(parent(index), index)
3.    index = parent(index)
```

**HeapifyDown Algorithm:**
```
heapifyDown(index):
1. while hasChildren(index):
2.    smallestChild = indexOfSmallestChild(index)
3.    if current(index) <= smallestChild: break
4.    swap(index, smallestChild)
5.    index = smallestChild
```

#### Memory Efficiency
- **Array Representation**: No pointers needed, cache-friendly
- **Complete Tree**: No wasted array slots
- **Dynamic Sizing**: Grows as needed with Go slices

#### Real-world Applications
- Priority queues (process scheduling)
- Dijkstra's shortest path algorithm
- Huffman coding
- k-smallest/largest element problems
- Heap sort algorithm
- Event simulation systems

---

### Max Heap

**File:** `/datastructures/maxheap/MaxHeap.go`

#### Overview and Properties
A Max Heap is the inverse of a Min Heap - each parent node is greater than or equal to its children. The largest element is always at the root. Uses the same array-based complete binary tree structure.

#### Memory Layout
```
Heap Array: [8, 7, 5, 3, 2, 1]
Index:       0  1  2  3  4  5

Tree Representation:
        8
       / \
      7   5
     / \ /
    3  2 1

Same index relationships as Min Heap
```

#### Implementation Analysis
- **Heap Property**: Parent ≥ children (max-heap)
- **Structure**: Identical to min heap, different comparison
- **Operations**: Insert (bubbleUp), Extract (bubbleDown)
- **Method Names**: bubbleUp/bubbleDown vs heapifyUp/heapifyDown

#### Operation Complexity Analysis
Same as Min Heap - all operations have identical complexity:
| Operation | Time Complexity | Space Complexity |
|-----------|----------------|------------------|
| Insert    | O(log n)      | O(1)            |
| Extract   | O(log n)      | O(1)            |
| Peek      | O(1)          | O(1)            |

#### Detailed Operation Traces

**Insert Operation (bubbleUp):**
```go
mh := &MaxHeap{}
Insert sequence: [5, 3, 8, 1, 2, 7]

Insert(5): [5]
           5

Insert(3): [5, 3]  (5 > 3, no swap needed)
           5
          /
         3

Insert(8): [8, 3, 5]  (8 > 5, bubble up to root)
Step 1: [5, 3, 8]  (insert at end)
Step 2: [8, 3, 5]  (8 > 5, swap with parent)
           8
          / \
         3   5

Insert(1): [8, 3, 5, 1]
           8
          / \
         3   5
        /
       1

Insert(2): [8, 3, 5, 1, 2]
           8
          / \
         3   5
        / \
       1   2

Insert(7): [8, 7, 5, 1, 2, 3]  (7 bubbles up)
Step 1: [8, 3, 5, 1, 2, 7]  (insert at end)
Step 2: [8, 7, 5, 1, 2, 3]  (7 > 3, swap with parent)
           8
          / \
         7   5
        / \ /
       1  2 3
```

**Extract Operation:**
```go
// Starting: [8, 7, 5, 1, 2, 3]
Extract(): returns 8

Step 1: [3, 7, 5, 1, 2]  (move last to root)
Step 2: Bubble down 3:
   3 < max(7, 5) = 7, swap: [7, 3, 5, 1, 2]
   3 < max(1, 2) = 2, swap: [7, 2, 5, 1, 3]

Final: [7, 2, 5, 1, 3]
       7
      / \
     2   5
    / \
   1   3
```

#### Key Differences from Min Heap
1. **Comparison Direction**: > instead of <
2. **Root Property**: Maximum vs minimum element
3. **Use Cases**: Typically different applications

#### Real-world Applications
- Priority queues (high priority first)
- Job scheduling (highest priority jobs)
- Finding k-largest elements
- Heap sort (descending order)
- Maximum sliding window problems
- Top-k frequent elements

#### Heap Comparison Summary
| Aspect | Min Heap | Max Heap |
|--------|----------|----------|
| Root | Minimum element | Maximum element |
| Parent-Child | Parent ≤ Child | Parent ≥ Child |
| Extract | Smallest first | Largest first |
| Use Case | k-smallest, ascending sort | k-largest, descending sort |

---

## Cross-Structure Analysis and Comparisons

### Performance Comparison Matrix

| Structure | Access | Insert | Delete | Search | Space |
|-----------|--------|--------|--------|--------|-------|
| Stack | O(1) top | O(1) | O(1) | O(n) | O(n) |
| Queue | O(1) front | O(1) | O(n)* | O(n) | O(n) |
| Circular Queue | O(1) front | O(1) | O(1) | O(n) | O(n) |
| Deque | O(1) ends | O(1)/O(n)** | O(1)/O(n)** | O(n) | O(n) |
| Singly Linked | O(n) | O(1) head | O(n) | O(n) | O(n) |
| Doubly Linked | O(n) | O(1) ends | O(1)*** | O(n) | O(n) |
| Circular Linked | O(n) | O(n) | O(n) | O(n) | O(n) |
| Binary Tree | O(n) | O(n) | O(n) | O(n) | O(n) |
| Binary Search Tree | O(log n)**** | O(log n)**** | O(log n)**** | O(log n)**** | O(n) |
| Trie | O(m) | O(m) | O(m) | O(m) | O(TOTAL) |
| Min/Max Heap | O(1) root | O(log n) | O(log n) | O(n) | O(n) |

*Simple queue implementation  
**Depends on end (front=O(n), back=O(1))  
***When node reference is known  
****Average case, O(n) worst case for unbalanced BST

### Use Case Decision Matrix

| Need | Recommended Structure | Alternative |
|------|----------------------|-------------|
| LIFO access | Stack | Deque |
| FIFO access | Circular Queue | Queue |
| Both ends access | Deque | Doubly Linked List |
| Ordered data | BST | Sorted Array |
| String prefixes | Trie | Hash Table |
| Priority processing | Heap | BST |
| Circular processing | Circular Linked List | Circular Queue |
| Memory efficiency | Array-based structures | Linked structures |
| Fast arbitrary access | Array | N/A for these structures |

### Memory Layout Comparison

**Contiguous Memory (Array-based):**
- Stack, Queue, Circular Queue, Deque, Heaps
- Better cache performance
- Fixed or dynamic sizing
- Memory efficient (no pointers)

**Non-contiguous Memory (Pointer-based):**
- All Linked Lists, Trees, Trie
- Dynamic sizing
- Memory overhead for pointers
- Flexible structure modifications

---

## Conclusion

This comprehensive analysis covers 12 fundamental data structures implemented in Go, providing insights into their design decisions, performance characteristics, and practical applications. Each structure serves specific use cases, and understanding their trade-offs is crucial for optimal algorithm design and system performance.

The implementations demonstrate core computer science concepts while highlighting Go-specific features like slices, maps, and pointer semantics. The analysis reveals both strengths and limitations of each approach, guiding developers in making informed architectural decisions.

Key takeaways:
1. **Choose structures based on access patterns** (LIFO/FIFO/random)
2. **Consider memory vs performance trade-offs** (arrays vs linked structures)
3. **Understand worst-case vs average-case complexity**
4. **Leverage structure-specific optimizations** (BST ordering, Trie prefixes, Heap properties)
5. **Account for real-world constraints** (cache performance, memory usage, simplicity)

These implementations provide a solid foundation for understanding data structures and can be extended with additional operations, optimizations, and variants as needed for specific applications.