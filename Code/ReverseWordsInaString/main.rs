fn main() {
    println!("{}", reverse_words("the sky is blue".to_string()));
}

fn reverse_words(s: String) -> String {
    let mut words: Vec<&str> = s.split_whitespace().collect();
    words.reverse();
    words.join(" ")
}