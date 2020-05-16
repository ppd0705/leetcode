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

#### recusion
```
def recursion(level, param1, param2, ...): 
    # recursion terminator  递归终止条件
    if level > MAX_LEVEL: 
	   process_result 
	   return 
    # process logic in current level  处理当前层逻辑
    process(level, data...) 
    # drill down  下探到下一层
    self.recursion(level + 1, p1, ...) 
    # reverse the current level status if needed  清理当前层
```