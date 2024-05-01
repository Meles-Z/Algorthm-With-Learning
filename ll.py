

class Node:
    def __init__(self, data):
        self.data=data
        self.next=None

class LinkedList:
    def __init__(self):
        self.head=None
    def is_empty(self):
        return self.head is None
    
    def insert_at_beagining(self, data):
        new_node=Node(data)
        if self.is_empty():
            self.head=new_node
        else:
            new_node.next=self.head
            self.head=new_node
    def insert_at_middle_after(self, data, position):
        new_node=Node(data)
        if self.is_empty():
            self.head=new_node
        else:
            curr = self.head
            count = 0 # if you want to change at given position change to count=1
            while curr and count < position:
                curr = curr.next
                count += 1
            if not curr:
                print("Invalid position")
                return
            new_node.next = curr.next
            curr.next = new_node

    def insert_at_middle_before(self, data, position):
        new_node = Node(data)
        if position < 1:
            print("Invalid position")
            return
        if self.is_empty():
            self.head = new_node
        else:
            prev = None
            curr = self.head
            count = 1
            while curr and count < position:
                prev = curr
                curr = curr.next
                count += 1
            if not curr:
                print("Invalid position")
                return
            new_node.next = curr
            if prev:
                prev.next = new_node
            else:
                self.head = new_node

        
    def insert_at_end(self, data):
        new_node = Node(data)
        if self.is_empty():
            self.head = new_node
        else:
            curr = self.head
            while curr.next: 
               curr = curr.next
            curr.next = new_node


    def travers_forward(self):
        curr=self.head
        while curr:
            print(curr.data, end=" ")  
            curr=curr.next
        print()    

ll=LinkedList()
ll.insert_at_beagining(4)
ll.insert_at_beagining(3)
ll.insert_at_beagining(6)
ll.insert_at_end(7)
ll.insert_at_middle_after(20,2)
ll.insert_at_middle_before(40,3)
ll.travers_forward()