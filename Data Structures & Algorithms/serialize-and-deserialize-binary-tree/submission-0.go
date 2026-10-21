/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    arr := []*int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil{
			arr = append(arr,nil)
			return
		}
		arr = append(arr,&root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	data, err := json.Marshal(arr)
	if err != nil {
		return ""
	}

	return string(data)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
   var arr []*int

	err := json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil
	}
	
	i:=0
	var dfs func()*TreeNode
	dfs = func()*TreeNode{
		if i >= len(arr){
			return nil
		}
		value := arr[i]
		i++
		if value == nil{
			return nil
		}
		root :=  &TreeNode{
			Val: *value,
		}
		
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	return dfs()
}
