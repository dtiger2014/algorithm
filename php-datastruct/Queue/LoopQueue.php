<?php

class LoopQueue
{
    private $data, $size, $front, $tail;
    private static $capacity;

    public function __construct()
    {
        $this->data = array();
        $this->size = 0;
        $this->front = 0;
        $this->tail = 0;

        self::$capacity = 10;
    }

    public function getSize()
    {
        return $this->size;
    }

    public function isEmpty()
    {
        return $this->front == $this->tail;
    }

    public function enqueue($e)
    {
        if (($this->tail + 1) % self::$capacity == $this->front)
            $this->resize(self::$capacity * 2);

        $this->data[$this->tail] = $e;
        $this->tail = ($this->tail + 1) % self::$capacity;
        $this->size++;
    }

    public function dequeue()
    {
        if ($this->isEmpty())
            throw new Exception("Loop Queue is empty");

        $delElement = $this->data[$this->front];
        $this->data[$this->front] = null;
        $this->front = ($this->front + 1) % self::$capacity;
        $this->size--;

        if ($this->size == self::$capacity / 4 && self::$capacity / 2 != 0)
            $this->resize(self::$capacity / 2);

        return $delElement;
    }

    public function getFront() {
        if ($this->isEmpty())
            throw new Exception("Loop Queue is empty");
        return $this->data[$this->front];
    }

    private function resize($cap)
    {
        $newData = array();

        for($i = 0 ; $i < $this->size ; $i++)
            $newData[$i] = $this->data[($i + $this->front) % self::$capacity];
        $this->data = $newData;

        self::$capacity *= $cap;
        $this->front = 0;
        $this->tail = $this->size;
    }

    public function toString()
    {
        $str = "front [";
        for ($i = $this->front; $i != $this->tail; $i = ($i + 1) % self::$capacity) {
            $str .= $this->data[$i];
            if (($i + 1) % self::$capacity != $this->tail)
                $str .= ", ";
        }
        $str .= "] tail\n";
        return $str;
    }
}


require_once ("../util/common.php");

$timer = new Timer();
$timer->start();

$mem = new Memorys();
$mem->start();

$count = 10000000;

$queue = new LoopQueue();
for ($i = 0; $i < $count; $i++) {
    $queue->enqueue($i);
    if ($i % 3 == 2) {
        $queue->dequeue();
    }
}


$timer->stop();
$mem->stop();

echo $timer->getTime();
echo $mem->getMem();

/**
 * 总结 ：
 * 循环队列，时间复杂度 是 O(1)，但是在 扩大array 容量的时候，会从新创建一个array 遍历赋值到新的array中。时间复杂度为O(1)均摊
 * 
 * 10000000 数据量
 * LoopQueue : Run time : 2.3538749218 s
 * 
 * 算法时间复杂度：
 * getSize ：O(1)
 * isEmpty : O(1)
 * enqueue : O(1) 均摊
 * dequeue  : O(1) 均摊
 * getFront : O(1)
 * 
 */

