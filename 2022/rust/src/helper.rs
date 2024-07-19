use std::{
    fs::File,
    io::{self, BufReader, BufRead},
};

pub fn read_lines(filename: &str) -> std::io::Lines<BufReader<File>> {
    let fh = File::open(filename).unwrap();
    let reader = io::BufReader::new(fh);

    reader.lines()
}