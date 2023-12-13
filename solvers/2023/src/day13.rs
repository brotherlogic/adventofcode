pub fn solve_day13_part1(data: String) -> i32 {
    let all = get_boards(data);

    let mut total = 0;
    for board in all {
        let v = symmetry(board.clone());
        let rboard = rotate(board.clone());

        println!("BOARD");
        for row in &board {
            println!("{}", row);
        }
        println!("RBOARD");
        for row in &rboard {
            println!("{}", row);
        }

        let h = symmetry(rboard);
        println!("GOT {} and {}", h, v);
        total += v*100 + h;
    }

    return total;
}

fn rotate(board: Vec<String>) -> Vec<String> {
    let mut nboard = Vec::new();

    for i in 0..board[0].len() {
        let mut nstr = "".to_string();
        for row in &board {
            nstr += &row[i..i+1];
        }

        nboard.push(nstr);
    }
    return nboard;
}

fn symmetry(board: Vec<String>) -> i32 {
    let mut v = 0;

    //Vertical
    for (y, _row) in board.iter().enumerate() {
        let mut m = true;
        if y > 0 {
            for t in 0..y {
                if (y+t) < board.len() && y-t-1 >= 0 && board[y+t] != board[y-t-1] {
                    m = false;
                    break;
                }
            }
        } else {
            m = false;
        }
        if m {
            v = y as i32;
            break;
        }
    }


    return  v;
}

fn get_boards(data: String) -> Vec<Vec<String>> {
    let mut all_boards = Vec::new();

    let mut curr_board = Vec::new();
    for line in data.split("\n") {
        if line.trim().len() == 0 {
            all_boards.push(curr_board);
            curr_board = Vec::new();
        } else {
         curr_board.push(line.trim().to_string());
        }
    }

    all_boards.push(curr_board);
    return all_boards;
}

#[cfg(test)]
mod testsca {
    use super::*;

#[test]
fn part1_test_first() {
   let test_case = "#.##..##.
   ..#.##.#.
   ##......#
   ##......#
   ..#.##.#.
   ..##..##.
   #.#.##.#.
   
   #...##..#
   #....#..#
   ..##..###
   #####.##.
   #####.##.
   ..##..###
   #....#..#".to_string();

   let score = solve_day13_part1(test_case);
   assert_eq!(score, 405)
}
}