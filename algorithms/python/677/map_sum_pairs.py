from typing import List


class Item:
    def __init__(self, key: str, value: int):
        self.key: str = key
        self.value: int = value

    def __repr__(self) -> str:
        return f"<Item: key: {self.key}, value: {self.value}>"


def bisect_left(a: List[Item], x: str) -> int:
    lo, hi = 0, len(a)
    while lo < hi:
        mid = (lo + hi) // 2
        if a[mid].key < x:
            lo = mid + 1
        else:
            hi = mid
    return lo


class MapSum:

    def __init__(self):
        self.items: List[Item] = []

    def search(self, key: str) -> int:
        return bisect_left(self.items, key)

    def insert(self, key: str, val: int) -> None:
        idx = self.search(key)
        if idx == len(self.items):
            self.items.append(Item(key, val))
            return

        if self.items[idx].key == key:
            self.items[idx].value = val
            return

        self.items.append(Item(key, val))
        self.items.sort(key=lambda x: x.key)

    def sum(self, prefix: str) -> int:
        res = 0
        for i in range(self.search(prefix), len(self.items)):
            if not self.items[i].key.startswith(prefix):
                break
            res += self.items[i].value
        return res


if __name__ == "__main__":
    m = MapSum()
    assert m.sum("ap") == 0
    m.insert("apple", 10)
    assert m.sum("bap") == 0
    assert m.sum("ap") == 10
    m.insert("apple", 20)
    assert m.sum("ap") == 20, f"got: {m.sum('ap')}"
    m.insert("bb", 10)
    m.insert("app", 1)
    assert m.sum("ap") == 21
