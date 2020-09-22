<?php

class Memorys
{
    private static $initMem;
    private static $finishMem;

    private $reset = false;

    public function __construct()
    {
        self::$initMem = 0;
        self::$finishMem = 0;
        $this->reset = true;
    }

    public function start()
    {
        if ($this->reset == false)
            return "Memory record not reset\n";

        self::$initMem = memory_get_usage();
        $this->reset = false;
        return "Memory record start!\n";
    }

    public function stop()
    {
        if ($this->reset == true)
            return "Memory record is reset\n";

        self::$finishMem = memory_get_usage();
        $this->reset = true;
        return "Memory record stop!\n";
    }

    public function getMem()
    {
        $useMem = self::$finishMem - self::$initMem;
        return "Init Memory is ".$this->memory_usage(self::$initMem).
                " | Finish Memory is ".$this->memory_usage(self::$finishMem).
                " | Memory Used : ".$this->memory_usage($useMem)."\n";
    }

    private function memory_usage($num) { 
        $memory = ( ! function_exists('memory_get_usage')) ? '0' : round($num/1024/1024, 2).'MB'; 
        return $memory; 
    } 
}