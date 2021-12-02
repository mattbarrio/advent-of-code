const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  //.map((i) => Number(i));

  //console.log("Starting problem_1");
  //problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

const problem_1 = (array) => {
  //
  // let previous_measurement = array[0];
  let depth = 0;
  let hortizontal_pos = 0;
  //console.log(array);
  for (i in array) {
    const dir_type = String(array[i].split(" ")[0]);
    const direction = Number(array[i].split(" ")[1]);
    switch (dir_type) {
      case 'forward':
        hortizontal_pos += direction;
        break
      case 'up':
        depth -= direction;
        break;
      case 'down':
        depth += direction;
        break;
      default:
        console.log("unknown " + dir_type +  " " + direction);
        break;
    }
  }

  console.log(
    "hor: %d, dep: %d = %d",
    hortizontal_pos,
    depth,
    hortizontal_pos * depth
  );
  // hor: 1971, dep: 830 = 1635930
};

const problem_2 = (array) => {
  // let previous_measurement = array[0];
  let depth = 0;
  let hortizontal_pos = 0;
  let aim = 0;
  // console.log(array);
  for (i in array) {
    const dir_type = String(array[i].split(" ")[0]);
    const direction = Number(array[i].split(" ")[1]);
    switch (dir_type) {
      case 'forward':
        hortizontal_pos += direction;
        depth += aim * direction;
        break
      case 'up':
        aim -= direction;
        break;
      case 'down':
        aim += direction;
        break;
      default:
        console.log("unknown " + dir_type +  " " + direction);
        break;
    }
  }

  console.log(
    "aim: %d, hor: %d, dep: %d = %d",
    aim,
    hortizontal_pos,
    depth,
    hortizontal_pos * depth
  );
  // aim: 830, hor: 1971, dep: 904018 = 1781819478
};
