pub fn solve_day12_part1(data: String) -> i32 {
    let mut st = 0;
    for line in data.split("\n") {
    st +=  run_calc(line.to_string(), 1);
    }
    return st;
}

pub fn solve_day12_part2(data: String) -> i32 {
    let mut st = 0;
    for line in data.split("\n") {
    st +=  run_calc(line.to_string(), 5);
    }
    return st;
}

fn minv(a: usize, b: usize) -> usize {
    if a < b {
        return a;
    }
    return b;
}

fn rep(st: String, num: usize) -> String {
    let mut fstr = "".to_string();
    for i in 0..num {
        fstr += &st;
    }
    return fstr.to_string();
}

fn run_calc(line: String, max: i32) -> i32 {
    println!("RUN {}", line);
    let mut elems = line.split_whitespace();
    let  base = elems.next().unwrap();
    let groups = elems.next().unwrap();

    let mut nums = Vec::new();
    for num in groups.split(",") {
        nums.push(num.parse::<usize>().unwrap());
    }

    let mut nnums = Vec::new();
    let mut nnstr = "".to_string();
    for i in 0..max {
        for num in &nums {
            nnums.push(*num);
        }
        if i < max-1 {
        nnstr += base;
        nnstr += "?";
        } else {
            nnstr += base;
        }
    }

    let mut success = 0;
    let mut process: Vec<(usize, String, String)> = Vec::new();
    process.push((0 as usize, nnstr.to_string(), "".to_string()));

    while process.len() > 0 {
        let (npointer, st, sofar) = process.pop().unwrap();
        //println!("STR {} ({}) {} ({:?}) {}", st, st.len(), npointer, nnums, sofar);

        // Winning case
        if st.len() == 0 && npointer == nnums.len() {
            //println!("WINNER: {}", sofar);
            success += 1;
        } else if npointer >= nnums.len() {
                //println!("WINNER F {} -> {}", st, sofar);
                success += 1;
        } else if st.len() > 0 {
            if st[0..1] == *"?" || st[0..1] == *"#" {
             //   println!("BOING {}", st);
                if st.len() >= nnums[npointer] && suitable(st[0..minv(nnums[npointer]+1, st.len())].to_string(), nnums[npointer] == st.len()) {
                    //println!("FOUND SUIT");
                    if nnums[npointer] == st.len() {
                        //println!("{} --> {} {} BLANK FROM {}",st,  "", npointer+1, sofar);
                        process.push((npointer+1, "".to_string(), sofar.clone() + &rep("#".to_string(), nnums[npointer])));
                    } else {
                        //println!("{} --> {} {} FROM {}", st, st[nnums[npointer]+1..].to_string(), npointer+1, sofar);
                        process.push((npointer+1, st[nnums[npointer]+1..].to_string(),sofar.clone() + &rep("#".to_string(), nnums[npointer]) + "."));
                    }
                }
            }
            //println!("{} --> {} {}", st, st[1..].to_string(), npointer);
            if st[0..1] == *"?" || st[0..1] == *"." {
                process.push((npointer, st[1..].to_string(), sofar.clone() + "."));
            }
        }
    }

    return success;
}

fn suitable(str: String, end: bool) -> bool {
    //println!("SUIT: {} {}", str, end);
    for c in str[0..str.len()-1].chars() {
        if c != '#' && c != '?' {
            //println!("PUSHING END {}", c);
            return false;
        }
    }

    if !end {
        //println!("CHECKING {}", str.chars().last().unwrap());
        return str.chars().last().unwrap() == '?' || str.chars().last().unwrap()  == '.';
    } else {
        return str.chars().last().unwrap() == '#' || str.chars().last().unwrap() == '?';
    }
}

fn all(s: String, c: char) -> bool {
    for ch in s.chars() {
        if ch != c {
            return false;
        }
    }
    return true;
}

fn any(s: String, c: char) -> bool {
    for ch in s.chars() {
        if ch == c {
            return true;
        }
    }
    return false;
}

fn build(c: &str, num: i32) -> String{
    let mut str = "".to_string();
    for i in [0..num] {
        str += c;
    }

    return str;
}


fn run_count(line: String) -> Vec<i32> {
    let mut r = Vec::new();
    let mut count = 0;
    for ch in line.chars() {
        if ch == '.' {
            if count > 0 {
                r.push(count);
                count = 0;
            }
        } else {
            count+=1;
        }
    }

    if count > 0 {
        r.push(count);
    }

    return r;
}

fn calculate_orgs(line: String, goals: &Vec<i32>) -> i32 {
    let mut total = 0;
    for (posx, ch) in line.chars().enumerate() {
        if ch == '?' {
            total += calculate_orgs(line[0..posx].to_string() + "." + &line[posx+1..], goals);
            total += calculate_orgs(line[0..posx].to_string() + "#" + &line[posx+1..], goals);
            return total;
         }
    }

    // If we've reached here then there's nothing left to fill
    let counts = run_count(line);
    if counts.len() != goals.len() {
        return 0;
    }

    for (i, val) in counts.iter().enumerate() {
        if *val != goals[i] {
            return 0;
        }
    }

    return 1;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_reduced() {
       let test_case = "???.### 1,1,3".to_string();
       let score = solve_day12_part2(test_case);
       assert_eq!(score, 1)
    }

#[test]
fn part1_test_first() {
   let test_case = "???.### 1,1,3
   .??..??...?##. 1,1,3
   ?#?#?#?#?#?#?#? 1,3,1,6
   ????.#...#... 4,1,1
   ????.######..#####. 1,6,5
   ?###???????? 3,2,1".to_string();

   let score = solve_day12_part1(test_case);
   assert_eq!(score, 21)
}

#[test]
fn part2_test_first() {
   let test_case = "???.### 1,1,3
   .??..??...?##. 1,1,3
   ?#?#?#?#?#?#?#? 1,3,1,6
   ????.#...#... 4,1,1
   ????.######..#####. 1,6,5
   ?###???????? 3,2,1".to_string();

   let score = solve_day12_part2(test_case);
   assert_eq!(score, 525152)
}
}