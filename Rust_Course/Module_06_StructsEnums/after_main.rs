struct Person {
    name: String,
    age: u8,
}

impl Person {
    fn greet(&self) {
        println!("Hi, I’m {} and I’m {} years old.", self.name, self.age);
    }
}

fn main() {
    let p = Person {
        name: String::from("Arpit"),
        age: 28,
    };
    p.greet();
}
