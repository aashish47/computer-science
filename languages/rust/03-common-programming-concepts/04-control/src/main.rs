fn main() {
    // infinite loop
    // loop{
        
    // }

    let mut counter = 0;
    let y = loop {
        counter += 1;
        if counter == 10 {
            break counter
        }
    };

    println!("{y}");

    // loop labels
    'loop1 : loop{
        loop{
            break 'loop1
        }
    }

    while counter < 12{
        counter += 1;
    }

    println!("{counter}");

    for elem in [1,2,3,4]{
        println!("{elem}");
    }

    for num in 1..4{
        println!("{num}");
    }

    // invalid
    // if counter{}

    if counter == 10{
        println!("{counter}")
    }else{
        println!("{counter}")
    }
}
