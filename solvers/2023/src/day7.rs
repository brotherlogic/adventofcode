use std::collections::HashMap;

#[derive(Debug)]
struct Hand {
    cards: String,
    bid: i64,
    ordering: String,
}

fn translate(cards: String) -> String {
    let mut nstr = "".to_string();
    for c in cards.chars() {
        match c {
            '2' => nstr += "A",
            '3' => nstr += "B",
            '4' => nstr += "C",
            '5' => nstr += "D",
            '6' => nstr += "E",
            '7' => nstr += "F",
            '8' => nstr += "G",
            '9' => nstr += "H",
            'T' => nstr += "I",
            'J' => nstr += "J",
            'Q' => nstr += "K",
            'K' => nstr += "L",
            'A' => nstr += "M",
            _   => nstr += "N",
        }
    }
    return nstr;
}

fn get_rank(cards: String) -> String {
    let mut mapper: HashMap<char, u32> = HashMap::new();
    for c in cards.chars() {
        if mapper.contains_key(&c) {
            mapper.insert(c, mapper.get(&c).unwrap()+1);
        } else {
            mapper.insert(c, 1);
        }
    }

    if mapper.keys().len() == 1 {
        return "G".to_string();
    }

    if mapper.keys().len() == 2 {
        for (key, val) in &mapper {
            if *val == 1 || *val == 4 {
                return "F".to_string();
            }
            return "E".to_string();
        }
    }

    if mapper.keys().len() == 3 {
        for (key, val) in &mapper {
            if *val == 3 {
                return "D".to_string();
            } else if *val == 2 {
              return "C".to_string();
            }
        }
    }

    if mapper.keys().len() == 4{
        return "B".to_string();
    }
        return "A".to_string();
}

fn get_ordering(cards: String) -> String {
    return get_rank(cards.clone()) + "-" + &translate(cards);
}

fn build_hands(data: String) -> Vec<Hand> {
    let mut hands = Vec::new();
    for line in data.split("\n") {
        let mut elems = line.trim().split_whitespace();
        let cards = elems.next().unwrap().to_string();
        let bid = elems.next().unwrap().parse::<i64>().unwrap();

        hands.push(Hand{cards: cards.clone(), bid: bid, ordering: get_ordering(cards)});
    }

    return hands;
}


pub fn solve_day7_part1(data: String) -> i64 {
    let mut hands = build_hands(data);

    // Sort the hands here
    hands.sort_by(|a, b| a.ordering.cmp(&b.ordering));

    let mut rank = 1;
    let mut total = 0;
    for hand in hands {
        println!("HAND {:?}", hand);
        total += rank * hand.bid;
        rank+=1;
    }

    return total;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
       let test_case = "32T3K 765
       T55J5 684
       KK677 28
       KTJJT 220
       QQQJA 483".to_string();

       let score = solve_day7_part1(test_case);
       assert_eq!(score, 6440)
    }
}