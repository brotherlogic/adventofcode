use std::collections::HashMap;

struct Scratch {
    card_number: u32,
    winners: Vec<u32>,
    hand: Vec<u32>,
}

fn build_scratch(data: String) -> Vec<Scratch> {
    let mut all = Vec::new();

    let  lines = data.split("\n");
    for line in lines {
        let mut s = Scratch{
            card_number: 0,
            winners: Vec::new(),
            hand: Vec::new(),
        };

        let mut bits = line.split(":");
        let fb = bits.next().unwrap();
        let mut first_bits = fb.split_whitespace();
        first_bits.next();
        let gn = first_bits.next().unwrap();
        s.card_number = gn.parse::<u32>().unwrap();

        let  other = bits.next().unwrap().trim();
        let mut halves = other.split("|");
        let  wins = halves.next().unwrap().trim();
        for w in wins.split_whitespace() {
            s.winners.push(w.parse::<u32>().unwrap());
        }
        let  haves = halves.next().unwrap().trim();
        for h in haves.split_whitespace() {
            s.hand.push(h.parse::<u32>().unwrap());
        }

        all.push(s);
    }

    return all;
}

pub fn solve_day4_part1(data: String) -> u32 {
    let mut total = 0;

    for game in build_scratch(data) {
        let mut gtotal = 0;
        for winner in &game.winners {
            for mine in &game.hand {
                if mine == winner {
                    println!("FOUND {}", mine);
                    if gtotal == 0 {
                    gtotal += 1;
                    } else {
                        gtotal *=2;
                    }
                    break;
                }
            }
        }
        total += gtotal;
    }

    return total;
}

pub fn solve_day4_part2(data: String) -> u32 {
    let mut total = 0;
    let mut mapper: HashMap<u32,u32> = HashMap::new();

    for game in build_scratch(data).iter().rev() {
        let mut gtotal = 0;
        for winner in &game.winners {
            for mine in &game.hand {
                if mine == winner {
                    gtotal += 1;
                    break;
                }
            }
        }

        let mut ntotal: u32 = 1;
        for gval in game.card_number+1..game.card_number+gtotal+1 {
            ntotal += mapper[&gval]
        }
        mapper.insert(game.card_number, ntotal);
 
    }

    for (_key, value) in mapper {
        total += value;
    }

    return total;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
        let mut test_cases: HashMap<String, u32> = HashMap::new();
        test_cases.insert("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53".to_string(), 8);
        test_cases.insert("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19".to_string(), 2);
        test_cases.insert("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1".to_string(), 2);
        test_cases.insert("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83".to_string(), 1);
        test_cases.insert("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36".to_string(), 0);
        test_cases.insert("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11".to_string(), 0);
        test_cases.insert("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11".to_string(), 13);

        for (case, answer) in &test_cases {
            assert_eq!(&solve_day4_part1(case.to_string()),answer);
        } 
    }

    #[test]
    fn part2_tests() {
        let case = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11".to_string();
        let answer: u32 = 30;
        assert_eq!(solve_day4_part2(case.to_string()), answer)
    }
}