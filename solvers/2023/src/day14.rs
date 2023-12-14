use std::collections::HashMap;

pub fn solve_day14_part1(data: String) -> i32 {
    let mut sum = 0;

    let mut rboard = rotate(get_board(data));
    for line in &rboard {
        println!("LINE {}", line);
    }

    for line in tilt(rboard) {
        println!("TILT {}", line);
        let mut mult =  line.len();
        for ch in line.chars() {
            if ch == 'O' {
                sum += mult as i32;
            }
            mult -= 1;
        }
    }


    return sum;
}

fn weight(data: Vec<String>) -> i32 {
    let mut sum = 0;
    for line in data {
    let mut mult = line.len();
    for ch in line.chars() {
        if ch == 'O' {
            sum += mult as i32;
        }
        mult -= 1;
    }
}
    return sum;
}

fn flatten(data: Vec<String>) -> String {
    let mut flat = "".to_string();
    for line in data {
        flat += &line;
    }
    return flat;
}

pub fn solve_day14_part2(data: String) -> i32 {
    let mut sum = 0;
    let mut seen_map: HashMap<String, i32>  = HashMap::new();
    let mut weight_map: HashMap<String, i32> = HashMap::new();

    let mut iboard = get_board(data);
    for i in 0..10000 {
        iboard = tilt(rotate(iboard));
        iboard = tilt(rotate(iboard));
        iboard = flip(tilt(flip(rotate(iboard))));
        iboard = flip(tilt(flip(rotate(iboard))));

        let flat = flatten(iboard.clone());
        if seen_map.contains_key(&flat) {
            println!("SEEN {} {} -> {}", i, seen_map.get(&flat).unwrap(), weight_map.get(&flat).unwrap());
            let value = (1000000000 - (i-1)) % (i - seen_map.get(&flat).unwrap());
            println!("VALUE = {} FROM {} and {}", value, i-1, seen_map.len());
            for (key, val) in &seen_map {
                if *val == value as i32{
                    return *weight_map.get(key).unwrap();
                }
            }
        } else {
            seen_map.insert(flat.clone(), i as i32);
            weight_map.insert(flat, weight(rotate(iboard.clone())));
        }
    }

    for line in tilt(iboard) {
        let mut mult =  line.len();
        for ch in line.chars() {
            if ch == 'O' {
                sum += mult as i32;
            }
            mult -= 1;
        }
    }

    return sum;
}

fn tilt(board: Vec<String>) -> Vec<String> {
    let mut nboard = Vec::new();
    for line in board {
        let mut ocount = 0;
        let mut dotcount = 0;
        let mut nline = "".to_string();

        for ch in line.chars() {
            if ch == 'O' {
                ocount += 1;
            } else if ch == '.' {
                dotcount += 1;
            } else {
                for i in 0..ocount {
                    nline += "O";
                }
                for i in 0..dotcount {
                    nline += ".";
                }
                nline += "#";
                ocount = 0;
                dotcount = 0;
            }
        }
        for i in 0..ocount {
            nline += "O";
        }
        for i in 0..dotcount {
            nline += ".";
        }

        nboard.push(nline);
    }

    return nboard;

}

fn get_board(data: String) -> Vec<String> {
   
    let mut curr_board = Vec::new();
    for line in data.split("\n") {
         curr_board.push(line.trim().to_string());
  }
    return curr_board
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

fn flip(board: Vec<String>) -> Vec<String> {
    let mut nboard = Vec::new();

    for line in board {
        nboard.push(line.chars().rev().collect::<String>())
    }
    return nboard;
}

#[cfg(test)]
mod testsca {
    use super::*;

#[test]
fn part1_test_first() {
   let test_case = "O....#....
   O.OO#....#
   .....##...
   OO.#O....O
   .O.....O#.
   O.#..O.#.#
   ..O..#O..O
   .......O..
   #....###..
   #OO..#....".to_string();

   let score = solve_day14_part1(test_case);
   assert_eq!(score, 136)
}

#[test]
fn part2_test_first() {
    let test_case = "O....#....
    O.OO#....#
    .....##...
    OO.#O....O
    .O.....O#.
    O.#..O.#.#
    ..O..#O..O
    .......O..
    #....###..
    #OO..#....".to_string();
 
    let score = solve_day14_part2(test_case);
    assert_eq!(score, 64)
 }
}