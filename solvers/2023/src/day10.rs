pub fn solve_day10_part1(data: String) -> i32 {
    let pipes = build_pipes(data);

    let (mut startx, mut starty) = (0,0);

    let mut best = Vec::new();

    for (posy, row) in  pipes.board.iter().enumerate() {
        let mut rbest = Vec::new();
        for (posx, value) in row.iter().enumerate() {
            rbest.push(i32::MAX);  
            if *value == 'S' {
                (startx,starty) = (posx as usize,posy as usize);
            }
        }
        best.push(rbest);
    }

    println!("FOUND START AT {} {}", startx,starty);
  
    best[starty][startx] = 0;
    let mut process = Vec::new();
    process.push((startx,starty,0));

    println!("HERE {:?}", best);


    while process.len() > 0 {
        let (currx,curry,val) = process.pop().unwrap();
        println!("TRAVERSE {},{}", currx,curry);
        match pipes.board[curry][currx] {
            'S' => {
                if pipes.board[curry-1][currx] == '|' || pipes.board[curry-1][currx] == 'F' || pipes.board[curry-1][currx] == '7' {
                    if best[curry-1][currx] > val+1 {
                    best[curry-1][currx] = val+1;
                    process.push((currx,curry-1, val+1));
                    }
                }
                if pipes.board[curry][currx+1] == '-' || pipes.board[curry][currx+1] == '7' || pipes.board[curry][currx] == 'J' {
                    if best[curry][currx+1] > val+1 {
                    best[curry][currx+1] = val+1;
                    process.push((currx+1,curry, val+1));
                    }
                }
                if pipes.board[curry+1][currx] == '|' || pipes.board[curry+1][currx] == 'L' || pipes.board[curry+1][currx] == 'J' {
                    if best[curry+1][currx] > val + 1 {
                    best[curry+1][currx] = val + 1;
                    process.push((currx, curry+1, val+1));
                    }
                }
                if pipes.board[curry][currx-1] == '-' || pipes.board[curry][currx-1] == 'L' || pipes.board[curry][currx-1] == 'F' {
                    if best[curry][currx-1] > val + 1 {
                    best[curry][currx-1] = val + 1;
                    process.push((currx, curry+1,val+1));
                    }
                }
            },
            '|' => {
                if best[curry-1][currx] > val +1  {
                    best[curry-1][currx] = val + 1;
                    process.push((currx, curry-1, val+1));
                }
                if best[curry+1][currx] > val + 1 {
                    best[curry+1][currx] = val + 1;
                    process.push((currx, curry-1, val+1));
                }
            },
            '-' => {
                if best[curry][currx+1] > val +1  {
                    best[curry][currx+1] = val + 1;
                    process.push((currx+1, curry, val+1));
                }
                if best[curry][currx-1] > val + 1 {
                    best[curry][currx-1] = val + 1;
                    process.push((currx-1, curry, val+1));
                }
            },
            'L' => {
                if best[curry-1][currx] > val +1  {
                    best[curry-1][currx] = val + 1;
                    process.push((currx, curry-1, val+1));
                }
                if best[curry][currx+1] > val + 1 {
                    best[curry][currx+1] = val + 1;
                    process.push((currx+1, curry, val+1));
                }
            },
            'J'=> {
                if best[curry-1][currx] > val +1  {
                    best[curry-1][currx] = val + 1;
                    process.push((currx, curry-1, val+1));
                }
                if best[curry][currx-1] > val + 1 {
                    best[curry][currx-1] = val + 1;
                    process.push((currx-1, curry, val+1));
                }
            }
            '7' => {
                if best[curry+1][currx] > val +1  {
                    best[curry+1][currx] = val + 1;
                    process.push((currx, curry+1, val+1));
                }
                if best[curry][currx-1] > val + 1 {
                    best[curry][currx-1] = val + 1;
                    process.push((currx-1, curry, val+1));
                }
            }
            'F' => {
                if best[curry+1][currx] > val +1  {
                    best[curry+1][currx] = val + 1;
                    process.push((currx, curry+1, val+1));
                }
                if best[curry][currx+1] > val + 1 {
                    best[curry][currx+1] = val + 1;
                    process.push((currx+1, curry, val+1));
                }
            },
            _ => println!("DONE"), //pass
        }
    }

    println!("RES {:?}", best);

    let mut bestv = 0;
    for row in best {
        for val in row {
            if val > bestv && val < i32::MAX{
                bestv = val;
            }
        }
    }

    return bestv;
}

fn build_pipes(data: String) -> Pipes {
    let mut pipes = Vec::new();
    for line in data.split("\n") {
        let mut nvec = Vec::new();
        for c in line.trim().chars() {
            nvec.push(c);
        }
        pipes.push(nvec);
    }

    return Pipes{board: pipes};
} 

struct Pipes {
    board: Vec<Vec<char>>,
}


#[cfg(test)]
mod testsca {
    use super::*;

#[test]
fn part1_test_first() {
   let test_case = ".....
   .S-7.
   .|.|.
   .L-J.
   .....".to_string();

   let score = solve_day10_part1(test_case);
   assert_eq!(score, 4)
}
}