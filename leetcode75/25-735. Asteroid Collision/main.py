"""
We are given an array asteroids of integers representing asteroids in a row. The indices of the asteriod in the array represent their relative position in space.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.



Example 1:

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

Example 2:

Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.

Example 3:

Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

Ideas:
1. Minus means moving left, + means moving right
+ve stack                   -ve stack
[                           [
    10                      -5
    5
]                           ]
2. Loop till eiehter one of them is empty
3. We pop from both, if we find a value, see which one is higher, and then inser the higher value into the result, ignore the lower value
4, We ignore abs(-5), we put 10, result = 10
5. No value in -ve, just append all thats left in +ve stack


Another example
[10,2,-5]
+ve stack    -ve stack
2           -5
10

2 and -5, 2 and abs(-5), 2 and 5, 5 is greater, so -5 stays

10 and -5
10 greater, 10 stays
final result = [10]


[-2,-1,1,2]

2       -1
1       -2
"""

"""

class Solution:
    def asteroidCollision(self, asteroids: list[int]) -> list[int]:
        positive_stack = []
        negative_stack = []

        for i in asteroids:
            if i >= 0:
                positive_stack.append(i)
            else:
                negative_stack.append(i)

        result = []
        while len(positive_stack) > 0 and len(negative_stack) > 0:
            top_positive = positive_stack[len(positive_stack) - 1]
            top_negative = negative_stack[len(negative_stack) - 1]

            neg = abs(top_negative)

            if neg == top_positive:
                negative_stack.pop(-1)
                positive_stack.pop(-1)
            elif neg < top_positive:
                negative_stack.pop(-1)
            elif top_positive > neg:
                positive_stack.pop(-1)
            else:
                print("non match")
                result.append(negative_stack.pop(-1))
                result.append(positive_stack.pop(-1))

        result = result + positive_stack + negative_stack
        return result

"""


class Solution:
    def asteroidCollision(self, asteroids: list[int]) -> list[int]:
        stack = []
        for a in asteroids:
            # print(a)
            # loop while there is some value in stack, and current value is less than 0 and previous value is greater than 0, because asteroid will collide this case
            # if current value is negative and previous value is positive we collide, else just append in stack
            while stack and a < 0 < stack[-1]:
                # print("condition matched for ", a, stack[-1])
                abs_last = abs(a)  # get absolute value
                # print(f"Absolute value is {abs_last}")
                if (
                    abs_last == stack[-1]
                ):  # if current value is equals to previous value, then both get destroyed, do not append current value, and pop current value, and break out of this whjile loop
                    # print(f"equals {stack[-1]} and {a}")
                    stack.pop()
                    break
                elif (
                    abs_last > stack[-1]
                ):  # if current value is greater then, previoois value is destroyed, do not append here, keep going to top of loop, becasue we have to check with next value,
                    stack.pop()  # the append will be handleded in the else statement of while loop
                else:  # this means a is less than previous asteroid so this gets destroyed
                    break
            else:  # so if cuurrent asterriod is not negative, or previous asteroid abnd current movinbg in same direction append
                # print("appending ", a)
                stack.append(a)

        return stack


s = Solution()
# print(s.asteroidCollision([5, 10, -5]))
# print(s.asteroidCollision([8, -8]))
print(s.asteroidCollision([10, 2, -5]))
# print(s.asteroidCollision([-2, -1, 1, 2]))
# print(s.asteroidCollision([-2, -2, -2, 1]))
# print(s.asteroidCollision([-2, 1, -2, -2]))
# print(s.asteroidCollision([-2, 2, -1, -2]))
