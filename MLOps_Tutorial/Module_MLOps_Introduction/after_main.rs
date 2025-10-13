use std::{thread, time::Duration};

fn main() {
    let steps = ["Training", "Logging", "Deploying", "Monitoring"];
    for step in steps {
        println!("{} step in progress...", step);
        thread::sleep(Duration::from_secs(1));
    }
    println!("✅ MLOps pipeline completed successfully!");
}
