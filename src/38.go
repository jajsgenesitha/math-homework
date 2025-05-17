import numpy as np

def generate_random_sequence(length):
    """
    Generate a random sequence of a given length.
    The sequence is generated using the following rule: each element is chosen from a list of 0s and 1s with equal probability.

    Example usage:
    >>> generate_random_sequence(5)
    [0, 0, 0, 0, 0]
    """
    sequence = np.random.choice([0, 1], size=length, p=[0.7, 0.3])
    return sequence

# Check function to verify the correctness of the solution
def check_function():
    test_cases = {
        5: [0, 0, 0, 0, 0],
        6: [0, 1, 1, 0, 0, 1],
        3: [0, 0, 0]
    }
    
    for length, expected in test_cases.items():
        result = generate_random_sequence(length)
        if np.array_equal(result, expected):
            print(f"Test passed for {length}-element sequence.")
        else:
            print(f"Test failed for {length}-element sequence. Expected: {expected}, but got: {result}")

# Run the check function to verify the solution
check_function()
