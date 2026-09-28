fn main() {
    let vec1:Vec<u32> = Vec::new(); 
    let vec2 = vec![1,2,3];

    let num1 = &vec2[2];
    println!("{num1}");
    let num:Option<&u32> = vec2.get(1);
    println!("{num:?}");

    for i in &vec2{
        println!("{i}")
    }

    enum SpreadSheetCell{
        Int(i32),
        Float(f64),
        Text(String),
    }

    let row = vec![SpreadSheetCell::Int(3),
    SpreadSheetCell::Float(10.12),
    SpreadSheetCell::Text(String::from("blue")),
    ];
}
