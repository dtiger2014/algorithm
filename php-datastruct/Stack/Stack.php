<?php

class Stack
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

    public function push($element)
    {
        array_push($this->data, $element);
        $this->size++;
    }

    public function pop()
    {
        if ($this->size <= 0){
            return null;
        }
        $this->size--;
        return array_pop($this->data);
    }

    public function peek()
    {
        return end($this->data);
    }

    public function toString()
    {
        $str = "Array Stack: [";
        for ($i = 0; $i < $this->size; $i++){
            $str .= $this->data[$i];
            if ($i != $this->size - 1){
                $str .= ",";
            }
        }
        return $str."] top\n";
    }
}

require_once ("../util/common.php");

TestStack();

function TestStack(){
    $timer = new Timer();
    
    $count = 5000000;
    
    echo $timer->start();
    //Stack
    $stack = new Stack();
    for($i = 0; $i < $count; $i++) {
        $stack->push($i);
    }
    for($i = 0; $i < $count; $i++) {
        $stack->pop();
    }
    echo $timer->stop();
    echo $timer->getTime();

    //array
    echo $timer->start();
    $arr = array();
    for ($i = 0; $i < $count; $i++){
        array_push($arr, $i);
    }
    for ($i = 0; $i < $count; $i++) {
        array_pop($arr);
    }
    echo $timer->stop();
    echo $timer->getTime();
}

/**
 * 总结 ：
 * php中栈（Stack）可以直接使用array( array_push(), array_pop(), end() )进行出入栈。效率会更好。
 * 
 * 5000000数据量
 * Stack : Run time : 0.959608078 s
 * array : Run time : 0.5113089085 s
 * 
 * 算法时间复杂度：
 * getSize ：O(1)
 * isEmpty : O(1)
 * push : O(1)
 * pop  : O(1)
 * peek : O(1)
 * 
 */
?>