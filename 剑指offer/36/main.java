package main

// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 36. 二叉搜索树与双向链表


// 二叉搜索树：右大左小
// 链表从小到大
// 搜索树，左中右，中序遍历，从左开始
class Solution {
    private Node head, pre;
    public Node treeToDoublyList(Node root) {
        if (root == null) return root;
        dfs(root);
        head.left = pre;
        pre.right = head;
        return head;
    }

    private void dfs(Node root) {
        if (root == null) return ;
        if (root.left != null) dfs(root.left);
        if (pre != null) {
            root.left = pre;
            pre.right = root;
        } else head = root;
        pre = root;
        if (root.right != null) dfs(root.right);
    }
}

class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};