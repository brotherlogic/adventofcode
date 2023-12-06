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

    let mut lines = data.split("\n");
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
    base: i64,
    end: i64,
}

struct SeedRange {
    base: i64,
    end: i64,
    stype: String,
}

fn process_range(s: SeedRange, mappers: &Vec<Mapper>) -> i64 {
    let curr := s.stype;
    let all_ranges: Vec<SeedRange> = Vec::new();
    all_ranges.push(s);

    while curr != "location" {
        let mut new_ranges = Vec::new();
        for mapper in mappers {
                for seed in all_ranges {
      
                if mapper.base == seed.stype {
                    // Current range 
                }
            }
        }
    }
    let ranges = expand_ranges(s)
}

fn path(data: String) -> i32 {
    let (seeds, mappers) = build_data(data);

    return run_range(Range{base:79, end: 79}, &mappers, "seed").try_into().unwrap();
}

fn overlap(r: Range, m: Mapper) -> Range {
    let mut base = 0;
    let mut end = 0;
    if r.base < m.map_start {
        base = m.map_start;
    }  else {
        base = r.base
    }

    if r.end > m.map_end {
        end = m.map_end;
    } else {
        end = r.end
    }

    if base < end {
        return Range{base: base, end: end};
    }
    return Range{base: 0, end: 0};
}

fn run_range(r: Range, mappers: &Vec<Mapper>, curr: &str) -> i64 {
    println!("RANGE {:?} {}", r, curr);
    let mut best = i64::MAX;
    for mapper in mappers {
        if mapper.base == curr {
            println!("RUNNING MAPPER {:?}", mapper);
            let mut base = 0;
            let mut end = 0;
            if r.base < mapper.map_start {
                base = mapper.map_start;
            }  else {
                base = r.base
            }
        
            if r.end > mapper.map_end {
                end = mapper.map_end;
            } else {
                end = r.end
            }

         
            let mut overlap = Range{base: r.base, end: r.end};
            if base <= end {
                overlap =  Range{base: base, end: end};
            }
             
            println!("OVERLAP {:?}", overlap);

            if overlap.base != 0 && overlap.end != 0 {
                let nbest = run_range(overlap, mappers, &mapper.result);
                if nbest < best {
                    best = nbest;
                }
            }
        }
    }

    println!("HERE {}", best);
    return best;
}

fn reverse_solve(data: String) -> i32 {
    let (seeds, mappers) = build_data(data);
    let mut start = 0;
    loop {
        let value = reverse(start, &mappers);
        for s in &seeds {
            if value == s.value {
                return start.try_into().unwrap();
            }
        }
        start+=1
    }
}

pub fn reverse_solve_part2(data: String) -> i32 {
    let (seeds, mappers) = build_data(data);
    let mut start = 0;
    loop {
        let value = reverse(start, &mappers);
        let mut first: i64 = 0;
        let mut second: i64 = 0;
        for s in &seeds {
            if first == 0 {
                first = s.value;
            } else {
                if value >= first && value <= first+s.value {
                    return start.try_into().unwrap();
                }
                first = 0;
                second = 0;
            }
        }
        start+=1
    }
}

fn reverse(result: i64, mappers: &Vec<Mapper>) -> i64 {
    let mut bs = Seed{stype: "location".to_string(), value: result};

    let mut lc = 0;
    while bs.stype != "seed" {
        let mut done = false;
        lc += 1;
        if lc > 100 {
            return 99;
        }
        for mapper in mappers {
            if mapper.result == bs.stype && mapper.map_start <= bs.value-mapper.adjustment && mapper.map_end >= bs.value-mapper.adjustment {
                bs = Seed{
                    stype: mapper.base.to_string(),
                    value: bs.value - mapper.adjustment,
                };
                done = true;
            }
        }
        if !done {
            for mapper in mappers {
                if mapper.result == bs.stype {
                    bs = Seed{
                        stype: mapper.base.to_string(),
                        value: bs.value,
                    };
                    break;
                }
            }
        }
    }

    return bs.value;
}

pub fn solve_day5_part2(data: String) -> i32 {
    println!("Starting");
    let (seeds, mappers) = build_data(data);

    let mut lowest = i64::MAX;
    let mut start: i64 = 0;
    let mut end: i64 = 0;
 
    for  tseed in seeds {
        if start == 0 {
            start = tseed.value;
            continue;
        }
        if end == 0 {
            end = tseed.value;
        }

        for sv in start..start+end {
            let mut seed = Seed{value: sv, stype: "seed".to_string()};
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

        start = 0;
        end = 0;
    }

    return lowest.try_into().unwrap();
}

pub fn solve_day5_part1(data: String) -> i32 {
    println!("Starting");
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
    let answer: i32 = 35;
    let ianswer: i32 = 79;
    assert_eq!(solve_day5_part1(data.to_string()), answer);
    assert_eq!(reverse_solve(data.to_string()), answer);
    assert_eq!(path(data.to_string()), ianswer);
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
    let answer: i32 = 46;
    assert_eq!(reverse_solve_part2(data.to_string()), answer)
}
}