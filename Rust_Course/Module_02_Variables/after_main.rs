fn main() {
    let name = "Arpit";
    let mut age = 28;
    const HEIGHT: f64 = 1.75;

    println!("Name: {}", name);
    println!("Age: {}", age);
    println!("Height: {} m", HEIGHT);

    age += 1;
    println!("Next year age: {}", age);
}
