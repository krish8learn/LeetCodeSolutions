use std::collections::HashSet;
struct Solution; 

impl Solution {
    pub fn find_difference(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<Vec<i32>> {
        let set1: HashSet<i32> = nums1.into_iter().collect();
        let set2: HashSet<i32> = nums2.into_iter().collect();

        let only1: Vec<i32> = set1.difference(&set2).cloned().collect();
        let only2: Vec<i32> = set2.difference(&set1).cloned().collect();
        vec![only1, only2]
    }
}

fn main() {
    let result = Solution::find_difference(vec![1, 2, 3], vec![2, 4, 6]);
    println!("{:?}", result);

    let result2 = Solution::find_difference(vec![1, 2, 3, 3], vec![1, 1, 2, 2]);
    println!("{:?}", result2);
}