#Hash
class hash_table:
    def __init__(self):
        self.table = {}
    def insert(self, key, value):
        self.table[key] = value
    def delete(self, key):
        del self.table[key]
    def search(self, key):
        return self.table.get(key, None)
#Tests

#Caso 1: insert()
h = hash_table()
h.insert("name", "Daniela")
if h.table["name"] == "Daniela":
    print("Caso 1: insert() - OK")

#Caso 2: search() existente
h = hash_table()
h.insert("age", 25)
if h.search("age") == 25:
    print("Caso 2: search() existente - OK")

#Caso 3: search() no existente
h = hash_table()
if h.search("city") == None:
    print("Caso 3: search() no existente - OK")

#Caso 4: delete()
h = hash_table()
h.insert("age", 25)
h.delete("age")
if h.search("age") == None:
    print("Caso 4: delete() - OK")