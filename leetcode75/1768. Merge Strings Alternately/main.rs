use std::collections::VecDeque;

pub fn merge_alternately(word1: String, word2: String) -> String {
    let mut word1Vec: VecDeque<String> = VecDeque::new();
    let mut word2Vec: VecDeque<String> = VecDeque::new();

    for w in word1.chars() {
        word1Vec.push_back(w.to_string());
    }

    for w in word2.chars() {
        word2Vec.push_back(w.to_string());
    }

    let mut finalVal: String = String::new();
    while word1Vec.len() > 0 || word2Vec.len() > 0 {
        if word1Vec.len() > 0 {
            let v = VecDeque::pop_front(&mut word1Vec).unwrap();
            finalVal = finalVal + &v;
        }

        if word2Vec.len() > 0 {
            let v = VecDeque::pop_front(&mut word2Vec).unwrap();
            finalVal = finalVal + &v;
        }
    }

    return finalVal;
}

pub fn optimized_sol(word1: String, word2: String) -> String {
    let mut merged = String::new();
    let mut w1_iter = word1.chars();
    let mut w2_iter = word2.chars();

    loop {
        if let Some(c1) = w1_iter.next() {
            merged.push(c1);
        }
        if let Some(c2) = w2_iter.next() {
            merged.push(c2);
        }

        if w1_iter.as_str().is_empty() && w2_iter.as_str().is_empty() {
            break;
        }
    }
    merged.extend(w1_iter);
    merged.extend(w2_iter);

    merged
}
fn main() {
    let result = merge_alternately(String::from("abc"), String::from("pqr"));
    println!("Value= {}", result);
}
