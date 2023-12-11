pub fn solve_day10_part1(data: String) -> i32 {
    let (best, _) = build_best(data);

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

fn has_seen(x: usize, y: usize, pipe_dir: char, seen: &Vec<(usize,usize, char)>) -> bool {
    for (xe, ye, c) in seen {
        if *xe == x && *ye == y && pipe_dir == *c {
            return true;
        }
    }

    return false;
}

fn can_escape(x: usize, y: usize, best: Vec<Vec<i32>>, map: &Vec<Vec<char>>) -> bool {
    println!("STARTING {} {}", x, y);
    let mut seen = Vec::new();
    let mut togo = Vec::new();

    togo.push((x, y, '.'));
    seen.push((x, y, '.'));

    while togo.len() > 0 {
        let (currx,curry, pipe_dir) = togo.pop().unwrap();

        if currx == 0 || currx == best[0].len()-1 || curry == 0 || curry == best.len()-1 {
            println!("REACHED END {},{}", currx, curry);
            return true;
        }

        // Go west
        if pipe_dir == '.' && currx > 0 && best[curry][currx-1] == i32::MAX {
            if !has_seen(currx-1, curry, pipe_dir, &seen) {
                println!("PUSHW {} {}", currx-1, curry);
                togo.push((currx-1, curry, '.'));
                seen.push((currx-1, curry, '.'));
            }
        }

        // Go east
        if pipe_dir == '.' && currx < best[0].len()-1 && best[curry][currx+1] == i32::MAX {
            if !has_seen(currx+1, curry,  pipe_dir, &seen) {
                println!("PUSHE {} {}", currx+1, curry);
                togo.push((currx+1, curry, '.'));
                seen.push((currx+1, curry, '.'));
            }
        }

        // Go north
        if pipe_dir == '.' && curry > 0 && best[curry-1][currx] == i32::MAX {
            if !has_seen(currx, curry-1,  pipe_dir, &seen) {
                println!("PUSHN {} {}", currx, curry-1);
                togo.push((currx, curry-1, '.'));
                seen.push((currx, curry-1, '.'));
            }
        }

        // Go north in-between
        if curry > 0 && currx < best[0].len() {
            let left = map[curry-1][currx];
            let right = map[curry-1][currx+1];

             if (left == 'J' || left == '|' || left == '7' || left == 'L' || left == 'F') && (right == 'J' || right == '|' || right == '7' || right == 'L' || right == 'F') {
                println!("IN BETWEEN {},{} -> {} {} ({:?}) -> {},{}", currx, curry, left, right, map[curry-1],currx, curry-1);

                if !has_seen(currx, curry-1,  pipe_dir, &seen) {
                    togo.push((currx, curry-1, '|'));
                    seen.push((currx, curry-1, '|'));
                }
            }
        }

        // Go south
        if pipe_dir == '.' && curry < best.len()-1 && best[curry+1][currx] == i32::MAX {
            if !has_seen(currx, curry+1,  pipe_dir, &seen) {
                println!("PUSHS {} {}", currx, curry+1);
                togo.push((currx, curry+1, '.'));
                seen.push((currx, curry+1, '.'));
            }
        }

        // Go south in-between
        if curry < best.len()-1 && currx < best[0].len(){
            let left = map[curry+1][currx];
            let right = map[curry+1][currx+1];

          
            if (left == 'J' || left == '|' || left == '7' || left == 'L' || left == 'F') && (right == 'J' || right == '|' || right == '7' || right == 'L' || right == 'F') {
                if !has_seen(currx, curry+1, pipe_dir,  &seen) {
                    togo.push((currx, curry+1, '|'));
                    seen.push((currx, curry+1, '|'));
                }
            } else if left == '.'  {
                if !has_seen(currx, curry+1, '.', &seen) {
                    togo.push((currx, curry+1, '.'));
                    seen.push((currx, curry+1, '.'));
                }
            } else if right == '.' {
                if !has_seen(currx+1, curry+1, '.', &seen) {
                    togo.push((currx+1, curry+1, '.'));
                    seen.push((currx+1, curry+1, '.'));
                }
            }
        }
    }

    return false;
}

pub fn solve_day10_part2(data: String) -> i32 {
    let (best, pipes) = build_best(data);

    let mut escapes = 0;

    for (posy, row) in  best.iter().enumerate() {
        for (posx, val) in row.iter().enumerate() {
            if *val == i32::MAX  {
               if !can_escape(posx, posy, best.clone(), &pipes.board) {
                println!("ESCAPE {} {}", posx, posy);
                escapes+=1;
               } else {
                println!("NO_ESCAPE {} {}", posx, posy);
               }
            }
        }
    }

    return escapes;
}

fn build_best(data: String) -> (Vec<Vec<i32>>, Pipes) {
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

  
    best[starty][startx] = 0;
    let mut process = Vec::new();
    process.push((startx,starty,0));

  

    while process.len() > 0 {
        let (currx,curry,val) = process.pop().unwrap();
        match pipes.board[curry][currx] {
            'S' => {
                if curry > 0 && (pipes.board[curry-1][currx] == '|' || pipes.board[curry-1][currx] == 'F' || pipes.board[curry-1][currx] == '7') {
                    if best[curry-1][currx] > val+1 {
                    best[curry-1][currx] = val+1;
                    process.push((currx,curry-1, val+1));
                    }
                }
                if currx < pipes.board[curry].len() && (pipes.board[curry][currx+1] == '-' || pipes.board[curry][currx+1] == '7' || pipes.board[curry][currx+1] == 'J') {
                    if best[curry][currx+1] > val+1 {
                    best[curry][currx+1] = val+1;
                    process.push((currx+1,curry, val+1));
                    }
                }
                if curry < pipes.board.len() && (pipes.board[curry+1][currx] == '|' || pipes.board[curry+1][currx] == 'L' || pipes.board[curry+1][currx] == 'J') {
                    if best[curry+1][currx] > val + 1 {
                    best[curry+1][currx] = val + 1;
                    process.push((currx, curry+1, val+1));
                    }
                }
                if currx > 0 && (pipes.board[curry][currx-1] == '-' || pipes.board[curry][currx-1] == 'L' || pipes.board[curry][currx-1] == 'F') {
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
                    process.push((currx, curry+1, val+1));
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

  return (best, pipes);
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

#[test]
fn part1_test_second() {
   let test_case = "..F7.
   .FJ|.
   SJ.L7
   |F--J
   LJ...".to_string();

   let score = solve_day10_part1(test_case);
   assert_eq!(score, 8)
}

#[test]
fn part2_test_first() {
    let test_case = "...........
    .S-------7.
    .|F-----7|.
    .||.....||.
    .||.....||.
    .|L-7.F-J|.
    .|..|.|..|.
    .L--J.L--J.
    ...........".to_string();
 
    let score = solve_day10_part2(test_case);
    assert_eq!(score, 4)
 }

 #[test]
 fn part2_test_second() {
     let test_case = "..........
     .S------7.
     .|F----7|.
     .||....||.
     .||....||.
     .|L-7F-J|.
     .|..||..|.
     .L--JL--J.
     ..........".to_string();
  
     let score = solve_day10_part2(test_case);
     assert_eq!(score, 4)
  }
}