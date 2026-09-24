fn add (a: i32, b:i32)-> i32{
    a + b
}

fn main() {
    let x = 2;
    let y = 3;

    println!("{}",add(x, y));

    let z = {
        x + y
    };

    println!("{z}");
}
