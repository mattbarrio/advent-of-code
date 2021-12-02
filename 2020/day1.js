const fs = require("fs");

fs.readFile("day1.txt", function (err, data) {
  if (err) throw err;
  const arr = data.toString().split("\n");
  const array = arr.map((i) => Number(i));
  console.log(array);
  for (i in array) {
    for (x in array) {
      for (y in array) {
        const checkSum = array[i] + array[x] + array[y];
        if (checkSum == 2020) {
          const answer = array[i] * array[x] * array[y];
          console.log(answer, array[i], array[x], array[y]);
        }
      }
    }
  }
});
