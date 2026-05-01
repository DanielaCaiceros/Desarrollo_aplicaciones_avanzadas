#Stack
class stack:
    def __init__(self):
        self.items = []
    def push(self, item):
        self.items.append(item)
    def pop(self):
        return self.items.pop()
    def is_empty(self):
        return len(self.items) == 0

#Tests

#Caso 1: Push elementos
s = stack()
s.push(10)
s.push(20)
if s.items == [10, 20]:
    print("Caso 1: Push elementos - OK")

#Caso 2: Pop elemento
s = stack()
s.push(10)
s.push(20)

if s.pop() == 20 and s.pop() == 10:
    print("Caso 2: Pop elemento - OK")

#Caso 3: Stack vacío después de pop
s = stack()
s.push(5)
s.pop()

if s.is_empty() == True:
    print("Caso 3: Stack vacío después de pop - OK")

#Caso 4: is_empty() en stack nuevo
s = stack()
if s.is_empty() == True:
    print("Caso 4: is_empty() en stack nuevo - OK")

