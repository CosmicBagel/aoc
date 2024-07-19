mod helper;
use helper::read_lines;

const SAMPLE_DATA: &str = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

pub fn p1() {
    let mut misplaced_item_priority_sum = 0;
    for line_result in read_lines("day3.txt") {
        let line = line_result.unwrap();
        // for line in SAMPLE_DATA.lines() {

        //two parallel arrays of 52 len one for each compartment
        //each slot corresponds with an item type
        //fill out the arrays
        let mut compartment_a = vec![false; 52];
        let mut compartment_b = vec![false; 52];

        let (line_a, line_b) = line.split_at(line.len() / 2);
        let line_a_bytes = line_a.as_bytes();
        let line_b_bytes = line_b.as_bytes();
        for i in 0..line_a.len() {
            let a_ind = (char_to_priority(line_a_bytes[i]) - 1) as usize;
            let b_ind = (char_to_priority(line_b_bytes[i]) - 1) as usize;
            compartment_a[a_ind] = true;
            compartment_b[b_ind] = true;
        }

        //afterwards go through the arrays in parallel and check if any
        // of the slots match
        for i in 0..compartment_a.len() {
            if compartment_a[i] && compartment_b[i] {
                misplaced_item_priority_sum += i + 1;
                break;
            }
        }
    }
    println!("{}", misplaced_item_priority_sum);
}

//take letter, convert it type item time / priority number
//only grabs first byte, no good with multi-byte characters
//only use with ascii characters
fn char_to_priority(c: u8) -> u8 {
    if c >= 65 && c <= 90 {
        c - 38
    } else {
        c - 96
    }
}

pub fn p2() {
    let mut misplaced_item_priority_sum = 0;

    let lines_vec: Vec<String> = read_lines("day3.txt").map(|lr| lr.unwrap()).collect();
    for lines_triplet in lines_vec.chunks(3) {
        let line_a = &lines_triplet[0];
        let line_b = &lines_triplet[1];
        let line_c = &lines_triplet[2];

        //two parallel arrays of 52 len one for each compartment
        //each slot corresponds with an item type
        //fill out the arrays
        let mut rucksack_a = vec![false; 52];
        let mut rucksack_b = vec![false; 52];
        let mut rucksack_c = vec![false; 52];

        let line_a_bytes = line_a.as_bytes();
        let line_b_bytes = line_b.as_bytes();
        let line_c_bytes = line_c.as_bytes();
        for i in 0..line_a_bytes.len() {
            let a_ind = (char_to_priority(line_a_bytes[i]) - 1) as usize;
            rucksack_a[a_ind] = true;
        }
        for i in 0..line_b_bytes.len() {
            let b_ind = (char_to_priority(line_b_bytes[i]) - 1) as usize;
            rucksack_b[b_ind] = true;
        }

        for i in 0..line_c_bytes.len() {
            let c_ind = (char_to_priority(line_c_bytes[i]) - 1) as usize;
            rucksack_c[c_ind] = true;
        }

        //afterwards go through the arrays in parallel and check if any
        // of the slots match
        for i in 0..rucksack_a.len() {
            if rucksack_a[i] && rucksack_b[i] && rucksack_c[i] {
                misplaced_item_priority_sum += i + 1;
                break;
            }
        }
    }
    println!("{}", misplaced_item_priority_sum);
}
