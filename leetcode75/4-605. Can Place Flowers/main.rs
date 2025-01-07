impl Solution {
    pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
        let mut count = 0;
        let mut flower = flowerbed.clone();

        for (i, _v) in flowerbed.iter().enumerate() {
            if flower[i] == 1 {
                // can not use the original flowerbed, becasue we are
                // changing flower not flowered to fix borowchecker complain
                continue;
            }

            let empty_left = (i == 0) || (flower[i - 1] == 0);
            let empty_right = (i == (flower.len() - 1)) || (flower[i + 1] == 0);

            if empty_right && empty_left {
                flower[i] = 1; // had to do this beacuse we can not change the one used in iteration
                               // because mutable and immutable reference is not allowed
                count += 1;
            }
        }
        return count >= n;
    }
}

