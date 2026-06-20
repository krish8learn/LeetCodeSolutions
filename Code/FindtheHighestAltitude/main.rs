fn main() {

}

struct Solution;

impl Solution {
    pub fn find_highest_altitude(gain: Vec<i32>) -> i32{
        let mut current: i32 = 0;
        let mut highest: i32 = 0;
        for g in gain {
            current += g;
            if current > highest {
                highest = current;
            }
        }
        highest
    }
}