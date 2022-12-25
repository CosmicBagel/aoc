mod helper;
use helper::read_lines;

const SAMPLE_DATA: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

pub fn p1() {
    let mut count = 0;
    // for line in SAMPLE_DATA.lines() {
    for line_result in read_lines("day4.txt") {
        let line = line_result.unwrap();
        let pair: Vec<&str> = line.split(",").collect();

        let range_a: Vec<&str> = pair[0].split("-").collect();
        let range_a_start: u32 = range_a[0].parse().unwrap();
        let range_a_end: u32 = range_a[1].parse().unwrap();

        let range_b: Vec<&str> = pair[1].split("-").collect();
        let range_b_start: u32 = range_b[0].parse().unwrap();
        let range_b_end: u32 = range_b[1].parse().unwrap();

        if range_a_start >= range_b_start && range_a_end <= range_b_end {
            count += 1;
            continue;
        }
        if range_b_start >= range_a_start && range_b_end <= range_a_end {
            count += 1;
            continue;
        }
    }
    println!("{}", count);
}

pub fn p2() {
    let mut count = 0;
    // for line in SAMPLE_DATA.lines() {
    for line_result in read_lines("day4.txt") {
        let line = line_result.unwrap();
        let pair: Vec<&str> = line.split(",").collect();

        let range_a: Vec<&str> = pair[0].split("-").collect();
        let range_a_start: u32 = range_a[0].parse().unwrap();
        let range_a_end: u32 = range_a[1].parse().unwrap();

        let range_b: Vec<&str> = pair[1].split("-").collect();
        let range_b_start: u32 = range_b[0].parse().unwrap();
        let range_b_end: u32 = range_b[1].parse().unwrap();

        if range_a_start >= range_b_start && range_a_start <= range_b_end {
            count += 1;
            continue;
        }
        if range_b_start >= range_a_start && range_b_start <= range_a_end {
            count += 1;
            continue;
        }
    }
    println!("{}", count);
}