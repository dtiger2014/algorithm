<?php

class Timer
{
    private static $startTime;
    private static $endTime;

    private $reset = false;

    public function __construct()
    {
        self::$startTime = 0;
        self::$endTime = 0;

        $this->reset = true;
    }

    public function start()
    {
        if ($this->reset == false)
            return "Timer not reset\n";

        self::$startTime = microtime(true);
        $this->reset = false;
        return "Timer start\n";
    }

    public function stop()
    {
        if ($this->reset == true)
            return "Timer not reset\n";

        self::$endTime = microtime(true);
        $this->reset = true;
        return "Timer stop\n";
    }

    public function restart(){
        if ($this->reset == true)
            return "Timer not reset\n";

        self::$endTime = 0;
        start();
        return "Timer restart\n";
    }

    public function getTime()
    {
        return "Run time : ".round(self::$endTime - self::$startTime, 10)." s\n";
    }

}