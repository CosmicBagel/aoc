//https://adventofcode.com/2022/day/1

use std::{
    fs::File,
    io::{self, BufRead},
};

pub fn p1() {
    let fh = File::open("day1.txt").unwrap();
    let reader = io::BufReader::new(fh);
    let lines = reader.lines();

    let mut sum = 0;
    let mut max: u32 = 0;
    for line_result in lines {
        let line = line_result.unwrap();
        if line.len() == 0 {
            if sum > max {
                max = sum;
            }
            sum = 0;
            continue;
        }

        let line_val: u32 = line.parse().unwrap();
        sum += line_val;
    }
    //data doesn't end with empty line, need one last check after loop
    if sum > max {
        max = sum;
    }

    println!("max calorie {}", max);
}

pub fn p2() {
    let fh = File::open("day1.txt").unwrap();
    let reader = io::BufReader::new(fh);
    let lines = reader.lines();

    let mut cals_per_elf: Vec<u32> = Vec::new();

    let mut sum = 0;
    for line_result in lines {
        let line = line_result.unwrap();
        if line.len() == 0 {
            cals_per_elf.push(sum);
            sum = 0;
            continue;
        }

        let line_val: u32 = line.parse().unwrap();
        sum += line_val;
    }
    //data doesn't end with empty line, need one last check after loop
    cals_per_elf.push(sum);

    // sort and sum top 3
    cals_per_elf.sort();
    cals_per_elf.reverse();

    let mut top_three_sum: u32 = 0;
    for val in &cals_per_elf[..3] {
        top_three_sum += val;
    }

    println!("top three sum {}", top_three_sum);
}
