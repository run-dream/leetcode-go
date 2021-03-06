### 定义
双指针，顾名思义，就是利用两个指针去遍历。可分为两类
1. 快慢指针 主要解决链表中的问题，比如典型的判定链表中是否包含环
2. 左右指针 主要解决数组（或者字符串）中的问题，比如二分查找。

### 应用
#### 快慢指针
1. 判定链表中是否含有环
  两个指针，一个跑得快，一个跑得慢。如果不含有环，跑得快的那个指针最终会遇到 null，说明链表不含环；如果含有环，快指针最终会超慢指针一圈，和慢指针相遇，说明链表含有环。
2. 已知链表中含有环，返回这个环的起始位置 
    当快慢指针相遇时，让其中任一个指针指向头节点，然后让它俩以相同速度前进，再次相遇时所在的节点位置就是环开始的位置
3. 寻找链表的中点
    快指针一次前进两步，慢指针一次前进一步，当快指针到达链表尽头时，慢指针就处于链表的中间位置。
4. 寻找链表的倒数第 k 个元素
    让快指针先走 k 步，然后快慢指针开始同速前进。这样当快指针走到链表末尾 null 时，慢指针所在的位置就是倒数第 k 个链表节点

#### 左右指针
右指针在数组中实际是指两个索引值，一般初始化为 left = 0, right = nums.length - 1 
1. 二分查找
2. 两数之和
3. 反转数组
4. 滑动窗口算法
