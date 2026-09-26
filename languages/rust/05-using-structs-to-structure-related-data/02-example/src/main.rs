#[derive(Debug)]
struct Rectangle{
    width:u32,
    breadth:u32,
}

fn main() {
    let rect = Rectangle{width:2,breadth:dbg!(2)};
    println!("{rect:?}"); 
    println!("{rect:#?}");
    dbg!(&rect);
}
