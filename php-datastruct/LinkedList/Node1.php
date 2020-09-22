<?php

class Node1
{
    public $element;
    public $next;
    public $prev;

    public function __construct($e=null, $next=null, $prev=null)
    {
        $this->element = $e;
        $this->next = $next;
        $this->prev = $prev;
    }
}
