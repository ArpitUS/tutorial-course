fn main() {
    let s1 = String::from("Rust");
    let s2 = &s1;

    println!("Borrowed reference: {}", s2);
    println!("Original still accessible: {}", s1);
}
