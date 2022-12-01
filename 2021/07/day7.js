const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  //.map((row) => row.replace(" -> ", ","));

  // console.log("Starting problem_1");
  // problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

const calculate_total = (crabs, current_position) => {
  let total = 0;
  crabs.forEach((position) => {
    let diff = Math.abs(position - current_position);
    total += ((1 + diff) * diff) / 2;
    // total += Math.abs(position - current_position) :::: part 1
  });
  return total;
};

const problem_1 = (array) => {
  let totals = [];
  let crabs = array[0].split(",").map((position) => Number(position));
  crabs.forEach((current_position, index) => {
    totals.push(calculate_total(crabs, current_position));
  });
  console.log(Math.min(...totals));
  // 348664
};

const problem_2 = (array) => {
  let totals = [];
  let crabs = array[0].split(",").map((position) => Number(position));
  crabs.forEach((current_position, index) => {
    totals.push(calculate_total(crabs, current_position));
  });
  console.log(Math.min(...totals));
  // 100220525
};
