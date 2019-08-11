# -*- coding: utf-8 -*-

def person(name , age , **kw):
    if 'city' in kw:
        #有city参数
        pass
    if 'job' in kw:
        #有job参数
        pass
    print('name:', name , 'age:', age , 'other:', kw )

person('Jack', 24, city='Beijing', addr='Chaoyang', zipcode=123456)
#用户依然可以无限制的传入参数

#所以应该定义如下

def person(name, age, *, city, job): #使用*号作为分隔符 后面的参数都属于命名参数
    print(name, age, city, job)

person('Jack', 24, 'Beijing', 'Engineer')

'''
命名关键字参数必须传入参数名，这和位置参数不同。如果没有传入参数名，调用将报错：
File "D:/python项目自学/My-learning/命名关键参数.py", line 20, in <module>
    person('Jack', 24, 'Beijing', 'Engineer')
TypeError: person() takes 2 positional arguments but 4 were given
'''

def person(name, age, *, city='beijing', job='Engineer'): #使用*号作为分隔符 后面的参数都属于命名参数
    print(name, age, city, job)

person('Jack', 24)