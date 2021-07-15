package skiplist

import (
	"fmt"
	"math/rand"
)

// 常量
const (
	// MaxLevel 最高层数
	MaxLevel = 32
	// 每一层的晋升概率
	Probability = 0.25
)

// LevelNode 索引节点
type LevelNode struct {
	forward *Node
	span    int
}

// 新建索引节点
func newLevelNode() *LevelNode {
	return &LevelNode{nil, 0}
}

// Node 跳表节点结构体
type Node struct {
	// 跳表保存的关键字
	ele string
	// 用于排序的分值
	score float64
	// 后续节点
	backward *Node
	// 节点的所有索引节点
	levels []*LevelNode
}

// 新建跳表节点
func newNode(level int, score float64, ele string) *Node {
	levels := make([]*LevelNode, MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		levels[i] = newLevelNode()
	}
	return &Node{ele, score, nil, levels}
}

// SkipList 跳表结构体
type SkipList struct {
	//跳表头节点和尾节点
	head, tail *Node
	//跳表当前层数
	level int
	//跳表长度
	length int
}

// NewSkipList 实例化跳表对象
func NewSkipList() *SkipList {
	//头节点，便于操作
	head := newNode(MaxLevel, 0.0, "")
	return &SkipList{head, nil, 1, 0}
}

// Length 获取跳表长度
func (this *SkipList) Length() int {
	return this.length
}

// Level 获取跳表层级
func (this *SkipList) Level() int {
	return this.level
}

// RandomLevel 获取一个随机的层级 返回值的范围是 [1, MaxLevel]
func randomLevel() int {
	level := 1
	for rand.Float64() < Probability && level < MaxLevel {
		level += 1
	}
	return level
}

// Insert 插入节点到跳表中
func (this *SkipList) Insert(ele string, score float64) *Node {
	if len(ele) == 0 {
		return nil
	}

	cur := this.head
	// 在各个层查找节点的插入位置，也就是目标节点的上一个节点
	update := make([]*Node, MaxLevel)
	// 用来计算span的值
	rank := make([]int, MaxLevel)
	// 遍历完以后, cur 就是之后要插入到第0层的位置的下一个节点
	for i := this.level - 1; i >= 0; i-- {
		if i < this.level-1 {
			rank[i] = rank[i+1]
		}

		for cur.levels[i].forward != nil && (cur.levels[i].forward.score < score || (cur.levels[i].forward.score == score && cur.levels[i].forward.ele < ele)) {
			rank[i] += cur.levels[i].span
			cur = cur.levels[i].forward
		}
		update[i] = cur
	}

	// 通过随机算法获取该节点层数
	level := randomLevel()

	// 如果新节点的层数比表中其他节点的层数都要大
	// 那么初始化表头节点中未使用的层，并将它们记录到 update 数组中
	// 将来也指向新节点
	if level > this.level {
		// 初始化未使用层
		for i := this.level; i < level; i++ {
			rank[i] = 0
			update[i] = this.head
			update[i].levels[i].span = this.length
		}

		// 更新表中节点最大层数
		this.level = level
	}

	// 创建一个新的跳表节点
	node := newNode(level, score, ele)
	// 将前面记录的指针指向新节点，并做相应的设置
	for i := 0; i < level; i++ {
		// 设置新节点的 forward 指针
		node.levels[i].forward = update[i].levels[i].forward
		// 将沿途记录的各个节点的 forward 指针指向新节点
		update[i].levels[i].forward = node
		// 计算新节点跨越的节点数量
		node.levels[i].span = update[i].levels[i].span - (rank[0] - rank[i])
		// 更新新节点插入之后，沿途节点的 span 值
		update[i].levels[i].span = (rank[0] - rank[i]) + 1
	}

	// 未接触的节点的 span 值也需要增一，这些节点直接从表头指向新节点
	for i := level; i < this.level; i++ {
		update[i].levels[i].span++
	}

	// 设置新节点的后退指针
	if update[0] != this.head {
		node.backward = update[0]
	} else {
		node.backward = nil
	}

	// 处理尾节点
	if node.levels[0].forward != nil {
		node.levels[0].forward.backward = node
	} else {
		this.tail = node
	}

	//更新跳表长度
	this.length++

	return node
}

// Delete 删除节点
func (this *SkipList) Delete(ele string, score float64) bool {
	cur := this.head
	// 在各个层查找节点的插入位置，也就是目标节点的上一个节点
	update := make([]*Node, MaxLevel)
	// 按照索引从上向下查找节点
	for i := this.level - 1; i >= 0; i-- {
		for cur.levels[i].forward != nil && (cur.levels[i].forward.score < score || (cur.levels[i].forward.score == score && cur.levels[i].forward.ele < ele)) {
			cur = cur.levels[i].forward
		}
		update[i] = cur
	}

	node := cur.levels[0].forward
	if node != nil && node.score == score && node.ele == ele {
		this.deleteNode(node, update)
		return true
	}
	return false
}

// 删除具体的节点
func (this *SkipList) deleteNode(node *Node, update []*Node) {
	// 更新所有和被删除节点有关的节点的指针，解除它们之间的关系
	for i := 0; i < this.level; i++ {
		// 合并节点
		if update[i].levels[i].forward == node {
			update[i].levels[i].span += node.levels[i].span - 1
			update[i].levels[i].forward = node.levels[i].forward
		} else {
			update[i].levels[i].span -= 1
		}
	}

	// 更新被删除节点的前进和后退指针
	if node.levels[0].forward != nil {
		node.levels[0].forward.backward = node.backward
	} else {
		// 最后一个节点则更新tail指针
		this.tail = node.backward
	}

	// 更新跳跃表最大层数（只在被删除节点是跳跃表中最高的节点时才执行）
	for this.level > 1 && this.head.levels[this.level-1].forward == nil {
		this.level--
	}

	this.length--
}

// isInRange 快速判断是否可能有节点
func (this *SkipList) isInRange(min, max float64) bool {
	if min > max {
		return false
	}
	if this.tail == nil || this.tail.score < min {
		return false
	}
	if first := this.head.levels[0].forward; first == nil || first.score > max {
		return false
	}
	return true
}

// GetRange 根据分值区间查找节点, 简化一下问题，只考虑闭区间
func (this *SkipList) GetRange(min, max float64) (begin, end *Node) {
	if !this.isInRange(min, max) {
		return nil, nil
	}
	cur := this.head.levels[0].forward
	for i := this.level - 1; i >= 0; i-- {
		for cur.levels[i].forward != nil && cur.levels[i].forward.score < min {
			cur = cur.levels[i].forward
		}
	}
	if cur.levels[0].forward == nil {
		begin = cur
	} else {
		begin = cur.levels[0].forward
	}
	cur = this.head.levels[0].forward
	for i := this.level - 1; i >= 0; i-- {
		for cur.levels[i].forward != nil && (cur.levels[i].forward.score <= max) {
			cur = cur.levels[i].forward
		}
	}
	if cur.levels[0].forward == nil && cur.score <= max {
		end = cur
	} else {
		end = cur.backward
	}
	return begin, end
}

func (this *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", this.level, this.length)
}
