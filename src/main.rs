fn main() {
    p1();
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
        // look at length of first row to deterimine vec reserve size
        // set bool false (not visible)
        // parse each char in line, push to row

    // assume all outer trees are visible, add length of border
    
    // left to right, top to bottom, right to left, bottom to top
    // check which trees are visible, mark as such if is visible
    // count visible trees of inner trees (skip first and bottom rows and well 
    // as first and last indices of rows)
    // 


}
