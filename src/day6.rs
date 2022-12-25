mod helper;
use std::collections::VecDeque;

use helper::read_lines;

const SAMPLE_DATA: &str = "mjqjpqmgbljsphdztnvjfqwrcgsmlb
bvwbjplbgvbhsrlpgdmjqwftvncz
nppdvjthqldpwncqszvftbrmjlhg
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw";

pub fn p1() {
    // for line in SAMPLE_DATA.lines() {
    for line_result in read_lines("day6.txt") {
        let line = line_result.unwrap();
        let chars: Vec<char> = line.chars().collect();

        for (count, window) in chars.windows(4).enumerate() {
            let mut set: HashSet<&char> = HashSet::new();
            let mut all_unique = true;
            for c in window {
                if !set.contains(c) {
                    set.insert(c);
                } else {
                    all_unique = false;
                    break;
                }
            }

            if all_unique {
                println!("{}", count + 4);
                break;
            }
        }
    }
}

pub fn p2() {
    // for line in SAMPLE_DATA.lines() {
    for line_result in read_lines("day6.txt") {
        let line = line_result.unwrap();
        let chars: Vec<char> = line.chars().collect();

        for (count, window) in chars.windows(14).enumerate() {
            let mut set: HashSet<&char> = HashSet::new();
            let mut all_unique = true;
            for c in window {
                if !set.contains(c) {
                    set.insert(c);
                } else {
                    all_unique = false;
                    break;
                }
            }

            if all_unique {
                // println!("{:?}", window);
                println!("{}", count + 14);
                break;
            }
        }
    }
}
