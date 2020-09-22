<?php

/**
 * 自动加载class
 */
require_once ("../util/common.php");
spl_autoload_register(function ($class_name) {
    require_once $class_name . '.php';
});

/**
 * 测试
 */
TestLinkedList();
// TestLinkedListStack();

function TestLinkedList() {
    $linkedList = new LinkedList();
    for ($i = 0; $i < 5; $i++) {
        $linkedList->addLast($i);
        echo $linkedList->toString();
    }

    $linkedList->add(2, 666);
    echo $linkedList->toString();

    $linkedList->remove(2);
    echo $linkedList->toString();

    $linkedList->removeFirst();
    echo $linkedList->toString();

    $linkedList->removeLast();
    echo $linkedList->toString();
}

function TestLinkedListStack()
{
    $stack = new LinkedListStack();
    for ($i = 0; $i < 5; $i++) {
        $stack->push($i);
        echo $stack->toString();
    }

    $stack->pop();
    echo $stack->toString();
}
