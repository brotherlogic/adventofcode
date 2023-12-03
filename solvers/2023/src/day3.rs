
struct Board {
    nums: Vec<Number>,
    symbols: Vec<Symbol>,
}

struct Number {
    value: u32,
    x: u32,
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
                    b.nums.push(Number{value: gnum, x: x-xval, y: y});
                    curr_num = "".to_string();
                }
                // pass
            } else {
                b.symbols.push(Symbol{symbol: c.to_string(), x: x, y: y});
            }

            x+=1
        }
        y+=1
    }

    return b;
}

pub fn solve_day3_part1(board: String) -> u32 {
    let mut snum = 0;
    let mut board = build_board(board);
    for num in board.nums {
        snum += num.value;
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