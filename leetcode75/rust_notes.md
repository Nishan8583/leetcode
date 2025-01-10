# Rust notes that might help in leetcode

1. Enumerate string

```rust
for c in val.chars() {
  c.to_string();
}
```

2. Queue implementation

```rust
use std::collections::VecDeque;

let mut word1Vec: VecDeque<String> = VecDeque::new();
word1Vec.push_back(w.to_string()); // push_back
 let v = VecDeque::pop_front(&mut word1Vec).unwrap();

```

3. Remove whitespace from front and back, reverse string and loop

```rust
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


```

4. Iterate in reverse

```rust
for (i,v) in vector.iter().enumerate().rev() {}
```
