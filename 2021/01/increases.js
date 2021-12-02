const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data
    .toString()
    .split("\n")
    .map((i) => Number(i));

  console.log("Starting problem_1");
  problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

const problem_1 = (array) => {
  // get the count of any time the previous number is less than the current number
  let previous_measurement = array[0];
  let counter = 0;

  for (i in array) {
    if (previous_measurement < array[i]) {
      counter += 1;
      // console.log(array[i]);
    }
  }

  console.log({ counter });
};

const problem_2 = (array) => {
  // get the total number of times the sum of the previous three numbers is less than the sum of the current three numbers
  let counter = 0;

  for (i in array) {
    let previous_measurement = array[i] + array[+i + 1] + array[+i + 2];
    let current_measurement = array[+i + 1] + array[+i + 2] + array[+i + 3];
    // console.log(previous_measurement, current_measurement);

    if (previous_measurement < current_measurement) {
      counter += 1;
    }
  }

  console.log({ counter });
};
