import sys
sys.path.append("..")
from math import sqrt, pi

def calculate_circle_area(radius):
    area = pi * radius**2
    return area

radius = 5.0
circle_area = calculate_circle_area(radius)
print("The area of the circle with radius", radius, "is:", circle_area)
