<?php
/**
 * 官方：
 * PHP 中的数组实际上是一个有序映射。映射是一种把 values 关联到 keys 的类型。
 * 此类型在很多方面做了优化，因此可以把它当成真正的数组，或列表（向量），散列表（是映射的一种实现），
 * 字典，集合，栈，队列以及更多可能性。由于数组元素的值也可以是另一个数组，树形结构和多维数组也是允许的。
 */

$array = array(
    "name" => "Tester",
    "Age"  => 20,
);
$array1 = [
    "name" => "Tester",
    "Age"  => 20,
];

// 常用 array 内置方法

//1.用回调函数过滤数组中的单元 返回：array
array_filter($array);

//2.交换数组中的键和值 返回：array
array_flip($array);

//3.检查数组里是否有指定的键名或索引 返回：bool
array_key_exists("name", $array);

//4.返回数组中部分的或所有的键名 返回：array
array_keys($array);

//5.合并一个或多个数组
array_merge($array, $array1);

//6.对多个数组或多维数组进行排序
array_multisort($array, $array1);

//7.弹出数组最后一个单元（出栈）
array_pop($array);

//8.将一个或多个单元压入数组的末尾（入栈）
array_push($array, "111");

//9.从数组中随机取出一个或多个单元
array_rand($array, 2);

//10.使用传递的数组替换第一个数组的元素
array_replace($array);

//11.返回单元顺序相反的数组
array_reverse($array);

//12.在数组中搜索给定的值，如果成功则返回首个相应的键名
array_search("111", $array);

//13.将数组开头的单元移出数组
array_shift($array);

//14.从数组中取出一段
array_slice($array, 0 , 10);

//15.去掉数组中的某一部分并用其它值取代
array_splice();

//16.对数组中所有值求和
array_sum($array);

//17.移除数组中重复的值
array_unique($array);

//18.在数组开头插入一个或多个单元
array_unshift($array, "add");

//19.返回数组中所有的值
array_values($array);


/**
 * 其他数组相关方法
 * 
 * arsort — 对数组进行逆向排序并保持索引关系
 * asort — 对数组进行排序并保持索引关系
 * compact — 建立一个数组，包括变量名和它们的值
 * count — 计算数组中的单元数目，或对象中的属性个数
 * current — 返回数组中的当前单元
 * each — 返回数组中当前的键／值对并将数组指针向前移动一步
 * end — 将数组的内部指针指向最后一个单元
 * extract — 从数组中将变量导入到当前的符号表
 * in_array — 检查数组中是否存在某个值
 * key_exists — 别名 array_key_exists
 * key — 从关联数组中取得键名
 * krsort — 对数组按照键名逆向排序
 * ksort — 对数组按照键名排序
 * list — 把数组中的值赋给一组变量
 * natcasesort — 用“自然排序”算法对数组进行不区分大小写字母的排序
 * natsort — 用“自然排序”算法对数组排序
 * next — 将数组中的内部指针向前移动一位
 * pos — current 的别名
 * prev — 将数组的内部指针倒回一位
 * range — 根据范围创建数组，包含指定的元素
 * reset — 将数组的内部指针指向第一个单元
 * rsort — 对数组逆向排序
 * shuffle — 打乱数组
 * sizeof — count 的别名
 * sort — 对数组排序
 * uasort — 使用用户自定义的比较函数对数组中的值进行排序并保持索引关联
 * uksort — 使用用户自定义的比较函数对数组中的键名进行排序
 * usort — 使用用户自定义的比较函数对数组中的值进行排序
 */
