
def containsDuplicate(nums: list[int]) -> bool: 
    hashmap: dict[int, bool] = {}
    for value in nums:
        if hashmap.get(value, False):
            return True
        hashmap.setdefault(value, True)
    return False
