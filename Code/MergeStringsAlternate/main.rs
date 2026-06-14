struct Solution; 

impl Solution {
    pub fn merge_alternately(word1: String, word2: String) -> String {
        let mut result = String::new();
        let mut index = 0;

        let chars1: Vec<char> = word1.chars().collect();
        let chars2: Vec<char> = word2.chars().collect();

        let len1 = chars1.len();
        let len2 = chars2.len();

        while index < len1 || index < len2 {
            if index < len1 {
                result.push(chars1[index]);
            }

            if index < len2 {
                result.push(chars2[index]);
            }

            index += 1;
        }

        result

    }
}

fn main() {

    let result = Solution::merge_alternately(
        String::from("abc"),
        String::from("pqr"),
    );
    println!("{}", result); // apbqcr
}

