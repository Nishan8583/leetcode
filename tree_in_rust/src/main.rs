struct node {
    val: i32,
    next: Option<Box<node>>,
}

fn main() {
    println!("Hello, world!");
    let mut head = Box::new(node { val: 1, next: None });
    head.as_mut().next = Some(Box::new(node {
        val: 2,
        next: Some(Box::new(node {
            val: 3,
            next: Some(Box::new(node { val: 4, next: None })),
        })),
    }));

    let head_ptr = Some(head);

    let mut current = head_ptr.as_ref(); // converts Option<Box<node>> to
                                         // Option<&Box<node>>

    while let Some(node) = current {
        println!("{}", node.val);
        current = node.next.as_ref();
    }
}
