<?php
/**
 * 自动加载class
 */
require_once ("../util/common.php");
spl_autoload_register(function ($class_name) {
    require_once $class_name . '.php';
});
