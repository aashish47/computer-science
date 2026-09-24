fn main() {
    //default immutable
    let x = 5;
    println!("{x}");

    //mutable
    let mut y = 6;
    println!("{y}");

    y = 7;
    println!("{y}");

    //constant
    const Z:i32 = 10;
    println!("{Z}");

    //shadowing
    let x = 10;
    println!("{x}");

    let spaces = "     ";
    let spaces = spaces.len();
    println!("total spaces: {spaces}");
}
