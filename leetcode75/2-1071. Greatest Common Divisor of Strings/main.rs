use std::cmp::min;
impl Solution {
    pub fn gcd_of_strings(str1: String, str2: String) -> String {
        let mut result = String::new();
        let c = str1.clone();
        let p = str2.clone();
        for i in 1..min(str1.len(), str2.len()) + 1 {
            let b = c.clone();
            let q = p.clone();
            let canditate = &c[..i];
            if Self::divides(canditate.to_string(), b) && Self::divides(canditate.to_string(), q) {
                result = canditate.to_string();
            }
        }
        return result;
    }

    pub fn divides(s: String, t: String) -> bool {
        if t.len() % s.len() != 0 {
            return false;
        }

        let product = t.len() / s.len();
        return s.repeat(product) == t;
    }
}

