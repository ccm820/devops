

# Shell

## 内置命令

- `awk`：用于模式扫描和处理的强大文本处理工具。
- `sed`：文本转换的流编辑器。
- `cut`：从输入中提取特定列或字段。
- `sort`：对文本文件的行进行排序。
- `uniq`：从排序文件中删除重复行。

### awk

![image-20240417235040994](C:\Users\Caiman\Documents\Docs\Shell.assets\image-20240417235040994.png)

![image-20240417235052046](C:\Users\Caiman\Documents\Docs\Shell.assets\image-20240417235052046.png)



> ARGC               命令行参数个数
> ARGV               命令行参数排列
> ENVIRON            支持队列中系统环境变量的使用
> FILENAME           awk浏览的文件名
> FNR                浏览文件的记录数
> FS                 设置输入域分隔符，等价于命令行 -F选项
> NF                 浏览记录的域的个数
> NR                 已读的记录数
> OFS                输出域分隔符
> ORS                输出记录分隔符
> RS                 控制记录分隔符

AWK数组
For (element in array ) print array[element]





> 列转成行

```shell
awk '{ for (i = 1; i <= NF; i++) a[i,NR] = $i }
     NF > p { p = NF }
     END { 
         for(j = 1; j <= p; j++) {
             str = a[j,1]
             for(i = 2; i <= NR; i++) {
                 str = str " " a[j,i]
             }
             print str
         }
     }' 1.txt

awk '{ for (i = 1; i <= NF; i++) a[i,NR] = $i }
     NF > p { p = NF }
     END { 
         for(j = 1; j <= p; j++) {
             for(k = 1; k <= NR; k++) {
                 printf "%s ",a[j,k]
             }
             print ""
         }
     }' 1.txt

awk '{for(i=1;i<=NF;i++) a[i,NR]=$i} NF>p{p=NF} END{for(j=1;j<=p;j++){ for(k=1;k<=NR;k++) {printf "%s ",a[j,k]} print ""}}' 1.txt
```

```shell
#!/usr/bin/bash

declare -A rows
while read -a line; do
    mylen=${#line[@]}
    echo "len:",$mylen
    ((i++))
    for j in `seq $mylen` ; do
        rows[$i,$j]=${line[$((j-1))]}
    done
done < 2.txt
echo $i $mylen
echo ${rows[*]}

printf "%20s\n"  | tr " " "#"
for m in `seq $mylen`;do
    for n in `seq $i`;do
        printf "%s\t" ${rows[$n,$m]}
    done
    printf "\n"
done
```

```shell
#!/usr/bin/bash
###################################################################
declare -A rows 
declare -A widths

while read -a fields 
do 
    ((i++))
    fnum=${#fields[@]}
    for j in `seq $fnum`
    do
        fld=${fields[$((j-1))]}
        rows[$i,$j]=$fld
        if (( ${#fld} > ${widths[$i]:-0} ))
        then 
            widths[$i]=${#fld}
        fi 
    done 
done < $1

for m in `seq $fnum`
do
    for n in `seq $i`
    do
        printf "%-${widths[$n]}s " ${rows[$n,$m]}
    done 
    printf "\n"
done
```



> 去重（除了用sort -u外）

```shell
awk '!a[$0]++' file
```

> 两个文件根据某个相同字段，求差集

```shell
awk -F","  'FNR==NR{a[$1]=1;next}{if(!(a[$1])) print $0}'  PAY_009_SGW000_020_20140725_dz_ocs PAY_009_SGW000_020_20140725.3001 | more
```



> 传入参数

```shell
find . -name "RBI_Z31_`date +'%Y%m%d'`*" | xargs cat | awk -F"|" -v OFS="|" -v ZeroTime="`date +'%Y%m%d'`000000" '{
                if( ($15=="0" && $11 != "5") || ($15 == "1" && $11 != "5" && $7 > ZeroTime)){
                        if($14 == "769"){
                                print $4,$14,$16 >> "unactive_cbp3.unl"
                        }else if($14 == "200" || $14 == "757"){
                          print $4,$14,$16 >> "unactive_cbp2.unl"
                        }else if($14 == "753" || $14 == "755" || $14 == "752"){
                          print $4,$14,$16 >> "unactive_cbp4.unl"
                        }else{
                                print $4,$14,$16 >> "unactive_cbp1.unl"
                        }
                }
        }' 
```



> 统计

```shell
ls | xargs cat | awk -F'|' 'BEGIN{total=0;}{if($10==757 && ($11==506262 || $11==506772)) total=total+$7;} END{printf " RETURN: %d\n",total}'
```



### sed

```
sed [-hnV][-e<script>][-f<script文件>][文本文件]
```

- a ：新增， a 的后面可以接字串，而这些字串会在新的一行出现(目前的下一行)～
- c ：取代， c 的后面可以接字串，这些字串可以取代 n1,n2 之间的行！
- d ：删除，因为是删除啊，所以 d 后面通常不接任何东东；
- i ：插入， i 的后面可以接字串，而这些字串会在新的一行出现(目前的上一行)；
- p ：打印，亦即将某个选择的数据印出。通常 p 会与参数 sed -n 一起运行～
- s ：取代，可以直接进行取代的工作哩！通常这个 s 的动作可以搭配正则表达式！例如 1,20s/old/new/g 就是啦！

```shell
#在 testfile 文件的第四行后添加一行，并将结果输出到标准输出，在命令行提示符下输入如下命令：
sed -e 4a\newLine testfile 

# 将 testfile 的内容列出并且列印行号，同时，请将第 2~5 行删除！
nl testfile | sed '2,5d'

# 要删除第 3 到最后一行：
nl testfile | sed '3,$d'

# 在第二行后(即加在第三行) 加上drink tea
nl testfile | sed '2a drink tea'
# 如果是要在第二行前
nl testfile | sed '2i drink tea' 
 
# 以行为单位的替换与显示
# 将第 2-5 行的内容取代成为 No 2-5 number
nl testfile | sed '2,5c No 2-5 number'

# 数据的搜寻并显示
# 搜索 testfile 有 oo 关键字的行:
nl testfile | sed -n '/oo/p'

# 数据的搜寻并删除
# 删除 testfile 所有包含 oo 的行，其他行输出
nl testfile | sed  '/oo/d'

# 数据的搜寻并执行命令
# 搜索 testfile，找到 oo 对应的行，执行后面花括号中的一组命令，每个命令之间用分号分隔，这里把 oo 替换为 kk，再输出这行：
nl testfile | sed -n '/oo/{s/oo/kk/;p;q}'  

# 数据的查找与替换
# 将 testfile 文件中每行第一次出现的 oo 用字符串 kk 替换，然后将该文件内容输出到标准输出:
sed -e 's/oo/kk/' testfile
# g 标识符表示全局查找替换，使 sed 对文件中所有符合的字符串都被替换，修改后内容会到标准输出，不会修改原文件(用-i就会修改原文件)
sed -e 's/oo/kk/g' testfile
#多点编辑
# 一条 sed 命令，删除 testfile 第三行到末尾的数据，并把 HELLO 替换为 RUNOOB :
nl testfile | sed -e '3,$d' -e 's/HELLO/RUNOOB/'

# 直接修改文件内容
# 利用 sed 直接在 regular_express.txt 最后一行加入 # This is a test:
sed -i '$a # This is a test' regular_express.txt
```

### cut

```
cut  [-bn] [file]
cut [-c] [file]
cut [-df] [file]
```

- -b ：以字节为单位进行分割。这些字节位置将忽略多字节字符边界，除非也指定了 -n 标志。
- -c ：以字符为单位进行分割。
- -d ：自定义分隔符，默认为制表符。
- -f ：与-d一起使用，指定显示哪个区域。
- -n ：取消分割多字节字符。仅和 -b 标志一起使用。如果字符的最后一个字节落在由 -b 标志的 List 参数指示的
  范围之内，该字符将被写出；否则，该字符将被排除



### sort

```
sort [-bcdfimMnr][-o<输出文件>][-t<分隔字符>][+<起始栏位>-<结束栏位>][--help][--verison][文件][-k field1[,field2]]
```

- -b 忽略每行前面开始出的空格字符。
- -c 检查文件是否已经按照顺序排序。
- -d 排序时，处理英文字母、数字及空格字符外，忽略其他的字符。
- -f 排序时，将小写字母视为大写字母。
- -i 排序时，除了040至176之间的ASCII字符外，忽略其他的字符。
- -m 将几个排序好的文件进行合并。
- -M 将前面3个字母依照月份的缩写进行排序。
- -n 依照数值的大小排序。
- -u 意味着是唯一的(unique)，输出的结果是去完重了的。
- -o<输出文件> 将排序后的结果存入指定的文件。
- -r 以相反的顺序来排序。
- -t<分隔字符> 指定排序时所用的栏位分隔字符。
- +<起始栏位>-<结束栏位> 以指定的栏位来排序，范围由起始栏位到结束栏位的前一栏位。
- --help 显示帮助。
- --version 显示版本信息。
- [-k field1[,field2]] 按指定的列进行排序。

```
sort -n -k 2 -t : facebook.txt
```

apple:10:2.5
orange:20:3.4
banana:30:5.5
pear:90:2.3
我们使用冒号作为间隔符，并针对第二列来进行数值升序排序，结果很令人满意。



其他的sort常用选项

-f会将小写字母都转换为大写字母来进行比较，亦即忽略大小写
-c会检查文件是否已排好序，如果乱序，则输出第一个乱序的行的相关信息，最后返回1
-C会检查文件是否已排好序，如果乱序，不输出内容，仅返回1
-M会以月份来排序，比如JAN小于FEB等等
-b会忽略每一行前面的所有空白部分，从第一个可见字符开始比较。



### read

- -a 后跟一个变量，该变量会被认为是个数组，然后给其赋值，默认是以空格为分割符。
- -d 后面跟一个标志符，其实只有其后的第一个字符有用，作为结束的标志。
- -p 后面跟提示信息，即在输入前打印提示信息。
- -e 在输入的时候可以使用命令补全功能。
- -n 后跟一个数字，定义输入文本的长度，很实用。
- -r 屏蔽\，如果没有该选项，则\作为一个转义字符，有的话 \就是个正常的字符了。
- -s 安静模式，在输入字符时不再屏幕上显示，例如login时输入密码。
- -t 后面跟秒数，定义输入字符的等待时间。
- -u 后面跟fd，从文件描述符中读入，该文件描述符可以是exec新开启的。

```shell
# 1、简单读取
#这里默认会换行  
echo "输入网站名: "  
#读取从键盘的输入  
read website  
echo "你输入的网站名是 $website"  

# 2、-p 参数，允许在 read 命令行中直接指定一个提示。
read -p "输入网站名:" website
echo "你输入的网站名是 $website" 

# 3. -t 参数指定 read 命令等待输入的秒数，当计时满时，read命令返回一个非零退出状态
if read -t 5 -p "输入网站名:" website
then
    echo "你输入的网站名是 $website"
else
    echo "\n抱歉，你输入超时了。"
fi

#4、除了输入时间计时，还可以使用 -n 参数设置 read 命令计数输入的字符。当输入的字符数目达到预定数目时，自动退出，并将输入的数据赋值给变量
read -n1 -p "Do you want to continue [Y/N]?" answer
case $answer in
Y | y)
      echo "fine ,continue";;
N | n)
      echo "ok,good bye";;
*)
     echo "error choice";;
esac


#5. -s 选项能够使 read 命令中输入的数据不显示在命令终端上（实际上，数据是显示的，只是 read 命令将文本颜色设置成与背景相同的颜色）
read  -s  -p "请输入您的密码:" pass
echo "\n您输入的密码是 $pass"

#6.读取文件
count=1    # 赋值语句，不加空格
cat test.txt | while read line      # cat 命令的输出作为read命令的输入,read读到>的值放在line中
do
   echo "Line $count:$line"
   count=$[ $count + 1 ]          # 注意中括号中的空格。
done
echo "finish"
```



### printf

**%s %c %d %f** 都是格式替代符，**％s** 输出一个字符串，**％d** 整型输出，**％c** 输出一个字符，**％f** 输出实数，以小数形式输出。

**%-10s** 指一个宽度为 10 个字符（**-** 表示左对齐，没有则表示右对齐），任何字符都会被显示在 10 个字符宽的字符内，如果不足则自动以空格填充，超过也会将内容全部显示出来。

**%-4.2f** 指格式化为小数，其中 **.2** 指保留2位小数。



```shell
# format-string为双引号
printf "%d %s\n" 1 "abc"

# 单引号与双引号效果一样
printf '%d %s\n' 1 "abc"

# 没有引号也可以输出
printf %s abcdef

# 格式只指定了一个参数，但多出的参数仍然会按照该格式输出，format-string 被重用
printf %s abc def

printf "%s\n" abc def

printf "%s %s %s\n" a b c d e f g h i j

# 如果没有 arguments，那么 %s 用NULL代替，%d 用 0 代替
printf "%s and %d \n"
```

### seq

```
       seq [OPTION]... LAST
       seq [OPTION]... FIRST LAST
       seq [OPTION]... FIRST INCREMENT LAST
```

-f, --format=FORMAT   use printf style floating-point FORMAT (default: %g)
-s, --separator=STRING  use STRING to separate numbers (default: \n)
-w, --equal-width    equalize width by padding with leading zeroes

```shell
seq -w 98 101
#098
#099
#100
#101

for i in $(seq 1 10)
#or 
for i in `seq 1 10`;
for i in `seq 10`
```

###  Alias

In Linux, **an alias is a shortcut** that references a command.

The `chmod` command in Linux is **used to modify the permissions and access mode of files and directories**. 

`SSH port forwarding`, also known as SSH tunneling, `is a technique used to securely transmit data between a local computer and a remote server over an encrypted SSH connection`. It allows you to establish a secure communication channel between two machines, forwarding traffic from a specific port on one machine to another port on the other machine through the SSH connection.

 A `zombie process`, also known as a *defunct process*, is a terminated process that has not been fully removed from the process table.

it's essential for the parent process to properly handle the termination of its child processes.

### set

set -x 　执行指令后，会先显示该指令及所下的参数。

set -e    如果任何语句的执行结果不是true则应该退出。如果要增加可读性，可以使用set -o errexit，它的作用与set -e相同。

```
command || { echo "command failed"; exit 1; }
```

　-a 　标示已修改的变量，以供输出至环境变量。
　-b 　使被中止的后台程序立刻回报执行状态。
　-C 　转向所产生的文件无法覆盖已存在的文件。
　-d 　Shell预设会用杂凑表记忆使用过的指令，以加速指令的执行。使用-d参数可取消。
　-e 　若指令传回值不等于0，则立即退出shell。　　
　-f　 　取消使用通配符。
　-h 　自动记录函数的所在位置。
　-H Shell 　可利用"!"加<指令编号>的方式来执行history中记录的指令。
　-k 　指令所给的参数都会被视为此指令的环境变量。
　-l 　记录for循环的变量名称。
　-m 　使用监视模式。
　-n 　只读取指令，而不实际执行。
　-p 　启动优先顺序模式。
　-P 　启动-P参数后，执行指令时，会以实际的文件或目录来取代符号连接。
　-t 　执行完随后的指令，即退出shell。
　-u 　当执行时使用到未定义过的变量，则显示错误信息。
　-v 　显示shell所读取的输入值。
　-x 　执行指令后，会先显示该指令及所下的参数。
　+<参数> 　取消某个set曾启动的参数。

## 变量

### 变量类型

**整数变量**： 在一些Shell中，你可以使用 **declare** 或 **typeset** 命令来声明整数变量。

这样的变量只包含整数值，例如：

```
declare -i my_integer=42
```



**字符串变量：** 在 Shell中，变量通常被视为字符串。

你可以使用单引号 **'** 或双引号 **"** 来定义字符串



**数组变量：** Shell 也支持数组，允许你在一个变量中存储多个值



**环境变量：** 这些是由操作系统或用户设置的特殊变量，用于配置 Shell 的行为和影响其执行环境。



### 只读变量

使用 readonly 命令可以将变量定义为只读变量，只读变量的值不能被改变。

```shell
myUrl="https://www.google.com"
readonly myUrl
myUrl="https://www.runoob.com"
```

### 变量赋值

```shell
a=23
# Assignment using 'let'
let a=16+5
let "counter++" # 最好是引起来
let "a=$a+6"
let "a=a+8"
let "d += 1"
```

### 删除变量

使用 unset 命令可以删除变量。语法：

```
unset variable_name
```

### $RANDOM

integer in the range 0~32767.



> 取500以内的随机数

```
RANGE=500
echo
number=$RANDOM
let "number %= $RANGE"
```



### Special Variable Types

**Local variables** ：Variables visible only within a code block or function 

```shell
local loc_var=23
```



### Quoting Variables

```shell
List="one two three"

for a in $List
do
echo "$a"
done
# one
# two
# three
# Splits the variable in parts at whitespace.
echo "---"
for a in "$List"
 # Preserves whitespace in a single variable.
do #
 ^
 ^
echo "$a"
done
# one two three
```



## 运算符



Shell 和其他编程语言一样，支持多种运算符，包括：

- 算数运算符
- 关系运算符
- 布尔运算符
- 字符串运算符
- 文件测试运算符

```shell
#!/bin/bash

val=`expr 2 + 2`
echo "两数之和为 : $val"
```

### 常用操作运算 let,双括号和expr

> `let`
>
> ```
> n=1; let --n && echo "True" || echo "False"
>  # False
> n=1; let n-- && echo "True" || echo "False"
>  # True
> ```



> `双括号  (( ... )) `
>
> ```shell
> a=$(( 5 + 3 ))
> (( var++ ))   # 跟let类似，是否要带上$都一样 (( $var++ ))
> ```



> `expr表达式` 
>
> ```
> 1、计算字串长度
> 
> > expr length “this is a test”
>  14
> 2、抓取字串
> 
> > expr substr “this is a test” 3 5
> is is
> 3、抓取第一个字符数字串出现的位置
> 
> > expr index "sarasara"  a
>  2
> 4、整数运算
> 
>  > expr 14 % 9
>  5
>  > expr 10 + 10
>  20
>  > expr 1000 + 900
>  1900
>  > expr 30 / 3 / 2
>  5
>  > expr 30 \* 3 (使用乘号时，必须用反斜线屏蔽其特定含义。因为shell可能会误解显示星号的意义)
>  90
>  > expr 30 * 3
>  expr: Syntax error
> 
> ```



### 算数运算符

| 运算符 | 说明                                          | 举例                          |
| :----- | :-------------------------------------------- | :---------------------------- |
| +      | 加法                                          | `expr $a + $b` 结果为 30。    |
| -      | 减法                                          | `expr $a - $b` 结果为 -10。   |
| *      | 乘法                                          | `expr $a \* $b` 结果为  200。 |
| /      | 除法                                          | `expr $b / $a` 结果为 2。     |
| %      | 取余                                          | `expr $b % $a` 结果为 0。     |
| =      | 赋值                                          | a=$b 把变量 b 的值赋给 a。    |
| ==     | 相等。用于比较两个数字，相同则返回 true。     | [ $a == $b ] 返回 false。     |
| !=     | 不相等。用于比较两个数字，不相同则返回 true。 | [ $a != $b ] 返回 true。      |

```shell
a=10
b=20

val=`expr $a + $b`
echo "a + b : $val"

val=`expr $a - $b`
echo "a - b : $val"

val=`expr $a \* $b`
echo "a * b : $val"

val=`expr $b / $a`
echo "b / a : $val"

val=`expr $b % $a`
echo "b % a : $val"

if [ $a == $b ]
then
   echo "a 等于 b"
fi
if [ $a != $b ]
then
   echo "a 不等于 b"
fi
```

### 关系运算符

| 运算符 | 说明                                                  | 举例                       |
| :----- | :---------------------------------------------------- | :------------------------- |
| -eq    | 检测两个数是否相等，相等返回 true。                   | [ $a -eq $b ] 返回 false。 |
| -ne    | 检测两个数是否不相等，不相等返回 true。               | [ $a -ne $b ] 返回 true。  |
| -gt    | 检测左边的数是否大于右边的，如果是，则返回 true。     | [ $a -gt $b ] 返回 false。 |
| -lt    | 检测左边的数是否小于右边的，如果是，则返回 true。     | [ $a -lt $b ] 返回 true。  |
| -ge    | 检测左边的数是否大于等于右边的，如果是，则返回 true。 | [ $a -ge $b ] 返回 false。 |
| -le    | 检测左边的数是否小于等于右边的，如果是，则返回 true。 | [ $a -le $b ] 返回 true。  |

### 布尔运算符

| 运算符 | 说明                                                | 举例                                     |
| :----- | :-------------------------------------------------- | :--------------------------------------- |
| !      | 非运算，表达式为 true 则返回 false，否则返回 true。 | [ ! false ] 返回 true。                  |
| -o     | 或运算，有一个表达式为 true 则返回 true。           | [ $a -lt 20 -o $b -gt 100 ] 返回 true。  |
| -a     | 与运算，两个表达式都为 true 才返回 true。           | [ $a -lt 20 -a $b -gt 100 ] 返回 false。 |

### 逻辑运算符

| 运算符 | 说明       | 举例                                       |
| :----- | :--------- | :----------------------------------------- |
| &&     | 逻辑的 AND | [[ $a -lt 100 && $b -gt 100 ]] 返回 false  |
| \|\|   | 逻辑的 OR  | [[ $a -lt 100 \|\| $b -gt 100 ]] 返回 true |

### 字符串运算符

| 运算符 | 说明                                         | 举例                     |
| :----- | :------------------------------------------- | :----------------------- |
| =      | 检测两个字符串是否相等，相等返回 true。      | [ $a = $b ] 返回 false。 |
| !=     | 检测两个字符串是否不相等，不相等返回 true。  | [ $a != $b ] 返回 true。 |
| -z     | 检测字符串长度是否为0，为0返回 true。        | [ -z $a ] 返回 false。   |
| -n     | 检测字符串长度是否不为 0，不为 0 返回 true。 | [ -n "$a" ] 返回 true。  |
| $      | 检测字符串是否不为空，不为空返回 true。      | [ $a ] 返回 true。       |

### 文件测试运算符

| 操作符  | 说明                                                         | 举例                      |
| :------ | :----------------------------------------------------------- | :------------------------ |
| -b file | 检测文件是否是块设备文件，如果是，则返回 true。              | [ -b $file ] 返回 false。 |
| -c file | 检测文件是否是字符设备文件，如果是，则返回 true。            | [ -c $file ] 返回 false。 |
| -d file | 检测文件是否是目录，如果是，则返回 true。                    | [ -d $file ] 返回 false。 |
| -f file | 检测文件是否是普通文件（既不是目录，也不是设备文件），如果是，则返回 true。 | [ -f $file ] 返回 true。  |
| -g file | 检测文件是否设置了 SGID 位，如果是，则返回 true。            | [ -g $file ] 返回 false。 |
| -k file | 检测文件是否设置了粘着位(Sticky Bit)，如果是，则返回 true。  | [ -k $file ] 返回 false。 |
| -p file | 检测文件是否是有名管道，如果是，则返回 true。                | [ -p $file ] 返回 false。 |
| -u file | 检测文件是否设置了 SUID 位，如果是，则返回 true。            | [ -u $file ] 返回 false。 |
| -r file | 检测文件是否可读，如果是，则返回 true。                      | [ -r $file ] 返回 true。  |
| -w file | 检测文件是否可写，如果是，则返回 true。                      | [ -w $file ] 返回 true。  |
| -x file | 检测文件是否可执行，如果是，则返回 true。                    | [ -x $file ] 返回 true。  |
| -s file | 检测文件是否为空（文件大小是否大于0），不为空返回 true。     | [ -s $file ] 返回 true。  |
| -e file | 检测文件（包括目录）是否存在，如果是，则返回 true。          | [ -e $file ] 返回 true。  |



## test 命令

### 数值测试

| 参数 | 说明           |
| :--- | :------------- |
| -eq  | 等于则为真     |
| -ne  | 不等于则为真   |
| -gt  | 大于则为真     |
| -ge  | 大于等于则为真 |
| -lt  | 小于则为真     |
| -le  | 小于等于则为真 |

```shell
num1=100
num2=100
if test $[num1] -eq $[num2]
then
    echo '两个数相等！'
else
    echo '两个数不相等！'
fi

# 设置
if [ "$a" -gt 0 ] && [ "$a" -lt 5 ]
then
	echo "The value of \"a\" lies somewhere between 0 and 5."
fi
```

### 字符串测试

| 参数      | 说明                     |
| :-------- | :----------------------- |
| =         | 等于则为真               |
| !=        | 不相等则为真             |
| -z 字符串 | 字符串的长度为零则为真   |
| -n 字符串 | 字符串的长度不为零则为真 |

```shell
num1="ru1noob"
num2="runoob"
if test $num1 = $num2
then
    echo '两个字符串相等!'
else
    echo '两个字符串不相等!'
fi
```

> What is the difference between [[ $string == "efg*" ]] and [[ $string == efg* ]] ?
>
> ```
> [[ $string == efg* ]] – checks if string begins with efg. 
> 
> [[ $string == "efg*" ]] – checks if string is efg. 
> ```



### 文件测试

| 参数      | 说明                                 |
| :-------- | :----------------------------------- |
| -e 文件名 | 如果文件存在则为真                   |
| -r 文件名 | 如果文件存在且可读则为真             |
| -w 文件名 | 如果文件存在且可写则为真             |
| -x 文件名 | 如果文件存在且可执行则为真           |
| -s 文件名 | 如果文件存在且至少有一个字符则为真   |
| -d 文件名 | 如果文件存在且为目录则为真           |
| -f 文件名 | 如果文件存在且为普通文件则为真       |
| -c 文件名 | 如果文件存在且为字符型特殊文件则为真 |
| -b 文件名 | 如果文件存在且为块特殊文件则为真     |

```shell
cd /bin
if test -e ./bash
then
    echo '文件已存在!'
else
    echo '文件不存在!'
fi
```



## 流程控制

### if else

#### fi

if 语句语法格式：

```
if condition
then
    command1 
    command2
    ...
    commandN 
fi
```

#### if else

```
if condition
then
    command1 
    command2
    ...
    commandN
else
    command
fi
```

#### if else-if else

if else-if else 语法格式：

```
if condition1
then
    command1
elif condition2 
then 
    command2
else
    commandN
fi
```

### for 循环

```
for var in item1 item2 ... itemN
do
    command1
    command2
    ...
    commandN
done
```

```shell
# Standard syntax.
for a in 1 2 3 4 5 6 7 8 9 10
do
	echo -n "$a "
done

# Using "seq" ...
for a in `seq 10`
do
	echo -n "$a "
done

# Using brace expansion ...
# Bash, version 3+.
for i in {1..100} ; do echo $i; done;

for a in {1..10}
do
	echo -n "$a "
done

## 使用双括号 (())
for((n=1; n<=10; n++))
# No do!
{
echo -n "* $n *"
}
```



### while 语句

```shell
#!/bin/bash
int=1
while(( $int<=5 ))
do
    echo $int
    let "int++"
    # int=`expr $int + 1`
    # `expr 1 + 1`
done
```

> 读取文件

```shell
count=1    # 赋值语句，不加空格
cat test.txt | while read line      # cat 命令的输出作为read命令的输入,read读到>的值放在line中
do
   echo "Line $count:$line"
   count=$[ $count + 1 ]          # 注意中括号中的空格。
done
echo "finish"
exit 0
```

### 无限循环

```
while :
do
    command
done
```

> a while loop can call a function

```shell
t=0
condition(){
    ((t++))

	if [ $t -lt 5 ]
    then
        return 0 # true
    else
    return 1 # false
    fi
}
while condition
#
 ^^^^^^^^^
#
 Function call -- four loop iterations.
do
	echo "Still going: t = $t"
done
```

## 数组

数组中可以存放多个值。Bash Shell 只支持一维数组（不支持多维数组），初始化时不需要定义数组大小（与 PHP 类似）。

与大部分编程语言类似，数组元素的下标由 0 开始。(注意：`zsh是下标是从1开始`)

Shell 数组用括号来表示，元素用"空格"符号分割开，语法格式如下：

```
array_name=(value1 value2 ... valuen)
```

### 读取数组

> ```
> ${数组名[下标]}
> echo ${array_name[@]} # 使用 @ 符号可以获取数组中的所有元素
> ```

```shell
#!/bin/bash

my_array=(A B "C" D)
echo "第一个元素为: ${my_array[0]}"
echo "第二个元素为: ${my_array[1]}"
echo "第三个元素为: ${my_array[2]}"
echo "第四个元素为: ${my_array[3]}"

#快速定义一个1-10的数组
ArrayVar=({1..10})  # 1 2 3 4 5 6 7 8 9 10
```

### **关联数组**

Bash 支持关联数组，可以使用任意的字符串、或者整数作为下标来访问数组元素。

关联数组使用 **[declare](https://www.runoob.com/linux/linux-comm-declare.html)** 命令来声明，语法格式如下：

```
declare -A array_name
```

```shell
declare -A site=(["google"]="www.google.com" ["runoob"]="www.runoob.com" ["taobao"]="www.taobao.com")

declare -A site
site["google"]="www.google.com"
site["runoob"]="www.runoob.com"
site["taobao"]="www.taobao.com"
```

### **获取数组元素、键、个数**

使用 **@** 或 ***** 可以获取数组中的所有元素，例如：

```shell
my_array[0]=A
my_array[1]=B
my_array[2]=C
my_array[3]=D

echo "数组的元素为: ${my_array[*]}"
echo "数组的元素为: ${my_array[@]}"

#在数组前加一个感叹号 ! 可以获取数组的所有键，例如：
echo "数组的键为: ${!site[*]}"    # 返回数组 site 的所有键，作为一个单一的字符串
echo "数组的键为: ${!site[@]}"	 # 返回数组 site 的所有键，每个键作为单独的单词
echo "数组的键为: ${!my_array[@]}"

#获取数组长度的方法与获取字符串长度的方法相同
echo "数组元素个数为: ${#my_array[*]}"
echo "数组元素个数为: ${#my_array[@]}"

str="Hello World"
length=${#str}
```

### 遍历数组

```shell
## 标准for循环
for(( i=0;i<${#array[@]};i++)) do
#${#array[@]}获取数组长度用于循环
	echo ${array[i]};
done;

## 遍历数组（不带数组下标）  (注意与awk的区别)
for element in ${array[@]}
#也可以写成for element in ${array[*]}
do
	echo $element
done

## 遍历（带数组下标）：
for i in "${!arr[@]}";   
do   
    printf "%s\t%s\n" "$i" "${arr[$i]}"  
done  
```



## 字符串

字符串是shell编程中最常用最有用的数据类型（除了数字和字符串，也没啥其它类型好用了），字符串可以用单引号，也可以用双引号，也可以不用引号。

```shell
your_name="runoob"
# 使用双引号拼接
greeting="hello, "$your_name" !"
greeting_1="hello, ${your_name} !"
echo $greeting  $greeting_1

# 使用单引号拼接
greeting_2='hello, '$your_name' !'
greeting_3='hello, ${your_name} !'
echo $greeting_2  $greeting_3

#获取字符串长度
string="abcd"
echo ${#string}   # 输出 4
echo ${#string[0]}   # 输出 4
# 变量为字符串时，${#string} 等价于 ${#string[0]}:

# 提取子字符串(启始位置从0开始，最后参数是长度不指定默认到字符串末尾)
string="runoob is a great site"
echo ${string:1:4} # 输出 unoo

# 查找子字符串
string="runoob is a great site"
echo `expr index "$string" io`  # 输出 4
```

### String Length

```shell
stringZ=abcABC123ABCabc
echo ${#stringZ}
echo `expr length $stringZ`
echo `expr "$stringZ" : '.*'`
###15
15
15
```

### Length of Matching Substring at Beginning of String

> `expr match "$string" '$substring'`
> 	$substring is a regular expression.
> `expr "$string" : '$substring'`
> 	$substring is a regular expression.

```shell
stringZ=abcABC123ABCabc
#
 |------|
#
 12345678
echo `expr match "$stringZ" 'abc[A-Z]*.2'`
echo `expr "$stringZ" : 'abc[A-Z]*.2'`
##8
##8
```

### Index

> expr index $string $substring
>
> ​	Numerical position in $string of first character in $substring that matches.

```shell
stringZ=abcABC123ABCabc
#
 123456 ...
echo `expr index "$stringZ" C12`
echo `expr index "$stringZ" 1c`
# 'c' (in #3 position) matches before '1'.
```



### Substring Extraction

> `${string:position}`
> 	Extracts substring from $string at $position.
> 	If the $string parameter is "*" or "@", then this extracts the positional parameters, [49] starting at $position.
> `${string:position:length}`
> 	Extracts $length characters of substring from $string at $position.

```shell
stringZ=abcABC123ABCabc
#
 0123456789.....
#
 0-based indexing.
echo ${stringZ:0}
echo ${stringZ:1}
echo ${stringZ:7}
echo ${stringZ:7:3}
###abcABC123ABCabc
# bcABC123ABCabc
# 23ABCabc
# 23A
# Three characters of substring.
```

### Here Document



> Here String 

```shell
# Instead of:
if echo "$VAR" | grep -q txt
# etc.
# if [[ $VAR = *txt* ]]
# Try:
if grep -q "txt" <<< "$VAR"
then
 #
 ^^^
echo "$VAR contains the substring sequence \"txt\""
fi
```

```shell
String="This is a string of words."
read -r -a Words <<< "$String"
# The -a option to "read"
#+ assigns the resulting values to successive members of an array.
echo "First word in String is: ${Words[0]}"
echo "Second word in String is: ${Words[1]}"
echo "Third word in String is: ${Words[2]}"
echo "Fourth word in String is: ${Words[3]}"
echo "Fifth word in String is: ${Words[4]}"
echo "Sixth word in String is: ${Words[5]}"
echo "Seventh word in String is: ${Words[6]}"
```

```shell
read -p "File: " file
 # -p arg to 'read' displays prompt.
if [ ! -e "$file" ]
then
 # Bail out if no such file.
echo "File $file not found."
exit $E_NOSUCHFILE
fi
read -p "Title: " title
cat - $file <<<$title > $file.new    # 在前面添加标题
```



> 结合数组处理行记录

```shell
OIFS=$IFS
IFS=', ' read -r -a array <<< "$string"

echo "${array[0]}"

for element in "${array[@]}"
do
    echo "$element"
done
IFS=$OIFS
```

> `IFS`
>
> ```shell
> echo "List of all users:"
> OIFS=$IFS; IFS=:
>  # /etc/passwd uses ":" for field separator.
> while read name passwd uid gid fullname ignore
> do
> 	echo "$name ($fullname)"
> done </etc/passwd
>  # I/O redirection.
> IFS=$OIFS
>  # Restore original $IFS.
> ```



## 函数

### shell中函数的定义

格式如下:

```
[ function ] funname [()]

{

    action;

    [return int;]

}
```



```shell
demoFun(){
    echo "这是我的第一个 shell 函数!"
}
echo "-----函数开始执行-----"
demoFun
echo "-----函数执行完毕-----"
```



### 带有 **return** 语句的函数

```shell
funWithReturn(){
    echo "这个函数会对输入的两个数字进行相加运算..."
    echo "输入第一个数字: "
    read aNum
    echo "输入第二个数字: "
    read anotherNum
    echo "两个数字分别为 $aNum 和 $anotherNum !"
    return $(($aNum+$anotherNum))
}
funWithReturn
echo "输入的两个数字之和为 $? !"
```

> 函数返回值在调用该函数后通过 **$?** 来获得。
>
> **注意：** **return** 语句只能返回一个介于 0 到 255 之间的整数，而两个输入数字的和可能超过这个范围。
>
> 要解决这个问题，您可以修改 return 语句，直接使用 echo 输出和而不是使用 return
>
> ```shell
> get_str()
> {
> 	echo "string"
> }
> 
> echo `get_str` #写法一
> echo $(get_str) #写法二
> ```



### 函数参数

```shell
funWithParam(){
    echo "第一个参数为 $1 !"
    echo "第二个参数为 $2 !"
    echo "第十个参数为 $10 !"
    echo "第十个参数为 ${10} !"
    echo "第十一个参数为 ${11} !"
    echo "参数总数有 $# 个!"
    echo "作为一个字符串输出所有参数 $* !"
}
funWithParam 1 2 3 4 5 6 7 8 9 34 73
```

| 参数处理 | 说明                                                         |
| :------- | :----------------------------------------------------------- |
| $#       | 传递到脚本的参数个数                                         |
| $*       | 以一个单字符串显示所有向脚本传递的参数。 如"$*"用「"」括起来的情况、以"$1 $2 … $n"的形式输出所有参数。 |
| $$       | 脚本运行的当前进程ID号                                       |
| $!       | 后台运行的最后一个进程的ID号                                 |
| $@       | 与$*相同，但是使用时加引号，并在引号中返回每个参数。 如"$@"用「"」括起来的情况、以"$1" "$2" … "$n" 的形式输出所有参数。 |
| $-       | 显示Shell使用的当前选项，与[set命令](https://www.runoob.com/linux/linux-comm-set.html)功能相同。 |
| $?       | 显示最后命令的退出状态。0表示没有错误，其他任何值表明有错误。 |

```shell
echo "Shell 传递参数实例！";
echo "第一个参数为：$1";

echo "参数个数为：$#";
echo "传递的参数作为一个字符串显示：$*";
```

> $* 与 $@ 区别：
>
> - 相同点：都是引用所有参数。
> - 不同点：只有在双引号中体现出来。假设在脚本运行时写了三个参数 1、2、3，则 " * " 等价于 "1 2 3"（传递了一个参数），而 "@" 等价于 "1" "2" "3"（传递了三个参数）。





### 问题求解

#### 打印三角形状的星号

> 输出N为5是则打印如下的形状
>
> ```
>        *
>       ***
>      *****
>     *******
> *********
> ```

```shell
#!/bin/bash

print_centered_triangle() {
    n=$1
    max_width=$((2 * n - 1))
    for ((i = 1; i <= n; i++)); do
        stars_count=$((2 * i - 1))
        left_spaces=$(( (max_width - stars_count) / 2 ))
        printf "%${left_spaces}s"  # 打印左边空格
        printf "%${stars_count}s\n" | tr " " "*"  # 打印星号并换行
        #seq -s* $((stars_count+1))|tr -d '[:digit:]'
    done
}
print_centered_triangle $1
```



## Parameter Expansion



1. **Use Default Values**
    `${parameter:-word}`
    If parameter is unset or null, the expansion of word is substituted. Otherwise, the value of parameter is substituted.

2. **Assign Default Values**
`${parameter:=word}`
If parameter is unset or null, the expansion of word is assigned to parameter.  
The value of parameter is then substituted. Positional parameters and special parameters may not be assigned to 
in this way.

3. `${parameter:?word}`
If parameter is null or unset, the expansion of word (or a message to that effect if word is not present) is written to 
the standard error and the shell, if it is not interactive, exits. Otherwise, the value of parameter is substituted.

4. `${parameter:+word}`
If parameter is null or unset, nothing is substituted, otherwise the expansion of word is substituted.

5. Substring Expansion
`${parameter:offset}`
`${parameter:offset:length}`

This is referred to as Substring Expansion. It expands to up to length characters of the value of parameter starting 
at the character specified by offset. If parameter is ‘@’, an indexed array subscripted by ‘@’ or ‘*’, or an associative 
array name, the results differ as described below. If length is omitted, it expands to the substring of the value of 
parameter starting at the character specified by offset and extending to the end of the value. length and offset are 
arithmetic expressions.

6. Indirect expansion
`${!prefix*}`
`${!prefix@}`
Expands to the names of variables whose names begin with prefix, separated by the first character of the IFS 
special variable. When ‘@’ is used and the expansion appears within double quotes, each variable name expands 
to a separate word.
eg：
echo ${!HO*}
output:
HOME HOSTNAME HOSTTYPE

7.
`${!name[@]}`
`${!name[*]}`

If name is an array variable, expands to the list of array indices (keys) assigned in name. 
If name is not an array, expands to 0 if name is set and null otherwise. 
When ‘@’ is used and the expansion appears within double quotes, each key expands to a separate word.


8. Parameter length
`${#parameter}`
The length in characters of the expanded value of parameter is substituted. If parameter is ‘*’ or ‘@’, 
the value substituted is the number of positional parameters.

9. Remove matching prefix/suffix pattern
`${parameter#word}`
`${parameter##word}`

The word is expanded to produce a pattern and matched according to the rules described below (see Pattern Matching). 
If the pattern matches the beginning of the expanded value of parameter, then the result of the expansion is 
the expanded value of parameter with the shortest matching pattern (the ‘#’ case) or the longest matching pattern 
(the ‘##’ case) deleted. If parameter is ‘@’ or ‘*’, the pattern removal operation is applied to each 
positional parameter in turn, and the expansion is the resultant list. If parameter is an array variable subscripted 
with ‘@’ or ‘*’, the pattern removal operation is applied to each member of the array in turn, and the expansion 
is the resultant list.
#(掐头)
eg:
echo $test => abcabcdf
echo ${test#abc}  => abcdf
echo ${test##abc} => abcdf

echo $test => aabc
echo ${test#*}  =>abc
echo ${test##*} =>

echo ${test#*a}  =>abc
echo ${test##*a} =>bc


${parameter%word}
${parameter%%word}
The word is expanded to produce a pattern and matched according to the rules described below (see Pattern Matching). 
If the pattern matches If the pattern matches a trailing portion of the expanded value of parameter, 
then the result of the expansion is the value of parameter with the shortest matching pattern (the ‘%’ case) or 
the longest matching pattern (the ‘%%’ case) deleted. If parameter is ‘@’ or ‘*’, 
the pattern removal operation is applied to each positional parameter in turn, and the expansion is the resultant list. 
If parameter is an array variable subscripted with ‘@’ or ‘*’, the pattern removal operation is applied to 
each member of the array in turn, and the expansion is the resultant list.
#(去尾)


10. Pattern  substitution
`${parameter/pattern/string}`

The pattern is expanded to produce a pattern just as in filename expansion. Parameter is expanded and the longest match 
of pattern against its value is replaced with string. The match is performed according to the rules described below 
(see Pattern Matching). If pattern begins with ‘/’, all matches of pattern are replaced with string. 
Normally only the first match is replaced. If pattern begins with ‘#’, it must match at the beginning of 
the expanded value of parameter. If pattern begins with ‘%’, it must match at the end of the expanded value of parameter. 
If string is null, matches of pattern are deleted and the / following pattern may be omitted. 
If the nocasematch shell option (see the description of shopt in The Shopt Builtin) is enabled, 
the match is performed without regard to the case of alphabetic characters. If parameter is ‘@’ or ‘*’, 
the substitution operation is applied to each positional parameter in turn, and the expansion is the resultant list. 
If parameter is an array variable subscripted with ‘@’ or ‘*’, the substitution operation is applied to each member of 
the array in turn, and the expansion is the resultant list.
eg:
echo ${HOME} => /home/ecaimch
echo ${HOME/m/s}  => /hose/ecaimch
echo ${HOME//m/s} => /hose/ecaisch


11. Case modification
${parameter^pattern}
${parameter^^pattern}
${parameter,pattern}
${parameter,,pattern}

This expansion modifies the case of alphabetic characters in parameter. 
The pattern is expanded to produce a pattern just as in filename expansion. 
Each character in the expanded value of parameter is tested against pattern, and, if it matches the pattern, 
its case is converted. The pattern should not attempt to match more than one character. 
The ‘^’ operator converts lowercase letters matching pattern to uppercase; 
the ‘,’ operator converts matching uppercase letters to lowercase. 
The ‘^^’ and ‘,,’ expansions convert each matched character in the expanded value; 
the ‘^’ and ‘,’ expansions match and convert only the first character in the expanded value. 
If pattern is omitted, it is treated like a ‘?’, which matches every character. 
If parameter is ‘@’ or ‘*’, the case modification operation is applied to each positional parameter in turn, 
and the expansion is the resultant list. If parameter is an array variable subscripted with ‘@’ or ‘*’, 
the case modification operation is applied to each member of the array in turn, and the expansion is the resultant list.


${parameter@operator}
The expansion is either a transformation of the value of parameter or information about parameter itself, 
depending on the value of operator. Each operator is a single letter:

Q
The expansion is a string that is the value of parameter quoted in a format that can be reused as input.

E
The expansion is a string that is the value of parameter with backslash escape sequences expanded as with the $'…' quoting mechanism.

P
The expansion is a string that is the result of expanding the value of parameter as if it were a prompt string (see Controlling the Prompt).

A
The expansion is a string in the form of an assignment statement or declare command that, if evaluated, will recreate parameter with its attributes and value.

a
The expansion is a string consisting of flag values representing parameter’s attributes.



Appendix
1).  Pattern Matching

*
Matches any string, including the null string. When the globstar shell option is enabled, 
and ‘*’ is used in a filename expansion context, two adjacent ‘*’s used as a single pattern will match all files 
and zero or more directories and subdirectories. If followed by a ‘/’, two adjacent ‘*’s will match only directories 
and subdirectories.

?
Matches any single character.

[…]
Matches any one of the enclosed characters.

?(pattern-list)
Matches zero or one occurrence of the given patterns.

*(pattern-list)
Matches zero or more occurrences of the given patterns.

+(pattern-list)
Matches one or more occurrences of the given patterns.

@(pattern-list)
Matches one of the given patterns.

!(pattern-list)
Matches anything except one of the given patterns

2) reference 
http://www.gnu.org/software/bash/manual/html_node/Pattern-Matching.html#Pattern-Matching

​		http://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html



## 时间处理

### 时间格式化

%% 输出%符号 a literal %
%a 当前域的星期缩写 locale’s abbreviated weekday name (Sun..Sat)
%A 当前域的星期全写 locale’s full weekday name, variable length (Sunday..Saturday)
%b 当前域的月份缩写 locale’s abbreviated month name (Jan..Dec)
%B 当前域的月份全称 locale’s full month name, variable length (January..December)
%c 当前域的默认时间格式 locale’s date and time (Sat Nov 04 12:02:33 EST 1989)
%C n百年 century (year divided by 100 and truncated to an integer) [00-99]
`%d 两位的天 day of month (01..31)`
%D 短时间格式 date (mm/dd/yy)
%e 短格式天 day of month, blank padded ( 1..31)
%F 文件时间格式 same as %Y-%m-%d
%g the 2-digit year corresponding to the %V week number
%G the 4-digit year corresponding to the %V week number
%h same as %b
`%H 24小时制的小时 hour (00..23)`
%I 12小时制的小时 hour (01..12)
%j 一年中的第几天 day of year (001..366)
%k 短格式24小时制的小时 hour ( 0..23)
%l 短格式12小时制的小时 hour ( 1..12)
%m 双位月份 month (01..12)
%M 双位分钟 minute (00..59)
%n 换行 a newline
%N 十亿分之一秒 nanoseconds (000000000..999999999)
%p 大写的当前域的上下午指示 locale’s upper case AM or PM indicator (blank in many locales)
%P 小写的当前域的上下午指示 locale’s lower case am or pm indicator (blank in many locales)
%r 12小时制的时间表示（时:分:秒,双位） time, 12-hour (hh:mm:ss [AP]M)
%R 24小时制的时间表示 （时:分,双位）time, 24-hour (hh:mm)
%s 自基础时间 1970-01-01 00:00:00 到当前时刻的秒数 seconds since 00:00:00 1970-01-01 UTC (a GNU extension)
%S 双位秒 second (00..60); the 60 is necessary to accommodate a leap second
%t 横向制表位(tab) a horizontal tab
`%T 24小时制时间表示 time, 24-hour (hh:mm:ss)`
`%u 数字表示的星期（从星期一开始 1-7）day of week (1..7); 1 represents Monday`
%U 一年中的第几周星期天为开始 week number of year with Sunday as first day of week (00..53)
%V 一年中的第几周星期一为开始 week number of year with Monday as first day of week (01..53)
%w 一周中的第几天 星期天为开始 0-6 day of week (0..6); 0 represents Sunday
%W 一年中的第几周星期一为开始 week number of year with Monday as first day of week (00..53)
%x 本地日期格式 locale’s date representation (mm/dd/yy)
%X 本地时间格式 locale’s time representation (%H:%M:%S)
%y 两位的年 last two digits of year (00..99)
%Y 年 year (1970…)
%z RFC-2822 标准时间格式表示的域 RFC-2822 style numeric timezone (-0500) (a nonstandard extension)
%Z 时间域 time zone (e.g., EDT), or nothing if no time zone is determinable



### 定时任务技巧

定进任务字段说明

分钟（0-59） 
小时（0-23） 
日期（1-31） 
月份（1-12） 
星期几（0-6，其中0代表星期日，好像7也代表星期日） 



> 每个月的最后一个周五定时跑任务

```
# 0 0 * * 5 [ "$(date +\%d -d tomorrow)" -le 7 ] && your_command
0 0 * * 5 [ "$(date -d "next Fri"  +\%d)" -le 7 ] && your_command
```

> 每个月的第一个周五定时跑任务

```
# 0 0 * * 5 [ "$(date +\%u)" -eq 5 ] && your_command  # not correct
0 0 * * 5 [ "$(date +\%d)" -le 7 ] && your_command
0 0 1-7 * * [ "$(date +\%u)" -eq 5 ] && your_command
0 0 1-7 * * [ "$(date +\%a)" = "Fri" ] && your_command
```



## 其他运维命令

### openssl命令

> 查看证书

```shell
openssl x509 -in cert.pem -text -noout
```



> 生成自签证书

```shell
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
```

- `req`：表示证书请求命令。
- `-x509`：表示生成自签名证书。
- `-newkey rsa:4096`：表示生成一个 RSA 密钥，长度为 4096 位。
- `-keyout key.pem`：指定私钥输出文件名。
- `-out cert.pem`：指定证书输出文件名。
- `-days 365`：证书有效期为 365 天。



> 生成证书请求（CSR）

```shell
openssl req -newkey rsa:2048 -keyout key.pem -out req.pem
```

- `req`：表示证书请求命令。
- `-newkey rsa:2048`：表示生成一个 RSA 密钥，长度为 2048 位。
- `-keyout key.pem`：指定私钥输出文件名。
- `-out req.pem`：指定证书请求输出文件名。



> 将证书转换为不同格式

```shell
# der转换为 PEM 格式
openssl x509 -in cert.der -outform PEM -out cert.pem
# PEM 转换为 DER 格式
openssl x509 -in cert.pem -outform DER -out cert.der
```



> openssl s_client

```
openssl s_client [-host host] [-port port] [-connect host:port] [-verify depth] [-cert filename] 
[-certform DER|PEM] [-key filename] [-keyform DER|PEM] [-pass arg] [-CApath directory] [-CAfile filename] 
[-reconnect][-pause] [-showcerts] [-debug] [-msg] [-state] [-nbio_test] [-nbio][-crlf] [-ign_eof] [-no_ign_eof] 
[-quiet] [-ssl2] [-ssl3] [-tls1_1] [-tls1_2] [-tls1] [-dtls1] [-no_ssl2][-no_ssl3] [-no_tls1] [-no_tls1_1] 
[-no_tls1_2] [-bugs] [-cipher cipherlist] [-starttls protocol] [-engine id] [-tlsextdebug] [-no_ticket] 
[-sess_out filename] [-sess_in filename] [-rand file(s)] 
```

```shell
#!/bin/bash

# 服务的域名或IP地址和端口
HOST="example.com"
PORT=443

# 检查证书是否在30天内过期的阈值
THRESHOLD_DAYS=30

# 获取证书信息
cert_info=$(echo | openssl s_client -servername $HOST -connect $HOST:$PORT 2>/dev/null | openssl x509 -noout -dates)

# 提取证书过期日期
expiry_date=$(echo "$cert_info" | grep "notAfter" | cut -d "=" -f 2)

# 转换日期格式为Unix时间戳
expiry_epoch=$(date -d "$expiry_date" +%s)

# 获取当前日期的Unix时间戳
current_epoch=$(date +%s)

# 计算证书过期天数
days_to_expiry=$(( ($expiry_epoch - $current_epoch) / 86400 ))

# 检查是否在30天内过期
if [ $days_to_expiry -le $THRESHOLD_DAYS ]; then
    echo "证书将在 $days_to_expiry 天内过期！"
else
    echo "证书在 $days_to_expiry 天后过期。"
fi
```

### strace 和 ltrace

这些系统调用主要分为几类：

- 文件和设备访问类 比如open/close/read/write/chmod等
- 进程管理类 fork/clone/execve/exit/getpid等
- 信号类 signal/sigaction/kill 等
- 内存管理 brk/mmap/mlock等
- 进程间通信IPC shmget/semget * 信号量，共享内存，消息队列等
- 网络通信 socket/connect/sendto/sendmsg 等
- 其他



System trace: diagnostic and debugging tool for tracing system calls and signals



常用选项：

> ```
> -tt 在每行输出的前面，显示毫秒级别的时间
> -T 显示每次系统调用所花费的时间
> -v 对于某些相关调用，把完整的环境变量，文件stat结构等打出来。
> -f 跟踪目标进程，以及目标进程创建的所有子进程
> -e 控制要跟踪的事件和跟踪行为,比如指定要跟踪的系统调用名称  
>     -e trace=file     跟踪和文件访问相关的调用(参数中有文件名)
>     -e trace=process  和进程管理相关的调用，比如fork/exec/exit_group
>     -e trace=network  和网络通信相关的调用，比如socket/sendto/connect
>     -e trace=signal    信号发送和处理相关，比如kill/sigaction
>     -e trace=desc  和文件描述符相关，比如write/read/select/epoll等
>     -e trace=ipc 进程见同学相关，比如shmget等
> -o 把strace的输出单独写到指定的文件
> -s 当系统调用的某个参数是字符串时，最多输出指定长度的内容，默认是32个字节
> -p 指定要跟踪的进程pid, 要同时跟踪多个pid, 重复多次-p选项即可
> ```

```shell
strace -tt -T -v -f -e trace=file -o /data/log/strace.log -s 1024 -p 23489
```

```
bash$ strace df
execve("/bin/df", ["df"], [/* 45 vars */]) = 0
uname({sys="Linux", node="bozo.localdomain", ...}) = 0
brk(0)
 = 0x804f5e4
...
```

```shell
# 系统调用的时间
strace -tt -f ./a.sh

# -c 统计每一系统调用的所执行的时间,次数和出错的次数等.
strace -c -o test.txt ./test
strace -c ./test  2>test.txt

# strace -p pid
strace -p pid

# 截断输出
#-s参数用于指定trace结果的每一行输出的字符串的长度，下面看看test程序中-s参数对结果有什么影响，现指定-s为20，然后在read的是是很我们输入一个超过20个字符的数字串
strace -s 20 ./test
```

Library trace: diagnostic and debugging tool that traces library calls invoked by a given command.

```
bash$ ltrace df
__libc_start_main(0x804a910, 1, 0xbfb589a4, 0x804fb70, 0x804fb68 <unfinished ...>:
setlocale(6, "")
 = "en_US.UTF-8"
bindtextdomain("coreutils", "/usr/share/locale") = "/usr/share/locale"
textdomain("coreutils")
 = "coreutils"
__cxa_atexit(0x804b650, 0, 0, 0x8052bf0, 0xbfb58908) = 0
getenv("DF_BLOCK_SIZE")
 = NULL
...
```

当发现进程或服务异常时，我们可以通过strace来跟踪其系统调用，“看看它在干啥”，进而找到异常的原因。熟悉常用系统调用，能够更好地理解和使用strace。

当然，万能的strace也不是真正的万能。当目标进程卡死在用户态时，strace就没有输出了。

这个时候我们需要其他的跟踪手段，比如gdb/perf/SystemTap等。



### top, vmstat 和 iostat

#### top

```
top 
```



%Cpu(s)

Values related to processor utilization are displayed on the third line. They provide insight into exactly what the CPUs are doing.

- `us` is the percent of time spent running user processes.
- `sy` is the percent of time spent running the kernel.
- `ni` is the percent of time spent running processes with manually configured [nice values](https://www.redhat.com/sysadmin/manipulate-process-priority).
- `id` is the percent of time idle (if low, CPU may be overworked).
- `wa` is the percent of wait time (if high, CPU is waiting for I/O access).
- `hi` is the percent of time managing hardware interrupts.
- `si` is the percent of time managing software interrupts.
- `st` is the percent of virtual CPU time waiting for access to physical CPU.



#### vmstat

```shell
vmstat 1 100
# The -a option will give us the active and inactive memory of the system
vmstat -a 1 10 
# The -f option will give us the number of forks since boot
vmstat -f
# The -d option gives you read/write stats for various disks
vmstat -d 
# The -t option gives us timestamp information with every update
vmstat -t 
```



```
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 1  0      0 10635100  93888 2776244    0    0     5    19    6   17  0  0 100  0  0
 0  0      0 10635100  93888 2776244    0    0     0     0   62  198  0  0 100  0  0
 0  0      0 10635100  93888 2776284    0    0     0     0   49  169  0  0 100  0  0
 0  0      0 10635352  93888 2776284    0    0     0     0   68  205  0  0 100  0  0
 0  0      0 10635100  93888 2776284    0    0     0     0   63  200  0  0 100  0  0
 0  0      0 10635100  93888 2776284    0    0     0     0   47  167  0  0 100  0  0
 
- Procs
     - r: The number of runnable processes (running or waiting for run times)
     - b: The number of processes in uninterruptible sleep. 
 - Memory
     - swpd: the amount of virtual memory used. 
     - free: the amount of idle memory
     - buff: the amount of memory used as buffers
     - cache: the amount of memory used as cache. 
     - inact: the amount of inactive memory. (-a option)
     - active: the amount of active memory. (-a option)
 - Swap
     - si: Amount of memory swapped in from disk (/s). 
     - so: Amount of memory swapped to a block device (/s). 
 - IO
     - bi: Blocks received from a block device (blocks/s). 
     - bo: Blocks sent to a block device (blocks/s). 
 - System
     - in: The number of interrupts per second, including the clock. 
     - cs: The number of context switcher per second. 
 - CPU
     - These are percentages of total CPU time. 
     - us: Time spent running non-kernel code. (user time, including nice time)
     - sy: Time spent running kernel code. (system time)
     - id: Time spent idle. Prior to Linux 2.5.41, this includes IO-wait time. 
     - wa: Time spent waiting for IO.  Before Linux 2.5.41, included in idle.
     - st: Time stolen from a virtual machine.  Prior to Linux 2.6.11, unknown.
```

> **r** 表示运行队列(就是说多少个进程真的分配到CPU)，我测试的服务器目前CPU比较空闲，没什么程序在跑，当这个值超过了CPU数目，就会出现CPU瓶颈了。这个也和top的负载有关系，一般负载超过了3就比较高，超过了5就高，超过了10就不正常了，服务器的状态很危险。top的负载类似每秒的运行队列。如果运行队列过大，表示你的CPU很繁忙，一般会造成CPU使用率很高。
>
> **b** 表示阻塞的进程,这个不多说，进程阻塞，大家懂的。
>
> **swpd** 虚拟内存已使用的大小，如果大于0，表示你的机器物理内存不足了，如果不是程序内存泄露的原因，那么你该升级内存了或者把耗内存的任务迁移到其他机器。
>
> **free**  空闲的物理内存的大小，我的机器内存总共8G，剩余3415M。
>
> **buff**  Linux/Unix系统是用来存储，目录里面有什么内容，权限等的缓存，我本机大概占用300多M
>
> **cache** cache直接用来记忆我们打开的文件,给文件做缓冲，我本机大概占用300多M(这里是Linux/Unix的聪明之处，把空闲的物理内存的一部分拿来做文件和目录的缓存，是为了提高 程序执行的性能，当程序使用内存时，buffer/cached会很快地被使用。)
>
> **si** 每秒从磁盘读入虚拟内存的大小，如果这个值大于0，表示物理内存不够用或者内存泄露了，要查找耗内存进程解决掉。我的机器内存充裕，一切正常。
>
> **so** 每秒虚拟内存写入磁盘的大小，如果这个值大于0，同上。
>
> **bi** 块设备每秒接收的块数量，这里的块设备是指系统上所有的磁盘和其他块设备，默认块大小是1024byte，我本机上没什么IO操作，所以一直是0，但是我曾在处理拷贝大量数据(2-3T)的机器上看过可以达到140000/s，磁盘写入速度差不多140M每秒
>
> **bo** 块设备每秒发送的块数量，例如我们读取文件，bo就要大于0。bi和bo一般都要接近0，不然就是IO过于频繁，需要调整。
>
> **in** 每秒CPU的中断次数，包括时间中断
>
> **cs** 每秒上下文切换次数，例如我们调用系统函数，就要进行上下文切换，线程的切换，也要进程上下文切换，这个值要越小越好，太大了，要考虑调低线程或者进程的数目,例如在apache和nginx这种web服务器中，我们一般做性能测试时会进行几千并发甚至几万并发的测试，选择web服务器的进程可以由进程或者线程的峰值一直下调，压测，直到cs到一个比较小的值，这个进程和线程数就是比较合适的值了。系统调用也是，每次调用系统函数，我们的代码就会进入内核空间，导致上下文切换，这个是很耗资源，也要尽量避免频繁调用系统函数。上下文切换次数过多表示你的CPU大部分浪费在上下文切换，导致CPU干正经事的时间少了，CPU没有充分利用，是不可取的。
>
> **us** 用户CPU时间，我曾经在一个做加密解密很频繁的服务器上，可以看到us接近100,r运行队列达到80(机器在做压力测试，性能表现不佳)。
>
> **sy** 系统CPU时间，如果太高，表示系统调用时间长，例如是IO操作频繁。
>
> **id** 空闲 CPU时间，一般来说，id + us + sy = 100,一般我认为id是空闲CPU使用率，us是用户CPU使用率，sy是系统CPU使用率。
>
> **wt** 等待IO CPU时间。





#### iostat

-C 显示CPU使用情况

-d 显示磁盘使用情况

-k 以 KB 为单位显示

-m 以 M 为单位显示

-N 显示磁盘阵列(LVM) 信息

-n 显示NFS 使用情况

-p[磁盘] 显示磁盘和分区的情况

-t 显示终端和CPU的信息

-x 显示详细信息

-V 显示版本信息



常用命令

```shell
# you can use the -x flag to view extended statistics
iostat -d -x -k 1 1
# -p allows you to specify a particular device to focus in on
iostat -p sda 
```

```
Device            r/s     rkB/s   rrqm/s  %rrqm r_await rareq-sz     w/s     wkB/s   wrqm/s  %wrqm w_await wareq-sz     d/s     dkB/s   drqm/s  %drqm d_await dareq-sz     f/s f_await  aqu-sz  %util
sda              0.14      9.52     0.05  26.57    0.38    66.46    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.01
sdb              0.01      0.29     0.00   0.00    0.19    21.84    0.00      0.00     0.00   0.00    1.50     2.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    1.00    0.00   0.00
sdc              1.78     72.07     1.04  36.92    0.37    40.40    1.77    294.69    14.59  89.18    9.47   166.46    0.08     45.28     0.00   2.30    0.39   552.76    0.28    0.46    0.02   0.29


avgqu-sz - average queue length of a request issued to the device
await - average time for I/O requests issued to the device to be served (milliseconds)
r_await - average time for read requests to be served (milliseconds)
w_await - average time for write requests to be served (milliseconds)
```

> **`%util` (Device Utilization)**
>
> - **Meaning**: The percentage of time the device was busy processing requests.
> - **What to Look For**: High values (close to or at 100%) indicate that the device is heavily utilized and may be a bottleneck.
>
> **`await` (Average Wait Time)**
>
> - **Meaning**: The average time (in milliseconds) that I/O requests spend waiting in the queue and being serviced.
> - **What to Look For**: High `await` times indicate latency in the storage system, suggesting that the disk is struggling to keep up with requests.
>
> **`svctm` (Service Time)**
>
> - **Meaning**: The average time (in milliseconds) taken by the device to service an I/O request.
> - **What to Look For**: Large differences between `svctm` and `await` suggest that a significant amount of time is being spent in the I/O queue rather than on actual disk operations.
>
> **`r/s` and `w/s` (Read/Write Requests per Second)**
>
> - **Meaning**: The number of read (`r/s`) and write (`w/s`) requests issued to the device per second.
> - **What to Look For**: High values could indicate heavy read or write activity, which may help in identifying workloads that are stressing the disk.
>
> **`rMB/s` and `wMB/s` (Read/Write Throughput)**
>
> - **Meaning**: The amount of data read from (`rMB/s`) or written to (`wMB/s`) the device per second.
> - **What to Look For**: High throughput might indicate whether the disk is under heavy read or write load.
>
> **`avgqu-sz` (Average Queue Size)**
>
> - **Meaning**: The average number of I/O requests in the queue.
> - **What to Look For**: A large queue size often means that the storage system is overwhelmed and struggling to process requests efficiently.
>
> **Example Analysis:**
>
> - If you notice high `await` times coupled with high `%util`, this often indicates that the disk is a bottleneck, possibly due to heavy I/O operations.
>
> - A high `avgqu-sz` combined with a high `await` suggests the disk is overwhelmed, potentially pointing to the need for performance tuning or hardware upgrades.

### 性能分析命令

#### perf

> CPU

使用 `perf stat` 和 `perf record` 分析 CPU 使用情况，找出 CPU 瓶颈。

示例：发现某个函数占用了过多的 CPU 时间，可以使用 `perf record` 记录其性能数据，然后用 `perf report` 分析。



> TOP

找出程序中的性能热点（耗时最多的部分）。

示例：用 `perf top` 实时监控性能热点，或者用 `perf record` 和 `perf report` 找出热点函数。



> I/O

程序频繁进行磁盘 I/O 操作，可以用 `perf trace` 跟踪这些系统调用，找出瓶颈。



> Cache

分析缓存命中率、缓存未命中等。

示例：程序运行缓慢，怀疑是缓存未命中导致的，可以用 `perf stat -e cache-misses ./my_program` 分析。



> 进程

分析进程间通信的性能。

示例：用 `perf record -e sched:sched_switch` 跟踪进程调度，分析进程切换的开销。



> 多线程和并发分析

分析多线程程序的性能问题，如线程争用、锁竞争等。

示例：用 `perf lock record` 记录锁使用情况，用 `perf lock report` 分析锁竞争。



> **示例**

首先，用 `perf stat` 查看整体的性能统计信息：

```
perf stat ./my_program
```

然后，用 `perf record` 记录详细的性能数据：

```
perf record ./my_program
```

最后，用 `perf report` 查看分析结果：

```
perf report
```



```
# Overhead  Command  Shared Object       Symbol
# ........  .......  ..................  ....................................
#
    25.45%  my_program  my_program         [.] main
    12.36%  my_program  libc.so.6          [.] __libc_start_main
    11.89%  my_program  my_program         [.] do_work
    ...
```



#### sar

```
-A：所有报告的总和
-u：输出CPU使用情况的统计信息
-v：输出inode、文件和其他内核表的统计信息
-d：输出每一个块设备的活动信息
-r：输出内存和交换空间的统计信息
-b：显示I/O和传送速率的统计信息-R：输出内存页面的统计信息
-y：终端设备活动情况
-w：输出系统交换活动信息
-B：显示换页状态；
-e：设置显示报告的结束时间
-f：从指定文件提取报告
-i：设状态信息刷新的间隔时间
-p：报告每个CPU的状态
-q：平均负载分析
```



> 平均负载统计分析

sar -q #查看平均负载：其中每间隔1秒钟统计一次总共统计三次 #

sar -q 1 3



> 内存统计分析

sar -r #查看内存使用情况，每间隔1秒钟统计一次总共统计三次：#

sar -r 1 3



> 磁盘IO

sar -b #查看I/O和传递速率的统计信息，每间隔1秒钟统计一次总共统计三次：#

sar -b 1 3

```
19:18:09          tps      rtps      wtps      dtps   bread/s   bwrtn/s   bdscd/s
19:18:10         0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:18:11         0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:18:12         0.00      0.00      0.00      0.00      0.00      0.00      0.00
Average:         0.00      0.00      0.00      0.00      0.00      0.00      0.00
```

> 磁盘使用情况

sar -d #磁盘使用详情统计，每间隔1秒钟统计一次总共统计三次：#

sar -d 1 3

```
19:17:48          DEV       tps     rkB/s     wkB/s     dkB/s   areq-sz    aqu-sz     await     %util
19:17:49          sda      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:17:49          sdb      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:17:49          sdc      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00

19:17:49          DEV       tps     rkB/s     wkB/s     dkB/s   areq-sz    aqu-sz     await     %util
19:17:50          sda      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:17:50          sdb      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:17:50          sdc      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00

19:17:50          DEV       tps     rkB/s     wkB/s     dkB/s   areq-sz    aqu-sz     await     %util
19:17:51          sda      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:17:51          sdb      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:17:51          sdc      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00

Average:          DEV       tps     rkB/s     wkB/s     dkB/s   areq-sz    aqu-sz     await     %util
Average:          sda      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
Average:          sdb      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
Average:          sdc      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
```

DEV 磁盘设备的名称，如果不加-p，会显示dev253-0类似的设备名称，因此加上-p显示的名称更直接
tps：每秒I/O的传输总数
rd_sec/s 每秒读取的扇区的总数
wr_sec/s 每秒写入的扇区的总数
avgrq-sz 平均每次次磁盘I/O操作的数据大小（扇区）
avgqu-sz 磁盘请求队列的平均长度
await 从请求磁盘操作到系统完成处理，每次请求的平均消耗时间，包括请求队列等待时间，单位是毫秒（1秒等于1000毫秒），等于寻道时间+队列时间+服务时间
svctm I/O的服务处理时间，即不包括请求队列中的时间
%util I/O请求占用的CPU百分比，值越高，说明I/O越慢



> 网络使用分析

sar -n #统计网络信息
sar -n选项使用6个不同的开关：DEV，EDEV，NFS，NFSD，SOCK，IP，EIP，ICMP，EICMP，TCP，ETCP，UDP，SOCK6，IP6，EIP6，ICMP6，EICMP6和UDP6 ，DEV显示网络接口信息，EDEV显示关于网络错误的统计数据，NFS统计活动的NFS客户端的信息，NFSD统计NFS服务器的信息，SOCK显示套接字信息，ALL显示所有5个开关。它们可以单独或者一起使用。



 sar -n DEV 1 1

```
19:26:08        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
19:26:09           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
19:26:09         eth0      1.00      0.00      0.21      0.00      0.00      0.00      1.00      0.00

Average:        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
Average:           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
Average:         eth0      1.00      0.00      0.21      0.00      0.00      0.00      1.00      0.00
```



> 进程，文件状态

sar -v #进程、inode、文件和锁表状态 ，每间隔1秒钟统计一次总共统计三次：#

sar -v 1 3



```
19:23:59    dentunusd   file-nr  inode-nr    pty-nr
19:24:00       128547      1696    132904         3
19:24:02       128547      1728    132904         3
19:24:03       128547      1728    132904         3
Average:       128547      1717    132904         3
```

dentunusd 在缓冲目录条目中没有使用的条目数量
file-nr 被系统使用的文件句柄数量
inode-nr 已经使用的索引数量
pty-nr 使用的pty数量





> 常用选项

```
默认监控: sar 5 5     //  CPU和IOWAIT统计状态 
(1) sar -b 5 5        // IO传送速率
(2) sar -B 5 5        // 页交换速率
(3) sar -c 5 5        // 进程创建的速率
(4) sar -d 5 5        // 块设备的活跃信息
(5) sar -n DEV 5 5    // 网路设备的状态信息
(6) sar -n SOCK 5 5   // SOCK的使用情况
(7) sar -n ALL 5 5    // 所有的网络状态信息
(8) sar -P ALL 5 5    // 每颗CPU的使用状态信息和IOWAIT统计状态 
(9) sar -q 5 5        // 队列的长度（等待运行的进程数）和负载的状态
(10) sar -r 5 5       // 内存和swap空间使用情况
(11) sar -R 5 5       // 内存的统计信息（内存页的分配和释放、系统每秒作为BUFFER使用内存页、每秒被cache到的内存页）
(12) sar -u 5 5       // CPU的使用情况和IOWAIT信息（同默认监控）
(13) sar -v 5 5       // inode, file and other kernel tablesd的状态信息
(14) sar -w 5 5       // 每秒上下文交换的数目
(15) sar -W 5 5       // SWAP交换的统计信息(监控状态同iostat 的si so)
(16) sar -x 2906 5 5  // 显示指定进程(2906)的统计信息，信息包括：进程造成的错误、用户级和系统级用户CPU的占用情况、运行在哪颗CPU上
(17) sar -y 5 5       // TTY设备的活动状态
(18) 将输出到文件(-o)和读取记录信息(-f)
```

###  network monitoring tools 

#### 选择工具的建议

- **实时流量监控**: 使用 ` `、`nload`、`bmon`。
- **详细流量分析**: 使用 `iptraf`、`tcpdump`、`wireshark`。
- **带宽测试**: 使用 `iperf3`。
- **延迟和连通性测试**: 使用 `ping`、`mtr`。
- **硬件参数调整**: 使用 `ethtool`。
- **网络安全扫描**: 使用 `nmap`。
- **长期趋势分析**: 使用 `sar`。

#### iptraf 

    -h, --help            show this help message
    
    -i <iface>            start the IP traffic monitor (use '-i all' for all interfaces)
    -d <iface>            start the detailed statistics facility on an interface
    -s <iface>            start the TCP and UDP monitor on an interface
    -z <iface>            shows the packet size counts on an interface
    -l <iface>            start the LAN station monitor (use '-l all' for all LAN interfaces)
    -g                    start the general interface statistics
    
    -B                    run in background (use only with one of the above parameters
    -f                    clear all locks and counters
    -t <n>                run only for the specified <n> number of minutes
    -L <logfile>          specifies an alternate log file

“iptraf -g” 显示每一个网卡上的流量



“iptraf -i eth0” 查看远程主机端



#### nethogs

**NetHogs** is an open-source command-line program (similar to Linux [top command](https://www.tecmint.com/find-linux-processes-memory-ram-cpu-usage/)) that is used to monitor real-time network traffic bandwidth used by each process or application in Linux.

```
yum install nethogs
```



### swap相关

`mkswap`
Creates a swap partition or file. The swap area must subsequently be enabled with swapon.
`swapon, swapoff`
Enable / disable swap partitition or file. These commands usually take effect at bootup and shutdown.



### grep 

grep -w "完整单词匹配"       #数字或字母。其字符就是单词的分隔符
