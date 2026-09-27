#[derive(Debug)]
enum Message{
    Hello,
    Bye,
    End,
}

enum Shape{
    Rectangle(u32, u32),
}

impl Message{
    fn call(&self){
        println!("{self:?}")
    }
}

fn main() {
    let hello = Message::Hello;
    let _bye = Message::Bye;
    let _end = Message::End;
    hello.call();
    

    let rectangle = Shape::Rectangle(5, 4);
    match rectangle {
        Shape::Rectangle(width, height) => {
            print!("{0} {1}", width, height);
        }
    }

    let num1 = Some(5);
    let num2: Option<u32> = None;
    println!("{num1:?} {num2:?}");
}
