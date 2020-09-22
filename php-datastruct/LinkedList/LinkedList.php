<?php
/**
 * LinkedList 链表。包含一个指针（next）单向遍历搜索。线性结构
 * 使用场景：
 *
 * 主要操作：
 * addFirst：O(1)
 * addLast：O(n)
 * add：O(n) avg
 * getFirst：O(1)
 * getLast：O(n)
 * get：O(n) avg
 * set：O(n) avg
 * contains：O(n) avg
 * removeFirst：O(1)
 * removeLast：O(n)
 * remove：O(n) avg
 * removeElement：O(n) avg
 * 
 * getSize ：O(1)
 * isEmpty ：O(1)
 * 
 * 总结：单向链表，线性结构。第一个元素的操作是O(1)，但是要获取后面的元素效率会大幅下降。
 * 可以扩展成 LinkedListStack 链表栈。跟数组栈效率差不多O(1)。
 * 链表是动态结构，可以自动分配内存大小。数组需要初始化大小才可以使用。
 */

class LinkedList
{
    private $dummyHead;
    private $size;

    public function __construct()
    {
        $this->dummyHead = new Node();
        $this->size = 0;
    }

    public function getSize()
    {
        return $this->size;
    }

    public function isEmpty()
    {
        return $this->size == 0;
    }
    
    public function addFirst($element)
    {
        $this->add(0, $element);
    }

    public function addLast($element)
    {
        $this->add($this->size, $element);
    }

    public function add($index, $element)
    {
        if ($index < 0 || $index > $this->size) {
            throw new Exception("Add failed. Illegal index.");
        }
        
        $prev = $this->dummyHead;
        for($i = 0; $i < $index; $i++) {
            $prev = $prev->next;
        }
        
        $prev->next = new Node($element, $prev->next);
        $this->size++;
    }

    public function getFirst()
    {
        return get(0);
    }

    public function getLast()
    {
        return get($this->size - 1);
    }

    public function get($index)
    {
        if ($index < 0 || $index >= $size) {
            throw new Exception("Get failed. Illegal index.");
        }

        $cur = $this->dummyHead->next;
        for ($i = 0; $i < $index; $i++) {
            $cur = $cur->next;
        }
        return $cur->element;
    }

    public function set($index, $element)
    {
        if ($index < 0 || $index >= $size) {
            throw new Exception("Set failed. Illegal index.");
        }

        $cur = $this->dummyHead->next;
        for ($i = 0; $i < $index; $i++) {
            $cur = $cur->next;
        }
        $cur->element = $element;
    }

    public function contains($element)
    {
        $cur = $this->dummyHead->next;
        while($cur != null) {
            if ($cur->element == $element) {
                return true;
            }
            $cur = $cur->next;
        }
        return false;
    }
    
    public function removeFirst() 
    {
        return $this->remove(0);
    }

    public function removeLast()
    {
        return $this->remove($this->size - 1);
    }

    public function remove($index)
    {
        if ($index < 0 || $index >= $this->size) {
            throw new Exception("Remove failed. Index is illegal.");
        }

        $prev = $this->dummyHead;
        for ($i = 0; $i < $index; $i++) {
            $prev = $prev->next;
        }

        $retNode = $prev->next;
        $prev->next = $retNode->next;
        $retNode->next = null;
        $this->size--;
        
        return $retNode->element;
    }

    public function removeElement($element)
    {
        $prev = $this->dummyHead;
        while ($prev->next != null) {
            if ($prev->next->element == $element) {
                break;
            }
            $prev = $prev->next;
        }

        if ($prev->next != null) {
            $delNode = $prev->next;
            $prev->next = $delNode->next;
            $delNode->next = null;
            $this->size--;
        }
    }

    public function toString()
    {
        $str = "";
        for ($cur = $this->dummyHead->next; $cur != null; $cur = $cur->next) {
            $str .= strval($cur->element)." -> ";
        }

        $str .= "NULL\n";
        return $str;
    }
}