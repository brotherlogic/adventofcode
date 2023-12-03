
struct Board {
    nums: Vec<Number>,
    symbols: Vec<Symbol>,
}

#[derive(Debug)]
struct Number {
    value: u32,
    x: u32,
    xe: u32,
    y: u32,
}

struct Symbol {
    symbol: String,
    x: u32,
    y: u32,
}

fn build_board(board: String) -> Board {
    let mut b = Board{
        nums:  Vec::new(),
        symbols:  Vec::new(),
    };

    let mut elems = board.split("\n");
    let mut y: u32 = 0;
    for elem in elems {
        let mut x: u32 = 0;
        let mut curr_num :String = "".to_string();
        for c in elem.chars() {
            if c.is_digit(10) {
                curr_num += &c.to_string()
            } else if c == '.' {
                if curr_num.len() > 0 {
                    let xval = u32::try_from(curr_num.len()).ok().unwrap();
                    let gnum = curr_num.parse::<u32>().unwrap();
                    b.nums.push(Number{value: gnum, x: x-xval, xe: x-1,y: y});
                    curr_num = "".to_string();
                }
                // pass
            } else {
                // we might have a symbol after a number
                if curr_num.len() > 0 {
                    let xval = u32::try_from(curr_num.len()).ok().unwrap();
                    let gnum = curr_num.parse::<u32>().unwrap();
                    b.nums.push(Number{value: gnum, x: x-xval, xe: x-1,y: y});
                    curr_num = "".to_string();
                }
                b.symbols.push(Symbol{symbol: c.to_string(), x: x, y: y});
            }

            x+=1
        }
        y+=1
    }

    return b;
}

fn safe_sub(val: u32) -> u32 {
    if val == 0 {
        return 0
    }
    return val-1
}

fn fits(num: &Number, board: &Board) -> bool {
    for symbol in &board.symbols {
        //is it to the left?
        if symbol.y == num.y && num.x != 0 && symbol.x == num.x-1 {
            return true
        }

        //is it to the right?
        if symbol.y == num.y && symbol.x == num.x+num.xe+1 {
            return true
        }

        // is it on the top part
        if num.y != 0 && symbol.y == num.y-1 && symbol.x >= safe_sub(num.x)  && symbol.x <= num.xe+1 {
            return true
        }

             // is it on the bottom part
             if symbol.y == num.y+1 && symbol.x >= safe_sub(num.x)  && symbol.x <= num.xe+1 {
                return true
            }
    }

    return false
}

pub fn solve_day3_part1(board: String) -> u32 {
    let mut snum = 0;
    let  board = build_board(board);
    for num in &board.nums {
        if fits(&num, &board) {
        snum += num.value;
        }
    }
    return snum;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
       let board = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..".to_string();
       let answer = solve_day3_part1(board);
       assert_eq!(answer, 4361);
    }



}