pub fn solve_day22_part1(data: String) -> i32 {
    let mut blocks = build_blocks(data).clone();

    let mut move_count = 0;
    loop {
        let cblock = blocks.remove(0);

        if can_fall(cblock.clone(), blocks.clone()) {
            move_count = 0;
            let nblock = Block {
                x: cblock.x,
                y: cblock.y,
                z: cblock.z - 1,
                wx: cblock.wx,
                wy: cblock.wy,
                wz: cblock.wz,
            };
            blocks.push(nblock);
        } else {
            move_count += 1;
            blocks.push(cblock);
        }

        if move_count > blocks.len() {
            break;
        }
    }

    let mut fall_count = 0;
    // Now count the non movers
    for _ in 0..blocks.len() {
        let tblock = blocks.remove(0);

        for _ in 0..blocks.len() {
            let mut found_fall = false;
            let ttblock = blocks.remove(0);
            if !found_fall && can_fall(ttblock.clone(), blocks.clone()) {
                fall_count += 1;
            }
            blocks.push(ttblock);
        }
    }

    return fall_count;
}

fn can_fall(b: Block, bs: Vec<Block>) -> bool {
    // Block has reached the ground
    println!("CHECKING {:?}", b);
    if b.z == 1 {
        println!("ALREADY AT GROUND");
        return false;
    }

    for bb in bs {
        if b.wx > 1 {
            for x in b.x..b.x + b.wx {
                if blocks(x, b.y, b.z, bb.clone()) {
                    println!("{:?} blocks on x {:?}", bb, b);
                    return false;
                }
            }
        }
        if b.wy > 1 {
            for y in b.y..b.y + b.wy {
                if blocks(b.x, y, b.z, bb.clone()) {
                    return false;
                }
            }
        }
        if b.wz > 1 {
            for z in b.z..b.z + b.wz {
                if blocks(b.x, b.y, z, bb.clone()) {
                    return false;
                }
            }
        }
    }

    return true;
}

fn blocks(x: i32, y: i32, z: i32, b: Block) -> bool {
    if b.wx > 1 {
        for bx in b.x..b.x + b.wx {
            if bx == x && b.y == y && b.z == z - 1 {
                return true;
            }
        }
    }
    if b.wy > 1 {
        for by in b.y..b.y + b.wy {
            if b.x == x && by == y && b.z == z - 1 {
                return true;
            }
        }
    }
    if b.wz > 1 {
        for bz in b.z..b.z + b.wz {
            if b.x == x && b.y == y && bz == z - 1 {
                return true;
            }
        }
    }

    return false;
}

#[derive(Clone, Debug)]
struct Block {
    x: i32,
    y: i32,
    z: i32,
    wx: i32,
    wy: i32,
    wz: i32,
}

fn build_blocks(data: String) -> Vec<Block> {
    let mut blocks = Vec::new();

    for line in data.split("\n") {
        let mut pieces = line.trim().split("~");
        let mut pos = pieces.next().unwrap().split(",");
        let mut siz = pieces.next().unwrap().split(",");

        let x = pos.next().unwrap().parse::<i32>().unwrap();
        let y = pos.next().unwrap().parse::<i32>().unwrap();
        let z = pos.next().unwrap().parse::<i32>().unwrap();

        let wx = 1 + siz.next().unwrap().parse::<i32>().unwrap() - x;
        let wy = 1 + siz.next().unwrap().parse::<i32>().unwrap() - y;
        let wz = 1 + siz.next().unwrap().parse::<i32>().unwrap() - z;

        blocks.push(Block {
            x: x,
            y: y,
            z: z,
            wx: wx,
            wy: wy,
            wz: wz,
        });
    }

    return blocks;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
        let test_case = "1,0,1~1,2,1
        0,0,2~2,0,2
        0,2,3~2,2,3
        0,0,4~0,2,4
        2,0,5~2,2,5
        0,1,6~2,1,6
        1,1,8~1,1,9"
            .to_string();

        let pulses = solve_day22_part1(test_case);
        assert_eq!(pulses, 5)
    }
}
