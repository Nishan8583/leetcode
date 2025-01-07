'''
MaxHeap to store the character with highest count on top, because doing that first will have least leastInterval
We pop from max heap, pop it, and it in queue with current time + n because we want to check at what time it will be avaiable
Keep on looping till either maxHeap or queue has value

'''


class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:

        count = Counter(tasks)  # creates a hashmap of characters and count

        # Python only allows min Heap, so to create min heap we createt negative values
        maxHeap = [ -cnt for cnt in count.values() ]
        heapq.heapify(maxHeap)

        time = 0
        q = deque()
        while maxHeap or q:
            time += 1
            if maxHeap:
                cnt = 1 + heapq.heappop(maxHeap)
                if cnt:
                    q.append([cnt,time+n])
            
            if q and q[0][1] == time:
                heapq.heappush(maxHeap,q.popleft()[0])
        
        return time

