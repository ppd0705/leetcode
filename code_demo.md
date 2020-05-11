代码片段收集
---

#### 前序遍历
```python
def preorder(self, root):
    if root:
        self.traverse_path.append(self.root.val)
        self.preorder(root.left)
        self.preorder(root.right)
```
#### 中序遍历
```python
def inorder(self, root):
    if root:
        self.inorder(root.left)
        self.traverse_path.append(self.root.val)
        self.inorder(root.right)
```

#### 后序遍历
```python
def postorder(self, root):
    if root:
        self.postorder(root.left)
        self.postorder(root.right)
        self.traverse_path.append(self.root.val)

```