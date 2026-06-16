class UnionFind:
    def __init__(self, n):
        self.parent = [i for i in range(n)]

    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])

        return self.parent[x]

    def union(self, x, y):
        x_root = self.find(x)
        y_root = self.find(y)

        if x_root == y_root:
            return False

        self.parent[x_root] = y_root
        return False


def main():
    uf = UnionFind(7)
    uf.union(2, 0)
    uf.union(4, 0)
    uf.union(3, 1)
    uf.union(5, 3)
    uf.union(6, 4)
    print(uf.parent)
    print(uf.find(5))
    print(uf.find(6))


if __name__ == "__main__":
    main()
