# -*- coding: utf-8 -*-

def my_abs(x):
    if x >= 0:
        return x
    else:
        return -x

print (my_abs(-99))


import math

print(math.sqrt(2))

def quadratic(a,b,c):
    y = b*b-4*a*c
    if y >= 0:
        temp = math.sqrt(y)
        x1 = (temp-b)/(2*a)
        x2 = -((temp+b)/(2*a))
        return x1,x2
    else:
        return
print(quadratic(2,3,1))
