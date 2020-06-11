/**You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
**/

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  sum:=0
  carry:=0
  type list struct{
      head *ListNode
      tail *ListNode
  }

  var x,y int
  p:=l1
  q:=l2
  l:= &list{} //Creating an empty list
  //To iterate through elements of the the list
  for
  {
      if (p==nil&&q==nil){break}
      if (p!=nil){x=p.Val}else{x=0}
      if (q!=nil){y=q.Val}else{y=0}

      sum=(x + y+ carry)
      carry=(sum)/10

      n:=&ListNode{
          Val: sum%10,
      }
      if l.head==nil{
          l.head=n
      }else{
          l.tail.Next=n
      }
      l.tail=n

      if p!=nil {p=p.Next}
      if q!=nil {q=q.Next}
    }
    if carry > 0{
      n:=&ListNode{
        Val: carry,
      }
      l.tail.Next=n
      l.tail=n
    }
    return l.head
}
