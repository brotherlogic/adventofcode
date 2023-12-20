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

    return run_split(nnstr, nnums);
}

fn find_highest(map: Vec<Vec<i32>>) -> (usize,usize) {
    let mut highest:i32 = -1;
    let mut hv = (0,0);
    for (xpos, row) in map.iter().enumerate() {
        for(pos,value) in row.iter().enumerate() {
            if value > &highest {
                hv = (xpos, pos);
                highest = *value;
            }
        }
    }
    return hv;
}

fn fits(st: String, sp: usize, len: usize) -> bool {
   
    if any(st[sp..sp+len].to_string(), '.') {
        return false;
    }
    
    // Return true if we're at the start of the string and the right most char is eligibl
    if sp == 0 {
        return st[sp+len..sp+len+1] != *"#";
    }

    if sp+len < st.len() && st[sp+len..sp+len+1] == *"#" {
        return false;
    }

    // value to the left must be a '.' or ?
    let val = &st[sp-1..sp];
    return !(st[sp-1..sp] == *"#")
}

fn run_split(st: String, nums: Vec<usize>) -> i32 {
    let mut supermap: Vec<Vec<i32>> = Vec::new();
    for _ in &nums {
        supermap.push(vec![0; st.len()]);
    }
    supermap.push(vec![0; st.len()]);
    supermap[nums.len()][st.len()-1] = 1;
   
    // Work backwards through the nums
    let mut npointer = nums.len()-1;
    while npointer >= 0 {
        let mut spoint = 0;
        for (xpos, entry) in supermap[npointer+1].iter().enumerate() {
            if entry > &0 { 
                spoint = xpos;
            }
        }
        if npointer != nums.len()-1 {
            spoint -= (1+nums[npointer]);
        } else {
            spoint -= (nums[npointer]-1);
        }

    
        while spoint >= 0 {
            let nstr = &st[spoint..];
            if fits(st.clone(), spoint, nums[npointer]) {
                if npointer != 0 || !any(st[0..spoint].to_string(), '#') {
                    supermap[npointer][spoint] = smap_sum(st.clone(), supermap.clone(), npointer+1, spoint+nums[npointer] + 1, npointer == nums.len()-1);
                }
            }
            //if st[spoint+nums[npointer]-1..spoint+nums[npointer]] ==  *"#" {
            //    break;
            //}
        if spoint == 0 {
            break;
        }
             spoint -= 1;
        }
       if npointer == 0 {
        break;
       }
       npointer -= 1;
    }

  

    let mut stotal = 0;
   for num in &supermap[0] {
stotal += *num;
   }

 
   return stotal;
}

fn smap_sum(st: String, smap: Vec<Vec<i32>>, row: usize, spoint: usize, mrow: bool) -> i32 {
    if mrow && spoint >= smap[0].len() {
        return 1;
    }

    let mut sval = 0;
    let mut in_hash = false;
    for (xpos, val) in smap[row].iter().enumerate() {
        if xpos >= spoint {
            if st[xpos..xpos+1] == *"#" {
                in_hash = true;
            } else if in_hash {
                return sval;
            }
           
            sval += val;
        }
    }
    return sval;
}

fn suitable(str: String, end: bool) -> bool {
    for c in str[0..str.len()-1].chars() {
        if c != '#' && c != '?' {
            return false;
        }
    }

    if !end {
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
       let test_case = "?###???????? 3,2,1".to_string();
       let score = solve_day12_part2(test_case);
       assert_eq!(score, 506250)
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

