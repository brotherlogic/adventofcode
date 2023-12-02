use std::collections::HashMap;

struct Game {
    number: u32,
    rounds: Vec<Round>,
}

struct Round {
    blue: u32,
    red: u32,
    green: u32,
}

fn build_game(data: &str) -> Game {
    let mut elems = data.split(":");

    let p1 = elems.next().unwrap();
    let mut fparts = p1.split(" ");
    //println!("{}", p1);
    fparts.next().unwrap();
    let p2 = fparts.next().unwrap();
    let gnumber = p2.parse::<u32>().unwrap();

    let mut arounds = Vec::new();
    let nelems = elems.next().unwrap().split(";");
    for nelem in nelems {
        let parts = nelem.split(",");
        let mut r = Round{green: 0, red: 0, blue: 0};
        for part in parts {
            let mut bits = part.trim().split(" ");
            let num = bits.next().unwrap();
            let gnum = num.parse::<u32>().unwrap();
            let colour = bits.next().unwrap();
            if colour == "red" {
                r.red = gnum
            } else if colour == "green" {
                r.green = gnum
            } else if colour == "blue" {
                r.blue = gnum
            }
        }

        arounds.push(r)
    }

    return Game{number: gnumber, rounds: arounds}
}

fn legal(data: &str, red: u32, green: u32, blue: u32) -> u32{
    let agame = build_game(data);
    for round in agame.rounds {
        if round.blue > blue || round.red > red || round.green > green {
            return 0
        }
    }
    return agame.number;
}

fn minimum_game(data: &str) -> u32{
    let agame = build_game(data);
    let mut red:u32 = 1;
    let mut green:u32 = 1;
    let mut blue:u32 = 1;
    for round in agame.rounds {
       if round.blue > blue {
        blue = round.blue
       }
       if round.green > green {
        green = round.green
       }
       if round.red > red {
        red = round.red
       }
    }
    return red*blue*green;
}

pub fn solve_day2_part1(data: String) -> u32 {
    let parts = data.split("\n");
    let mut total = 0;
    for part in parts {
        if part.len() > 0 {
        total += legal(part, 12, 13, 14);
        }
    }
    return total;
}

pub fn solve_day2_part2(data: String) -> u32 {
    let parts = data.split("\n");
    let mut total = 0;
    for part in parts {
        if part.len() > 0 {
        total += minimum_game(part);
        }
    }
    return total;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
        let mut test_cases: HashMap<String, u32> = HashMap::new();
        test_cases.insert("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green".to_string(), 1);
        test_cases.insert("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue".to_string(), 2);
        test_cases.insert("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red".to_string(), 0);
        test_cases.insert("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red".to_string(), 0);
        test_cases.insert("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green".to_string(), 5);
        test_cases.insert("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green\n".to_string(), 8);

        for (case, answer) in &test_cases {
            assert_eq!(&solve_day2_part1(case.to_string()),answer);
        } 
    }

    #[test]
    fn part2_tests() {
        let mut test_cases: HashMap<String, u32> = HashMap::new();
        test_cases.insert("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green".to_string(), 48);
        test_cases.insert("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue".to_string(),12);
        test_cases.insert("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red".to_string(), 1560);
        test_cases.insert("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red".to_string(), 630);
        test_cases.insert("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green".to_string(), 36);
        test_cases.insert("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green\n".to_string(), 2286);

        for (case, answer) in &test_cases {
            assert_eq!(&solve_day2_part2(case.to_string()),answer);
        } 
    }

}