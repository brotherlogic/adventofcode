#[derive(Debug)]
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
    for t in iter1 {
        println!("HERE");
        let d = iter2.next();
        races.push(Race{
            time: t.parse::<i64>().unwrap(), 
            distance: d.unwrap().parse::<i64>().unwrap(),
        })
    }

    return races;
}

fn solve_race(race: Race) -> i32 {
    let a: f64 = -1_f64;
    let b: f64 = race.time as f64;
    let c: f64 = -1_f64 * race.distance as f64;
    let lower_bound = (-b + f64::sqrt(b*b - 4_f64*a*c)) / (2_f64*a);
    let upper_bound = (-b - f64::sqrt(b*b - 4_f64*a*c)) / (2_f64*a);

    println!("HERE {} {}", upper_bound.ceil(), lower_bound.ceil());
    let val =  (upper_bound.floor() - lower_bound.ceil()) as i32;
    if upper_bound.fract() == 0.0 {
        return val-1;
    }
    return 1+val;
}

pub fn solve_day6_part1(data: String) -> i32 {
    let mut solution = 1;

    let races = build_races(data);
    println!("Running {:?}", races);
    for race in races {

        solution *= solve_race(race);
    }

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
