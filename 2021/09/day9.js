const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  array.forEach((row, index) => {
    array[index] = row.split("").map((i) => Number(i));
  });
  //console.log(array);
  //.map((row) => row.replace(" -> ", ","));

  console.log("Starting problem_1");
  problem_1(array);

  console.log("Starting problem_2");
  //problem_2(array);
});

const is_low_point = (location, location_index, row_index, array) => {
  let short_circuit = false;
  let top_number = array?.[row_index - 1]?.[location_index];
  let left_number = array?.[row_index]?.[location_index - 1];
  let right_number = array?.[row_index]?.[location_index + 1];
  let bottom_number = array?.[row_index + 1]?.[location_index];
  console.log(row_index, location_index);
  if (location < top_number || top_number === undefined) {
    console.log("top ", location, top_number);
  } else {
    return;
  }
  if (location < left_number || left_number === undefined) {
    console.log("left ", location, left_number);
  } else {
    return;
  }
  if (location < right_number || right_number === undefined) {
    console.log("right ", location, right_number);
  } else {
    return;
  }
  if (location < bottom_number || bottom_number === undefined) {
    console.log("bottom ", location, bottom_number);
  } else {
    return;
  }
  return location;
};

const problem_1 = (array) => {
  let low_points = [];
  array.forEach((row, row_index) => {
    row.forEach((location, location_index) => {
      let low_point = is_low_point(location, location_index, row_index, array);
      if (typeof low_point === "number") {
        low_points.push(low_point);
      }
    });
  });
  console.log(low_points);
  let total = 0;
  low_points.forEach((num) => {
    total += num;
  });
  total += low_points.length;
  console.log(total);
};

const problem_2 = (array) => {
  if (row_index === 0) {
    console.log("first row");
  } else if (row_index === 99) {
    console.log("last row");
  } else if (location_index < row.length) {
    //console.log("row %d", row_index);
  } else {
    console.log(location);
  }
  console.log();
};
