const fs = require('fs');

fs.readFile('day2.txt', function(err, data) {
    if(err) throw err;
    const array = data.toString().split("\n");
    //const array = arr.map((i) => Number(i));
    console.log(array);
    var matched = 0;
    for(i in array) {
        //console.log(array[i]);
        const min = parseInt(array[i].split("-")[0]);
        const max = parseInt(array[i].split("-").pop().split(" ")[0]);
        const char = array[i].split(" ")[1].split(":")[0];
        const pass = array[i].split(": ")[1].split('');

        // var re = new RegExp(char,"g");
        // const stripedPass = (pass.match(re) || []).length;
        if (pass[min-1] === char && pass[max-1] !== char) {
            matched += 1;
            console.log("min " + min, "max " + max, "char " + char, "pass " + pass, "pass+1 " + pass[min-1], "pass+1 " + pass[max-1]);
        } else if (pass[min-1] !== char && pass[max-1] === char) {
            matched += 1;
            console.log("min " + min, "max " + max, "char " + char, "pass " + pass, "pass+1 " + pass[min-1], "pass+1 " + pass[max-1]);
        }
    }
    console.log("matched: " + matched);
});