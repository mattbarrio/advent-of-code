const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  // .split(",");
  //.map((row) => row.replace(" -> ", ","));

  console.log("Starting problem_1");
  problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

const recalculate_list = (array) => {
  array.forEach((fish, index) => {
    if (fish === 0) {
      array[index] = 6;
      array.push(8);
    } else {
      array[index]--;
    }
  });
  return array;
};

const problem_1 = (array) => {
  // 0 is a valid day
  // after 0, create new lanternfish with a timer of 8 and reset the lanternfish timer to 6
  let lanternfish = array[0].split(",").map((i) => Number(i));
  for (let i = 1; i <= 80; i++) {
    lanternfish = recalculate_list(lanternfish);
  }

  // How many lanternfish would there be after 80 days?
  console.log(lanternfish.length);
  // 180028 (too low) fish < 0
  // 404194 (too high) fish === 0
  // 388739 let i = 1
};

const problem_2 = (array) => {
  // 0 is a valid day
  // after 0, create new lanternfish with a timer of 8 and reset the lanternfish timer to 6

  // The solution for part one wont work, even increasing the heap I get this error Fatal JavaScript invalid size error 144230630 on loop 144
  // New thought process, count all the fish in certain days, then decrement the entire pool of lanternfish in said day and add new fish, rinse and repeat
  let lanternfish = array[0].split(",").map((i) => Number(i));
  const days = new Array(9).fill(0);
  lanternfish.forEach((day) => {
    days[day] += 1;
  });

  for (let i = 1; i <= 256; i++) {
    console.log("starting loop %d", i);
    // calculate day 0 lanternfish before shifting the array
    day0 = days[0];
    console.log(days);
    // shift the lanternfish in the array
    for (let day = 0; day <= 8; day++) {
      days[day] = days[day + 1];
    }
    // add new lanternfish and update day 6 counter with day 0 lanternfish
    days[6] += day0;
    days[8] = day0;
    console.log(days);
  }

  // How many lanternfish would there be after 256 days?
  let total_lanternfish = 0;
  days.forEach((day) => {
    total_lanternfish += day;
  });
  console.log(total_lanternfish);
  // 46846243753462 (too high)
  // 1741362314973 precalculate day0 lanternfish
};
