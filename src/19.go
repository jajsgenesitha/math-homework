def find_max_product(*args):
    max_num = max(args)
    product = 1

    for num in args:
        if num <= 0:
            continue
        
        current_product = num * max_num
        max_num = max(current_product, max(num, abs(max_num) // num))
        product *= current_product

    return product

max_product = find_max_product(5, -2, 3, -1)
print("The maximum product is:", max_product)
