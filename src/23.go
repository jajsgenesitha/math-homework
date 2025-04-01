import random

def make_random_array(length):
    """
    Generate an array of length 'length' containing random integers between -10 and 10.
    """
    return [random.randint(-10, 10) for _ in range(length)]

def main():
    # Example usage:
    num_elements = int(input("Enter the number of elements you want: "))
    arr = make_random_array(num_elements)
    
    print("\nGenerated Array:")
    for i in range(len(arr)):
        print(f"Element {i+1}: {arr[i]}")

if __name__ == "__main__":
    main()
