const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  //.map((i) => Number(i));

  console.log("Starting problem_1");
  problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

const problem_1 = (array) => {
  let main_obj = {};
  let sub_obj = {};
  let g_rate = "";
  let e_rate = "";

  for (i in array) {
    //011010111110
    main_obj = array[i].toString().split("");
    //[ '0', '1', '1', '0', '1', '0', '1', '1', '1', '1', '1', '0' ]
    for (n in main_obj) {
      if (!isNaN(Number(sub_obj[n]))) {
        sub_obj[n] += Number(main_obj[n]);
      } else {
        sub_obj[n] = Number(main_obj[n]);
      }
    }
  }

  for (i in sub_obj) {
    if (sub_obj[i] <= 500) {
      g_rate += "0";
      e_rate += "1";
    } else {
      g_rate += "1";
      e_rate += "0";
    }
  }

  console.log(parseInt(g_rate, 2) * parseInt(e_rate, 2));
};

const scrub_array = (array, position, search_for) => {
  // console.log("scrub_array... pos %d with %d", position, search_for);
  let scrubed_array = [];
  for (i in array) {
    split_line = array[i].toString().split("");
    if (split_line[position] == search_for) {
      scrubed_array.push(array[i]);
    }
  }
  // console.log(scrubed_array);
  return scrubed_array;
};

const recalculate_array = (array, position, total, type) => {
  // console.log("recalculate_array...");
  // console.log(array.length, position, total, type);
  switch (type) {
    case "o2":
      if (total < array.length / 2) {
        array = scrub_array(array, position, 0);
        return array;
      } else {
        array = scrub_array(array, position, 1);
        return array;
      }
    case "co2":
      if (total < array.length / 2) {
        array = scrub_array(array, position, 1);
        return array;
      } else {
        array = scrub_array(array, position, 0);
        return array;
      }
  }
};

const recalculate_totals = (array) => {
  // console.log("recalculate_totals...");
  let main_obj = {};
  let sub_obj = {};
  for (i in array) {
    // console.log("recalculate_totals iteration %d", i);
    // 011010111110
    main_obj = array[i].toString().split("");
    // [ '0', '1', '1', '0', '1', '0', '1', '1', '1', '1', '1', '0' ]
    for (n in main_obj) {
      if (!isNaN(Number(sub_obj[n]))) {
        sub_obj[n] += Number(main_obj[n]);
      } else {
        sub_obj[n] = Number(main_obj[n]);
      }
    }
  }
  //console.log(sub_obj);
  return sub_obj;
};

const problem_2 = (array) => {
  let o2_sub_obj = {};
  let o2_rating = 0;
  let o2_scrubed_array = array;
  let co2_sub_obj = {};
  let co2_rating = 0;
  let co2_scrubed_array = array;
  const iteration = 12;
  for (let i = 0; i <= iteration; i++) {
    // console.log("loop iteration %d", i);
    if (co2_scrubed_array.length != 1) {
      co2_sub_obj = recalculate_totals(co2_scrubed_array);
      co2_scrubed_array = recalculate_array(
        co2_scrubed_array,
        i,
        co2_sub_obj[i],
        "co2"
      );
    } else {
      co2_rating = parseInt(co2_scrubed_array[0], 2);
    }
    // console.log("----------------------");
    if (o2_scrubed_array.length != 1) {
      o2_sub_obj = recalculate_totals(o2_scrubed_array);
      o2_scrubed_array = recalculate_array(
        o2_scrubed_array,
        i,
        o2_sub_obj[i],
        "o2"
      );
    } else {
      o2_rating = parseInt(o2_scrubed_array[0], 2);
    }
    // console.log("======================");
  }
  console.log(o2_rating * co2_rating);
};

// wrong 4786320
// wrong 4797452
// wrong 1384152
// 4790390
