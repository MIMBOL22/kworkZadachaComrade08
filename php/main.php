<?php
$b = file_get_contents("../message.txt");

$c = explode("\n", $b);

$arr = [];
$max_key = 0;
foreach($c as $v){
	$d = explode(" ",$v);
	$ky = (int)$d[0];
	$arr[$ky] = $d[1];
	
	if($ky > $max_key){
		$max_key = $ky;
	}
}

$key = 0;
$i = 1;
$result1 = [];
while($key < $max_key){
    $key = ($i+1)*$i/2;
    array_push($result1, $arr[$key]);
    $i+=1;
}

$res1_str = implode(" ", $result1);
echo $res1_str;
