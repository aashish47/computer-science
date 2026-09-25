fn main() {
    {
        let s1 = "hello";
        println!("{s1}");
    }
    
    let s2 = String::from("Hello World");
    println!("{s2}");

    let s3 = "asda";
    let mut s4 = String::from(s3);

    println!("{s3}");

    // ownership moved
    // let _s5 = s4;
    
    println!("{s4}");

    s4 = String::from("aaaa");
    println!("{s4}");

    
    // takes_ownership(s4);
    // println!("{s4}");

    s4 = takes_and_gives_ownership(s4);

    let (_s4, _len) = length(s4);
    
}

fn _takes_ownership(_s4: String){

}

fn takes_and_gives_ownership(s4: String)->String{
    s4
}

fn length(s4:String)->(String, usize){
    let length = s4.len();
    (s4, length)
}
