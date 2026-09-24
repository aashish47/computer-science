fn main() {
    let a = 5;
    let b = "a";
    let c = true;
    let d = 1.5;

    println!("{a} {b} {c} {d}");

    let tup = (1,2,"abc");
    let (x,y,z) = tup;
    
    println!("{:#?}", tup);
    println!("{x} {y} {z}");
    println!("{}",tup.0);

    let arr = [1,2,3,4];

    println!("{}",arr[0]);
    println!("{:#?}",arr);

    let a: [i32; 5] = [1, 2, 3, 4, 5];
    println!("{:?}",a);
    let a = [3; 5];
    println!("{:?}",a);
}
