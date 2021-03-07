class Solution {
    public int minimumLengthEncoding(String[] words) {
        TrieNode trie = new TrieNode();
        HashMap<TrieNode, Integer> nodes = new HashMap<>();
        
        for (int i = 0; i < words.length; i++) {
            String word = words[i];
            TrieNode current = trie;
            for (int j = word.length() - 1; j >= 0; j--) {
                current = current.get(word.charAt(j));
            }
            nodes.put(current, i);
        }
        
        int ans = 0;
        for (TrieNode node: nodes.keySet()) {
            if (node.count == 0) {
                ans += words[nodes.get(node)].length() + 1;
            }
        }
        return ans;
    }
}

class TrieNode {
    TrieNode[] children;
    int count;
    
    TrieNode() {
        count = 0;
        children = new TrieNode[26];
    }
    
    public TrieNode get(char c) {
        if (children[c-'a'] == null) {
            children[c-'a'] = new TrieNode();
            count++;
        }
        return children[c-'a'];
    }
}
