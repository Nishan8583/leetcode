class Solution:
    def islandsAndTreasure(self, grid: list[list[int]]) -> None:
        memo = [[-2 for _ in range(len(grid[0]))] for _ in range(len(grid))]

        for row_index, row in enumerate(grid):
            for col_index,val in enumerate(row):
                INF=2**32-1
                visited = set()
                #print(row_index,col_index)
                v = self.get_reach_treasure(row_index,col_index,grid,1,INF,visited,memo)
                if v != INF:
                    grid[row_index][col_index] =v
                    memo[row_index][col_index] = v
                else:
                    grid[row_index][col_index] = -1
                    


                


    def get_reach_treasure(self,row,col,grid,sum,inf,visited,memo):
        #print(row,col)
        # 1st base condition
        # check if row and col is out of grid, if out of grid, traverse no more
        if row >= len(grid) or row < 0 or col >= len(grid[0]) or col < 0:
            #print("Out of bound")
            return inf
        
        if (row,col) in visited:
            return inf

        if memo[row][col] != -2:
            print("Value was cached",row,col,memo[row][col])
            return memo[row][col]

        # 2nd base condition, water can not be traversed, traverse no more
        if grid[row][col] == -1:
            #print("Found water")
            return inf
        # if we reached the treasure, return 0, since it takes 0 to get to itself
        if grid[row][col] == 0:
            return 0

        print("New",row,col)
        # prevent recursion
        visited.add((row,col))
        
        # Cheeck if we have treasure in the adjacent grids
        up = row -1 
        left = col - 1 
        down = row + 1 
        right = col +1

        # traverse up
        #print("Traversing up")
        up_sum = self.get_reach_treasure(up,col,grid,sum,inf,visited,memo)
        # traverse left
        #print("Traversing left")
        left_sum = self.get_reach_treasure(row,left,grid,sum,inf,visited,memo)
        # traverse below
        #print("Traversgin down")
        down_sum = self.get_reach_treasure(down,col,grid,sum,inf,visited,memo)
        # traverse right
        #print("right")
        right_sum = self.get_reach_treasure(row,right,grid,sum,inf,visited,memo)

        visited.remove((row,col))
        return min(up_sum,left_sum,down_sum,right_sum) + 1


s = Solution()
print("---TEST CASE 1----------")
testCase = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
s.islandsAndTreasure(testCase)

print("---TEST CASE 2----------")
testCasse2=[[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
s.islandsAndTreasure(testCasse2)
print(testCase)
