const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data
    .toString()
    .split("\n")
    .map((row) => row.replace(" -> ", ","));

  //console.log("Starting problem_1");
  //problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

const largest_number = (array) => {
  max_array = [];
  for (r in array) {
    int_array = array[r].split(",").map((r) => Number(r));
    max_array[r] = Math.max(...int_array);
  }
  return Math.max(...max_array);
};

const problem_1 = (array) => {
  max_value = largest_number(array);
  // we need to build an array with the largest number of possible coordinates, then create a sub-array inside each row, again with the largest number of possible coordinates
  diagram = new Array(max_value);
  for (let i = 0; i <= max_value; i++) {
    diagram[i] = new Array(max_value).fill(0);
  }

  for (r in array) {
    const [x1, y1, x2, y2] = array[r].split(",");
    // only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2
    if (x1 === x2 || y1 === y2) {
      let xStart = Math.min(x1, x2);
      let yStart = Math.min(y1, y2);
      let xEnd = Math.max(x1, x2);
      let yEnd = Math.max(y1, y2);
      // mark diagram
      for (let x = xStart; x <= xEnd; x++) {
        for (let y = yStart; y <= yEnd; y++) {
          diagram[x][y] += 1;
        }
      }
    }
  }

  // iterate over every item in the array to find overlaps
  let overlaps = 0;
  diagram.forEach((row) => {
    row.forEach((point) => {
      if (point >= 2) overlaps += 1;
    });
  });
  console.log(overlaps);
  // 6225
};

const problem_2 = (array) => {
  max_value = largest_number(array);
  console.log(max_value);
  // we need to build an array with the largest number of possible coordinates, then create a sub-array inside each row, again with the largest number of possible coordinates
  diagram = new Array(max_value);
  for (let i = 0; i <= max_value; i++) {
    diagram[i] = new Array(max_value).fill(0);
  }

  for (r in array) {
    const [x1, y1, x2, y2] = array[r].split(",").map((r) => Number(r));
    for (let x = 0; x <= Math.abs(x1 - x2); x += 1) {
      for (let y = 0; y <= Math.abs(y1 - y2); y += 1) {
        if (x1 === x2 || y1 === y2 || x === y) {
          // calculate x position
          if (x1 < x2) {
            xPos = x1 + x;
          } else {
            xPos = x1 - x;
          }
          // calculate y postition
          if (y1 < y2) {
            yPos = y1 + y;
          } else {
            yPos = y1 - y;
          }
          diagram[xPos][yPos] += 1;
        }
      }
    }
  }

  // iterate over every item in the array to find overlaps
  let overlaps = 0;
  diagram.forEach((row) => {
    row.forEach((point) => {
      if (point >= 2) overlaps += 1;
    });
  });
  console.log(overlaps);
  // 22116
};
