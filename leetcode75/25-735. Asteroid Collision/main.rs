pub fn asteroid_collision(asteroids: Vec<i32>) -> Vec<i32> {
    let mut stack: Vec<i32> = vec![];

    for i in asteroids {
        let mut append_value = false;
        while stack.len() > 0 && (i < 0 && 0 < stack[stack.len() - 1]) {
            let abs_value = i.abs();

            if abs_value == stack[stack.len() - 1] {
                stack.pop();
                append_value = false; // in this condition, both asteroids got destroyed, so no
                                      // append
                break;
            } else if abs_value > stack[stack.len() - 1] {
                // in this case, prev value got destroyed, so we want to append
                append_value = true;
                stack.pop();
            } else {
                append_value = false; // since this value was small, we do not append
                break;
            }
        }
        if append_value {
            stack.push(i);
        }
    }

    return stack;
}

fn main() {
    println!("{:?}", asteroid_collision(vec![10, 2, -5]))
}
