use crate::aoc2015_01::{part1, part2};
use std::fs;

mod aoc2015_01;

fn get_parsed_file(file_path: String) -> String {
    let content = fs::read_to_string(file_path).expect("Failed to read the file");
    return content;
}

fn main() {
    let file_path = String::from("./inputs/aoc2015/day01.txt");
    let result = get_parsed_file(file_path);
    let characters = result.split("");
    let result = part1(characters.clone());
    print!("result parte 1: {result}");
    let result2 = part2(characters.clone());
    print!("result parte 2: {result2}");
}
