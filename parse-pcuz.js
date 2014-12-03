'use strict';
var fs = require('fs');
var req = require('request');
var ch = require('cheerio');

var  url = 'http://www.pc.uz/trade/orgs/cat1013';
var  a = url.split("/");
// file name to save json data
var fname = "pcuz-cat-parse-" + a[a.length - 1] + "-js.json";
// get all orgs in one go
url += "?&sort=0&limit=10000";

console.log("Let's start loading and parsing" + url);

// get html page and process
req(url, function(err, res, html) {

    if(err) {
        console.log(err);
    }
    var $ = ch.load(html);
    var start = new Date();
    var orgs = [];
    var names = [];
    var tels = [];

    // get org names
    $('table[width="98%"]').filter(function() {
        var data = $(this);
        data.find('a[style="font-size:11pt; text-decoration:none;"]').filter(function() {
            names.push($(this).text());
        });
    });

    // get all phone numbers
    $('td.line_about').filter(function() {
        var data = $(this);
        var tel = data.children('[style="padding-bottom:1px"]').first().text();
        tels.push(tel);
    });

    // zip results
    for (var i = 0; i < names.length; i++) {
        orgs.push({name: names[i], tel: tels[i]});
    }

    // write to file
    fs.writeFile(fname, JSON.stringify(orgs, null, 4), function(err) {
        console.log(err);
    });

    var end = new Date() - start;
    console.log("html page parsed and saved as json object");
    console.info("time taken %dms", end)

});