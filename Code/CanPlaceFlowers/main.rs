
fn main() {
    let flowerbed: Vec<i32> = vec![1, 0, 0, 0, 1];
    let n =1;
    println!("{}", Solution::can_place_flowers(flowerbed, n))
}

struct  Solution;

impl Solution {
    pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
        let mut mutflowerbed = flowerbed;
        let len = mutflowerbed.len();
        let mut count = 0;

        for i in 0..len {
            if mutflowerbed[i] == 0
                && (i == 0 || mutflowerbed[i - 1] == 0)
                && (i == len - 1 || mutflowerbed[i + 1] == 0)
            {
                mutflowerbed[i] = 1;
                count += 1;
            }
        }

        count >= n
    }
}