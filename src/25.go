def fibonacci(n):
    if n <= 1:
        return n
    else:
        return fibonacci(n-1) + fibonacci(n-2)

# Example usage
print(fibonacci(8)) # Output: 21

def prime_numbers(limit):
    primes = []
    for num in range(2, limit + 1):
        if all(num % i != 0 for i in range(2, int(num**0.5) + 1)):
            primes.append(num)
    return primes

# Example usage
print(prime_numbers(10)) # Output: [2, 3, 5, 7]
