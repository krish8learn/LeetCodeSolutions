struct RecentCounter {
    calls: Vec<i32>
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl RecentCounter {

    fn new() -> Self {
        RecentCounter { calls: Vec::new() }
    }
    
    fn ping(&mut self, t: i32) -> i32 {
        self.calls.push(t);
        while !self.calls.is_empty() && self.calls[0] < t - 3000 {
            self.calls.remove(0);
        }
        self.calls.len() as i32
    }
}

fn main() {
    let mut recent_counter = RecentCounter::new();
    println!("{}", recent_counter.ping(1));     // returns 1
    println!("{}", recent_counter.ping(100));   // returns 2
    println!("{}", recent_counter.ping(3001));  // returns 3
    println!("{}", recent_counter.ping(3002));  // returns 3
}