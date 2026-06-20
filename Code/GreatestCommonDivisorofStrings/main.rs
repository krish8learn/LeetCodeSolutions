fn main() {
    println!("{}", gcd_of_strings("ABCABC".to_string(), "ABC".to_string()));
}

fn gcd_of_strings(str1: String, str2: String) -> String {
    // Check compatibility: str1+str2 must equal str2+str1
    if str1.clone() + &str2 != str2.clone() + &str1 {
        return "".to_string();
    }

    // Find gcd of lengths
    let divisor = gcd(str1.len(), str2.len());

    // Return substring of str1 up to gcd length
    str1[..divisor].to_string()
}

fn gcd(a: usize, b: usize) -> usize {
    if b == 0 {
        return a;
    }
    gcd(b, a % b)
}
