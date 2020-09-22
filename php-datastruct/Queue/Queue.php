<?php

class Queue
{
    private $data;
    private $size;

    public function __construct()
    {
        $this->data = array();
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

    public function enqueue($element)
    {
        $this->size++;
        return array_push($this->data, $element);
    }

    public function dequeue()
    {
        $this->size--;
        return array_shift($this->data);
    }

    public function getFront()
    {
        return current($this->data);
    }

    public function toString()
    {
        $str = "Queue: Front [";
        for ($i = 0; $i < $this->size; $i++) {
            $str .= $i;
            if ($i != $this->size - 1)
                $str .= ", ";
        }
        $str .= "] tail\n";
        return $str;
    }
}

require_once ("../util/common.php");

TestQueue();

function TestQueue(){
    $timer = new Timer();
    
    $count = 100000;

    $timer->start();
    $queue = new Queue();
    for ($i = 0; $i < $count; $i++) {
        $queue->enqueue($i);
        // if ($i % 3 == 2) {
        //     $queue->dequeue();
        //     // echo $queue->toString();
        // }
    }
    $timer->stop();
    echo $timer->getTime();

    $timer->start();
    $queue1 = array();
    for ($i = 0; $i < $count; $i++) {
        array_push($queue1, $i);
        // if ($i % 3 == 2) {
        //     array_shift($queue1);
        // }
    }
    $timer->stop();
    echo $timer->getTime();
}

/**
 * 总结 ：
 * php中队列（Queue）可以直接使用array （array_push, array_shift, current()或者array[0] )来实现。
 * 出队效率不好，可以选择使用顺序队列
 * 
 * 5000000数据量
 * Queue : Run time : 2.670691967 s
 * array ：Run time : 2.6555628777 s
 * 
 * 算法时间复杂度：
 * getSize ：O(1)
 * isEmpty : O(1)
 * enqueue : O(1)
 * dequeue  : O(n) -> 普通队列 效率低下，因为需要每次需要 删除第一个元素之后，需要遍历所有的元素，往左挪一个位置。
 * getFront : O(1)
 * 
 */
