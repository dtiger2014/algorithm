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
        $this->data[$this->tail] = $e;
        $this->tail++;
        $this->size++;
    }

    public function dequeue()
    {
        if ($this->isEmpty())
            throw new Exception("Loop Queue is empty");

        $delElement = $this->data[$this->front];
        unset($this->data[$this->front]);
        $this->front++;
        $this->size--;

        return $delElement;
    }

    public function getFront() {
        if ($this->isEmpty())
            throw new Exception("Loop Queue is empty");
        return $this->data[$this->front];
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
echo $timer->getTime();

$mem->stop();
echo $mem->getMem();

/**
 * 总结 ：
 * 循环队列，时间复杂度 是 O(1)。 php数组是映射表，所以 只需要控制 front 和tail  就可以 达到O(1)的时间复杂度的LoopQueue。
 * 而不需要重新遍历赋值。效率会高一些。
 * 
 * 10000000 数据量
 * LoopQueue : Run time : 1.6590909958 s
 * 
 * 算法时间复杂度：
 * getSize ：O(1)
 * isEmpty : O(1)
 * enqueue : O(1)
 * dequeue  : O(1)
 * getFront : O(1)
 * 
 */