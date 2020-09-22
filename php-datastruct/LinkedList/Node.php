<?php

class Node 
{
    public $element;
    public $next;

    public function __construct($e=null, $next=null)
    {
        $this->element = $e;
        $this->next = $next;
    }
}
