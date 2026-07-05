fn main() {
    println!("{}", Solution::length_of_last_word("Hello World".into()));
}

struct  Solution;

impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        s.trim()
            .split_whitespace()
            .last()
            .map_or(0, |word| word.len() as i32)      
    }
}

