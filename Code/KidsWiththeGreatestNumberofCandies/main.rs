fn main() {
    let candies: Vec<i32> = vec![1, 0, 0, 0, 1];
    let extra_candies = 1;

    // Use {:?} to print Vec<bool>
    println!("{:?}", Solution::kids_with_candies(candies, extra_candies));
}
struct Solution;

impl Solution {
    pub fn kids_with_candies(candies: Vec<i32>, extra_candies: i32) -> Vec<bool> {
        let max_value = *candies.iter().max().unwrap();
        // let mut result: Vec<bool> = Vec::new();

        // for &num in &candies {
        //     let new_val = num + extra_candies;
        //     if new_val >= max_value {
        //         result.push(true);   // condition met
        //     } else {
        //         result.push(false);  // condition not met
        //     }
        // }

        // result

        candies.iter().map(|&candy| candy + extra_candies >= max_value).collect()
    }
}
