<?php
/**
 * 链表实现队列
 *
 * 主要操作：
 * enqueue ：入队 O(1) or O(n)
 * dequeue ：出队 O(1) or O(n)
 * getFront ：获取队首元素 O(1) or O(n)
 * 
 * getSize ：O(1)
 * isEmpty ：O(1)
 * 
 * 总结：
 * 链表实现队列两种办法：
 * 1. enqueue：addLast() O(n) | dequeue: removeFirst() O(1) | getFront: getFrist() O(1)
 * 2. enqueue: addFirst() O(1) | dequeue: removeLast() O(n) | getFront: getLast() O(n)
 * 
 * 第一种方法：入队 慢， 出队 快
 * 第二种方法：入队 块， 出队 慢
 */
class LinkedListQueue extends LinkedList
{
    public function __construct($e=null, $next=null)
    {
        parent::__construct();
    }

    public function getSize()
    {
        return $this->size;
    }

    public function isEmpty()
    {
        return $this->size == 0;
    }

    public function enqueue($element)
    {
        $this->addLast($element);

        // 第二种方法
        // $this->addFirst($element);
    }

    public function dequeue()
    {
        return $this->removeFirst();
        
        // 第二种方法
        // return $this->removeLast();
    }

    public function getFirst()
    {
        return $this->getFrist();

        // 第二种方法
        // return $this->getLast();
    }

    public function toString()
    {
        $str = "LinkedListQueue: top ";
        $str .= parent::toString();
        return $str;
    }
}