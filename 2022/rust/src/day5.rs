mod helper;
use std::collections::VecDeque;

use helper::read_lines;

/*
[T] [V]                     [W]
[V] [C] [P] [D]             [B]
[J] [P] [R] [N] [B]         [Z]
[W] [Q] [D] [M] [T]     [L] [T]
[N] [J] [H] [B] [P] [T] [P] [L]
[R] [D] [F] [P] [R] [P] [R] [S] [G]
[M] [W] [J] [R] [V] [B] [J] [C] [S]
[S] [B] [B] [F] [H] [C] [B] [N] [L]
 1   2   3   4   5   6   7   8   9
*/
const SAMPLE_DATA: &str = "move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";
//expected result is CMZ

pub fn p1() {
    // let mut stacks: Vec<Vec<&str>> = vec![
    //     vec!["Z", "N"],
    //     vec!["M", "C", "D"],
    //     vec!["P"]];
    let mut stacks: Vec<Vec<&str>> = vec![
        vec!["T", "V", "J", "W", "N", "R", "M", "S"],
        vec!["V", "C", "P", "Q", "J", "D", "W", "B"],
        vec!["P", "R", "D", "H", "F", "J", "B"],
        vec!["D", "N", "M", "B", "P", "R", "F"],
        vec!["B", "T", "P", "R", "V", "H"],
        vec!["T", "P", "B", "C"],
        vec!["L", "P", "R", "J", "B"],
        vec!["W", "B", "Z", "T", "L", "S", "C", "N"],
        vec!["G", "S", "L"],
    ];
    // fucked up and put everything in backwards so have to reverse them first
    for s in &mut stacks[..] {
        s.reverse();
    }

    // for line in SAMPLE_DATA.lines() {
    for line_result in read_lines("day5.txt") {
        let line = line_result.unwrap();

        let words: Vec<&str> = line.split(" ").collect();
        let quantity: u32 = words[1].parse().unwrap();
        let from: u32 = words[3].parse().unwrap();
        let to: u32 = words[5].parse().unwrap();

        let mut crane: VecDeque<&str> = VecDeque::new();
        //put things on crane
        for _ in 0..quantity {
            let container_result = stacks[(from - 1) as usize].pop();
            match container_result {
                Some(container) => crane.push_back(container),
                _ => (),
            };
        }

        //offload things
        for _ in 0..quantity {
            let container_result = crane.pop_front();
            match container_result {
                Some(container) => stacks[(to - 1) as usize].push(container),
                _ => (),
            };
        }
    }

    for s in stacks {
        if s.len() > 0 {
            print!("{}", s[s.len() - 1 as usize]);
        }
    }

    println!("");
}

pub fn p2() {
    // let mut stacks: Vec<Vec<&str>> = vec![
    //     vec!["Z", "N"],
    //     vec!["M", "C", "D"],
    //     vec!["P"]];
    let mut stacks: Vec<Vec<&str>> = vec![
        vec!["T", "V", "J", "W", "N", "R", "M", "S"],
        vec!["V", "C", "P", "Q", "J", "D", "W", "B"],
        vec!["P", "R", "D", "H", "F", "J", "B"],
        vec!["D", "N", "M", "B", "P", "R", "F"],
        vec!["B", "T", "P", "R", "V", "H"],
        vec!["T", "P", "B", "C"],
        vec!["L", "P", "R", "J", "B"],
        vec!["W", "B", "Z", "T", "L", "S", "C", "N"],
        vec!["G", "S", "L"],
    ];
    // fucked up and put everything in backwards so have to reverse them first
    for s in &mut stacks[..] {
        s.reverse();
    }

    // for line in SAMPLE_DATA.lines() {
    for line_result in read_lines("day5.txt") {
        let line = line_result.unwrap();

        let words: Vec<&str> = line.split(" ").collect();
        let quantity: u32 = words[1].parse().unwrap();
        let from: u32 = words[3].parse().unwrap();
        let to: u32 = words[5].parse().unwrap();

        let mut crane: Vec<&str> = Vec::new();
        //put things on crane
        for _ in 0..quantity {
            let container_result = stacks[(from - 1) as usize].pop();
            match container_result {
                Some(container) => crane.push(container),
                _ => (),
            };
        }

        //offload things
        for _ in 0..quantity {
            let container_result = crane.pop();
            match container_result {
                Some(container) => stacks[(to - 1) as usize].push(container),
                _ => (),
            };
        }
    }

    for s in stacks {
        if s.len() > 0 {
            print!("{}", s[s.len() - 1 as usize]);
        }
    }

    println!("");

}