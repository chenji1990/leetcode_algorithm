from typing import List


class Solution:
    def permutation(self, s: str) -> List[str]:
        array: List[str] = [ x for x in s ]
        array.sort()
        length = len(array)
        temp: List[str] = []
        vis: List[bool] = [False] * length
        res: List[str] = []
        def backtrack(i: int):
            nonlocal temp
            if i == length:
                res.append("".join(temp))
                return
            for j in range(len(vis)):
                if vis[j] or (j > 0 and not vis[j - 1] and array[j - 1] == array[j]):
                    continue
                vis[j] = True
                temp.append(array[j])

                backtrack(i + 1)

                vis[j] = False
                temp = temp[:len(temp) - 1]
        
        backtrack(0)
        return res

