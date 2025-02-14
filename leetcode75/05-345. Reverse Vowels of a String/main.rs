use std::{collections::HashSet, hash::Hash, result};

    pub fn reverse_vowels(s: String) -> String {
        //let vowel_set = HashSet::new()
        // vowel_set.insert("a")
        let vower_set: HashSet<char> = HashSet::from(['a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U']);
        let mut stack = vec![];
        for c in s.chars() {
            if vower_set.contains(&c) {
                stack.push(c);
            }
        }

        let mut result = String::new();
        for c in s.chars(){
            if vower_set.contains(&c) {
                result.push(stack.pop().unwrap());
            } else {
                result.push(c);
            }
        }
        return result;
    }


fn main() {
    println!("Reverse {}",reverse_vowels(String::from("IceCreAm")))
}
