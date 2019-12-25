class MyCircularDeque:

    def __init__(self, k: int):
        """
        Initialize your data structure here. Set the size of the deque to be k.
        """
        self._max_size = k + 1
        self._front = self._rear = 0

        self._arr = [0] * self._max_size

    def insertFront(self, value: int) -> bool:
        """
        Adds an item at the front of Deque. Return true if the operation is successful.
        """
        if self.isFull():
            return False
        if self.isEmpty():
            self._arr[self._front] = value
            self._rear = (self._rear + 1) % self._max_size
        else:
            self._front = (self._front - 1) % self._max_size
            self._arr[self._front] = value
        return True

    def insertLast(self, value: int) -> bool:
        """
        Adds an item at the rear of Deque. Return true if the operation is successful.
        """
        if self.isFull():
            return False
        self._arr[self._rear] = value
        self._rear = (self._rear + 1) % self._max_size
        return True

    def deleteFront(self) -> bool:
        """
        Deletes an item from the front of Deque. Return true if the operation is successful.
        """
        if self.isEmpty():
            return False
        self._front = (self._front + 1) % self._max_size
        return True

    def deleteLast(self) -> bool:
        """
        Deletes an item from the rear of Deque. Return true if the operation is successful.
        """
        if self.isEmpty():
            return False
        self._rear = (self._rear - 1) % self._max_size
        return True

    def getFront(self) -> int:
        """
        Get the front item from the deque.
        """
        if self.isEmpty():
            return -1
        return self._arr[self._front]

    def getRear(self) -> int:
        """
        Get the last item from the deque.
        """
        if self.isEmpty():
            return -1
        return self._arr[self._rear - 1]

    def isEmpty(self) -> bool:
        """
        Checks whether the circular deque is empty or not.
        """
        return self._front == self._rear

    def isFull(self) -> bool:
        """
        Checks whether the circular deque is full or not.
        """
        return (self._rear + 1) % self._max_size == self._front
