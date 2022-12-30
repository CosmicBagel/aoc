#![allow(dead_code)]
use crate::helper::read_lines;

mod helper;

fn main() {
    p2();
}

// all outer trees visible (16)
// 5 inner trees visible
// 21 total visible trees
const SAMPLE_DATA: &str = "30373
25512
65332
33549
35390";

fn p1() {
    // read in full grid of heights
    // vec<vec<(u8, bool)>>
    // determine width and height
    // look at length of first row to determine vec reserve size
    // set bool false (not visible)
    // parse each char in line, push to row

    let mut height_grid: Vec<Vec<(u8, bool)>> = Vec::new();
    // for line in SAMPLE_DATA.lines() {
    for line_result in read_lines("day8.txt") {
        let line = line_result.unwrap();

        let line_chars = line.chars();
        let mut row: Vec<(u8, bool)> = Vec::new();
        for c in line_chars {
            let val = c.to_digit(10).unwrap() as u8;
            row.push((val, false));
        }
        height_grid.push(row);
    }

    // assume all outer trees are visible, add length of border

    // left to right, top to bottom, right to left, bottom to top
    // check which trees are visible, mark as such if is visible
    // count visible trees of inner trees (skip first and bottom rows and well
    // as first and last indices of rows)

    let width = height_grid[0].len();
    let height = height_grid.len();

    // left-to-right
    for y in 1..(height - 1) {
        let row = &mut height_grid[y];
        let mut highest: u8 = row[0].0;
        for x in 1..(width - 1) {
            if row[x].0 > highest {
                highest = row[x].0;
                row[x].1 = true;
            }
        }
    }

    // top-to-bottom
    for x in 1..(width - 1) {
        let mut highest = height_grid[0][x].0;
        for y in 1..(height - 1) {
            let cell = &mut height_grid[y][x];
            if cell.0 > highest {
                highest = cell.0;
                cell.1 = true;
            }
        }
    }

    // right-to-left
    for y in 1..(height - 1) {
        let mut highest: u8 = height_grid[y][width - 1].0;
        for x in (0..(width - 1)).rev() {
            let cell = &mut height_grid[y][x];
            if cell.0 > highest {
                highest = cell.0;
                cell.1 = true;
            }
        }
    }

    // bottom-to-top
    for x in 1..(width - 1) {
        let mut highest = height_grid[height - 1][x].0;
        for y in (0..(height - 1)).rev() {
            let cell = &mut height_grid[y][x];
            if cell.0 > highest {
                highest = cell.0;
                cell.1 = true;
            }
        }
    }

    // print visible trees
    let mut count = 0;
    for y in 1..(height - 1) {
        // print!("{}:", y);
        for x in 1..(width - 1) {
            let cell = height_grid[y][x];
            if cell.1 {
                // print!("{}", cell.0);
                count += 1;
            } else {
                // print!(" ");
            }
        }
        // print!("\n");
    }
    //add perimeter trees
    count += width * 2;
    count += (height - 2) * 2;
    println!("visible trees: {}", count);
}

fn p2() {
    let mut tree_height_grid: Vec<Vec<u8>> = Vec::new();
    // for line in SAMPLE_DATA.lines() {
        for line_result in read_lines("day8.txt") {
        let line = line_result.unwrap();

        let line_chars = line.chars();
        let mut row: Vec<u8> = Vec::new();
        for c in line_chars {
            let val = c.to_digit(10).unwrap() as u8;
            row.push(val);
        }
        tree_height_grid.push(row);
    }

    let height = tree_height_grid.len();
    let width = tree_height_grid[0].len();

    let mut highest_score = 0u32;
    for proposed_y in 1..height {
        for proposed_x in 2..width {
            let proposed_tree = tree_height_grid[proposed_y][proposed_x];

            let mut right_count = 0u32;
            let mut left_count = 0u32;
            let mut bottom_count = 0u32;
            let mut top_count = 0u32;

            let house_height: u8 = proposed_tree;

            // right side trees
            for adjacent_x in (proposed_x + 1)..width {
                right_count += 1;

                let adjacent_tree = tree_height_grid[proposed_y][adjacent_x];
                if adjacent_tree >= house_height {
                    break;
                }
            }

            // left side trees
            for adjacent_x in (0..proposed_x).rev() {
                left_count += 1;

                let adjacent_tree = tree_height_grid[proposed_y][adjacent_x];
                if adjacent_tree >= house_height {
                    break;
                }
            }

            // bottom side trees
            for adjacent_y in (proposed_y + 1)..height {
                bottom_count += 1;

                let adjacent_tree = tree_height_grid[adjacent_y][proposed_x];
                if adjacent_tree >= house_height {
                    break;
                }
            }

            // top side trees
            for adjacent_y in (0..proposed_y).rev() {
                top_count += 1;

                let adjacent_tree = tree_height_grid[adjacent_y][proposed_x];
                if adjacent_tree >= house_height {
                    break;
                }
            }

            // might have to change this to not multiply zero counts, not sure if that's intentional
            // for the scoring
            let score = right_count * left_count * bottom_count * top_count;
            if score > highest_score {
                highest_score = score;
            }
        }
    }

    println!("highest score: {}", highest_score);
}
