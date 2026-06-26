
fn main(){
    println!("{}", Solution::reverse_vowels("hello".to_string()));
}
struct Solution;

impl Solution {
    pub fn reverse_vowels(s: String) -> String {
        let mut chars: Vec<char> = s.chars().collect();
        let mut left = 0;
        let mut right = chars.len() - 1;
        while left < right {
            while left < right && !Self::is_vowel(chars[left]) {
                left += 1;
            }
            while left < right && !Self::is_vowel(chars[right]) {
                right -= 1;
            }
            if left < right {
                chars.swap(left, right);
                left += 1;
                right -= 1;
            }
        }
        chars.into_iter().collect()
    }

    fn is_vowel(c: char) -> bool {
        matches!(c.to_ascii_lowercase(), 'a' | 'e' | 'i' | 'o' | 'u')
    }
}