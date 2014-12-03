<?php
include('simple_html_dom.php');

// set url to 
// get all in one go
$url = "http://www.pc.uz/trade/orgs/cat1013?&sort=0&limit=10000";

// file name to store json array";
$fname = "pcuz-parse-php.json";
// get page
$html = file_get_html($url);

// start stopwatch
$start = microtime(true);

$orgs = [];
$names = [];
$tels = [];

// fine orgs names
foreach($html->find('table a[style="font-size:11pt; text-decoration:none;"] strong') as $name) {
    array_push($names, $name->innertext); 
}
// find orgs tels
foreach($html->find('td.line_about') as $div) {
    array_push($tels, $div->find('div[style="padding-bottom:1px"] span', 0)->innertext);
}
// create std obj to serialize it
$org = new stdClass();
// zip names with tels
foreach($names as $k => $name) {
    $orgs[] = array("name"=> $name, "tel" => $tels[$k]);
}

// decode to json with utf-8
$json_data = json_encode($orgs, JSON_UNESCAPED_UNICODE);

// open/create file and write data to it
$file = fopen($fname, "w") or die('Unable to open file');
fwrite($file, $json_data);
fclose($file);

// measure execution tiem in ms
$time = (microtime(true) - $start)*1000;
print("Time taken" + $time + "ms");
