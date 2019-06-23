

def get_func(val):
    a = 1
    def next():
        nonlocal a
        res = a * val
        a += 1
        return res
    return next

tmp = get_func(5)

for cell in tmp.__closure__:
    print(cell.cell_contents, end=" ")
print()
print(tmp())
for cell in tmp.__closure__:
    print(cell.cell_contents, end=" ")
print()
print(tmp())
for cell in tmp.__closure__:
    print(cell.cell_contents, end=" ")
print()
print(tmp())