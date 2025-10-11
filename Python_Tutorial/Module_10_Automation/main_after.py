names = ["Arpit", "John", "Sara"]

# Write names to file
with open("names.txt", "w") as f:
    for name in names:
        f.write(name + "\n")

# Read and greet
with open("names.txt", "r") as f:
    for line in f:
        print(f"Hello {line.strip()}!")
