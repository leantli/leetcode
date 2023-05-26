package main;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // 不管咋样，marshal 这步得用层序遍历
    public String serialize(TreeNode root) {
        if (root == null)
            return "[]";
        StringBuilder bs = new StringBuilder("[");
        LinkedList<TreeNode> queue = new LinkedList<TreeNode>();
        queue.offer(root);
        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();
            if (node == null)
                bs.append("null,");
            else {
                queue.offer(node.left);
                queue.offer(node.right);
                bs.append(node.val).append(",");
            }
        }
        bs.deleteCharAt(bs.length() - 1);
        System.out.println(bs.toString());
        return bs.toString() + "]";
    }

    // 解析可以继续用层序遍历，不能使用 map，map 和层序遍历得到的数组还是不符合的
    public TreeNode deserialize(String data) {
        if (data.equals("[]"))
            return null;
        String[] ss = data.substring(1, data.length() - 1).split(",");
        System.out.println(Arrays.toString(ss));
        // 再用 hashmap 搞一下
        HashMap<Integer, TreeNode> map = new HashMap<>();
        for (int i = 0; i < ss.length; i++) {
            if (!ss[i].equals("null")) {
                System.out.println("i=" + i + " ss[i]=" + ss[i]);
                map.put(i, new TreeNode(Integer.parseInt(ss[i])));
            }
        }
        System.out.println(ss.length);
        for (int i = 0; i < ss.length; i++) {
            TreeNode node = map.get(i);
            if (node != null) {
                node.left = map.get(i * 2 + 1);
                node.right = map.get(i * 2 + 2);
            }
        }
        System.out.println(map);
        return map.get(0);

        // 层序遍历
        // TreeNode root = new TreeNode(Integer.parseInt(ss[0]));
        // LinkedList<TreeNode> queue = new LinkedList<TreeNode>();
        // queue.offer(root);
        // int i = 1;
        // while (!queue.isEmpty()) {
        // TreeNode node = queue.poll();
        // if (!ss[i].equals("null")) {
        // node.left = new TreeNode(Integer.parseInt(ss[i]));
        // queue.add(node.left);
        // }
        // i++;
        // if(!ss[i].equals("null")) {
        // node.right = new TreeNode(Integer.parseInt(ss[i]));
        // queue.add(node.right);
        // }
        // i++;
        // }
        // return root;
    }

    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));