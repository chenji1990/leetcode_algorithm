from heapq import *


class MedianFinder:

    def __init__(self):
        self.maxHeap = []
        self.minHeap = []

    def addNum(self, num: int) -> None:
        if len(self.maxHeap) == len(self.minHeap):
            heappush(self.maxHeap, -heappushpop(self.minHeap, -num))
        else:
            heappush(self.minHeap, -heappushpop(self.maxHeap, num))

    def findMedian(self) -> float:
        if len(self.maxHeap) == len(self.minHeap):
            return float(self.maxHeap[0] - self.minHeap[0]) / 2
        return float(self.maxHeap[0])