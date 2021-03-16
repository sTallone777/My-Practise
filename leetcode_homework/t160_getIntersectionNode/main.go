func getIntersectionNode(headA, headB *ListNode) *ListNode {
    ha := headA
    hb := headB
    for ha != hb {
        if ha == nil{
            ha = headB
        }else{
            ha = ha.Next
        }
        if hb == nil{
            hb = headA
        }else{
            hb = hb.Next
        }
    } 
    return ha
}