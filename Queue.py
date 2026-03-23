
#Queue
class queue:
    def __init__(self):
        self.items = []
    def enqueue(self, item):
        self.items.append(item)
    def dequeue(self):
        return self.items.pop(0)
    def is_empty(self):
        return len(self.items) == 0 

#Tests

#Caso 1: enqueue elementos
q = queue()
q.enqueue(10)
q.enqueue(20)

if q.items == [10, 20]:
    print("Caso 1: Push elementos - OK")

#Caso 2: dequeue elemento
q = queue()
q.enqueue(10)
q.enqueue(20)

if q.dequeue() == 10:
    print("Caso 2: Dequeue elemento - OK")

#Caso 3: múltiples dequeue
q = queue()
q.enqueue(10)
q.enqueue(20)

q.dequeue()
if q.dequeue() == 20:
    print("Caso 3: Múltiples dequeue - OK")

#Caso 4: queue vacía después de dequeue
q = queue()
q.enqueue(5)
q.dequeue()

if q.is_empty() == True:
    print("Caso 4: Queue vacía después de dequeue - OK")

