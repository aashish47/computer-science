use std::collections::HashMap;

fn main() {
    let mut scores = HashMap::new();
    let key = String::from("Blue");
    scores.insert(&key, 10);
    scores.get(&key).copied().unwrap_or(0);

    for (key, value) in &scores{
        println!("{key}: {value}");
    }

    scores.entry(&String::from("yellow")).or_insert(50);
}
