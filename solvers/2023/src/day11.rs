pub fn solve_day11_part2(data: String, adder: i64) -> i64 {
    let mut sgalaxies = Vec::new();

    let mut max_row: i64 = 0;
    let mut max_col: i64 = 0;
    for (posy, line) in data.split("\n").enumerate() {
        for (posx, c) in line.trim().chars().enumerate() {
            if c == '#' {
                sgalaxies.push((posx as i64, posy as i64));
                if posy as i64 > max_row {
                    max_row = posy as i64;
                }
                if posx as i64 > max_col {
                    max_col = posx as i64;
                }
            }
        }
    }

    println!("Found {} galaxies", sgalaxies.len());

    let mut galaxies: Vec<(i64,i64)> = Vec::new();
    let mut rows = Vec::new();
    let mut cols = Vec::new();
    for row in 0..max_row {
        let mut count = 0;
        for (_gx,gy) in &sgalaxies {
            if *gy == row {
                count+=1;
            }
        }
        if count == 0 {
            rows.push(row);
        }
    }

    for col in 0..max_col{
        let mut count = 0;
        for (gx,_gy) in &sgalaxies {
            if *gx == col {
                count+=1;
            }
        }
        if count == 0 {
            cols.push(col);
        }
    }

    for (gx,gy) in &sgalaxies {
        let mut ngx = *gx;
        let mut ngy = *gy;

        for col in &cols {
            if col < gx {
                ngx+=adder-1;
            }
        }

        for row in &rows {
            if row < gy {
                ngy+=adder-1;
            }
        }

        galaxies.push((ngx,ngy));
    }

    let mut tdist = 0;
    for (i, (xv,yv)) in galaxies.iter().enumerate() {
        for (xo,yo) in &galaxies[i..galaxies.len()] {
            tdist += abs(xv-xo) + abs(yv-yo);
        }
    }

    return tdist;
}

pub fn solve_day11_part1(data: String) -> i64 {
    let mut sgalaxies = Vec::new();

    let mut max_row: i64 = 0;
    let mut max_col: i64 = 0;
    for (posy, line) in data.split("\n").enumerate() {
        for (posx, c) in line.trim().chars().enumerate() {
            if c == '#' {
                sgalaxies.push((posx as i64, posy as i64));
                if posy as i64 > max_row {
                    max_row = posy as i64;
                }
                if posx as i64 > max_col {
                    max_col = posx as i64;
                }
            }
        }
    }

    println!("Found {} galaxies", sgalaxies.len());

    let mut galaxies: Vec<(i64,i64)> = Vec::new();
    let mut rows = Vec::new();
    let mut cols = Vec::new();
    for row in 0..max_row {
        let mut count = 0;
        for (_gx,gy) in &sgalaxies {
            if *gy == row {
                count+=1;
            }
        }
        if count == 0 {
            rows.push(row);
        }
    }

    for col in 0..max_col{
        let mut count = 0;
        for (gx,_gy) in &sgalaxies {
            if *gx == col {
                count+=1;
            }
        }
        if count == 0 {
            cols.push(col);
        }
    }

    for (gx,gy) in &sgalaxies {
        let mut ngx = *gx;
        let mut ngy = *gy;

        for col in &cols {
            if col < gx {
                ngx+=1;
            }
        }

        for row in &rows {
            if row < gy {
                ngy+=1;
            }
        }

        galaxies.push((ngx,ngy));
    }

    let mut tdist = 0;
    for (i, (xv,yv)) in galaxies.iter().enumerate() {
        for (xo,yo) in &galaxies[i..galaxies.len()] {
            tdist += abs(xv-xo) + abs(yv-yo);
        }
    }

    return tdist;
}

fn abs(val: i64) -> i64 {
    if val > 0 {
        return val;
    }
    return 0-val;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
    let test_case = "...#......
    .......#..
    #.........
    ..........
    ......#...
    .#........
    .........#
    ..........
    .......#..
    #...#.....".to_string();

    let score = solve_day11_part1(test_case);
    assert_eq!(score, 374)
    }

    #[test]
    fn part2_test_first() {
    let test_case = "...#......
    .......#..
    #.........
    ..........
    ......#...
    .#........
    .........#
    ..........
    .......#..
    #...#.....".to_string();

    let score = solve_day11_part2(test_case, 100);
    assert_eq!(score, 8410)
    }

    #[test]
    fn part2_test_second() {
    let test_case = "...#......
    .......#..
    #.........
    ..........
    ......#...
    .#........
    .........#
    ..........
    .......#..
    #...#.....".to_string();

    let score = solve_day11_part2(test_case, 10);
    assert_eq!(score, 1030)
    }
}