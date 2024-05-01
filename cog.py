class Node:
    def __init__(self, data):
        self.data=data
        self.next=None
        self.prev=None
class CirularLinkedList:
    def __init__(self):
        self.head=None
        self.tail=None
     
    def is_empty(self):
        return self.head is None
    def insert_at_begining(self, data):
        new_node=Node(data)
        if self.is_empty():
            self.head=self.tail=new_node
        else:
            self.head.prev=new_node
            new_node.next=self.head
            self.head=new_node
    def insert_after(self, data, position):
        new_node = Node(data)
        if self.is_empty():
            self.head = self.tail = new_node
        else:
            curr = self.head
            count = 1
            while curr and count < position:
                curr = curr.next
                count += 1
            if not curr:
                print("Invalid index")
                return
            if curr == self.tail:
                new_node.prev = self.tail
                self.tail.next = new_node
                self.tail = new_node
            else:
                new_node.next = curr.next
                new_node.prev = curr
                curr.next.prev = new_node
                curr.next = new_node



            

    def insert_at_end(self, data):
        new_node=Node(data)
        if self.is_empty():
            self.head=self.tail=new_node
        else:
            self.tail.next=new_node
            new_node.prev=self.tail
            self.tail=new_node
    def delete_at_begining(self):
        if self.is_empty():
            return None
        data=self.head.data
        if self.head==self.tail:
            self.head=self.tail=None
        else:
            self.head=self.head.next
            self.head.prev=None
        return data
            
    def delete_at_end(self):
        
        if self.is_empty():
            return None
        data=self.head.data
        if self.head==self.tail:
            self.head=self.tail=None
        else:
            self.tail=self.tail.prev
            self.tail.next=None
        return data
    def display_forward(self):
        curr=self.head
        while curr:
            print(curr.data, end=" ")
            curr=curr.next
        print()
    def display_backward(self):
        curr=self.tail
        while curr:
            print(curr.data, end=" ")
            curr=curr.prev
        print()
        
cll=CirularLinkedList()
cll.insert_at_begining(1)
cll.insert_at_begining(4)
cll.insert_at_begining(9)
cll.insert_at_end(8)
cll.insert_at_end(5)
cll.insert_after(10,2)
cll.delete_at_end()
cll.delete_at_begining()
cll.display_forward()
cll.display_backward()