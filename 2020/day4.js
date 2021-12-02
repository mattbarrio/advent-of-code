const fs = require("fs");

const fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];
const byr = [1920, 2002];
const iyr = [2010, 2020];
const eyr = [2020, 2030];
const hgtCm = [150, 193];
const hgtIn = [59, 76];
const ecl = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];

fs.readFile("day4.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  var compiledArr = [];
  var matched = 0;

  for (i in array) {
    // console.log("row " + i + " " + array[i]);
    // console.log(compiledArr);
    if (array[i] == "" || i == array.length - 1) {
      var passportItems = [];
      var passportFields = [];
      var passportValues = [];
      var passportValid = false;
      passportItems = compiledArr.join(" ").split(" ");
      if (passportItems.length < 7) {
        //console.log("row " + i, "skipping ", passportItems);
        compiledArr = [];
        continue;
      }
      //console.log(passportItems);
      for (passportItem in passportItems) {
        //console.log("item: " + passportItems[passportItem]);
        var passportValid = false;
        var item = passportItems[passportItem].split(":")[0];
        var val = passportItems[passportItem].split(":")[1];
        switch (item) {
          case "byr":
            if (val.length == 4 && val >= byr[0] && val <= byr[1]) {
              passportValid = true;
            }
            // console.log("byr: " + passportValid);
            break;
          case "iyr":
            if (val.length == 4 && val >= iyr[0] && val <= iyr[1]) {
              passportValid = true;
            }
            // console.log("iyr: " + passportValid);
            break;
          case "eyr":
            if (val.length == 4 && val >= eyr[0] && val <= eyr[1]) {
              passportValid = true;
            }
            // console.log("eyr: " + passportValid);
            break;
          case "hgt":
            if (val.includes("cm")) {
              var cm = val.split("cm")[0];
              if (cm >= hgtCm[0] && cm <= hgtCm[1]) {
                passportValid = true;
              }
            } else if (val.includes("in")) {
              var inches = val.split("in")[0];
              if (inches >= hgtIn[0] && inches <= hgtIn[1]) {
                passportValid = true;
              }
            }
            // console.log("hgt: " + passportValid);
            break;
          case "hcl":
            if (val.includes("#")) {
              var regex = /[0-9a-f]{6}/g;
              if (val.split("#")[1].match(regex)) {
                passportValid = true;
              }
            }
            // console.log("hcl: " + passportValid);
            break;
          case "ecl":
            if (val.length == 3 && ecl.indexOf(val) !== -1) {
              passportValid = true;
            }
            //console.log("ecl: " + val + " " + passportValid);
            break;
          case "pid":
            if (val.length == 9) {
              var regex = /[0-9]{9}/g;
              if (val.match(regex)) {
                passportValid = true;
              }
            }
            // console.log("pid: " + passportValid);
            break;
          case "cid":
            passportValid = true;
            // console.log("cid: " + passportValid);
            break;
          default:
            console.log("ERR: " + item);
        }
        if (passportValid !== true) {
          break;
        }
        passportValues = passportValues.concat(
          passportItems[passportItem].split(":")[1]
        );
        passportFields = passportFields.concat(item);
      }
      //console.log("row " + i, "passportValues: " + passportValues, "passportFields: " + passportFields);
      if (fields.every((item) => passportFields.includes(item))) {
        matched += 1;
        //console.log(passportItems);
        console.log(
          "row " + i,
          "passportValues: " + passportValues,
          "passportFields: " + passportFields
        );
      }
      // else {
      //    console.log("row " + i, "passport invalid " + passportValues, "passportFields: " + passportFields);
      // }
      compiledArr = [];
      continue;
    }
    compiledArr = compiledArr.concat(array[i]);
    //console.log(compiledArr);
  }
  console.log("matched: " + matched);
});
