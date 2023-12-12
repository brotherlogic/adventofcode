pub fn solve_day10_part1(data: String) -> i32 {
    let (best, _, _) = build_best(data);

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

fn has_seen(x: usize, y: usize, seen: &Vec<(usize,usize)>) -> bool {
    for (xe, ye) in seen {
        if *xe == x && *ye == y {
            return true;
        }
    }

    return false;
}


pub fn solve_day10_part2(data: String) -> i32 {
    let (best, pipes, s_rep) = build_best(data);

    let mut filled = Vec::new();
    for (posy, row) in best.iter().enumerate() {
        let mut temp_vec = Vec::new();
        for (posx, val) in row.iter().enumerate() {
            if *val < i32::MAX {
                temp_vec.push(pipes.board[posy][posx]);
            } else {
                temp_vec.push('.');
            }
        }
        filled.push(temp_vec);
    }


    //Find an entry point
    let mut investigate = Vec::new();
    for (posy, row) in filled.iter().enumerate() {
        for (posx, c) in row.iter().enumerate() {
            if posy == 0 || posy == filled.len()-1 || posx == 0 || posx == filled[0].len() - 1 {
                if *c == '.' {
                    investigate.push((posx,posy));
                }
            }
        }
    }

    // Outside fill
    let mut seen = Vec::new();
    while investigate.len() > 0 {
        let (posx, posy) = investigate.pop().unwrap();
        seen.push((posx, posy));
        filled[posy][posx] = 'O';
        if posx > 0 && filled[posy][posx-1] == '.' {
            if !has_seen(posx-1, posy, &seen) {
                investigate.push((posx-1,posy));
            }
        }
        if posy > 0 && filled[posy-1][posx] == '.' {
            if !has_seen(posx, posy-1, &seen) {
                investigate.push((posx,posy-1));
            }
        }
        if posx < filled[0].len()-1 && filled[posy][posx+1] == '.' {
            if !has_seen(posx+1, posy, &seen) {
             investigate.push((posx+1,posy));
            }
        }
        if posy < filled.len()-1 && filled[posy+1][posx] == '.' {
            if !has_seen(posx, posy+1, &seen) {
            investigate.push((posx,posy+1));
            }
        }
    }

    // Now run traces
    let mut start = (0,0);
    for (posy, row) in pipes.board.iter().enumerate() {
        for (posx, c) in row.iter().enumerate() {
            if *c == '|' {
                start = (posx, posy);
                break;
            }
            if filled[posy][posx] != 'O' {
                break;
            }
        }
    }

    if start == (0,0) {
        println!("NOT FOUND!");
        println!("{:?}", filled);
    } else {
        let istart = (start.0, start.1);
        let mut currdir = "north";
        start.1 -= 1;

        let mut steps = 0;
        while start != istart {
            println!("TRACE {:?} -> {:?} {}", start, pipes.board[start.1][start.0], currdir);
            steps+=1;
            if steps > 40 {
                break;
            }

            let mut currpos = pipes.board[start.1][start.0];
            if currpos == 'S' {
                currpos = s_rep;
                println!("S REP {}", s_rep);
            }

            if currdir == "north" {
                if start.0 > 0 && filled[start.1][start.0-1] == '.' {
                    filled[start.1][start.0-1] = 'O';
                }
                if currpos == '|' {
                    start.1 -=1;
                } else if currpos == 'F' {
                    start.0 += 1;
                    currdir = "east";
                } else if currpos == '7' {
                    start.0 -= 1;
                    currdir = "west";
                }
            } else if currdir == "south" {
                if start.0 < filled[0].len()-1 && filled[start.1][start.0+1] == '.' {
                    filled[start.1][start.0+1] = 'O';
                }
                if currpos == '|' {
                    start.1 +=1;
                } else if currpos == 'J' {
                    start.0 -= 1;
                    currdir = "west";
                } else if currpos == 'L' {
                    start.0 += 1;
                    currdir = "east";
                }
            } else if currdir == "east" {
                if start.1 >0 && filled[start.1-1][start.0] == '.' {
                    filled[start.1-1][start.0] = 'O';
                }
                if currpos == '-' {
                    start.0 +=1;
                } else if currpos == '7' {
                    start.1 += 1;
                    currdir = "south";
                } else if currpos == 'J' {
                    start.1 -= 1;
                    currdir = "north";
                }
            } else if currdir == "west" {
                if start.1 < filled.len()-1 && filled[start.1+1][start.0] == '.' {
                    filled[start.1+1][start.0] = 'O';
                }
                if currpos == '-' {
                    start.0 -=1;
                } else if currpos == 'F' {
                    start.1 += 1;
                    currdir = "south";
                } else if currpos == 'L' {
                    start.1 -= 1;
                    currdir = "north";
                }
            }
        }
    }

    // Final clean
    for (posy, row) in filled.iter().enumerate() {
        for (posx, c) in row.iter().enumerate() {
            if *c == 'O' {
                investigate.push((posx, posy));
            }
        }
    }

    while investigate.len() > 0 {
        let (posx, posy) = investigate.pop().unwrap();
        seen.push((posx, posy));
        filled[posy][posx] = 'O';
        if posx > 0 && filled[posy][posx-1] == '.' {
            if !has_seen(posx-1, posy, &seen) {
                investigate.push((posx-1,posy));
            }
        }
        if posy > 0 && filled[posy-1][posx] == '.' {
            if !has_seen(posx, posy-1, &seen) {
                investigate.push((posx,posy-1));
            }
        }
        if posx < filled[0].len()-1 && filled[posy][posx+1] == '.' {
            if !has_seen(posx+1, posy, &seen) {
             investigate.push((posx+1,posy));
            }
        }
        if posy < filled.len()-1 && filled[posy+1][posx] == '.' {
            if !has_seen(posx, posy+1, &seen) {
            investigate.push((posx,posy+1));
            }
        }
    }

    for row in &filled {
    println!("{:?}", row);
    }

    let mut found = 0;
    for row in filled {
        for c in row {
            if c == '.' {
                found+=1
            }
        }
    }

    return found;
}

fn build_best(data: String) -> (Vec<Vec<i32>>, Pipes, char) {
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
    let mut s_replace = ' ';

    while process.len() > 0 {
        let (currx,curry,val) = process.pop().unwrap();
        match pipes.board[curry][currx] {
            'S' => {
                let mut north = false;
                let mut east = false;
                let mut west = false;
                let mut south = false;
                if curry > 0 && (pipes.board[curry-1][currx] == '|' || pipes.board[curry-1][currx] == 'F' || pipes.board[curry-1][currx] == '7') {
                    if best[curry-1][currx] > val+1 {
                    best[curry-1][currx] = val+1;
                    process.push((currx,curry-1, val+1));
                    }
                    north = true;
                }
                if currx < pipes.board[curry].len() && (pipes.board[curry][currx+1] == '-' || pipes.board[curry][currx+1] == '7' || pipes.board[curry][currx+1] == 'J') {
                    if best[curry][currx+1] > val+1 {
                    best[curry][currx+1] = val+1;
                    process.push((currx+1,curry, val+1));
                    }
                    east = true;
                }
                if curry < pipes.board.len() && (pipes.board[curry+1][currx] == '|' || pipes.board[curry+1][currx] == 'L' || pipes.board[curry+1][currx] == 'J') {
                    if best[curry+1][currx] > val + 1 {
                    best[curry+1][currx] = val + 1;
                    process.push((currx, curry+1, val+1));
                    }
                    south = true;
                }
                if currx > 0 && (pipes.board[curry][currx-1] == '-' || pipes.board[curry][currx-1] == 'L' || pipes.board[curry][currx-1] == 'F') {
                    if best[curry][currx-1] > val + 1 {
                    best[curry][currx-1] = val + 1;
                    process.push((currx, curry+1,val+1));
                    }
                    west = true;
                }

                if north && south {
                    s_replace = '|';
                }
                if north && east {
                    s_replace = 'L';
                }
                if north && west {
                    s_replace = 'J';
                }
                if south && east {
                    s_replace = 'F';
                }
                if south && west {
                    s_replace = '7';
                }
                if east && west {
                    s_replace = '-';
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

  return (best, pipes, s_replace);
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

  #[test]
  fn part2_test_third() {
      let test_case = ".F----7F7F7F7F-7....
      .|F--7||||||||FJ....
      .||.FJ||||||||L7....
      FJL7L7LJLJ||LJ.L-7..
      L--J.L7...LJS7F-7L7.
      ....F-J..F7FJ|L7L7L7
      ....L7.F7||L7|.L7L7|
      .....|FJLJ|FJ|F7|.LJ
      ....FJL-7.||.||||...
      ....L---J.LJ.LJLJ...".to_string();
   
      let score = solve_day10_part2(test_case);
      assert_eq!(score, 8)
   }

   #[test]
   fn part2_test_fourth() {
       let test_case = "FF7FSF7F7F7F7F7F---7
       L|LJ||||||||||||F--J
       FL-7LJLJ||||||LJL-77
       F--JF--7||LJLJ7F7FJ-
       L---JF-JLJ.||-FJLJJ7
       |F|F-JF---7F7-L7L|7|
       |FFJF7L7F-JF7|JL---7
       7-L-JL7||F7|L7F-7F7|
       L.L7LFJ|||||FJL7||LJ
       L7JLJL-JLJLJL--JLJ.L".to_string();
    
       let score = solve_day10_part2(test_case);
       assert_eq!(score, 10)
    }
}



