fn main() {
    let config = Some(4);
    match config{
        Some(x) =>{},
        None =>{},
    }

    // can be written as
    if let Some(x) =  config{}
    else{}

    let Some(x) = config else{
        return
    };
}
