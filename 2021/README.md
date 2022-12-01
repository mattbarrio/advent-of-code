Day template:

```javascript
const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");
  //.map((row) => row.replace(" -> ", ","));

  console.log("Starting problem_1");
  problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

const problem_1 = (array) => {
  console.log();
};

const problem_2 = (array) => {
  console.log();
};
```
