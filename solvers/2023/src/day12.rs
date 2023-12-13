pub fn solve_day12_part1(data: String) -> i32 {
    let mut st = 0;
    for line in data.split("\n") {
    st +=  run_calc(line.to_string());
    }
    return st;
}

fn run_calc(line: String) -> i32 {
    let mut elems = line.split_whitespace();
    let  base = elems.next().unwrap();
    let groups = elems.next().unwrap();

    let mut nums = Vec::new();
    for num in groups.split(",") {
        nums.push(num.parse::<i32>().unwrap());
    }

    return calculate_orgs(base.to_string(), &nums);
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
}