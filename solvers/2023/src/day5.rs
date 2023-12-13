#[derive(Debug)]
struct Mapper {
    base: String,
    result: String,
    map_start: i64,
    map_end: i64,
    adjustment: i64,
}

#[derive(Debug)]
struct Seed {
    stype: String,
    value: i64,
}

fn build_data(data: String) -> (Vec<Seed>, Vec<Mapper>) {
    let mut seeds: Vec<Seed>  = Vec::new();
    let mut mappers: Vec<Mapper> = Vec::new();

    let mut base = "";
    let mut result = "";

    let lines = data.split("\n");
    for line in lines {
        if line.trim().len() == 0 {
            continue;
        } else  if line.starts_with("seeds: ") {
            let parts = line.split_whitespace();
            for part in parts {
                if !part.starts_with("seeds") {
                seeds.push(Seed{stype: "seed".to_string(), value: part.parse::<i64>().unwrap()});
                }
            }
        } else if line.trim().ends_with("map:") {
            let mut parts = line.split_whitespace();
            let fp = parts.next().unwrap();
            let mut bits = fp.split("-");
            let fb = bits.next().unwrap();
            bits.next();
            let sb = bits.next().unwrap();
            base = fb;
            result = sb;
        } else {
            let mut nums = line.split_whitespace();
            let fp = nums.next().unwrap();
            let sp = nums.next().unwrap();
            let tp = nums.next().unwrap();

            mappers.push(Mapper{
                base: base.to_string(),
                result: result.to_string(),
                map_start: sp.parse::<i64>().unwrap(),
                map_end: sp.parse::<i64>().unwrap() + tp.parse::<i64>().unwrap(),
                adjustment: fp.parse::<i64>().unwrap() - sp.parse::<i64>().unwrap(),
            });
        }
    }

    return (seeds, mappers);
}

#[derive(Debug)]
struct Range {

}
#[derive(Debug)]
struct SeedRange {
    base: i64,
    end: i64,
    stype: String,
}

fn process_range(s: SeedRange, mappers: &Vec<Mapper>) -> i64 {
    let mut all_ranges: Vec<SeedRange> = Vec::new();
    all_ranges.push(s);
    let mut lowest = i64::MAX;

    while all_ranges.len() > 0 {
     //   println!("LEN {}", all_ranges.len());
        let c = all_ranges.pop().unwrap();
        if c.stype == "location" {
            if c.base < lowest {
                lowest = c.base;
            }
        } else {
            let mut processed = false;
            for mapper in mappers {
                if mapper.base == c.stype {
                    if c.base >= mapper.map_start  && c.end <= mapper.map_end {
                         // Mapper fully encloses range
                        all_ranges.push(SeedRange{
                            stype: mapper.result.clone(),
                            base: c.base + mapper.adjustment,
                            end: c.end + mapper.adjustment,
                        });
                        processed = true;
                        break;
                    } else if c.base >= mapper.map_start && c.base <= mapper.map_end {
                        // Range pops out of end of map
                        all_ranges.push(SeedRange{
                            stype: mapper.result.clone(),
                            base: c.base + mapper.adjustment,
                            end: mapper.map_end + mapper.adjustment,
                        });
                        all_ranges.push(SeedRange{
                            stype: mapper.base.clone(),
                            base: mapper.map_end+1,
                            end: c.end,
                        });
                        processed = true;
                        break;
                    } else if c.end >= mapper.map_start && c.end <= mapper.map_end {
                        // Range pops out of start of map
                        all_ranges.push(SeedRange{
                            stype: mapper.result.clone(),
                            base: mapper.map_start+mapper.adjustment,
                            end: c.end + mapper.adjustment,
                        });
                        all_ranges.push(SeedRange{
                            stype: mapper.base.clone(),
                            base: c.base,
                            end: mapper.map_start-1,
                        });
                        processed = true;
                        break;
                    }
                }   
            }

            if !processed {
                for mapper in mappers {
                    if mapper.base == c.stype {
                        all_ranges.push(SeedRange{
                            stype: mapper.result.clone(),
                            base: c.base,
                            end: c.end,
                        });
                        break;
                    }
                }
            }
        }
    }

    return lowest;
}

pub fn path_part_1(data: String) -> i64 {
    let (seeds, mappers) = build_data(data);
    let mut lowest: i64 = i64::MAX;
    for _seed in seeds {
        let nl = process_range(SeedRange{stype: "seed".to_string(),base: 13, end: 13}, &mappers);
        if nl < lowest {
            lowest = nl;
        }
    }
    return lowest;
}

pub fn path_part_2(data: String) -> i64 {
    let (seeds, mappers) = build_data(data);
    let mut lowest: i64 = i64::MAX;
    let mut first = 0;
    let  _second = 0;
    for seed in seeds {
        if first == 0 {
            first = seed.value
        } else {
            println!("HERE {} -> {}", first, first+seed.value);
            let nl = process_range(SeedRange{stype: "seed".to_string(),base: first, end: first+seed.value}, &mappers);
            if nl < lowest {
                lowest = nl;
            }
            first = 0;
        }
    }
    return lowest;
}

pub fn solve_day5_part1(data: String) -> i32 {
    let (seeds, mappers) = build_data(data);

    let mut lowest = i64::MAX;
 
    for mut seed in seeds {
        while seed.stype != "location" {
            let mut done = false;
            for mapper in &mappers {
                if mapper.base == seed.stype && mapper.map_start <= seed.value && mapper.map_end >= seed.value {
                    seed = Seed{
                        stype: mapper.result.to_string(),
                        value: seed.value + mapper.adjustment,
                    };
                    done = true
                }
            }
            if !done {
                for mapper in &mappers {
                    if mapper.base == seed.stype {
                        seed = Seed{
                            stype: mapper.result.to_string(),
                            value: seed.value,
                        };
                        break;
                    }
                }
            }
        }

        if seed.value < lowest {
            lowest = seed.value;
        }
    }

    return lowest.try_into().unwrap();
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
    let answer: i64 = 35;
    let _ianswer: i32 = 79;
    assert_eq!(path_part_1(data.to_string()), answer);
}

#[test]
fn part2_tests() {
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
    let answer: i64 = 46;
    assert_eq!(path_part_2(data.to_string()), answer);
}



}


