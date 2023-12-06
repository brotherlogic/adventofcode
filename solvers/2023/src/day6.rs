struct Race {
    time: i64,
    distance: i64,
}

fn build_races(data: String) -> Vec<Race> {
    let mut elems = data.split("\n");
    let mut iter1 = elems.next().unwrap().trim().split_whitespace();
    let mut iter2 = elems.next().unwrap().trim().split_whitespace();
    
    // Skip the prefix
    iter1.next();
    iter2.next();

    let mut races = Vec::new();
    for t in iter1.next() {
        let d = iter2.next();
        races.push(Race{
            time: t.parse::<i64>().unwrap(), 
            distance: d.unwrap().parse::<i64>().unwrap(),
        })
    }

    return races;
}


pub fn solve_day6_part1(data: String) -> i32 {
    let mut solution = 1;

    let races = build_races(data);

    return solution;
}


#[cfg(test)]
mod testsca {
    use super::*;

#[test]
fn part1_tests() {
    let data = "Time:      7  15   30
    Distance:  9  40  200".to_string();
    let answer = 288;
    let solution = solve_day6_part1(data);
    assert_eq!(solution, answer);
}
}
