# Python Basic

## IF

```python
# Python 当然也支持 else 语句， 语法如下：
if expression:
	if_suite
ELSE:
	else_suite

# Python 还支持 elif （意指 “else-if ”）语句，语法如下:
if expression1:
	if_suite
ELIF EXPRESSION2:
	elif_suite
else:
	else_suite
```



## Loops

**While**

```python
while counter < 3:
		print 'loop #%d' % (counter)
```



**for 循环和range()内建函数** 

range([start,]stop[,step])

range(start, end, step =1)
range(end)
range(start, end)

```python
for item in ['e-mail', 'net-surfing', 'homework','chat']:
	print item

squared = [x ** 2 for x in range(4)]
for i in squared:
	print i
```

> - break语句 只能用于循环体内。其效果是直接结束并退出**当前**循环，剩下的未循环的工作全部被忽略和取消。注意当前两个字，Python的break只能退出一层循环，对于多层嵌套循环，不能全部退出。
>
> - continue语句 与break不同，continue语句用于跳过当前循环的剩余部分代码，直接开始下一轮循环。它不会退出和终止循环，只是提前结束当前轮次的循环。同样的，continue语句只能用在循环内。



## 列表和元组

### 通用序列操作

| 运算                   | 结果：                                                       | 备注   |
| :--------------------- | :----------------------------------------------------------- | :----- |
| `x in s`               | 如果 *s* 中的某项等于 *x* 则结果为 `True`，否则为 `False`    | (1)    |
| `x not in s`           | 如果 *s* 中的某项等于 *x* 则结果为 `False`，否则为 `True`    | (1)    |
| `s + t`                | *s* 与 *t* 相拼接                                            | (6)(7) |
| `s * n` 或 `n * s`     | 相当于 *s* 与自身进行 *n* 次拼接                             | (2)(7) |
| `s[i]`                 | *s* 的第 *i* 项，起始为 0                                    | (3)    |
| `s[i:j]`               | *s* 从 *i* 到 *j* 的切片                                     | (3)(4) |
| `s[i:j:k]`             | *s* 从 *i* 到 *j* 步长为 *k* 的切片                          | (3)(5) |
| `len(s)`               | *s* 的长度                                                   |        |
| `min(s)`               | *s* 的最小项                                                 |        |
| `max(s)`               | *s* 的最大项                                                 |        |
| `s.index(x[, i[, j]])` | *x* 在 *s* 中首次出现项的索引号（索引号在 *i* 或其后且在 *j* 之前） | (8)    |
| `s.count(x)`           | *x* 在 *s* 中出现的总次数                                    |        |

### 可变序列操作

| 运算                      | 结果：                                                       | 备注 |
| :------------------------ | :----------------------------------------------------------- | :--- |
| `s[i] = x`                | 将 *s* 的第 *i* 项替换为 *x*                                 |      |
| `s[i:j] = t`              | 将 *s* 从 *i* 到 *j* 的切片替换为可迭代对象 *t* 的内容       |      |
| `del s[i:j]`              | 等同于 `s[i:j] = []`                                         |      |
| `s[i:j:k] = t`            | 将 `s[i:j:k]` 的元素替换为 *t* 的元素  （`t的长度必须与切片长度一致`） | (1)  |
| `del s[i:j:k]`            | 从列表中移除 `s[i:j:k]` 的元素                               |      |
| `s.append(x)`             | 将 *x* 添加到序列的末尾 (等同于 `s[len(s):len(s)] = [x]`)    |      |
| `s.clear()`               | 从 *s* 中移除所有项 (等同于 `del s[:]`)                      | (5)  |
| `s.copy()`                | 创建 *s* 的浅拷贝 (等同于 `s[:]`)                            | (5)  |
| `s.extend(t)` 或 `s += t` | 用 *t* 的内容扩展 *s* (基本上等同于 `s[len(s):len(s)] = t`)  |      |
| `s *= n`                  | 使用 *s* 的内容重复 *n* 次来对其进行更新                     | (6)  |
| `s.insert(i, x)`          | 在由 *i* 给出的索引位置将 *x* 插入 *s* (等同于 `s[i:i] = [x]`)   （注意return） |      |
| `s.pop()` 或 `s.pop(i)`   | 提取在 *i* 位置上的项，并将其从 *s* 中移除                   | (2)  |
| `s.remove(x)`             | 删除 *s* 中第一个 `s[i]` 等于 *x* 的项目。                   | (3)  |
| `s.reverse()`             | 就地将列表中的元素逆序。                                     | (4)  |

注释：

1. *t* 必须与它所替换的切片具有相同的长度。

2. 可选参数 *i* 默认为 `-1`，因此在默认情况下会移除并返回最后一项。

3. 当在 *s* 中找不到 *x* 时 `remove()` 操作会引发 [`ValueError`](https://docs.python.org/zh-cn/3/library/exceptions.html#ValueError)。

4. 当反转大尺寸序列时 `reverse()` 方法会原地修改该序列以保证空间经济性。 为提醒用户此操作是通过间接影响进行的，它并不会返回反转后的序列。

5. 包括 `clear()` 和 `copy()` 是为了与不支持切片操作的可变容器 (例如 [`dict`](https://docs.python.org/zh-cn/3/library/stdtypes.html#dict) 和 [`set`](https://docs.python.org/zh-cn/3/library/stdtypes.html#set)) 的接口保持一致。 `copy()` 不是 [`collections.abc.MutableSequence`](https://docs.python.org/zh-cn/3/library/collections.abc.html#collections.abc.MutableSequence) ABC 的一部分，但大多数具体的可变序列类都提供了它。

   *在 3.3 版本加入:* `clear()` 和 `copy()` 方法。

6. *n* 值为一个整数，或是一个实现了 [`__index__()`](https://docs.python.org/zh-cn/3/reference/datamodel.html#object.__index__) 的对象。 *n* 值为零或负数将清空序列。 序列中的项不会被拷贝；它们会被多次引用，正如 [通用序列操作](https://docs.python.org/zh-cn/3/library/stdtypes.html#typesseq-common) 中有关 `s * n` 的说明。

### **列表基本操作**

- **创建列表**

aList=[123,’abc’,4.56,[‘a’,’b’]]

- **访问列表**

print aList

索引操作符[]和切片操作符[:] ,从0开始算，[m:n]从下标m至n（不包含=n的元素）

print aList[3]

print aList[:2]  # [123,’abc’]

print aList[1:2]  # abc



> 循环访问

```python
for value in values:
    print(value)

for index in range(len(values)):
     value = values[index]
     print(index, value)

for count, value in enumerate(values):
	print(count, value)
```



-  **更新列表**

aList[3]=’z’

- **删除元素或整个列表**

del aList[3]

del aList  

pop() 方法可以删除并从列表中返回一个特定的对象(从后面取)

### **操作符**

- **标符操作符 (比较大小 < , > )**

> 直接以内建函数cmp作操作，按ascii码大小从第一个元素开始作比较
>

-  **切片（[] 和 [:] ）**

> 扩展切片指的是这样的切片形式：`a[start:stop:step]`，其中`step`是一个非零整数，即比简单切片多了调整**步长**的功能，此时切片的行为可概括为：从`start`对应的位置出发，以`step`为步长索引序列，直至越过`stop`对应的位置，且不包括`stop`本身
>
> 按下标索引值来访问列表元素，下标从0开始，如 [m:n], 下标起始m至n的元素, 结束下标n的元素不包含
>

- **成员关系操作 ( in , not in )**

> 检查一个对象是否是一个列表（或元组）的成员 ， 在的话，返回True, 否则返回False
>

- **连接符操作 (+)**

> 直接将一个列表B添加到另一个列表A， è A+B 即可，也可用extend替代‘+’
>

- **重复操作符 (*)**

> 重复操作更多用在字符串，将列表乘以N，相当于N个列表本身相加，组成新的列表,长度扩展了N倍
>

- **生成器表达式**

> **(expr for iter_var in iterable if cond_expr)**

| 列表函数                            | 作用                                                         |
| ----------------------------------- | ------------------------------------------------------------ |
| **List.append(obj)**                | 向列表中添加对象obj，无返回值                                |
| **List.count(obj)**                 | 返回一个对象在list中出列的次数                               |
| **List.extend(seq)**                | 将序列seq添加到List中，无返回值 ： `用于添加多个，参数是可迭代对象序列` ，也可用+也可但是不改变原list，会生成新list |
| **List.index(obj,i=0,j=len(list))** | List的索引 i~j（不包含j）范围内搜索对象obj,返回找到的索引值，如无则触发ValueError异常 |
| **List.insert(index,obj)**          | 在索引为index的地方插入对象obj                               |
| **List.pop(index=-1)**              | 删除并返回该对象，默认是最后一个对象                         |
| **List.remove(obj)**                | 从列表中删除对象,从0开始只删除一个                           |
| **List.reverse()**                  | 原列表翻转成列表 ,但reversed(List)返回一个迭代器，原列表不变 |
| **List.sort([func,key,reverse])**   | 以指定方式排序列表中的成员，reverse为true则反序，默认是false |

> `extend`与`+=`是等价的，会扩展原有的列表，`+`只能用来连接列表，且不改变原有的列表，会返回一个新列表，`append`会往原有列表中添加一个新的元素。

## 字典

```python
dict1 = {}
dict2 = {'name': 'earth', 'port': 80}

thisdict = dict(name = "John", age = 36, country = "Norway")
print(thisdict)
```

可以用一个很方便的内建方法**fromkeys()** 来创建一个"默认"字典, 字

典中元素具有相同的值 (如果没有给出， 默认为None)

```
>>> ddict = {}.fromkeys(('x', 'y'), -1)
>>> ddict
{'y': -1, 'x': -1}
# 字典的长度
>>> len(ddict)
```

> 访问，增加，删除
>
> 方法has_key()和 in 以及 not in 操作符都是布尔类型的  #  3.x不支持 has_key

```Python
x = thisdict.get("model")

for key in dict2:
	print 'key=%s, value=%s' % (key, dict2[key])

# Get Keys
x = thisdict.keys()
# Get Values
x = thisdict.values()
# get Items
x = thisdict.items()
# Check if Key Exists
if "model" in thisdict:
  print("Yes, 'model' is one of the keys in the thisdict dictionary")

# update或add
thisdict["color"] = "red"
# The update() method will update the dictionary with the items from a given argument. If the item does not exist, the item will be added.
thisdict.update({"color": "red"})
# The pop() method removes the item with the specified key name:
thisdict.pop("model")
# The popitem() method removes the last inserted item (in versions before 3.7, a random item is removed instead):
thisdict.popitem()
# 
del thisdict["model"]
# delete the dictionary completely:
del thisdict 

# **Loop dictionary**
for x in thisdict:
  print(thisdict[x])

for x in thisdict.values():
  print(x)

for x in thisdict.keys():
  print(x)

for x, y in thisdict.items():
  print(x, y)

# Dict Copy
thisdict.copy()
mydict = dict(thisdict)

# 按值来排序
d = {'k':299, 'c':100,'a':999}
print(sorted(d.items(),key=lambda x:x[1]))
```



> methods

| Method                                                       | Description                                                  |
| :----------------------------------------------------------- | :----------------------------------------------------------- |
| [clear()](https://www.w3schools.com/python/ref_dictionary_clear.asp) | Removes all the elements from the dictionary                 |
| [copy()](https://www.w3schools.com/python/ref_dictionary_copy.asp) | Returns a copy of the dictionary                             |
| [fromkeys()](https://www.w3schools.com/python/ref_dictionary_fromkeys.asp) | Returns a dictionary with the specified keys and value       |
| [get()](https://www.w3schools.com/python/ref_dictionary_get.asp) | Returns the value of the specified key                       |
| [items()](https://www.w3schools.com/python/ref_dictionary_items.asp) | Returns a list containing a tuple for each key value pair    |
| [keys()](https://www.w3schools.com/python/ref_dictionary_keys.asp) | Returns a list containing the dictionary's keys              |
| [pop()](https://www.w3schools.com/python/ref_dictionary_pop.asp) | Removes the element with the specified key                   |
| [popitem()](https://www.w3schools.com/python/ref_dictionary_popitem.asp) | Removes the last inserted key-value pair                     |
| [setdefault()](https://www.w3schools.com/python/ref_dictionary_setdefault.asp) | Returns the value of the specified key. If the key does not exist: insert the key, with the specified value |
| [update()](https://www.w3schools.com/python/ref_dictionary_update.asp) | Updates the dictionary with the specified key-value pairs    |
| [values()](https://www.w3schools.com/python/ref_dictionary_values.asp) | Returns a list of all the values in the dictionary           |

> 只支持两个字典是否相等==（是比较的key以及对应的值（与key顺序无关），不是比较id（id可能不一样）），不支持< , >的比较（会报错）

## 集合

**交集 &** : **x&y**，返回一个新的集合，包括同时在集合 x 和y中的共同元素。

**并集 |** : **x|y**，返回一个新的集合，包括集合 x 和 y 中所有元素。

**差集 -** : **x-y**，返回一个新的集合,包括在集合 x 中但不在集合 y 中的元素。

**补集 ^** : **x^y**，返回一个新的集合，包括集合 x 和 y 的非共同元素。



添加元素   

```python
s.add(x)     # 增加单个
s.update(y)  # y是可迭代对象（增加多个）
```

移除元素  

```python
s.remove( x )     # 将元素 x 从集合 s 中移除，如果元素不存在，则会发生错误。
s.discard( x )    # 将元素 x 从集合 s 中移除，如果元素不存在，则`不会发生错误`。
```

计算集合元素个数 ，清除，判断元素是否存在集合中

```python
len(s)   # length of s
s.clear()  # clear all element in s

x in s     # check if element x in s
```



| 方法                                                         | 描述                                                         |
| :----------------------------------------------------------- | :----------------------------------------------------------- |
| [add()](https://www.runoob.com/python3/ref-set-add.html)     | 为集合添加元素                                               |
| [clear()](https://www.runoob.com/python3/ref-set-clear.html) | 移除集合中的所有元素                                         |
| [copy()](https://www.runoob.com/python3/ref-set-copy.html)   | 拷贝一个集合                                                 |
| [difference()](https://www.runoob.com/python3/ref-set-difference.html) | 返回多个集合的差集                                           |
| [difference_update()](https://www.runoob.com/python3/ref-set-difference_update.html) | 移除集合中的元素，该元素在指定的集合也存在。                 |
| [discard()](https://www.runoob.com/python3/ref-set-discard.html) | 删除集合中指定的元素                                         |
| [intersection()](https://www.runoob.com/python3/ref-set-intersection.html) | 返回集合的交集                                               |
| [intersection_update()](https://www.runoob.com/python3/ref-set-intersection_update.html) | 返回集合的交集。                                             |
| [isdisjoint()](https://www.runoob.com/python3/ref-set-isdisjoint.html) | 判断两个集合是否包含相同的元素，如果没有返回 True，否则返回 False。 |
| [issubset()](https://www.runoob.com/python3/ref-set-issubset.html) | 判断指定集合是否为该方法参数集合的子集。                     |
| [issuperset()](https://www.runoob.com/python3/ref-set-issuperset.html) | 判断该方法的参数集合是否为指定集合的子集                     |
| [pop()](https://www.runoob.com/python3/ref-set-pop.html)     | 随机移除元素（返回移除的元素）                               |
| [remove()](https://www.runoob.com/python3/ref-set-remove.html) | 移除指定元素（指定的元素不存时报错）                         |
| [symmetric_difference()](https://www.runoob.com/python3/ref-set-symmetric_difference.html) | 返回两个集合中不重复的元素集合。                             |
| [symmetric_difference_update()](https://www.runoob.com/python3/ref-set-symmetric_difference_update.html) | 移除当前集合中在另外一个指定集合相同的元素，并将另外一个指定集合中不同的元素插入到当前集合中。 |
| [union()](https://www.runoob.com/python3/ref-set-union.html) | 返回两个集合的并集                                           |
| [update()](https://www.runoob.com/python3/ref-set-update.html) | 给集合添加元素                                               |
| [len()](https://www.runoob.com/python3/python3-string-len.html) | 计算集合元素个数                                             |

> 集合作比较时不会报错，< ，>运算比较时，是比较一个集合是否是另一个集合的子集，不是子集就返回false
>
> ==时是比较两个集合的元素是否一样（set是排序的, ==的集合，id可能是不一样的）

## build-in Function

### all()

Check if all items in a list are True:

```python
mylist = [True, True, True]
x = all(mylist)
```

Check if all items in a list are True:

```python
mylist = [0, 1, 1]
x = all(mylist)
```



Check if all items in a *tuple* are True:

```python
mytuple = (0, True, False)
x = all(mytuple)
```



Check if all items in a *set* are True:

```python
myset = {0, 1, 0}
x = all(myset)
```



Check if all items in a *dictionary* are True:

```python
# key 全为非0才返回true
mydict = {0 : "Apple", 1 : "Orange"}
x = all(mydict)
```



### slice()

```python
In [164]: l=[1,2,3,4,5,6]
In [165]: a=slice(1,5,2)

In [174]: l[1:5:2]   # a=slice(1,5,2) ， 相当于 l[a]
Out[174]: [2, 4]

In [175]: l[a]
Out[175]: [2, 4]
```



## 字符串操作

### 字符串常用的方法

| Method                                                       | Description                                                  |
| :----------------------------------------------------------- | :----------------------------------------------------------- |
| [capitalize()](https://www.w3schools.com/python/ref_string_capitalize.asp) | Converts the first character to upper case                   |
| [casefold()](https://www.w3schools.com/python/ref_string_casefold.asp) | Converts string into lower case                              |
| [center()](https://www.w3schools.com/python/ref_string_center.asp) | Returns a centered string                                    |
| [count()](https://www.w3schools.com/python/ref_string_count.asp) | Returns the number of times a specified value occurs in a string |
| [encode()](https://www.w3schools.com/python/ref_string_encode.asp) | Returns an encoded version of the string                     |
| [endswith()](https://www.w3schools.com/python/ref_string_endswith.asp) | Returns true if the string ends with the specified value     |
| [expandtabs()](https://www.w3schools.com/python/ref_string_expandtabs.asp) | Sets the tab size of the string                              |
| [find()](https://www.w3schools.com/python/ref_string_find.asp) | Searches the string for a specified value and returns the position of where it was found |
| [format()](https://www.w3schools.com/python/ref_string_format.asp) | Formats specified values in a string                         |
| format_map()                                                 | Formats specified values in a string                         |
| [index()](https://www.w3schools.com/python/ref_string_index.asp) | Searches the string for a specified value and returns the position of where it was found |
| [isalnum()](https://www.w3schools.com/python/ref_string_isalnum.asp) | Returns True if all characters in the string are alphanumeric |
| [isalpha()](https://www.w3schools.com/python/ref_string_isalpha.asp) | Returns True if all characters in the string are in the alphabet |
| [isascii()](https://www.w3schools.com/python/ref_string_isascii.asp) | Returns True if all characters in the string are ascii characters |
| [isdecimal()](https://www.w3schools.com/python/ref_string_isdecimal.asp) | Returns True if all characters in the string are decimals    |
| [isdigit()](https://www.w3schools.com/python/ref_string_isdigit.asp) | Returns True if all characters in the string are digits      |
| [isidentifier()](https://www.w3schools.com/python/ref_string_isidentifier.asp) | Returns True if the string is an identifier                  |
| [islower()](https://www.w3schools.com/python/ref_string_islower.asp) | Returns True if all characters in the string are lower case  |
| [isnumeric()](https://www.w3schools.com/python/ref_string_isnumeric.asp) | Returns True if all characters in the string are numeric     |
| [isprintable()](https://www.w3schools.com/python/ref_string_isprintable.asp) | Returns True if all characters in the string are printable   |
| [isspace()](https://www.w3schools.com/python/ref_string_isspace.asp) | Returns True if all characters in the string are whitespaces |
| [istitle()](https://www.w3schools.com/python/ref_string_istitle.asp) | Returns True if the string follows the rules of a title      |
| [isupper()](https://www.w3schools.com/python/ref_string_isupper.asp) | Returns True if all characters in the string are upper case  |
| [join()](https://www.w3schools.com/python/ref_string_join.asp) | Joins the elements of an iterable to the end of the string   |
| [ljust()](https://www.w3schools.com/python/ref_string_ljust.asp) | Returns a left justified version of the string               |
| [lower()](https://www.w3schools.com/python/ref_string_lower.asp) | Converts a string into lower case                            |
| [lstrip()](https://www.w3schools.com/python/ref_string_lstrip.asp) | Returns a left trim version of the string                    |
| [maketrans()](https://www.w3schools.com/python/ref_string_maketrans.asp) | Returns a translation table to be used in translations       |
| [partition()](https://www.w3schools.com/python/ref_string_partition.asp) | Returns a tuple where the string is parted into three parts  |
| [replace()](https://www.w3schools.com/python/ref_string_replace.asp) | Returns a string where a specified value is replaced with a specified value |
| [rfind()](https://www.w3schools.com/python/ref_string_rfind.asp) | Searches the string for a specified value and returns the last position of where it was found |
| [rindex()](https://www.w3schools.com/python/ref_string_rindex.asp) | Searches the string for a specified value and returns the last position of where it was found |
| [rjust()](https://www.w3schools.com/python/ref_string_rjust.asp) | Returns a right justified version of the string              |
| [rpartition()](https://www.w3schools.com/python/ref_string_rpartition.asp) | Returns a tuple where the string is parted into three parts  |
| [rsplit()](https://www.w3schools.com/python/ref_string_rsplit.asp) | Splits the string at the specified separator, and returns a list |
| [rstrip()](https://www.w3schools.com/python/ref_string_rstrip.asp) | Returns a right trim version of the string                   |
| [split()](https://www.w3schools.com/python/ref_string_split.asp) | Splits the string at the specified separator, and returns a list |
| [splitlines()](https://www.w3schools.com/python/ref_string_splitlines.asp) | Splits the string at line breaks and returns a list          |
| [startswith()](https://www.w3schools.com/python/ref_string_startswith.asp) | Returns true if the string starts with the specified value   |
| [strip()](https://www.w3schools.com/python/ref_string_strip.asp) | Returns a trimmed version of the string                      |
| [swapcase()](https://www.w3schools.com/python/ref_string_swapcase.asp) | Swaps cases, lower case becomes upper case and vice versa    |
| [title()](https://www.w3schools.com/python/ref_string_title.asp) | Converts the first character of each word to upper case      |
| [translate()](https://www.w3schools.com/python/ref_string_translate.asp) | Returns a translated string                                  |
| [upper()](https://www.w3schools.com/python/ref_string_upper.asp) | Converts a string into upper case                            |
| [zfill()](https://www.w3schools.com/python/ref_string_zfill.asp) | Fills the string with a specified number of 0 values at the beginning |

### 拆分字符串



> `re`
>
> - split(): This method is used to split a given string into a list.
> - sub(): This method is used to find a substring where a regex pattern matches, and then it replaces the matched substring with a different string.
> - subn(): This method is similar to the sub() method, but it returns the new string, along with the number of replacements.



```python
def mySplit(s,ds):
    res = [s]
    print(res)
    for d in ds:
        t = []
        # In Python 2 map() returns a list while in Python 3 it returns an iterator.
        # map(lambda x: t.extend(x.split(d)),res)  # only work in Python 2
        for x in res:
            t.extend(x.split(d))
        ####################
        res = t
    return [ x for x in res if x ] 

s = 'ab;bbcc,effg|mm\topt;r,\ttxt'
print(mySplit(s,';,|\t'))   # ['ab', 'bbcc', 'effg', 'mm', 'opt', 'r', 'txt']

# use re.split
entries = re.split(r'[;,|\t]+', s)
print(entries)
```

> 注意re.split的模式里有分组，则分组匹配的也在在拆分的列表里
>
> ```python
> s = 'ab;bbcc,effg|mm\topt;r,\ttxt'
> entries = re.split(r'([;,|\t]+)', s)   # 有分组，分组也在在其中
> values = entries[::2]  # 隔两个取值
> ```
>
> 注意需要确认分组是非捕获分组,形如(`?:...`)
>
> ```
> fileds = re.split(r'(?:[;,|\t]+)', s)
> ```

### 调整字符串的文本格式

```python
astr='2024-03-23'
print(re.sub('(\d{4})-(\d{2})-(\d{2})',r'\2/\3/\1',astr)) # 03/23/2024
print(re.sub('(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})',r'\g<month>/\g<day>/\g<year>',astr))  # 03/23/2024
```



### 以字符串开头或结尾

> **startswith**

 def startswith(
    *__prefix*: *str* | tuple[*str*, ...],
    *__start*: SupportsIndex | None = ...,
    *__end*: SupportsIndex | None = ...,
    /
) -> *bool*



> endswith

 def endswith(
    *__suffix*: *str* | tuple[*str*, ...],
    *__start*: SupportsIndex | None = ...,
    *__end*: SupportsIndex | None = ...,
    /
) -> *bool*

```python
    a = 'a.sh'
    print(a.endswith(('.txt','.py')))
```

### fnmatch匹配Shell中常用的通配符

> fnmatch模块提供了两个函数—— fnmatch() 和 fnmatchcase()

```python
from fnmatch import fnmatch, fnmatchcase
fnmatch('foo.txt', '*.txt')
```

### 统一unicode表示方法

```python
import unicodedata

s1 = 'Spicy Jalape\u00f1o'
s2 = 'Spicy Jalapen\u0303o'

t1 = unicodedata.normalize('NFC', s1)
t2 = unicodedata.normalize('NFC', s2)
print(t1 == t2)
```



### 匹配多行

```python
comment = re.compile(r'/\*((?:.|\n)*?)\*/')
comment = re.compile(r'/\*(.*?)\*/', re.DOTALL)
```



### 左右中对齐

```python
s.ljust(20)  # 左对齐，在右边以空格填充
s.ljust(20,'=')  # 左对齐，在右边以=填充
s.rjust(20)
s.center(20)
```

使用内置`format`

```python
format(s,'<20')
format(s,'>20')
format(s,'^20')
```

### 格式化输出

> **%格式化字符串**
>
> ```python
> print("DAY %s 格式化字符串 %s " % (value1,value2))
> ```



> **format()**
>
> ```python
> >>>"{} {}".format("hello", "world")    # 不设置指定位置，按默认顺序
> 'hello world'
>  
> >>> "{0} {1}".format("hello", "world")  # 设置指定位置
> 'hello world'
>  
> >>> "{1} {0} {1}".format("hello", "world")  # 设置指定位置
> ```



> **f-string**
>
> ```python
> print(f'my name is {name}, this year is {date:%Y},Next year, I\'m {age+1}')  # my name is zings, this year is 2019,Next year, I'm 18
> 
> In [2]: for i in range(1,10):
>    ...:     for j in range(1,i+1):
>    ...:         print(f'{j}*{i}={i*j:<2}',end=" ")
>    ...:     print()
>    ...:
> 1*1=1
> 1*2=2  2*2=4
> 1*3=3  2*3=6  3*3=9
> 1*4=4  2*4=8  3*4=12 4*4=16
> 1*5=5  2*5=10 3*5=15 4*5=20 5*5=25
> 1*6=6  2*6=12 3*6=18 4*6=24 5*6=30 6*6=36
> 1*7=7  2*7=14 3*7=21 4*7=28 5*7=35 6*7=42 7*7=49
> 1*8=8  2*8=16 3*8=24 4*8=32 5*8=40 6*8=48 7*8=56 8*8=64
> 1*9=9  2*9=18 3*9=27 4*9=36 5*9=45 6*9=54 7*9=63 8*9=72 9*9=81
> ```



## Exception

### Many Exceptions

```python
try:
  print(x)
except NameError:
  print("Variable x is not defined")
except:
  print("Something else went wrong")
```



### Else

You can use the `else` keyword to define a block of code to be executed if no errors were raised:

```python
try:
  print("Hello")
except:
  print("Something went wrong")
else:
  print("Nothing went wrong")
```

### Finally

The `finally` block, if specified, will be executed regardless if the try block raises an error or not.

```python
try:
  f = open("demofile.txt")
  try:
    f.write("Lorum Ipsum")
  except:
    print("Something went wrong when writing to the file")
  finally:
    f.close()
except:
  print("Something went wrong when opening the file")
```

> 注意，在函数里，对于finally里有return时，它会优先于try和except里的return。

### **Raise exception**

```python
x = -1

if x < 0:
  raise Exception("Sorry, no numbers below zero")
```



## **内建函数apply()、filter()、map()、reduce()**

The `filter()` function returns an iterator where the items are filtered through a function to test if the item is accepted or not.

The `map()` function executes a specified function for each item in an iterable. The item is sent to the function as a parameter.



```python
numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9]
evens = list(filter(lambda x: x % 2 == 0, numbers))
print(evens)  # Output: [2, 4, 6, 8]
```





```python
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, numbers))
print(squared)  # Output: [1, 4, 9, 16, 25]
```



> Python 3.x returns a `generator object`

```python
# Define a list
my_list = [1, 2, 3]

res_list = []
# Use map() to apply a lambda function that extends the list with squared values
map(lambda x: res_list.extend([x**2]), my_list)        # only working in python2 as expected
################################
list(map(lambda x: res_list.extend([x**2]), my_list))  # require to use list to transfer into list in python3
[g for g in map(lambda x: res_list.extend([x**2]), my_list)] # current one is also working 
    
# Print the modified list
print(res_list)
```



```python
from functools import reduce
# reduce(func, [1, 2, 3]) = func(func(1, 2), 3)

numbers = [1, 2, 3, 4, 5]
# reduce(lambda x, y: x+y, [1, 2, 3, 4, 5]) calculates ((((1+2)+3)+4)+5). 
# reduce(lambda x, y: x*y, [1, 2, 3, 4, 5]) calculates ((((1*2)*3)*4)*5). 
product = reduce(lambda x, y: x * y, numbers)
print(product)  # Output: 120 (1 * 2 * 3 * 4 * 5)
```



> **列表、字典、集合解析进行数据筛选**
>
> ```python
> filter(lambda x:x>0,data)             		## 列表data大于0的元素
> [ x for in data if x > 0 ]                  ## 列表解析，同上
> { k:v for k,v in d.items() if v > 90 }       ## 字典解析，字典d中的value大于90的键值对，python2用 iteritems	
> {x for x in s if x%3==0 }             		##集合解析，集合s中所有被3整除的元素
> ```



> **lambda**
>
> ```python
> funcs = [lambda x: x+n for n in range(5)]
> #funcs = [lambda x,n=n: x+n for n in range(5)]
> for f in funcs:
> 	print(f(0))
> ## output as below:
> 4
> 4
> 4
> 4
> 4
> ```

## Collections

`collections` 模块是 Python 标准库中提供了一些有用的集合数据类型的模块。以下是一些常用的 `collections` 模块中的数据类型和用法：

1. **`namedtuple`（命名元组）:**

   - `namedtuple` 是一个工厂函数，用于创建具有命名字段的元组子类。这使得元组更易读和自文档化。

   ```python
   from collections import namedtuple
   
   # 创建一个命名元组类型
   Point = namedtuple('Point', ['x', 'y'])
   
   # 创建命名元组实例
   p = Point(1, 2)
   
   # 访问字段
   print(p.x, p.y)  # 输出: 1 2
   ```

2. **`Counter`（计数器）:**

   - `Counter` 是一个用于计数可哈希对象的子类，通常用于统计元素出现的次数。

   ```python
   from collections import Counter
   
   # 创建一个计数器
   c = Counter(['a', 'b', 'a', 'c', 'b', 'a'])
   
   print(c.keys())  #  dict_keys(['a', 'b', 'c'])
   # 访问元素计数
   print(c['a'])  # 输出: 3
   print(c.most_common(2))  # [('a', 3), ('b', 2)]
   ```

3. **`defaultdict`（默认字典）:**

   - `defaultdict` 是一个字典的子类，它允许给定默认值以供访问不存在的键。

   ```python
   from collections import defaultdict
   
   # 创建一个默认字典，初始值为 int 类型的 0
   d = defaultdict(int)
   
   # 访问不存在的键，会使用默认值 0
   print(d['key'])  # 输出: 0
   ```

4. **`deque`（双端队列）:**

   - `deque` 是一个双端队列，支持从两端高效地添加和删除元素。

   ```python
   from collections import deque
   
   # 创建一个双端队列
   dq = deque([1, 2, 3])
   
   # 从右侧添加元素
   dq.append(4)
   
   # 从左侧添加元素
   dq.appendleft(0)
   
   # 弹出右侧元素
   popped_element = dq.pop()
   
   print(dq)  # 输出: deque([0, 1, 2, 3])
   print(popped_element)  # 输出: 4
   ```



这只是 `collections` 模块提供的一小部分功能。根据需要，你还可以探索其他功能，如 `OrderedDict`、`ChainMap` 等。这些数据类型在不同的情况下提供了更多的灵活性和性能优.



| [**namedtuple()**](https://docs.python.org/2/library/collections.html#collections.namedtuple) | factory function for  creating tuple subclasses with named fields | *New in version 2.6.* |
| ------------------------------------------------------------ | ------------------------------------------------------------ | --------------------- |
| [**deque**](https://docs.python.org/2/library/collections.html#collections.deque) | list-like container with  fast appends and pops on either end | *New in version 2.4.* |
| [**Counter**](https://docs.python.org/2/library/collections.html#collections.Counter) | dict subclass for  counting hashable objects                 | *New in version 2.7.* |
| [**OrderedDict**](https://docs.python.org/2/library/collections.html#collections.OrderedDict) | dict subclass that  remembers the order entries were added   | *New in version 2.7.* |
| [**defaultdict**](https://docs.python.org/2/library/collections.html#collections.defaultdict) | dict subclass that calls  a factory function to supply missing values | *New in version 2.5.* |

> `namedtuple`是一个函数，它用来创建一个自定义的`tuple`对象，并且规定了`tuple`元素的个数，并可以用属性而不是索引来引用`tuple`的某个元素。
>
> 这样一来，我们用`namedtuple`可以很方便地定义一种数据类型，它具备tuple的不变性，又可以根据属性来引用，使用十分方便。

```python
from collections import namedtuple

if __name__ == '__main__':
  Student = namedtuple('Student',['name','age','sex'])
  s = Student('james',39,'male')
  print(s.name,s[1])    # output : james 39
```



> `deque`是为了高效实现插入和删除操作的双向列表，适合用于队列和栈：
>
> `deque`除了实现list的`append()`和`pop()`外，还支持`appendleft()`和`popleft()`，这样就可以非常高效地往头部添加或删除元素。
>
> (使用`list`存储数据时，按索引访问元素很快，但是插入和删除元素就很慢了，因为`list`是线性存储，数据量大的时候，插入和删除效率很低)

```PYTHON
from collections import deque

if __name__ == '__main__':

  d = deque([1,2,3,4,5])
  d.appendleft('x')
  print(d)
```







## 时间

1. 在Python中，通常有这几种方式来表示时间：

   1）时间戳 

   2）格式化的时间字符串 

   3）元组（struct_time）共九个元素。由于Python的time模块实现主要调用C库，所以各个平台可能有所不同。

2. UTC（Coordinated Universal Time，世界协调时）亦即格林威治天文时间，世界标准时间。在中国为UTC+8。DST（Daylight Saving Time）即夏令时。

3. 时间戳（timestamp）的方式：通常来说，时间戳表示的是从***\*1970年1月1日00:00:00\****开始按秒计算的偏移量。我们运行“type(time.time())”，返回的是float类型。返回时间戳方式的函数主要有time()，clock()等。

4. 元组（struct_time）方式：struct_time元组共有9个元素，返回struct_time的函数主要有gmtime()，localtime()，strptime()。下面列出这种方式元组中的几个元素：

| ***\*索引（Index）\**** | ***\*属性（Attribute）\**** | ***\*值（Values）\**** |
| ----------------------- | --------------------------- | ---------------------- |
| 0                       | tm_year（年）               | 比如2011               |
| 1                       | tm_mon（月）                | 1 - 12                 |
| 2                       | tm_mday（日）               | 1 - 31                 |
| 3                       | tm_hour（时）               | 0 - 23                 |
| 4                       | tm_min（分）                | 0 - 59                 |
| 5                       | tm_sec（秒）                | 0 - 61                 |
| 6                       | tm_wday（weekday）          | 0 - 6（0表示周日）     |
| 7                       | tm_yday（一年中的第几天）   | 1 - 366                |
| 8                       | tm_isdst（是否是夏令时）    | 默认为-1               |

| time.***\*time()\****                                        | 返回当前时间戳                                               |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| time.***\*localtime([secs])\****                             | 将一个时间戳转换为当前时区的struct_time。secs参数未提供，则以当前时间为准。 |
| time.***\*gmtime([\*******\**secs\**\******\*])\****         | 和localtime()方法类似，gmtime()方法是将一个时间戳转换为UTC时区（0时区）的struct_tim |
| time.***\*mktime(\*******\**t\**\******\*)\****              | 将一个struct_time转化为时间戳                                |
| time.***\*asctime([\*******\**t\**\******\*])\****           | 把一个表示时间的元组或者struct_time表示为这种形式：***\*'Sun Jun 20 23:21:05 1993'\****。如果没有参数，将会将time.localtime()作为参数传入 |
| time.***\*ctime([\*******\**secs\**\******\*])\****          | 把一个时间戳（按秒计算的浮点数）转化为time.asctime()的形式。如果参数未给或者为None的时候，将会默认time.time()为参数。它的作用相当于time.asctime(time.localtime(secs)) |
| time.***\*clock()\****                                       | 这个需要注意，在不同的系统上***\*含义不同\****。在UNIX系统上，它返回的是“进程时间”，它是用秒表示的浮点数（时间戳）。而在WINDOWS中，第一次调用，返回的是进程运行的实际时间。而第二次之后的调用是自第一次调用以后到现在的运行时间。（实际上是以WIN32上QueryPerformanceCounter()为基础，它比毫秒表示更为精确） |
| time.***\*strftime(\*******\**format\**\******\*[,\**** ***\**t\**\******\*])\**** | 把一个代表时间的元组或者struct_time（如由time.localtime()和time.gmtime()返回）转化为格式化的时间字符串。如果t未指定，将传入time.localtime()。如果元组中任何一个元素越界，ValueError的错误将会被抛出 |
| time.***\*clock()\****                                       | 这个需要注意，在不同的系统上***\*含义不同\****。在UNIX系统上，它返回的是“进程时间”，它是用秒表示的浮点数（时间戳）。而在WINDOWS中，第一次调用，返回的是进程运行的实际时间。而第二次之后的调用是自第一次调用以后到现在的运行时间。（实际上是以WIN32上QueryPerformanceCounter()为基础，它比毫秒表示更为精确） |
| time.***\*sleep(\*******\**secs\**\******\*)\****            | 线程推迟指定的时间运行。单位为秒                             |

![img](C:\Users\Caiman\Documents\Docs\A技术大全准备.assets\wps1.jpg)



## 装饰

```python
import time

def timing_decorator(func):
 def wrapper(*args, **kwargs):
     start_time = time.time()
     result = func(*args, **kwargs)
     end_time = time.time()
     elapsed_time = end_time - start_time
     print(f"{func.__name__} 执行时间: {elapsed_time:.4f} 秒")
     return result
 return wrapper

# 使用装饰器
@timing_decorator
def example_function():
 # 模拟一个耗时操作
 time.sleep(2)
 print("Function executed")

# 调用被装饰的函数
example_function()
```



## 类

### 继承和多态

类的静态方法(staticmethod)无法访问类属性；类方法(classmethod)是可以访问类属性的。
都可以通过类名或者实例名来调用方法（只能通过类名加 ‘.’ 的方式调用），但是不都不能间接调用



`Method Resolution Order (MRO) C3线性化算法`

Python3的继承机制不同于Python2。其核心原则是下面两条，请谨记！

- 子类在调用某个方法或变量的时候，首先在自己内部查找，如果没有找到，则开始根据继承机制在父类里查找。
- 根据父类定义中的顺序，以**深度优先**的方式逐一查找父类！



### 派生内置不可变类型并修改实例化行为

> __new__方法主要是当你继承一些不可变的class时(比如int, str, tuple)， 提供给你一个自定义这些类的实例化过程的途径。 

```python
class TestObj(object):
    def __init__(self, *args, **kwargs):
        print("in init...")
    def __new__(cls, *args, **kwargs):
        print("in new...")
        return object.__new__(cls, *args, **kwargs)  

class IntTuple(tuple):
    def __init__(self,iterable):
        # super(IntTuple,self).__init__(iterable)
        pass
    def __new__(cls,iterable):
        g = (i for i in iterable if isinstance(i, int) and i > 0)
        return super(IntTuple,cls).__new__(cls, g)

t = IntTuple([1,'x','pzfdf',[1,2],['x'],6,3])   
print(t)      # output (1, 6, 3)
```



### 使用描述符实例对类型检查

描述符的定义很简单，实现了下列*任意一个方法*的 Python 对象就是一个描述符（descriptor）：

- `__get__(self, obj, type=None)`
- `__set__(self, obj, value)`
- `__delete__(self, obj)`



对象属性的访问顺序：

①.实例属性

②.类属性

③.父类属性

④.__getattr__()方法



**描述符可以用来控制对属性的访问行为，实现计算属性、懒加载属性、属性访问控制等功能**

```python
class AssignValue(object):
    def __init__(self, name, mtype):
        self.name = name 
        self.mtype = mtype

    def __get__(self, instance,cls):
        print(f'in __get__ {instance},{cls}')
        return instance.__dict__[self.name]

    def __set__(self, instance,value):
        print(f'in __set__ {instance},{value}')
        if not isinstance(value,self.mtype):
            raise TypeError(f"excepted an '{self.mtype}', but given type '{type(value)}'")
        instance.__dict__[self.name] = value
        pass 
    def __delete__(self, instance):
        print(f'in __delete__ {instance}')
        del instance.__dict__[self.name]
        pass

class Person(object):
    name = AssignValue('name',str)
    age =  AssignValue('age',int)
    height = AssignValue('height',float)
    
p = Person()
p.name = 1
```

比较一下下面的方式

```python
class Person:
	def __init__(self, first_name):
		self.first_name = first_name
	# Getter function
	@property
	def first_name(self):
		return self._first_name	
	# Setter function
	@first_name.setter
	def first_name(self, value):
		if not isinstance(value, str):
			raise TypeError('Expected a string')
		self._first_name = value
	# Deleter function (optional)
	@first_name.deleter
	def first_name(self):
		raise AttributeError("Can't delete attribute")
```



### 上下文管理器（Context Manager）

上下文管理器是指在一段代码执行之前执行一段代码，用于一些预处理工作；执行之后再执行一段代码，用于一些清理工作。



## 线程

[`queue`](https://docs.python.org/zh-cn/3.10/library/queue.html#module-queue) 提供了一个线程安全的接口用来在运行中的线程之间交换数据。

[`asyncio`](https://docs.python.org/zh-cn/3.10/library/asyncio.html#module-asyncio) 提供了一个替代方式用来实现任务层级的并发而不要求使用多个操作系统线程。



**CPython 实现细节：** 在 CPython 中，由于存在 [全局解释器锁](https://docs.python.org/zh-cn/3.10/glossary.html#term-global-interpreter-lock)，同一时刻只有一个线程可以执行 Python 代码（虽然某些性能导向的库可能会去除此限制）。 如果你想让你的应用更好地利用多核心计算机的计算资源，推荐你使用 [`multiprocessing`](https://docs.python.org/zh-cn/3.10/library/multiprocessing.html#module-multiprocessing) 或 [`concurrent.futures.ProcessPoolExecutor`](https://docs.python.org/zh-cn/3.10/library/concurrent.futures.html#concurrent.futures.ProcessPoolExecutor)。 但是，如果你想要同时运行多个 I/O 密集型任务，则多线程仍然是一个合适的模型。





## Padans

### Pandas Series

A Pandas Series is like a column in a table.

```python
a = [1, 7, 2]

myvar = pd.Series(a, index = ["x", "y", "z"])

print(myvar)
```

```python
pandas.Series(data=None, index=None, dtype=None, name=None, copy=False, fastpath=False)
```

- `data`：Series 的数据部分，可以是列表、数组、字典、标量值等。如果不提供此参数，则创建一个空的 Series。
- `index`：Series 的索引部分，用于对数据进行标记。可以是列表、数组、索引对象等。如果不提供此参数，则创建一个默认的整数索引。
- `dtype`：指定 Series 的数据类型。可以是 NumPy 的数据类型，例如 `np.int64`、`np.float64` 等。如果不提供此参数，则根据数据自动推断数据类型。
- `name`：Series 的名称，用于标识 Series 对象。如果提供了此参数，则创建的 Series 对象将具有指定的名称。
- `copy`：是否复制数据。默认为 False，表示不复制数据。如果设置为 True，则复制输入的数据。
- `fastpath`：是否启用快速路径。默认为 False。启用快速路径可能会在某些情况下提高性能

```python
s = pd.Series({'a': 1, 'b': 2, 'c': 3, 'd': 4})

# 使用切片语法来访问 Series 的一部分
print(s['a':'c'])  # 返回索引标签 'a' 到 'c' 之间的元素
print(s[:3])  # 返回前三个元素

# 使用 del 删除指定索引标签的元素。
del s['a']  # 删除索引标签 'a' 对应的元素

# 使用 drop 方法删除一个或多个索引标签，并返回一个新的 Series。
s_dropped = s.drop(['b'])  # 返回一个删除了索引标签 'b' 的新 Series

# 算术运算
result = series * 2  # 所有元素乘以2

# 过滤
filtered_series = series[series > 2]  # 选择大于2的元素

# 数学函数
import numpy as np
result = np.sqrt(series)  # 对每个元素取平方根

# 获取索引
index = s.index

# 获取值数组
values = s.values

# 获取描述统计信息
stats = s.describe()

# 获取最大值和最小值的索引
max_index = s.idxmax()
min_index = s.idxmin()

# 其他属性和方法
print(s.dtype)   # 数据类型
print(s.shape)   # 形状
print(s.size)    # 元素个数
print(s.head())  # 前几个元素，默认是前 5 个
print(s.tail())  # 后几个元素，默认是后 5 个
print(s.sum())   # 求和
print(s.mean())  # 平均值
print(s.std())   # 标准差
print(s.min())   # 最小值
print(s.max())   # 最大值

df1 = pd.Series([2, 4, 8, 10, 12])
df2 = pd.Series([8, 12, 10, 15, 16])
df1=df1[~df1.isin(df2)]
print(df1)
```



### DataFrame

A Pandas DataFrame is a 2 dimensional data structure, like a 2 dimensional array, or a table with rows and columns.

```python
import pandas as pd

data = {
  "calories": [420, 380, 390],
  "duration": [50, 40, 45]
}

#load data into a DataFrame object:
df = pd.DataFrame(data)

print(df)   # 1

#refer to the row index:
print(df.loc[0])

#use a list of indexes:
print(df.loc[[0, 1]])

pd.options.display.max_rows = 9999

```



> ```
>      calories  duration
>   0       420        50
>   1       380        40
>   2       390        45
>   
>   
>   # print(df.loc[0])
>   calories    420
>   duration     50
>   Name: 0, dtype: int64
> 
>   #print(df.loc[[0, 1]])
>      calories  duration
>   0       420        50
>   1       380        40
> ```



### Analyzing DataFrames

```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.head(10))

# Print the first 5 rows of the DataFrame:
print(df.head())

print(df.tail()) 

print(df.info()) 
```

```
  <class 'pandas.core.frame.DataFrame'>
  RangeIndex: 169 entries, 0 to 168
  Data columns (total 4 columns):
   #   Column    Non-Null Count  Dtype  
  ---  ------    --------------  -----  
   0   Duration  169 non-null    int64  
   1   Pulse     169 non-null    int64  
   2   Maxpulse  169 non-null    int64  
   3   Calories  164 non-null    float64
  dtypes: float64(1), int64(3)
  memory usage: 5.4 KB
  None
```



### Cleaning data

#### empty cells

```python
import pandas as pd

df = pd.read_csv('data.csv')

# remove rows that contains empty cells.
new_df = df.dropna()

print(new_df.to_string())

# change original DataFrame
df.dropna(inplace = True)

print(df.to_string())

# replace empty values
df.fillna(130, inplace = True)
```

#### Wrong and duplicated data

> wrong format

```python
import pandas as pd

df = pd.read_csv('data.csv')

# convert all cells in the 'Date' column into dates
df['Date'] = pd.to_datetime(df['Date'])

print(df.to_string())

# Remove rows with a NULL value in the "Date" column
df.dropna(subset=['Date'], inplace = True)
```

> wrong data

```python
# replace value
# Set "Duration" = 45 in row 7:
df.loc[7, 'Duration'] = 45

# Loop through all values in the "Duration" column.
# If the value is higher than 120, set it to 120:
for x in df.index:
  if df.loc[x, "Duration"] > 120:
    df.loc[x, "Duration"] = 120

# Delete rows where "Duration" is higher than 120:
for x in df.index:
  if df.loc[x, "Duration"] > 120:
    df.drop(x, inplace = True)
```

> duplicated data

```python
# Returns True for every row that is a duplicate, otherwise False:
print(df.duplicated())

# Removing Duplicates
df.drop_duplicates(inplace = True)
```

## Coding

### 列转行输出

```python
text='''A B C
18 29 33
F  M  F'''

rows = [[word for word in line.split()] for line in text.split('\n') ]
columns = [[col[i] for col in rows] for i in range(len(rows[0]))]
 
 for c in columns:
	print(" ".join(c) + "\t")
```

### 打印居中的*形成三角型

```python
In [24]: def print_center_tangle(n):
    ...:     max_width = 2*n -1
    ...:     for i in range(1,n+1):
    ...:         star_count = 2*i-1
    ...:         head_spaces = (max_width - star_count)//2
    ...:         print(" "*head_spaces,end="")
    ...:         print("*" * star_count)
```

```
In [25]: print_center_tangle(3)
  *
 ***
*****

In [26]: print_center_tangle(4)
   *
  ***
 *****
*******

In [27]: print_center_tangle(5)
    *
   ***
  *****
 *******
```

> 倒形三角型

```python
In [32]: def print_reversed_center_tangle(n):
    ...:     max_width = 2*n -1
    ...:     for i in range(n,0,-1):
    ...:         star_count = 2*i-1
    ...:         head_spaces = (max_width - star_count)//2
    ...:         print(" "*head_spaces,end="")
    ...:         print("*" * star_count)
```

### climb floors

```python
from functools import lru_cache

@lru_cache
def climb(n,steps):
    count = 0
    if n == 0 :
        count = 1
    elif n > 0 :
        for step in steps:
            count = count + climb(n-step,steps)
    return count
```



