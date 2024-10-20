# **Groovy**

## Keywords

### Reserved Keywords

| abstract     | assert     | break      | case       |
| ------------ | ---------- | ---------- | ---------- |
| catch        | class      | const      | continue   |
| def          | default    | do         | else       |
| enum         | extends    | final      | finally    |
| for          | goto       | if         | implements |
| import       | instanceof | interface  | native     |
| new          | null       | non-sealed | package    |
| public       | protected  | private    | return     |
| static       | strictfp   | super      | switch     |
| synchronized | this       | threadsafe | throw      |
| throws       | transient  | try        | while      |

## Identifiers

## Strings

```groovy
def number = 1 
def eagerGString = "value == ${number}"
def lazyGString = "value == ${ -> number }"

assert eagerGString == "value == 1" 
assert lazyGString ==  "value == 1" 

number = 2 
assert eagerGString == "value == 1" 
assert lazyGString ==  "value == 2" 
```



```groovy
String sample = "Hello world"; 
println(sample[4]); // Print the 5 character in the string

//Print the 1st character in the string starting from the back 
println(sample[-1]); 
println(sample[1..2]);//Prints a string starting from Index 1 to 2 
println(sample[4..2]);//Prints a string starting from Index 4 back to 2 


// Concatenation of Strings
String a = "Hello";
String b = "World";

println("Hello" + "World");
println(a + b);

// String Repetition
println("Hello"*3); 
println(a*3); 

// String Length
println(a.length());

// Regular Expressions
'Groovy' =~ 'Groovy' 
'Groovy' =~ 'oo' 
'Groovy' ==~ 'Groovy'   // fully matched
'Groovy' ==~ 'oo'       // not fully 
'Groovy' =~ '∧G' 
'Groovy' =~ 'G$' 
'Groovy' =~ 'Gro*vy' 
'Groovy' =~ 'Gro{2}vy'
```



| S.No. | Methods & Description                                        |
| ----- | ------------------------------------------------------------ |
| 1     | [center()](https://www.tutorialspoint.com/groovy/groovy_center.htm)Returns a new String of length numberOfChars consisting of the recipient padded on the left and right with space characters. |
| 2     | [compareToIgnoreCase()](https://www.tutorialspoint.com/groovy/groovy_comparetoignorecase.htm)Compares two strings lexicographically, ignoring case differences. |
| 3     | [concat()](https://www.tutorialspoint.com/groovy/groovy_concat.htm)Concatenates the specified String to the end of this String. |
| 4     | [eachMatch()](https://www.tutorialspoint.com/groovy/groovy_eachmatch.htm)Processes each regex group (see next section) matched substring of the given String. |
| 5     | [endsWith()](https://www.tutorialspoint.com/groovy/groovy_endswith.htm)Tests whether this string ends with the specified suffix. |
| 6     | [equalsIgnoreCase()](https://www.tutorialspoint.com/groovy/groovy_equalsignorecase.htm)Compares this String to another String, ignoring case considerations. |
| 7     | [getAt()](https://www.tutorialspoint.com/groovy/groovy_getat.htm)It returns string value at the index position |
| 8     | [indexOf()](https://www.tutorialspoint.com/groovy/groovy_indexof.htm)Returns the index within this String of the first occurrence of the specified substring. |
| 9     | [matches()](https://www.tutorialspoint.com/groovy/groovy_matches.htm)It outputs whether a String matches the given regular expression. |
| 10    | [minus()](https://www.tutorialspoint.com/groovy/groovy_minus.htm)Removes the value part of the String. |
| 11    | [next()](https://www.tutorialspoint.com/groovy/groovy_next.htm)This method is called by the ++ operator for the class String. It increments the last character in the given String. |
| 12    | [padLeft()](https://www.tutorialspoint.com/groovy/groovy_padleft.htm)Pad the String with the spaces appended to the left. |
| 13    | [padRight()](https://www.tutorialspoint.com/groovy/groovy_padright.htm)Pad the String with the spaces appended to the right. |
| 14    | [plus()](https://www.tutorialspoint.com/groovy/groovy_plus.htm)Appends a String |
| 15    | [previous()](https://www.tutorialspoint.com/groovy/groovy_previous.htm)This method is called by the -- operator for the CharSequence. |
| 16    | [replaceAll()](https://www.tutorialspoint.com/groovy/groovy_replaceall.htm)Replaces all occurrences of a captured group by the result of a closure on that text. |
| 17    | [reverse()](https://www.tutorialspoint.com/groovy/groovy_strings_reverse.htm)Creates a new String which is the reverse of this String. |
| 18    | [split()](https://www.tutorialspoint.com/groovy/groovy_split.htm)Splits this String around matches of the given regular expression. |
| 19    | [subString()](https://www.tutorialspoint.com/groovy/groovy_substring.htm)Returns a new String that is a substring of this String. |
| 20    | [toUpperCase()](https://www.tutorialspoint.com/groovy/groovy_touppercase.htm)Converts all of the characters in this String to upper case. |
| 21    | [toLowerCase()](https://www.tutorialspoint.com/groovy/groovy_tolowercase.htm)Converts all of the characters in this String to lower case. |



## ranges

- 1..10 - An example of an inclusive Range
- 1..<10 - An example of an exclusive Range
- ‘a’..’x’ – Ranges can also consist of characters
- 10..1 – Ranges can also be in descending order
- ‘x’..’a’ – Ranges can also consist of characters and be in descending order.



```groovy
def rint = 1..10; 

println(rint.contains(2)); 
println(rint.contains(11));
      
println(rint.get(2)); 
println(rint.get(4));


```



| Sr.No. | Methods & Description                                        |
| ------ | ------------------------------------------------------------ |
| 1      | [contains()](https://www.tutorialspoint.com/groovy/groovy_contains.htm)Checks if a range contains a specific value |
| 2      | [get()](https://www.tutorialspoint.com/groovy/groovy_get.htm)Returns the element at the specified position in this Range. |
| 3      | [getFrom()](https://www.tutorialspoint.com/groovy/groovy_getfrom.htm)Get the lower value of this Range. |
| 4      | [getTo()](https://www.tutorialspoint.com/groovy/groovy_getto.htm)Get the upper value of this Range. |
| 5      | [isReverse()](https://www.tutorialspoint.com/groovy/groovy_isreverse.htm)Is this a reversed Range, iterating backwards |
| 6      | [size()](https://www.tutorialspoint.com/groovy/groovy_size.htm)Returns the number of elements in this Range. |
| 7      | [subList()](https://www.tutorialspoint.com/groovy/groovy_sublist.htm)Returns a view of the portion of this Range between the specified fromIndex, inclusive, and toIndex, exclusive |

## Numbers

| byte   | -128 to 127                                              |
| ------ | -------------------------------------------------------- |
| short  | -32,768 to 32,767                                        |
| int    | -2,147,483,648 to 2,147,483,647                          |
| long   | -9,223,372,036,854,775,808 to +9,223,372,036,854,775,807 |
| float  | 1.40129846432481707e-45 to 3.40282346638528860e+38       |
| double | 4.94065645841246544e-324d to 1.79769313486231570e+308d   |

### Integral  literals

- `byte`
- `char`
- `short`
- `int`
- `long`
- `java.math.BigInteger`

### Decimal literals

- `float`
- `double`
- `java.math.BigDecimal`



### Math Operation 

Division and power binary operations aside (covered below),

- binary operations between `byte`, `char`, `short` and `int` result in `int`
- binary operations involving `long` with `byte`, `char`, `short` and `int` result in `long`
- binary operations involving `BigInteger` and any other integral type result in `BigInteger`
- binary operations involving `BigDecimal` with `byte`, `char`, `short`, `int` and `BigInteger` result in `BigDecimal`
- binary operations between `float`, `double` and `BigDecimal` result in `double`
- binary operations between two `BigDecimal` result in `BigDecimal`

|                | byte | char | short | int  | long | BigInteger | float  | double | BigDecimal |
| :------------- | :--- | :--- | :---- | :--- | :--- | :--------- | :----- | :----- | :--------- |
| **byte**       | int  | int  | int   | int  | long | BigInteger | double | double | BigDecimal |
| **char**       |      | int  | int   | int  | long | BigInteger | double | double | BigDecimal |
| **short**      |      |      | int   | int  | long | BigInteger | double | double | BigDecimal |
| **int**        |      |      |       | int  | long | BigInteger | double | double | BigDecimal |
| **long**       |      |      |       |      | long | BigInteger | double | double | BigDecimal |
| **BigInteger** |      |      |       |      |      | BigInteger | double | double | BigDecimal |
| **float**      |      |      |       |      |      |            | double | double | double     |
| **double**     |      |      |       |      |      |            |        | double | double     |
| **BigDecimal** |      |      |       |      |      |            |        |        | BigDecimal |

### Number Methods

| S.No. | Methods & Description                                        |
| ----- | ------------------------------------------------------------ |
| 1     | [xxxValue()](https://www.tutorialspoint.com/groovy/groovy_xxxvalue.htm)This method takes on the Number as the parameter and returns a primitive type based on the method which is invoked. |
| 2     | [compareTo()](https://www.tutorialspoint.com/groovy/groovy_compareto.htm)The compareTo method is to use compare one number against another. This is useful if you want to compare the value of numbers. |
| 3     | [equals()](https://www.tutorialspoint.com/groovy/groovy_equals.htm)The method determines whether the Number object that invokes the method is equal to the object that is passed as argument. |
| 4     | [valueOf()](https://www.tutorialspoint.com/groovy/groovy_valueof.htm)The valueOf method returns the relevant Number Object holding the value of the argument passed. |
| 5     | [toString()](https://www.tutorialspoint.com/groovy/groovy_tostring.htm)The method is used to get a String object representing the value of the Number Object. |
| 6     | [parseInt()](https://www.tutorialspoint.com/groovy/groovy_parseint.htm)This method is used to get the primitive data type of a certain String. parseXxx() is a static method and can have one argument or two. |
| 7     | [abs()](https://www.tutorialspoint.com/groovy/groovy_abs.htm)The method gives the absolute value of the argument. The argument can be int, float, long, double, short, byte. |
| 8     | [ceil()](https://www.tutorialspoint.com/groovy/groovy_ceil.htm)The method ceil gives the smallest integer that is greater than or equal to the argument. |
| 9     | [floor()](https://www.tutorialspoint.com/groovy/groovy_floor.htm)The method floor gives the largest integer that is less than or equal to the argument. |
| 10    | [rint()](https://www.tutorialspoint.com/groovy/groovy_rint.htm)The method rint returns the integer that is closest in value to the argument. |
| 11    | [round()](https://www.tutorialspoint.com/groovy/groovy_round.htm)The method round returns the closest long or int, as given by the methods return type. |
| 12    | [min()](https://www.tutorialspoint.com/groovy/groovy_min.htm)The method gives the smaller of the two arguments. The argument can be int, float, long, double. |
| 13    | [max()](https://www.tutorialspoint.com/groovy/groovy_max.htm)The method gives the maximum of the two arguments. The argument can be int, float, long, double. |
| 14    | [exp()](https://www.tutorialspoint.com/groovy/groovy_exp.htm)The method returns the base of the natural logarithms, e, to the power of the argument. |
| 15    | [log()](https://www.tutorialspoint.com/groovy/groovy_log.htm)The method returns the natural logarithm of the argument. |
| 16    | [pow()](https://www.tutorialspoint.com/groovy/groovy_pow.htm)The method returns the value of the first argument raised to the power of the second argument. |
| 17    | [sqrt()](https://www.tutorialspoint.com/groovy/groovy_sqrt.htm)The method returns the square root of the argument. |
| 18    | [sin()](https://www.tutorialspoint.com/groovy/groovy_sin.htm)The method returns the sine of the specified double value. |
| 19    | [cos()](https://www.tutorialspoint.com/groovy/groovy_cos.htm)The method returns the cosine of the specified double value. |
| 20    | [tan()](https://www.tutorialspoint.com/groovy/groovy_tan.htm)The method returns the tangent of the specified double value. |
| 21    | [asin()](https://www.tutorialspoint.com/groovy/groovy_asin.htm)The method returns the arcsine of the specified double value. |
| 22    | [acos()](https://www.tutorialspoint.com/groovy/groovy_acos.htm)The method returns the arccosine of the specified double value. |
| 23    | [atan()](https://www.tutorialspoint.com/groovy/groovy_atan.htm)The method returns the arctangent of the specified double value. |
| 24    | [atan2()](https://www.tutorialspoint.com/groovy/groovy_atan2.htm)The method Converts rectangular coordinates (x, y) to polar coordinate (r, theta) and returns theta. |
| 25    | [toDegrees()](https://www.tutorialspoint.com/groovy/groovy_numbers_todegrees.htm)The method converts the argument value to degrees. |
| 26    | [radian()](https://www.tutorialspoint.com/groovy/groovy_radian.htm)The method converts the argument value to radians. |
| 27    | [random()](https://www.tutorialspoint.com/groovy/groovy_random.htm)The method is used to generate a random number between 0.0 and 1.0. The range is: 0.0 =< Math.random < 1.0. Different ranges can be achieved by using arithmetic. |



## Operators

### Overview

- Arithmetic operators

  | Operator | Description                                                  | Example                     |
  | -------- | ------------------------------------------------------------ | --------------------------- |
  | +        | Addition of two operands                                     | 1 + 2 will give 3           |
  | −        | Subtracts second operand from the first                      | 2 − 1 will give 1           |
  | *        | Multiplication of both operands                              | 2 * 2 will give 4           |
  | /        | Division of numerator by denominator                         | 3 / 2 will give 1.5         |
  | %        | Modulus Operator and remainder of after an integer/float division | 3 % 2 will give 1           |
  | ++       | Incremental operators used to increment the value of an operand by 1 | int x = 5;x++;x will give 6 |
  | --       | Incremental operators used to decrement the value of an operand by 1 | int x = 5;x--;x will give 4 |

- Relational operators

  | Operator | Description                                                  | Example               |
  | -------- | ------------------------------------------------------------ | --------------------- |
  | ==       | Tests the equality between two objects                       | 2 == 2 will give true |
  | !=       | Tests the difference between two objects                     | 3 != 2 will give true |
  | <        | Checks to see if the left objects is less than the right operand. | 2 < 3 will give true  |
  | <=       | Checks to see if the left objects is less than or equal to the right operand. | 2 <= 3 will give true |
  | >        | Checks to see if the left objects is greater than the right operand. | 3 > 2 will give true  |
  | >=       | Checks to see if the left objects is greater than or equal to the right operand. | 3 >= 2 will give true |

- Logical operators

  | Operator | Description                        | Example                       |
  | -------- | ---------------------------------- | ----------------------------- |
  | &&       | This is the logical “and” operator | true && true will give true   |
  | \|\|     | This is the logical “or” operator  | true \|\| true will give true |
  | !        | This is the logical “not” operator | !false will give true         |

- Bitwise operators

  | Sr.No | Operator & Description                                  |
  | ----- | ------------------------------------------------------- |
  | 1     | **&**This is the bitwise “and” operator                 |
  | 2     | **\|**This is the bitwise “or” operator                 |
  | 3     | **^**This is the bitwise “xor” or Exclusive or operator |
  | 4     | **~**This is the bitwise negation operator              |

- Assignment operators

  | Operator | Description                                                  | Example                        |
  | -------- | ------------------------------------------------------------ | ------------------------------ |
  | +=       | This adds right operand to the left operand and assigns the result to left operand. | def A = 5A+=3Output will be 8  |
  | -=       | This subtracts right operand from the left operand and assigns the result to left operand | def A = 5A-=3Output will be 2  |
  | *=       | This multiplies right operand with the left operand and assigns the result to left operand | def A = 5A*=3Output will be 15 |
  | /=       | This divides left operand with the right operand and assigns the result to left operand | def A = 6A/=3Output will be 2  |
  | %=       | This takes modulus using two operands and assigns the result to left operand | def A = 5A%=3Output will be 2  |

- Range Operators

  ```groovy
  def range = 0..5 
  
  class Example { 
     static void main(String[] args) { 
        def range = 5..10; 
        println(range); 
        println(range.get(2)); 
     } 
  }
  ```

### Operator Precedence

| Sr.No | Operators & Names                                            |
| ----- | ------------------------------------------------------------ |
| 1     | **++ -- + -**pre increment/decrement, unary plus, unary minus |
| 2     | *** / %**multiply, div, modulo                               |
| 3     | **+ -**addition, subtraction                                 |
| 4     | **== != <=>**equals, not equals, compare to                  |
| 5     | **&**binary/bitwise and                                      |
| 6     | **^**binary/bitwise xor                                      |
| 7     | **\|**binary/bitwise or                                      |
| 8     | **&&**logical and                                            |
| 9     | **\|\|**logical or                                           |
| 10    | **= \**= \*= /= %= += -= <<= >>= >>>= &= ^= \|=**Various assignment operators |

## List

```groovy
def numbers = [1, 2, 3]         

assert numbers instanceof List  
assert numbers.size() == 3      

// In the above example, we used a homogeneous list, but you can also create lists containing values of heterogeneous types:
def heterogeneous = [1, "a", true] 

// As lists can be heterogeneous in nature, lists can also contain other lists to create multidimensional lists:
def multi = [[0, 1], [2, 3]]     
assert multi[1][0] == 2          
```

addAll :  add all elements from one collection to another

| 1    | [add()](https://www.tutorialspoint.com/groovy/groovy_add.htm)Append the new value to the end of this List.    add(int index, Object value)   Append the new value to a particular position in the List |
| ---- | ------------------------------------------------------------ |
| 2    | [contains()](https://www.tutorialspoint.com/groovy/groovy_lists_contains.htm)Returns true if this List contains the specified value. |
| 3    | [get()](https://www.tutorialspoint.com/groovy/groovy_lists_get.htm)Returns the element at the specified position in this List. |
| 4    | [isEmpty()](https://www.tutorialspoint.com/groovy/groovy_isempty.htm)Returns true if this List contains no elements |
| 5    | [minus()](https://www.tutorialspoint.com/groovy/groovy_lists_minus.htm)Creates a new List composed of the elements of the original without those specified in the collection. |
| 6    | [plus()](https://www.tutorialspoint.com/groovy/groovy_lists_plus.htm)Creates a new List composed of the elements of the original together with those specified in the collection. |
| 7    | [pop()](https://www.tutorialspoint.com/groovy/groovy_pop.htm)Removes the last item from this List |
| 8    | [remove()](https://www.tutorialspoint.com/groovy/groovy_remove.htm)Removes the element at the specified position in this List. |
| 9    | [reverse()](https://www.tutorialspoint.com/groovy/groovy_reverse.htm)Create a new List that is the reverse the elements of the original List |
| 10   | [size()](https://www.tutorialspoint.com/groovy/groovy_lists_size.htm)Obtains the number of elements in this List. |
| 11   | [sort()](https://www.tutorialspoint.com/groovy/groovy_sort.htm)Returns a sorted copy of the original List. |

## Arrays

> Groovy reuses the list notation for arrays, but to make such literals arrays, you need to explicitly define the type of the array through coercion or type declaration.

```groovy
String[] arrStr = ['Ananas', 'Banana', 'Kiwi'] 
assert arrStr instanceof String[]    
assert !(arrStr instanceof List)

def numArr = [1, 2, 3] as int[]      
assert numArr instanceof int[]       
assert numArr.size() == 3
```



## Maps

```groovy
def colors = [red: '#FF0000', green: '#00FF00', blue: '#0000FF']   

assert colors['red'] == '#FF0000'    
assert colors.green  == '#00FF00'    

colors['pink'] = '#FF00FF'           
colors.yellow  = '#FFFF00'           

assert colors.pink == '#FF00FF'
assert colors['yellow'] == '#FFFF00'
assert colors instanceof java.util.LinkedHashMap

// use key var 
def key = 'name'
person = [(key): 'Guillaume']        

assert person.containsKey('name')    
assert !person.containsKey('key')    

// loop
for (entry in colors) {
    println "Hex Code: $entry.key = Color Name: $entry.value"
}
```

| 1    | [containsKey()](https://www.tutorialspoint.com/groovy/groovy_containskey.htm)Does this Map contain this key? |
| ---- | ------------------------------------------------------------ |
| 2    | [get()](https://www.tutorialspoint.com/groovy/groovy_maps_get.htm)Look up the key in this Map and return the corresponding value. If there is no entry in this Map for the key, then return null. |
| 3    | [keySet()](https://www.tutorialspoint.com/groovy/groovy_keyset.htm)Obtain a Set of the keys in this Map. |
| 4    | [put()](https://www.tutorialspoint.com/groovy/groovy_put.htm)Associates the specified value with the specified key in this Map. If this Map previously contained a mapping for this key, the old value is replaced by the specified value. |
| 5    | [size()](https://www.tutorialspoint.com/groovy/groovy_maps_size.htm)Returns the number of key-value mappings in this Map. |
| 6    | [values()](https://www.tutorialspoint.com/groovy/groovy_values.htm)Returns a collection view of the values contained in this Map. |

https://www.jenkins.io/doc/pipeline/steps/workflow-cps/



https://www.jenkins.io/doc/book/pipeline/shared-libraries/	





## Loops

```groovy
def count=0
while(count<5) {
 println(count);
 count++;
}

for(int i = 0;i<5;i++) {
    println(i);
}

for(int i in array) { 
    println(i); 
} 
```

```groovy
for(String str : list) {
    println(str);
}
```

```groovy
for (entry in map) {
    println "Hex Code: $entry.key = Color Name: $entry.value"
}
```

## Decision Making

### if/else

```groovy
int a = 2

//Check for the boolean condition 
if (a<100) { 
    //If the condition is true print the following statement 
    println("The value is less than 100"); 
} else { 
    //If the condition is false print the following statement 
    println("The value is greater than 100"); 
} 
```

### switch

```groovy
int a = 2

//Evaluating the expression value 
switch(a) {            
    //There is case statement defined for 4 cases 
    // Each case statement section has a break condition to exit the loop 

    case 1: 
    println("The value of a is One"); 
    break; 
    case 2: 
    println("The value of a is Two"); 
    break; 
    case 3: 
    println("The value of a is Three"); 
    break; 
    case 4: 
    println("The value of a is Four"); 
    break; 
    default: 
    println("The value is unknown"); 
    break; 
}
```



## Closures

> in methods

```groovy
def str1 = "Hello";
def clos = { param -> println "${str1} ${param}" }
clos.call("World");
```



> each Iteration

```groovy
map.each { println "Hex Code: $it.key = Color Name: $it.value" }

map.each { entry -> println "Hex Code: $entry.key = Color Name: $entry.value" }

map.each { key, val ->
    println "Hex Code: $key = Color Name $val"
}

def mp = ["TopicName" : "Maps", "TopicDescription" : "Methods in Maps"]             
mp.each {println it}
mp.each {println "${it.key} maps to: ${it.value}"}
```

> eachWithIndex

```groovy
map.eachWithIndex { entry, index ->
    def indent = ((index == 0 || index % 2 == 0) ? "   " : "")
    println "$index Hex Code: $entry.key = Color Name: $entry.value"
}

map.eachWithIndex { key, val, index ->
    def indent = ((index == 0 || index % 2 == 0) ? "   " : "")
    println "$index Hex Code: $key = Color Name: $val"
}
```



> takewhile 和 dropwhile  (出现第一个不满足条件就终止并返回)

```groovy
list = [1,2,3]
b = list.takeWhile {it < 2 } 

a = list.takeWhile {num -> num < 2 } 
println(a)   # [1]

b = list.dropWhile {num -> num < 2 } 
println(b)   # # [2, 3]
```



```groovy
def lst = [1,2,3,4];
def value;

value = lst.find {element -> element > 2}
println(value);    // 3 

		
// Is there any value above 2
value = lst.any{element -> element > 2}
println(value);

// Is there any value above 4
value = lst.any{element -> element > 4}
println(value);
```



> **`collect`**  vs  **`collectEntries`**
>
> **`collect`**: Use this when you want to transform each element of a collection into a new form and get a list of the results.
>
> **`collectEntries`**: Use this when you want to transform each element into key-value pairs and get a map of these pairs.

```groovy
def lst = [1,2,3,4];
def newlst = [];
newlst = lst.collect {element -> return element * element}
println(newlst);
      
def serializableData = jsonData.collectEntries { [(it.key): it.value] }
echo "Serializable Data: ${serializableData}"
```



> list (每两两一对)-> map 

```groovy
(1..10).collate(2).collectEntries {
	k,v -> [(k):v]
	//// or 
	// [(it[0]):it[1]]
}    // ==> result in map  [1:2, 3:4, 5:6, 7:8, 9:10]
```



| Sr.No. |                    Methods & Description                     |
| :----- | :----------------------------------------------------------: |
| 1      | [find()](https://www.tutorialspoint.com/groovy/groovy_find.htm)The find method finds the first value in a collection that matches some criterion. |
| 2      | [findAll()](https://www.tutorialspoint.com/groovy/groovy_findall.htm)It finds all values in the receiving object matching the closure condition. |
| 3      | [any() & every()](https://www.tutorialspoint.com/groovy/groovy_any_every.htm)Method any iterates through each element of a collection checking whether a Boolean predicate is valid for at least one element. |
| 4      | [collect()](https://www.tutorialspoint.com/groovy/groovy_collect.htm)The method collect iterates through a collection, converting each element into a new value using the closure as the transformer. |

> find  & findAll  is similar to filter in Python
>
> collect is similar to map function in Python

## Exception

```groovy
try {
    def arr = new int[3];
    arr[5] = 5;
} catch (ArrayIndexOutOfBoundsException ex) {
    println(ex.toString());
    println(ex.getMessage());
    println(ex.getStackTrace());  
} catch (Exception ex) {
    println("Catching the exception");
} finally {
    println("The final block");
}
```



## Json

**JsonSlurper**: JsonSlurper is a class that parses JSON text or reader content into Groovy data

>  [Jenkins Pipeline NotSerializableException: groovy.json.internal.LazyMap](https://stackoverflow.com/questions/37864542/jenkins-pipeline-notserializableexception-groovy-json-internal-lazymap)
>
> - use JsonSlurperClassic instead of JsonSlurper 
>
> - def ret = jsonSlurper.parseText(teststr)   (`def`  omitting will raise above error)
>
> - covert to hash map 
>
>   ```groovy
>   def jsonData = jsonSlurper.parseText(jsonString)
>   
>   // 转换为可序列化的 HashMap
>   def serializableData = jsonData.collectEntries { [(it.key): it.value] }
>   echo "Serializable Data: ${serializableData}"
>   ```
>
>   

```groovy
import groovy.json.JsonSlurper 
class Example {

   static void main(String[] args) {
      def jsonSlurper = new JsonSlurper()
      def obj = jsonSlurper.parseText ''' {"Integer": 12, "fraction": 12.55, "double": 12e13}'''
		
      println(obj.Integer);
      println(obj.fraction);
      println(obj.double); 
   } 
}
```



**JsonOutput** : This method is responsible for serialising Groovy objects into JSON strings.

```groovy
import groovy.json.JsonOutput 
class Example {
   static void main(String[] args) {
      def output = JsonOutput.toJson([name: 'John', ID: 1])
      println(output);  
   }
}
```



## Cases in pipeline

### return this

**`return this`**：虽然没有显式定义类，但由于 Jenkins 将 `vars/mySharedLib.groovy` 视为一个对象，你可以通过 `return this` 返回当前脚本实例，并继续调用其他方法。

**链式调用**：通过 `mySharedLib('Jenkins').anotherMethod()`，可以在调用 `call()` 方法之后继续调用 `anotherMethod()`，实现链式调用的效果。



```groovy
def call(String message) {
    echo "Hello, ${message}!"
    return this // 返回当前脚本实例，实现链式调用
}

def anotherMethod() {
    echo "This is another method!"
    return this // 返回当前脚本实例，支持继续链式调用
}
```

**Jenkinsfile**：

```groovy
@Library('my-shared-library') _

pipeline {
    agent any
    stages {
        stage('Test') {
            steps {
                script {
                    // 使用链式调用
                    mySharedLib('Jenkins').anotherMethod()
                }
            }
        }
    }
}
```

### load vs library

**`load vs libraray`**

> Load: Load individual scripts from workspace
>
> limit:  Modular, reusable code across jobs

```groov
// @Library annotation cannot evaluate dynamic content or variables.
// def branchName = 'feature-branch'
// @Library("my-shared-library@$branchName") _

properties([parameters([string(name: 'LIB_VERSION', defaultValue: 'master')])])
library "my-shared-library@${params.LIB_VERSION}"

script {
    def lib = library("my-shared-library@$branchName")
    lib.someFunction()
}
```



### Exception handling

**`catchError`**: 适用于你希望 Pipeline 继续执行，即使某个步骤失败。它捕获错误，但不会终止 Pipeline。

**`try-catch`**: 提供更灵活的错误处理机制，允许你根据需要决定是否继续或停止 Pipeline。可以捕获异常并根据情况处理，但如果使用 `error` 步骤，会终止 Pipeline。

This step `catchError` is most useful when used in Declarative Pipeline or with the options to set the stage result or ignore build interruptions. Otherwise, consider using plain `try`-`catch`(-`finally`) blocks.



### Elvis 操作符

```groovy
jsonPayload = null
def repoName = jsonPayload?.repository?.name ?: "unknown"
println(repoName)  // 输出: unknown
```



### split vs tokenize

```groovy
def str = "a,,b,c"
println str.tokenize(",")  // 输出: [a, b, c] （忽略了空值）
println str.split(",")     // 输出: [a, , b, c] （包含空值）
```

### add (<<)

`keys << path` 语法涉及向 `keys` 集合中添加 `path` 值。它使用了 Groovy 的操作符 `<<`，这是一个快捷方式，等效于调用 `add()` 方法，常用于添加元素到集合或列表中。

```groovy
def keys = []
def path = "/some/path"
keys << path
println keys  // 输出: [/some/path]
```



