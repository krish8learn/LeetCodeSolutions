fn main() {
    println!("{}", Solution::to_lower_case("Hello".into()));
}

struct  Solution;
impl Solution {
    pub fn to_lower_case(s: String) -> String {
        s.to_lowercase()
    }
}