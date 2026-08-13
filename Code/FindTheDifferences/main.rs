fn main() {
    let s = "abcd";
    let t = "abcde";
    println!("{}", find_the_difference(s, t));
}

fn find_the_difference(s: &str, t: &str) -> char {
    let mut x: u8 = 0;
    for b in s.bytes() {
        x ^= b;
    }
    for b in t.bytes() {
        x ^= b;
    }
    x as char
}