pub fn reverse_words(s: String) -> String {
    let mut r = s.trim_start().trim_end().to_string();
    let mut vals = r.split_whitespace().rev();
    let mut result = String::new();

    for v in vals {
        if v.is_empty() {
            continue;
        }
        result = result + v + " ";
    }
    return result.trim_end().to_string();
}

fn main() {
    println!("{}", reverse_words(String::from("a good   example")));
}
