#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub fn delete_middle(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    // if head itself is none or head.next is none, return
    if head.is_none() || head.as_ref()?.next.is_none() {
        return head;
    }

    let mut slow = &mut head.as_ref(); // as_ref takes a refeerence
                                       // to inner value without taking ownership of it
    let mut fast = &mut head.as_ref().unwrap().next;

    let mut prev = None;

    while let Some(f) = fast {
        fast = f.next.as_ref().and_then(|n| n.next.as_ref());
        prev = slow.as_mut();
        slow = &mut slow.as_mut().unwrap().next;
    }

    if let Some(prev_node) = prev {
        prev_node.next = prev_node.next.as_mut().unwrap().next.take();
    }

    head
}

fn main() {}
