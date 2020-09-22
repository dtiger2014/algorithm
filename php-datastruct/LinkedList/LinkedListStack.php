<?php

class LinkedListStack extends LinkedList
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

    public function push($element)
    {
        $this->addFirst($element);
    }

    public function pop()
    {
        return $this->removeFirst();
    }

    public function peek()
    {
        return $this->getFrist();
    }

    public function toString()
    {
        $str = "LinkedListStack: top ";
        $str .= parent::toString();
        return $str;
    }
}