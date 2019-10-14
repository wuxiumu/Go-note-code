# 最新PHP 面试、笔试题汇总(code happy)
读在前面：
面向对象编程和面向对象设计的五个基本原则「SOLID」

- 单一职责原则（Single Responsibility Principle）
- 开放封闭原则（Open Closed Principle）
- 里氏替换原则（Liskov Substitution Principle）
- 接口隔离原则（Interface Segregation Principle）
- 依赖反转原则（Dependency Inversion Principle）

1. 单一职责原则规定一个类有且仅有一个理由使其改变。换句话说，一个类的边界和职责应当是十分狭窄且集中的。我很喜欢的一句话"在类的职责问题上，无知是福"。一个类应当做它该做的事，并且不应当被它的任何依赖的变化所影响。
2. 开放封闭原则，又称开闭原则，规定代码对扩展是开放的，对修改是封闭的。
3. 里氏替换原则规定对象可以被其子类的实例所替换，并且不会影响到程序的正确性。
4. 接口隔离原则规定，不应该强制接口的实现依赖于它不使用的方法。
5. 依赖反转原则，它规定高层次的代码不应该依赖低层级的代码。换句话说，高层次的代码应该依赖抽象接口，抽象接口就像是「中间人」一样，负责连接着高层次和低层次代码。这个原则的另一层意思是，抽象接口不应该依赖具体实现，但具体实现应该依赖抽象接口

## 一、秒杀（商品超卖，高并发，同一用户多次抢购）
### 后端：redis+队列
Redis相关知识点汇总：https://blog.csdn.net/weixin_39815001/article/details/99586988

redis队列实现，三个队列（库存队列，排队队列，抢购结果队列）

用户先进入排队队列，先进先出，判断是否已经在抢购结果队列，如果在，则直接下一个，如果不在，将用户信息加入抢购结果队列，库存-1，等待数据库空闲时，将抢购结果写入数据库

### 前端：
面对高并发的抢购活动，前端常用的三板斧是【扩容】【静态化】【限流】

扩容：加机器，这是最简单的方法，通过增加前端池的整体承载量来抗峰值。

静态化：将活动页面上的所有可以静态的元素全部静态化，并尽量减少动态元素。通过CDN来抗峰值。

限流：一般都会采用IP级别的限流，即针对某一个IP，限制单位时间内发起请求数量。或者活动入口的时候增加游戏或者问题环节进行消峰操作。

有损服务：在接近前端池承载能力的水位上限的时候，随机拒绝部分请求来保护活动整体的可用性。

### 二、订单模块（同一订单，多家商户结算问题）
订单拆分：用户支付后，将订单拆分，生成子订单

### 三、用户下单
先判断有没有登录

点击下单，生成唯一订单号，状态为未支付

## 四、接口安全
使用HTTP的POST方式,对固定参数+附加参数进行数字签名,使用的是md5加密,

比如:我想通过标题获取一个信息,在客户端使用 信息标题+日期+双方约定好的一个key通过md5加密生成一个签名(sign),然后作为参数传递到服务器端,

服务器端使用同样的方法进行校验,如何接受过来的sign和我们通过算法算的值相同，证明是一个正常的接口请求，我们才会返回相应的接口数据。

## 五、如何处理负载、高并发
### 1、HTML静态化
其实大家都知道，效率最高、消耗最小的就是纯静态化的html页面，所以我们尽可能使我们的 网站上的页面采用静态页面来实现，这个最简单的方法其实也是最有效的方法。

### 2、图片服务器分离
把图片单独存储，尽量减少图片等大流量的开销，可以放在一些相关的平台上，如骑牛等

### 3、数据库集群和库表散列及缓存
数据库的并发连接为100，一台数据库远远不够，可以从读写分离、主从复制，数据库集群方面来着手。另外尽量减少数据库的访问，可以使用缓存数据库如memcache、redis。

### 4、镜像：
尽量减少下载，可以把不同的请求分发到多个镜像端。

### 5、负载均衡：
Apache的最大并发连接为1500，只能增加服务器，可以从硬件上着手，如F5服务器。当然硬件的成本比较高，我们往往从软件方面着手。
负载均衡 （Load Balancing） 建立在现有网络结构之上，它提供了一种廉价有效透明的方法扩展网络设备和服务器的带宽、增加吞吐量、加强网络数据处理能力，同时能够提高网络的灵活性和可用性。目前使用最为广泛的负载均衡软件是Nginx、LVS、HAProxy。我分别来说下三种的优缺点:
ps: 负载均衡 Nginx、LVS、HAProxy

## 六、修改会话的生存时间

1. 在php.ini中 设置 session.gc_maxlifetime = 1440 //默认时间
2. 代码实现；
    ```
	$lifeTime = 24 * 3600; //保存一天
	session_set_cookie_params($lifeTime); 
	session_start();
    ```
 
## 七、PHP的垃圾收集机制
PHP可以自动进行内存管理，清除不再需要的对象。

PHP使用了引用计数(referencecounting)这种单纯的垃圾回收(garbagecollection)机制。

每个对象都内含一个引用计数器，每个reference连接到对象，计数器加1。当reference离开生存空间或被设为NULL，计数器减1。

当某个对象的引用计数器为零时，PHP知道你将不再需要使用这个对象，释放其所占的内存空间

## 八、正则的引擎
正则引擎主要可以分为两大类：一种是DFA，一种是NFA。

一般而论，DFA引擎则搜索更快一些。但是NFA以表达式为主导，更容易操纵，因此一般程序员更偏爱NFA引擎！

可以使用是否支持忽略优先量词和分组捕获来判断引擎类型：支持 NFA,不支持 DFA

## 九、对一个大文件进行逐行遍历，如下方法性能较高的是？
写一个实现了IteratorAggregate 接口的类，通过该类使用foreach遍历。
（使用 IteratorAggregate 可将文件打开后通过移动指针的方式逐行遍历，不受文件大小影响。使用 file_get_contents 处理大文件很容易导致PHP内存溢出；调用exec 会产生额外的进程，影响性能；其他人写的类库质量不一定高。）

## 十、读取文件加锁和解锁
```
$fp = fopen("lock.txt","w+");
    if (flock($fp,LOCK_EX)) {
        //获得写锁，写数据
        fwrite($fp, "write something");
        // 解除锁定
        flock($fp, LOCK_UN);
    } else {
        echo "file is locking...";
    }
    fclose($fp);
```

## 十一、array_merge() 数组合并函数
定义：array_merge() 函数把一个或多个数组合并为一个数组。（您可以向函数输入一个或者多个数组。）

注释：如果两个或更多个数组元素有相同的键名，则最后的元素会覆盖其他元素。
 
如果两个数组都是索引数组，则不会覆盖

如果您仅向 array_merge() 函数输入一个数组，且键名是整数，则该函数将返回带有整数键名的新数组，其键名以 0 开始进行重新索引。

该函数与 array_merge_recursive() 函数之间的不同是在处理两个或更多个数组元素有相同的键名的情况。array_merge_recursive() 不会进行键名覆盖，而是将多个相同键名的值递归组成一个数组。

示例；
    ```
	$a1=array("red","green");
    $a2=array("blue","yellow");
    $a3=array("CC","DD");

    $b1=array("a"=>"sa","b"=>"sb");
    $b2=array("a"=>"qa","b"=>"qb");
    $b3=array("a"=>"wa","c"=>"ww");
    
    print_r(array_merge($a1,$a2)); //Array ( [0] => red [1] => green [2] => blue [3] => yellow )
    print_r(array_merge($a1,$a2,$a3));  //Array ( [0] => red [1] => green [2] => blue [3] => yellow [4] => CC [5] => DD )
    print_r(array_merge($b1,$b2));  //Array ( [a] => qa [b] => qb )
    print_r(array_merge($b1,$b2,$a3));  // Array ( [a] => qa [b] => qb [0] => CC [1] => DD )
    print_r(array_merge($b1,$b2,$b3));  // Array ( [a] => wa [b] => qb [c] => ww )
    ```

## 十二、获取文件扩展名
```
//plan A
function  get_ext(string $url){
	//$url = 'http://www.sina.com.cn/abc/de/fg.html?id=1&ajksfg&aakzsdfj';
	$a = parse_url($url); //Array ( [scheme] => http [host] => www.sina.com.cn [path] => /abc/de/fg.html [query] => id=1&ajksfg&aakzsdfj )
	$file = basename($a['path']);  //fg.html
	$b = explode('.',$file);
	return array_pop($b);
}
//plan B
function  get_ext(string $url){
	//$url = 'http://www.sina.com.cn/abc/de/fg.html?id=1&ajksfg&aakzsdfj';
	$a = basename($url); //fg.html?id=1&ajksfg&aakzsdfj
	$b = explode('?',$a);;  //Array ( [0] => fg.html [1] => id=1&ajksfg&aakzsdfj )
	$ext = explode('.',$b[0]);
	return array_pop($ext);
}
```

## 十三、遍历一个文件夹下的所有文件和子文件夹
```
function my_scandir($dir){
	$files = array();
	if(is_dir($dir)){
		if ($handle = opendir($dir)){
			while (($file = readdir($handle))!= false){
				if ($file != "." && $file != "..") {
					if (is_dir($dir."/".$file)){
						$files[$file] = my_scandir($dir."/".$file);
					} else{
						$files[] = $dir."/".$file;
					}
				}
			}
			closedir($handle);
			return $files;
		}
	}
}
```

## 十四、编写一个函数，递归遍历，实现无限分类
```
function tree($arr,$pid=0,$level=0){
		static $list = array();
		foreach ($arr as $v) {
			//如果是顶级分类，则将其存到$list中，并以此节点为根节点，遍历其子节点
			if ($v['parent_id'] == $pid) {
				$v['level'] = $level;
				$list[] = $v;
				tree($arr,$v['cat_id'],$level+1);
			}
		}
		return $list;
	}
```

## 十五、获取上月的最后一天
```
function get_last_month_last_day($date = ''){
		if ($date != '') {
			$time = strtotime($date);
		} else {
			$time = time();
		}
		$day = date('j',$time);//获取该日期是当前月的第几天
		return date('Y-m-d',strtotime("-{$day} days",$time));
	}
```
 
## 十六、php中WEB上传文件的原理是什么，如何限制上传文件的大小？
上传文件的表单使用post方式，并且要在form中添加enctype=‘multipart/form-data’。

一般可以加上隐藏域：，位置在file域前面。

value的值是上传文件的客户端字节限制。可以避免用户在花时间等待上传大文件之后才发现文件过大上传失败的麻烦。

使用file文件域来选择要上传的文件，当点击提交按钮之后，文件会被上传到服务器中的临时目录，在脚本运行结束时会被销毁，所以应该在脚本结束之前，将其移动到服务器上的某个目录下，可以通过函数move_uploaded_file()来移动临时文件，要获取临时文件的信息，使用$_FILES。

限制上传文件大小的因素有：

客户端的隐藏域MAX_FILE_SIZE的数值（可以被绕开）。

服务器端的upload_max_filesize，post_max_size和memory_limit。这几项不能够用脚本来设置。

自定义文件大小限制逻辑。即使服务器的限制是能自己决定，也会有需要个别考虑的情况。所以这个限制方式经常是必要的。

## 十七、双引号和单引号的区别
双引号解释变量，单引号不解释变量

双引号里插入单引号，其中单引号里如果有变量的话，变量解释

双引号的变量名后面必须要有一个非数字、字母、下划线的特殊字符，或者用{}讲变量括起来，否则会将变量名后面的部分当做一个整体，引起语法错误

双引号解释转义字符，单引号不解释转义字符，但是解释’\和\
能使单引号字符尽量使用单引号，单引号的效率比双引号要高（因为双引号要先遍历一遍，判断里面有没有变量，然后再进行操作，而单引号则不需要判断）

## 十八、常用的超全局变量
```
$_GET ----->get传送方式
$_POST ----->post传送方式
$_REQUEST ----->可以接收到get和post两种方式的值
$GLOBALS ----->所有的变量都放在里面
$_FILES ----->上传文件使用
$_SERVER ----->系统环境变量
$_SESSION ----->会话控制的时候会用到
$_COOKIE ----->会话控制的时候会用到
```

## 十九、echo、print_r、print、var_dump之间的区别
* echo、print是php语句，var_dump和print_r是函数
* echo 输出一个或多个字符串，中间以逗号隔开，没有返回值是语言结构而不是真正的函数，因此不能作为表达式的一部分使用
* print也是php的一个关键字，有返回值 只能打印出简单类型变量的值(如int，string)，如果字符串显示成功则返回true，否则返回false
* print_r 可以打印出复杂类型变量的值(如数组、对象）以列表的形式显示，并以array、object开头，但print_r输出布尔值和NULL的结果没有意义，因为都是打印"\n"，因此var_dump()函数更适合调试
* var_dump() 判断一个变量的类型和长度，并输出变量的数值

## 二十、对于大流量网站，采用什么方法来解决访问量的问题
确认服务器硬件是否能够支持当前的流量

数据库读写分离，优化数据表

优化SQL语句

禁止外部盗链

控制大文件的下载

使用不同主机分流主要流量

使用流量分析统计

## 二十一、语句include和require的区别
require是无条件包含，也就是如果一个流程里加入require，无论条件成立与否都会先执行require，当文件不存在或者无法打开的时候，会提示错误，并且会终止程序执行

include有返回值，而require没有(可能因为如此require的速度比include快)，如果被包含的文件不存在的化，那么会提示一个错误，但是程序会继续执行下去

注意:包含文件不存在或者语法错误的时候require是致命的，而include不是
require_once，include_once表示了只包含一次，避免了重复包含

## 二十二、php中传值与传引用的区别，并说明传值什么时候传引用
变量默认总是传值赋值，那也就是说，当将一个表达式的值赋予一个变量时，整个表达式的值被赋值到目标变量，这意味着：当一个变量的赋予另外一个变量时，改变其中一个变量的值，将不会影响到另外一个变量
php也提供了另外一种方式给变量赋值：引用赋值。这意味着新的变量简单的引用(换言之，成为了其别名或者指向)了原始变量。改动的新的变量将影响到原始变量，反之亦然。

使用引用赋值，简单地将一个&符号加到将要赋值的变量前(源变量)
对象默认是传引用

对于较大的数据，可以考虑传引用，这样可以节省内存的开销

## 二十三、PHP 不使用第三个变量实现交换两个变量的值
```
//方法一
$a.=$b;
$b=str_replace($b,"",$a);
$a=str_replace($b,"",$a);

//方法二
list($b,$a)=array($a,$b);
var_dump($a,$b);
``` 


## 二十四、mysql优化
MySQL查询SQL优化

## 二十五、redis 和 memache 缓存的区别
1. 数据类型
    
    redis支持多种数据类型（5种）：hash string list set zset
    
    memcache 只支持key-value
2. 持久性
    
    redis 支持两种持久化方式 RDB、AOF
    
    memcache 不支持持久化
3. 分布式存储
    
    redis支持master-slave复制模式
    
    memcache可以使用一致性hash做分布式
4. value大小不同
    
    memcache是一个内存缓存，key的长度小于250字符，单个item存储要小于1M，不适合虚拟机使用
5. 线程模型
    
    memcache是master+worker的线程模型，其中master完成网络监听后投递到worker线程，由worker线程处理
    
    redis是单进程单线程模型，即单个线程完成所有的事情

    这两种实现造成下面的差异，即redis更容易实现多种数据结构，类似列表，集合，hash，有序集合等，由于是单线程的，如果单实例部署redis，不能全面用到服务器多核的优势，通常部署时，都会通过多实例的方式去部署

6. 内存管理
    
    redis：redis没有自己得内存池，而是直接使用时分配，即什么时候需要什么时候分配，内存管理的事交给内核，自己只负责取和释放，直接malloc和free即可。内存管理没有什么特殊的算法，通过使用google的jmalloc库来做内存管理（申请，释放）
    
    memcache：memcached是有自己得内存池的，即预先分配一大块内存，然后接下来分配内存就从内存池中分配，这样可以减少内存分配的次数，提高效率，这也是大部分网络服务器的实现方式，只不过各个内存池的管理方式根据具体情况而不同。使用了类似linux的内存管理，即slab内存管理方式。
7. 其他
redis支持事务，频道（发布-订阅），集群；memcache不支持

ps：https://cloud.tencent.com/developer/article/1004377

## 二十六、apche 和 nginx 的优缺
nginx轻量级，比apache占用更少的内存及资源，抗并发

nginx处理请求是异步非阻塞的，而apache 则是阻塞型的，在高并发下nginx 能保持低资源低消耗高性能。

apache 相对于nginx 的优点：
rewrite比nginx 的rewrite 强大，少bug，稳定。（需要性能用nginx，求稳定就apache）。

## 二十七、一个函数的参数不能是对变量的引用，除非在php.ini中把 allow_call_time_pass_reference 设为on。

在PHP函数调用的时候，基本数据类型默认会使用值传递，而不是引用传递。allow_call_time_pass_reference 选项的作用为是否启用在函数调用时强制参数被按照引用传递。如果把allow_call_time_pass_reference 配置为on，那么在函数调用的时候会默认使用引用传值。但是不推荐使用这种方法，原因是该方法在未来的版本中很可能不再支持。如果想使用引用传递，那么推荐在函数调用的时候显式地使用&进行引用传递。

## 二十八、什么是内存管理？
内存管理主要是指程序运行时对计算机内存资源的分配、使用和释放等技术，内存管理的目标是高效、快速地分配内存同时及时地释放和回收内存资源。内存管理主要包括是否有足够的内存供程序使用，从内存池中获取可用内存，使用后及时销毁并重新分配给其他程序使用。

在PHP开发过程中，如果遇到大数组等操作，那么可能会造成内存溢出等问题。一些常见的处理方法如下：

1）通过ini_set(‘memory_limit’,‘64M’)方法重置php可以使用的内存大小，一般在远程主机上是不能修改php.ini文件的，只能通过程序设置。注：在safe_mode（安全模式）下，ini_set会失效。

2）另一方面可以对数组进行分批处理，及时销毁无用的变量，尽量减少静态变量的使用，在需要数据重用时，可以考虑使用引用（&）。同时对于数据库、文件操作完要及时关闭，对象使用完要及时调用析构函数等。

3）及时使用unset()函数释放变量，使用时需要注意以下两点：
① unset()函数只能在变量值占用内存空间超过256字节时才会释放内存空间。
② 只有当指向该变量的所有变量都销毁后，才能成功释放内存。

## 二十九、Memcache的特征和特性
```
1）协议简单。
2）基于libevent的事件处理。
3）内置内存存储方式。
4）Memcached不互相通信的分布式。
（1）单个item 最大的数据为1MB。
（2）单进程最大的使用内存为2GB，需要更多内存时可开多个端口。
（3）Memcached是多线程，非阻塞io复用的网络模型，Redis是单线程。
（4）键长最大为250字节。
```
## 三十、共享Session的方式
```
1）基于NFS的Session共享。NFS（Network File System）最早由Sun公司为解决Unix网络主机间的目录共享而研发。仅需将共享目录服务器mount到其他服务器的本地session目录即可。
2）基于数据库的Session共享。
3）基于Cookie的Session共享。原理是将全站用户的Session信息加密、序列化后以Cookie的方式，统一种植在根域名下（如：.host.com），利用浏览器访问该根域名下的所有二级域名站点时，会传递与之域名对应的所有Cookie内容的特性，从而实现用户的Cookie化Session 在多服务间的共享访问。
4）基于缓存（Memcache）的Session共享。Memcache是一款基于Libevent多路异步I/O技术的内存共享系统，简单的key + value数据存储模式使得代码逻辑小巧高效，因此在并发处理能力上占据了绝对优势，目前能达到2000/s平均查询，并且服务器CPU消耗依然不到10%。
```
## 三十一、memcache或redis雪崩如何解决？
造成原因：通常，在一个网站里，mysql数据库处理的请求比较少（20%），负载80%，缓存技术处理大多数请求（80%）

如果memcache或redis挂掉，所有请求都会在mysql处理，数据库的处理能力不足会直接宕机。这时候就算重启缓存和mysql也是无济于事的，因为缓存重启后，数据已经丢失，数据请求还是会走mysql，mysql还是会死掉（死循环）

解决方法：

缓存预热
```
1：先启动缓存，再启动数据库。（但是此时不提供对外服务）
2：通过一个PHP脚本把常用的key写入缓存中
3：开放对外服务【热点数据已经缓存，请求会被缓存处理，减轻mysql压力】
```

## 三十二、Redis持久化的方式？
1. Aof(append only file)

    redis执行命令时，会把我们执行的命令通过日志形式进行追加。安全性高，但是影响性能。
2. Rdb
    
    按照制定规则进行持久化
    save 900 1 (900s内1次redis操作 会做一次持久化)
    save 300 10 (300s内10次redis操作 会做一次持久化)
    save 60 10000 (60s内10000次redis操作 会做一次持久化)
    但是可能会存在数据丢失，比如：12:00做过一次持久化，正常的话，12:15会再做持久化，如果12:14缓存死掉，那么14分钟的数据会丢失。不大安全，但是性能比aof好很多

## 三十三、Linux系统中，进程间通信的方式。
管道：

管道分为有名管道和无名管道

无名管道是一种半双工的通信方式,数据只能单向流动,而且只能在具有亲缘关系的进程间使用.进程的亲缘关系一般指的是父子关系。无明管道一般用于两个不同进程之间的通信。当一个进程创建了一个管道,并调用fork创建自己的一个子进程后,父进程关闭读管道端,子进程关闭写管道端,这样提供了两个进程之间数据流动的一种方式。

有名管道也是一种半双工的通信方式,但是它允许无亲缘关系进程间的通信。

消息队列：

消息队列是消息的链表,存放在内核中并由消息队列标识符标识.消息队列克服了信号传递信息少,管道只能承载无格式字节流以及缓冲区大小受限等特点.消息队列是UNIX下不同进程之间可实现共享资源的一种机制,UNIX允许不同进程将格式化的数据流以消息队列形式发送给任意进程.对消息队列具有操作权限的进程都可以使用msget完成对消息队列的操作控制.通过使用消息类型,进程可以按任何顺序读信息,或为消息安排优先级顺序.

信号：

信号是一种比较复杂的通信方式,用于通知接收进程某个事件已经发生.

信号量：

信号量是一个计数器,可以用来控制多个线程对共享资源的访问.,它不是用于交换大批数据,而用于多线程之间的同步.它常作为一种锁机制,防止某进程在访问资源时其它进程也访问该资源.因此,主要作为进程间以及同一个进程内不同线程之间的同步手段.

共享内存：

共享内存就是映射一段能被其他进程所访问的内存,这段共享内存由一个进程创建,但多个进程都可以访问.共享内存是最快的IPC(进程间通信)方式,它是针对其它进程间通信方式运行效率低而专门设计的.它往往与其他通信机制,如信号量,配合使用,来实现进程间的同步与通信.

socket：

可用于不同及其间的进程通信
文件，互斥量等，不过我在swoole源码中看到了通过eventfd这种方式做进程通信的

## 三十四、HTTP Header 详解

## 三十五、TCP 三次握手和四次挥手

## 三十六、 Select、Poll、Epoll
大话 Select、Poll、Epoll

## 三十七、海量数据处理相关总结
1、海量日志数据，提取出某日访问百度次数最多的那个IP。

算法思想：分而治之+Hash

IP地址最多有2^32=4G种取值情况，所以不能完全加载到内存中处理
可以考虑采用“分而治之”的思想，按照IP地址的Hash(IP)%1024值，把海量IP日志分别存储到1024个小文件中。这样，每个小文件最多包含4MB个IP地址
对于每一个小文件，可以构建一个IP为key，出现次数为value的Hash map，同时记录当前出现次数最多的那个IP地址
可以得到1024个小文件中的出现次数最多的IP，再依据常规的排序算法得到总体上出现次数最多的IP
2… PS：https://blog.csdn.net/v_JULY_v/article/details/6279498

## 三十八、两台mysql服务器，其中一台挂了，怎么让业务端无感切换，并保证正常情况下讲台服务器的数据是一致的

不是核心业务的话，先停写，把备机拉起来，查看两台机器的日志，进行数据补偿，开写。

如果是核心业务的话，现在所有的写操作都在正常的状态机器上。把好的这台机器的备机拉起来，当主机。

备机的数据不一致怎么办？

你要勇敢怼回去，你们每秒多少写入操作。按照百万级表，每秒1000的写入效率，正常的设计是，分布在2台机器上每台500。这个级别的数据同步，出现差异的概率 可以忽略不计的。有一台出现问题，另一台也可以抗住。

## 三十九、redis是如何进行同步的，同步的方式，同步回滚怎么办，数据异常怎么办

redis 集群主从同步的简单原理

Redis的复制功能是基于内存快照的持久化策略基础上的，也就是说无论你的持久化策略选择的是什么，只要用到了Redis的复制功能，就一定会有内存快照发生

　　当Slave启动并连接到Master之后，它将主动发送一个SYNC命令( 首先Master会启动一个后台进程，将数据快照保存到文件中[rdb文件] Master 会给Slave 发送一个

Ping命令来判断Slave的存活状态 当存活时 Master会将数据文件发送给Slave 并将所有写命令发送到Slave )。Slave首先会将数据文件保存到本地之后再将数据加载到内存中。

　　当第一次链接或者是故障后，重新连接都会先判断Slave的存活状态再做全部数据的同步，之后只会同步Master的写操作(将命令发送给Slave)
问题：

　　当 Master 同步数据时 若数据量较大 而Master本身只会启用一个后台进程 来对多个Slave进行同步 ， 这样Master就会压力过大 ， 而且Slave 恢复的时间也会很慢！

redis 主从复制的优点：

(1)在一个Redis集群中，master负责写请求，slave负责读请求，这么做一方面通过将读请求分散到其他机器从而大大减少了master服务器的压力，另一方面slave专注于提供

读服务从而提高了响应和读取速度。

(2)在一个Redis集群中，如果master宕机，slave可以介入并取代master的位置，因此对于整个Redis服务来说不至于提供不了服务，这样使得整个Redis服务足够安全。

(3)水平增加Slave机器可以提高性能

## 四十、如何解决跨域
JSONP

添加响应头，允许跨域

代理的方式


## 四十一、写出以下输出
```
Q: "aa" == 1, "aa" == 0, 1 == "1", 1==="1", "12asdsad" + 1, "asdjkfgj12"+1
A: false, true, true, false, 13, 1
```
why:

php中 字符串==0 恒成立

php中 字符串和数字相加，如果字符串开头是数字，则等于字符串开头的数字（字符串第一个位置开始，到第一个非数字和.的位置截止）+数字

### 四十二、什么是服务容器、控制反转（IoC）、依赖注入（DI）
服务容器是用来管理类依赖与运行依赖注入的工具。Laravel框架中就是使用服务容器来实现 控制反转 和 依赖注入 。

控制反转(IoC) 就是说把创建对象的 控制权 进行转移，以前创建对象的主动权和创建时机是由自己把控的，而现在这种权力转移到第三方，也就是 Laravel 中的容器。

依赖注入（DI）则是帮助容器实现在运行中动态的为对象提供提依赖的资源。

## 四十三、Composer自动加载原理
composer加载核心思想是通过composer的配置文件在引用入口文件(autoload.php)时,将类和路径的对应关系加载到内存中,最后将具体加载的实现注册到spl_autoload_register函数中.最后将需要的文件包含进来.

## 四十四、一个请求到PHP，Nginx的主要过程。完整描述整个网络请求过程，原理。

1)、FastCGI进程管理器php-fpm自身初始化，启动主进程php-fpm和启动start_servers个CGI 子进程。主进程php-fpm主要是管理fastcgi子进程，监听9000端口。fastcgi子进程等待来自Web Server的连接。

2)、当客户端请求到达Web Server Nginx是时，Nginx通过location指令，将所有以php为后缀的文件都交给127.0.0.1:9000来处理，即Nginx通过location指令，将所有以php为后缀的文件都交给127.0.0.1:9000来处理。

3）FastCGI进程管理器PHP-FPM选择并连接到一个子进程CGI解释器。Web server将CGI环境变量和标准输入发送到FastCGI子进程。

4)、FastCGI子进程完成处理后将标准输出和错误信息从同一连接返回Web Server。当FastCGI子进程关闭连接时，请求便告处理完成。

5)、FastCGI子进程接着等待并处理来自FastCGI进程管理器（运行在 WebServer中）的下一个连接。

## 四十五、PHP的魔术方法
```
__set() // 在给不可访问属性赋值时，__set()会被调用
__get() // 读取不可访问属性的值时，__get()会被调用
__isset() //当对不可访问属性调用isset()或empty()，__isset()会被调用
__unset() // 当对不可访问属性调用unset()时，__unset()会被调用
__call() // 在对象中调用一个不可访问方法时，__call()会被调用
__callStatic() // 在静态上下文中调用一个不可访问的方法时，__callStatic会被调用
__construct() // 构造函数的类会在每次创建新对象时先调用此方法，所以非常适合在使用对象之前做一些初始化工作。
__destruct() // 析构函数会在到某个对象的所有引用都被删除或者当对象被显式销毁时执行。
__sleep() // serialize()函数会检查类中是否存在一个魔术方法__sleep()，如果存在，该方法会先被调用，然后再执行序列化操作。此功能可以用于清理对象，并返回一个包含对象中所有应被序列化的变量名称的数组。如果该方法未返回任何内容，则 NULL 被序列化，并产生一个 E_NOTICE 级别的错误。
__wakeup() // unserialize()函数会检查是否存在一个__wakeup()方法，如果存在，则会先调用该方法，然后再执行反序列化操作。__wakeup() 经常用在反序列化操作中，例如重新建立数据库连接，或执行其它初始化操作。
```

## 四十六、字符编码UTF8、GBK、GB2312的区别。
utf8是国际编码。通用性较好。

gbk是国内编码。通用型较utf8差，但是占用数据库比utf8小。

gb2312是一个简体中文字符集的中国国家标准，共收录6763个汉字。

## 四十七、MySQL默认的排序方式是什么
MyIsam存储引擎：在没有任何删除，修改的操作下，执行select不带order by那么会按照插入的顺序下进行排序。

InnDB存储引擎：在相同的情况下，select不带order by会根据主键来排序，从小到大。

## 四十八、OSI七层网络模型
物理层：建立、维护、断开物理连接

数据链路层：建立逻辑链接、进行硬件地址寻址、差错校验等功能（SDLC、HDLC、PPP、STP）

网络层：进行逻辑地址寻址，实现不同网络之间的路径选择（IP、IPX、OSPF）

传输层：定义传输数据的协议端口号，以及流程和差错校验（TCP,UDP）数据包一旦离开网卡即进入网络传输层

会话层：建立、管理、终止会话

表示层：数据的表示、安全、压缩

应用层：网络服务与最终用户的一个接口；协议有：HTTP、FTP、TFTP、SMTP、DNS、TELNET、HTTPS、POP3、DHCP

## 四十九、找出数组中出现一次的元素。10 10 11 11 12 13 12 13 16 只出现一次的数字。要求时间复杂度尽可能低
```
// 方法一
function onlyOne($arr) {
	$res = 0;
	for ($i = 0; $i < count($arr); $i++) {
		$res ^= $arr[$i];
	}
	
   return $res;
}
// 方法二
function onlyOne2($arr) {
	$m = array_count_values($arr);
	foreach ($m as $k => $v) {
	  if ($v == 1) {
		return $k;
	  }
	}
	return 0;
}
```  

## 五十、LRU算法
如果一个 数据在最近一段时间没有被访问到，那么在将来它被访问的可能性也很小（https://www.twblogs.net/a/5b7f0b662b717767c6ad6c42/zh-cn）

杂项：

获取客户端IP

没有使用代理

$_SERVER[‘REMOTE_ADDR’] 或者 getenv(‘REMOTE_ADDR’)

使用透明代理

$_SERVER[‘HTTP_X_FORWARDED_FOR’];

获取服务端IP

$_SERVER[‘SERVER_ADDR’] 或者 gethostbyname(‘www.baidu.com’);

将IP地址转换成int

ip2long($ip);  


好处：存储时可以直接存有符号int型，只需要4字节(节约空间)

存char类型需要15个字节

int转换成ip

long2ip($int);


获取当前时间戳

time()


打印前一天的时间

date('Y-m-d H:i:s',strtotime('-1 day'))


GB2312格式的字符串装换成UTF-8格式

iconv('GB2312','UTF-8','悄悄是别离的笙箫');


字符串转数组

explode(',',$str)


创建一个每周三01:00~04:00每3分钟执行执行一次的crontab指令
```
*/3 1-4 * * 3 /bin/bash /home/sijiaomao/ok.sh
```

php两数组相加
两个数组相加，若数组中存在相同键值的元素，则只保留第一个数组的元素

设置时区

	date_default_timezone_set("Asia/Shanghai");
 
在url中用get传值的时候，若中文出现乱码，应该用哪个函数对中文进行编码？

urlencode() 解码用urldecode()

PHP常用扩展

curl、iconv、json、mbstring、mysqli、PDO、hash、openssl、sockets、redis

php八种数据类型
```        
数据类型分为三种：
标量数据类型：boolean、string、integer、double
复合数据类型：array、object
特殊数据类型：resource、null
```
php进程模型，php怎么支持多个并发

守护进程模型：https://www.jianshu.com/p/542935a3bfa8

nginx的进程模型，怎么支持多个并发

https://www.zhihu.com/question/22062795

php-fpm各配置含义，fpm的daemonize模式

http://www.4wei.cn/archives/1002061

查看PHP进程当前使用的内存

memory_get_usage()

查看系统分配的内存

memory_get_peak_usage()

unset()可以释放内存，当处理完数据后 unset 掉，只能释放当前使用的内存，系统分配的内存并不会变小

内存被分划为， “已使用” 和 “空闲”， unset 只会把 “已使用” 变为 “空闲”， 下次内存请求时会先去"空闲"里取，

程序结束， GC 才会释放全部内存

参数绑定可以避免 SQL 注入攻击

		例如：$users = DB::select('select * from users where username = ? and passwd = ?', [$username,$passwd]);


除了使用 ? 占位符来代表参数绑定外，还可以使用命名绑定来执行查询：

		$results = DB::select('select * from users where id = :id', ['id' => 1]);


thinkphp5.0链接数据库

使用db助手函数默认每次都会重新连接数据库，而使用Db::name或者Db::table方法的话都是单例的

db函数如果需要采用相同的链接，可以传入第三个参数，例如：
db(‘user’,[],false)->where(‘id’,1)->find();

db(‘user’,[],false)->where(‘status’,1)->select();

上面的方式会使用同一个数据库连接，第二个参数为数据库的连接参数，留空表示采用数据库配置文件的配置。

PHP预定义变量（详见：https://www.php.net/manual/zh/reserved.variables.php）

   超全局变量 — 超全局变量是在全部作用域中始终可用的内置变量
```   
	$GLOBALS — 引用全局作用域中可用的全部变量
	$_SERVER — 服务器和执行环境信息
	$_GET — HTTP GET 变量
	$_POST — HTTP POST 变量
	$_FILES — HTTP 文件上传变量
	$_REQUEST — HTTP Request 变量
	$_SESSION — Session 变量
	$_ENV — 环境变量
	$_COOKIE — HTTP Cookies
	$php_errormsg — 前一个错误信息
	$HTTP_RAW_POST_DATA — 原生POST数据
	$http_response_header — HTTP 响应头
	$argc — 传递给脚本的参数数目
	$argv — 传递给脚本的参数数组
```
 