use std::env;
use std::time::{SystemTime, UNIX_EPOCH};

fn main() {
    let start = SystemTime::now();
    let since_the_epoch = start
        .duration_since(UNIX_EPOCH)
        .expect("Time went backwards");

    let name = env::var("INPUT_WHO_TO_GREET").expect("Missing env var: INPUT_WHO_TO_GREET");

    println!("Hello, {}!", name);
    println!("::set-output name=time::{:?}", since_the_epoch);
}
