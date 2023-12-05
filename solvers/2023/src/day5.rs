#[derive(Debug)]
struct Mapper {
    base: String,
    result: String,
    map_start: i32,
    map_end: i32,
    adjustment: i32,
}

#[derive(Debug)]
struct Seed {
    stype: String,
    value: i32,
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
        log.Println!("Parsing {}", line);
            let parts = line.split_whitespace();
            for part in parts {
                if !part.starts_with("seeds") {
                log.Println!("Seedv {}", part)
                    seeds.push(Seed{stype: "seed".to_string(), value: part.parse::<i32>().unwrap()});
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
                map_start: sp.parse::<i32>().unwrap(),
                map_end: sp.parse::<i32>().unwrap() + tp.parse::<i32>().unwrap(),
                adjustment: fp.parse::<i32>().unwrap() - sp.parse::<i32>().unwrap(),
            });
        }
    }

    return (seeds, mappers);
}

pub fn solve_day5_part1(data: String) -> i32 {
    println!("Starting");
    let (seeds, mappers) = build_data(data);

    let mut lowest = 99999999;
 
    for mut seed in seeds {
        while seed.stype != "location" {
            println!("HERE {:?}", seed);
            let mut done = false;
            for mapper in &mappers {
                if mapper.base == seed.stype && mapper.map_start <= seed.value && mapper.map_end >= seed.value {
                    println!("APPLYING {:?}", mapper);
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
                        println!("APPLYING {:?}", mapper);
                        seed = Seed{
                            stype: mapper.result.to_string(),
                            value: seed.value,
                        };
                        break;
                    }
                }
            }

            println!("RESULT {:?}", seed);
        }

        if seed.value < lowest {
            lowest = seed.value;
        }
    }

    return lowest;
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
    assert_eq!(solve_day5_part1(data.to_string()), answer)
}
}