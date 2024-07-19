mod helper;
use helper::read_lines;

const SAMPLE_DATA: &str = "A Y
B X
C Z";

pub fn p1() {
    let mut score_sum = 0;
    for line_result in read_lines("day2.txt") {
        let line = line_result.unwrap();

        // for line in SAMPLE_DATA.lines() {
        let columns: Vec<&str> = line.split(' ').collect();

        let move_score = match columns[1] {
            "X" => 0, //rock
            "Y" => 1, //paper
            "Z" => 2, //scissors
            _ => 0,
        };

        let opponent_score = match columns[0] {
            "A" => 0, //rock
            "B" => 1, //paper
            "C" => 2, //scissors
            _ => 0,
        };

        let win_draw_lose_score = if opponent_score - move_score == 0 {
            3
        } else if (opponent_score + 1) % 3 == move_score {
            6
        } else {
            0
        };

        score_sum += win_draw_lose_score + move_score + 1;
    }

    println!("{}", score_sum);
}

pub fn p2() {
    let mut score_sum = 0;
    for line_result in read_lines("day2.txt") {
        let line = line_result.unwrap();

        // for line in SAMPLE_DATA.lines() {
        let columns: Vec<&str> = line.split(' ').collect();

        let opponent_score = match columns[0] {
            "A" => 0, //rock
            "B" => 1, //paper
            "C" => 2, //scissors
            _ => 0,
        };

        let move_score = match columns[1] {
            "X" => {
                //lose
                let response = opponent_score - 1;
                if response < 0 {
                    2
                } else {
                    response
                }
            }
            "Y" => opponent_score,           //draw
            "Z" => (opponent_score + 1) % 3, //win
            _ => 0,
        };

        let win_draw_lose_score = if opponent_score - move_score == 0 {
            3
        } else if (opponent_score + 1) % 3 == move_score {
            6
        } else {
            0
        };

        score_sum += win_draw_lose_score + move_score + 1;
    }

    println!("{}", score_sum);
}
