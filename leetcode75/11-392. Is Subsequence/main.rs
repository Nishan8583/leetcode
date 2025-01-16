pub fn is_subsequence(s: String, t: String) -> bool {
    if s.len() == 0 {
        return true;
    }
    if t.len() < s.len() {
        return false;
    }
    let s_clone: Vec<char> = s.chars().collect();
    let mut left = 0;
    for char in t.chars() {
        if s_clone[left] == char {
            left += 1;
        }
        if left >= s.len() {
            return true;
        }
    }

    return left >= s.len();
}

fn main() {
    println!(
        "{}",
        is_subsequence(String::from("abc"), String::from("ahbgdc"))
    );
    println!(
        "{}",
        is_subsequence("axc".to_string(), "ahbgdc".to_string())
    );
}
