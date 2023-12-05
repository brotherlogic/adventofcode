struct Mapper {
    base: String,
    result: String,
    map_start: u32,
    map_end: u32,
    adjustment: u32,
}

struct Seed {
    stype: String,
    value: u32,
}

pub fn solve_day5_part1(data: String) -> u32 {
    return 0;
}


#[cfg(test)]
mod testsca {
    use super::*;

#[test]
fn part1_tests() {
    let data = "seeds: 79 14 55 13
seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4".to_string();
    let answer: u32 = 35;
    assert_eq!(solve_day5_part1(data.to_string()), answer)
}
}