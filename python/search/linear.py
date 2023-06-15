def linear_search(sequence:list,target:int):
    for index,item in enumerate(sequence):
        if item == target:
            return index
    return -1
    


def recursive_linear_search(sequence:list,low:int,high:int,target:int):
    if not (0<=high<len(sequence) and 0<= low <len(sequence)):
        raise Exception("Invalid upper or lower bound")
    if high < low:
        return -1
    if sequence[low] == target:
        return low
    if sequence[high] == target:
        return high
    return recursive_linear_search(sequence,low+1,high+1,target)