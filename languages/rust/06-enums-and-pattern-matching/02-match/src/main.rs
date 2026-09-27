#[derive(Debug)]
enum UsState{
    Miami,
    LA,
}

enum Coin{
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}

fn main() {
    let coin = Coin::Quarter(UsState::Miami) ;
    let amt = match coin{
        Coin::Penny =>1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(usstate)=>{
            println!("{usstate:?}");
            25
        }
    };

    let dice = 5;
    match dice{
        1 => println!("One!"),
        2 | 3 | 5 | 7 => println!("A small prime number!"),
        10..=20 => println!("Between 10 and 20"),
        _ => println!("Something else"), // Catch-all is required here
    }
}
