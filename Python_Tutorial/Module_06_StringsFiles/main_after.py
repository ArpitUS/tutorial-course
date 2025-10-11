name = "Arpit"
print(name[:3])

with open("data.txt", "w") as f:
    f.write("Learning Python file handling!\n")

with open("data.txt", "r") as f:
    content = f.read()
    print("File content:", content)
