struct Race {
    time: i64,
    distance: i64,
}

fn build_races(data: String) -> Vec<Race> {
    elems = data.split("\n")
    iter1 = elems.next().unwrap().trim().split_whitespace();
    iter2 = elems.next().unwrap().trim().split_whitespace();;
    
    // Skip the prefix
    iter1.next();
    iter2.next();;

    let races = Vec::new();
    for t in iter1.next() {
        d = iter2.next()
        races.push(Race{time: t.parse::<i64>().unwrap(), distance: d.parse:<i64>().unwrap()})
    }

    return races;
}


pub fn solve_day6_part1(data: String) -> u64 {
    let mut solution = 1

    let races = build_races(data)
}


#[cfg(test)]
mod testsca {
    use super::*;

#[test]
fn part1_tests() {
    let data = "Time:      7  15   30
    Distance:  9  40  200".to_string()
    let answer = 288
    solution = solve_day6_part1(data)
    asser_eq!(solution, answer)
}
}
