/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type item struct {
	Val int
	Ptr *ListNode
}

type hp []item

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(item)) }

func (h *hp) Pop() interface{} { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func mergeKLists(lists []*ListNode) *ListNode {
	h := &hp{}

	for _, p := range lists {
		if p == nil {
			continue
		}
		heap.Push(h, item{
			Val: p.Val,
			Ptr: p,
		})
	}

	head := &ListNode{}
	p := head
	for h.Len() > 0 {
		q := heap.Pop(h).(item).Ptr
		if q.Next != nil {
			heap.Push(h, item{
				Val: q.Next.Val,
				Ptr: q.Next,
			})
		}
		p.Next = q
		p = p.Next
		p.Next = nil
	}

	return head.Next
}