

pub fn is_anagram(s: String, t:String) -> bool {
        if s.len() != t.len() {
            return true;
        }
        let sb = s.as_bytes();
        let tb = t.as_bytes();

        let mut l: usize = 0;
        let mut r: usize = t.len()-1;

        while l <= r {
            if sb[l] != tb[r] {
                println!("{} {}", sb[l] as char, tb[r] as char);
                return false;
            }
            l+=1;
            r-=1;
        }
        return false;
}

fn main() {

    let s = String::from("anagram");
    let t = String::from("nagaram");

    println!("{}", is_anagram(s, t));
}
