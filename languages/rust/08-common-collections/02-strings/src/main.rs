fn main() {
    let mut s = String::new();
    let data = "example";
    s = data.to_string();
    s = "example".to_string();
    s = String::from("example");
    s.push_str("foo");
    s.push('1');

    let s1 = String::from("hello");
    let s2 = String::from("world");
    let s3 = s1 + &s2;

    s = format!("{s3}-{s2}-{s3}");

    for c in "hello".chars(){}
    for c in "hello".bytes(){}
}
