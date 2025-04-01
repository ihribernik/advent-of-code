use std::str::Split;

pub fn part1(characters: Split<'_, &str>) -> i32 {
    let mut result = 0;
    for char in characters {
        result = match char {
            "(" => result + 1,
            ")" => result - 1,
            _ => result,
        };
    }
    return result;
}

pub fn part2(characters: Split<'_, &str>) -> usize {
    let mut result = 0;
    let mut index_floor = 0;
    for (index, char) in characters.enumerate() {
        (result, index_floor) = match char {
            "(" => (result + 1, index),
            ")" => (result - 1, index),
            _ => (result, index),
        };
        if result < 0 {
            break;
        }
    }
    println!("stoped at: {index_floor}");
    return index_floor;
}