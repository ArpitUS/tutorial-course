use std::fs::File;

fn main() {
    match File::open("data.txt") {
        Ok(_) => println!("File opened successfully!"),
        Err(e) => println!("Error opening file: {}", e),
    }
}
