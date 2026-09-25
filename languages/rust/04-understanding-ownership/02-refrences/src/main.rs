fn main() {
    let mut s1 = String::from("hello");
    borrows(&s1);
    change(&mut s1);
    println!("{s1}");

    // multiple mutable refrence not allowed in same scope
    // mutable refrence and immutable refrence cannot exist together
    // can have multiple immutable refrences
    // scope starts from where it is defined till last use
    // let s2 = &mut s1;
    // let s3 = &mut s1;
    // println!("{s2} {s3}");

}

fn borrows(_s1:&String){}

fn change(_s1: &mut String){}

// fn dangling() -> &String{
//     let s1 = "asdasd";
//     &s1
// }