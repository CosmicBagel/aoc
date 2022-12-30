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
        let mut highest = height_grid[height-1][x].0;
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
    // read in full grid of heights
    // vec<vec<(u8, bool)>>
    // determine width and height
    // look at length of first row to determine vec reserve size
    // set bool false (not visible)
    // parse each char in line, push to row

    let mut height_grid: Vec<Vec<(u8, bool)>> = Vec::new();
    for line in SAMPLE_DATA.lines() {
	// for line_result in read_lines("day8.txt") {
        // let line = line_result.unwrap();

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
        let mut highest = height_grid[height-1][x].0;
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