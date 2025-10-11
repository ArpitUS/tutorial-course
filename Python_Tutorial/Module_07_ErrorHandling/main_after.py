try:
    number = int(input("Enter a number: "))
    print(f"The square is {number ** 2}")
except ValueError:
    print("That’s not a valid number!")
finally:
    print("Program finished.")
