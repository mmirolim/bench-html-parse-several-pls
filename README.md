##Benchmark html parsing with several programming languages

1. Go
2. JavaScript(V8)
3. PHP(PHP-FPM)

####Configuration
1. Go
    - go compiler required [golang.org](www.golang.org)
    - ```go run parse-pcuz.go```

2. JavaScript
    - nodejs required [nodejs.org](www.nodejs.org)
    - chreerio and request npm packages should be installed
    - ```node parse-pcuz.js```

3. PHP
    - PHP-FPM
    - "simple_html_dom.php" parse library used and included
    - ```php parse-pcuz.php```

Parsing target is http://www.pc.uz/trade/orgs/cat*.

By default whole list in one go parsed from

http://www.pc.uz/trade/orgs/cat1013?&sort=0&limit=10000

###Results
1. Go - 30ms (go 1.3.3)
2. JavaScript - 27ms (node 10.32)
3. PHP - 20ms (php 5.4) (traversing and saving)

Most of the time spent in file write and php version timer starts after
loading and parsin done.

*benchmark does not test multithreaded execution.

###@TODO
1. Modify html parsing php lib to measure execution time more precisely
2. Add Scala version
3. Measure separately parsing, json encoding and file write
