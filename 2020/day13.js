const fs = require("fs");
// I never finished this day... totally fucked it up
fs.readFile("day13.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  console.log(array);
  var isOpen = 0;
  //var right = [1, 3, 5, 7, 1];
  //var down = [1, 1, 1, 1, 2];
  var iterationAnswer = [];
  //for (r in right) {
  var isTree = 0;
  var startingPos = 0;
  var nextRow = 0;
  const startTime = array[0] - 1;
  console.log(startTime);
  var row = array[1].split(",");
  console.log(row);
  var lowestVal = 0;
  var busNum = 0;
  var nextTime = parseInt(startTime);

  function runner(startTime, busArr) {
    for (i in busArr) {
      if (busArr[i] === "x") {
        nextTime += 1;
        continue;
      }

      do {
        nextTime += 1;
        evaluatedNum = nextTime / row[i];
        console.log(
          "busNum: " + row[i],
          "nextTime: " + nextTime,
          "evaluatedNum: " + evaluatedNum
        );
      } while (!Number.isInteger(evaluatedNum));

      if (i == 0) {
        lowestVal = nextTime;
        busNum = row[i];
      } else {
        if (nextTime < lowestVal) {
          lowestVal = nextTime;
          busNum = row[i];
        }
      }
    }
  }

  if ((times = runner())) {
  }

  const waitTime = (lowestVal - (startTime + 1)) * busNum;
  console.log(
    "startTime: " + startTime,
    "busNum: " + busNum,
    "lowestVal: " + lowestVal,
    "waitTime: " + waitTime
  );
});
