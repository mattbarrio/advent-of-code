const fs = require("fs");

fs.readFile("input.txt", function (err, data) {
  if (err) throw err;
  const array = data.toString().split("\n");

  //console.log("Starting problem_1");
  //problem_1(array);

  console.log("Starting problem_2");
  problem_2(array);
});

// board is 5x5

const generate_bingo_cards = (array) => {
  let bingo_cards = [];
  let row_number = 0;
  let column_number = 0;
  let next_card = 0;
  for (i in array) {
    // remove extra whitespaces and split the array
    row = array[i].trim().split(/\s+/);
    // hack to check for empty line
    if (row.length == 1) {
      // increment the next card var
      next_card += 1;
      // reset row number when new card
      row_number = 0;
    } else {
      // build bingo card object
      for (n in row) {
        bingo_cards.push({
          card: next_card,
          row: row_number,
          column: column_number,
          number: row[n],
          drawn: false,
        });
        // increment column number
        column_number += 1;
      }
      // reset column number
      column_number = 0;
      // increment row number
      row_number += 1;
    }
  }
  return bingo_cards;
};

const mark_cards = (bingo_cards, number) => {
  console.log("marking cards with %d", number);
  const updated_cards = bingo_cards.map((card) =>
    card.number === number ? { ...card, drawn: true } : card
  );
  return updated_cards;
};

const check_cards = (bingo_cards) => {
  let winners = 0;
  // check all cards
  for (let index = 0; index < 101; index++) {
    console.log("checking card %d", index);
    // check all 5 rows and columns
    for (let i = 0; i < 6; i++) {
      let row = bingo_cards.filter(
        (card) => card.card === index && card.drawn === true && card.row === i
      );
      if (row.length === 5) {
        const undrawn_numbers = bingo_cards.filter(
          (card) => card.card === index && card.drawn === false
        );
        //1 console.log(row, undrawn_numbers);
        let score = 0;
        for (r in undrawn_numbers) {
          score += Number(undrawn_numbers[r].number);
        }
        //1 return score;
        winners += 1;
        console.log(undrawn_numbers, score, winners);
        if (winners === 100) {
          return bingo_cards.filter((card) => card.card === index);
        }
        break;
      }
      let column = bingo_cards.filter(
        (card) =>
          card.card === index && card.drawn === true && card.column === i
      );
      if (column.length === 5) {
        const undrawn_numbers = bingo_cards.filter(
          (card) => card.card === index && card.drawn === false
        );
        //1 console.log(column, undrawn_numbers);
        let score = 0;
        for (r in undrawn_numbers) {
          score += Number(undrawn_numbers[r].number);
        }
        //1 return score;
        winners += 1;
        console.log(undrawn_numbers, score, winners);
        if (winners === 100) {
          return bingo_cards.filter((card) => card.card === index);
        }
        break;
      }
    }
  }
};

const problem_1 = (array) => {
  let bingo_numbers = array[0].split(",");
  array.splice(0, 1);
  let bingo_cards = generate_bingo_cards(array);
  let winning_number;
  let current_number;
  for (i in bingo_numbers) {
    current_number = bingo_numbers[i];
    bingo_cards = mark_cards(bingo_cards, current_number);
    winning_number = check_cards(bingo_cards);
    if (typeof winning_number === "number") {
      break;
    }
  }

  console.log(winning_number * current_number); //.filter((cards) => cards.drawn === false));
  // 7658
  // 38594
};

const problem_2 = (array) => {
  let bingo_numbers = array[0].split(",");
  array.splice(0, 1);
  let bingo_cards = generate_bingo_cards(array);
  let winning_number;
  let current_number;
  for (i in bingo_numbers) {
    current_number = bingo_numbers[i];
    bingo_cards = mark_cards(bingo_cards, current_number);
    winning_number = check_cards(bingo_cards);
    if (typeof winning_number === "object") {
      console.log(winning_number, i);
      break;
    }
    // console.log(winning_number * current_number);
  }
};
