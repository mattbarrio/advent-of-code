const fs = require("fs");

fs.readFile("day3.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  console.log(array);
  var isOpen = 0;
  var right = [1, 3, 5, 7, 1];
  var down = [1, 1, 1, 1, 2];
  var iterationAnswer = [];
  for (r in right) {
    var isTree = 0;
    var startingPos = 0;
    var nextRow = 0;
    for (i in array) {
      const thisPos = startingPos;
      if (i != nextRow) {
        // console.log("skipping row " + i);
        continue;
      }
      if (i == 0) {
        startingPos += right[r];
        var row = array[i].split("");
        var val = row[startingPos];
        // console.log("thisPos " + thisPos, "pos: " + startingPos, "val: " + val);
      } else {
        var row = array[i].split("");
        var val = row[startingPos];
        // console.log(val);
        do {
          row = row.concat(row);
          val = row[startingPos];
          //console.log(row.length);
        } while (val === undefined);
        if (val === "#") {
          isTree += 1;
        } else if (val === ".") {
          isOpen += 1;
        }
        //console.log("thisPos " + thisPos, "pos: " + startingPos, "val: " + val);
        startingPos += right[r];
      }
      nextRow += down[r];
    }
    // console.log(right[r]);
    iterationAnswer = iterationAnswer.concat(isTree);
  }
  var answer = iterationAnswer.reduce(function (a, b) {
    return a * b;
  });
  console.log("iterations: " + iterationAnswer, "answer: " + answer);
});
