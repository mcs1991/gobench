**gobench测试**

**准备测试环境:**

指定数据库gobench(-db,没有会自动创建)，设置并发数4(-t)，readwrite模式(-m)，创建10张表(-tbcount)，每张表1万条数据(-tbsize)。

    ./gobench -u root -p 123456 -H 192.168.56.2 -P 3310 -db gobench -t 4 -tbsize 10000 -tbcount 10 -m readwrite -c prepare

**一般测试：**

启动一次压测，设置并发数为4(-t)，指定gobench库(-db)，压测10张表(-tbcount)，每张表1万数据(-tbsize)，readwrite模式(-m)，持续20分钟(-time),压测结果保存至当前路径下gobench.log中(-l)。

    ./gobench -u root -p 123456 -H 192.168.56.2 -P 3310 -db gobench -t 4 -tbsize 10000 -tbcount 10 -m readwrite -c run -time 1200 -l ./gobench.log

**自动化测试:**

**扩展功能1：**

启动自动压测模式(-f),设置并发数4(-t),连续压测3次(-C),每次压测持续20分钟(-time),压测间隔时间60秒(-s)，指定gobench库(-db)，压测10张表(-tbcount)，每张表1万数据(-tbsize)，readwrite模式(-m)。压测结果保存至当前路径下gobench_1.log中(-l)。

    ./gobench -u root -p 123456 -H 192.168.56.2 -P 3310 -db gobench -t 4 -tbsize 10000 -tbcount 10 -m readwrite -c run -time 1200 -C 3 -f -s 60 -l ./gobench_1.log

**扩展功能2:**

启动自动压测模式(-f),设置并发数分别为2,4,8(-t,以","分隔),每一个并发压测3次(-C)，一共进行9次压测，每次压测持续20分钟(-time)，压测间隔时间60秒(-s)。指定gobench库(-db)，压测10张表(-tbcount)，每张表1万数据(-tbsize)，readwrite模式(-m)。压测结果保存至当前路径下gobench_2.log中(-l)

    ./gobench -u root -p 123456 -H 192.168.56.2 -P 3310 -db gobench -t 2,4,8 -tbsize 10000 -tbcount 10 -m readwrite -c run -time 1200 -C 3 -f -s 60 -l ./gobench_3.log

**扩展功能3:**

启动自动压测模式(-f),设置并发数分别为2,4,8(-t,以","分隔)，每个并发数分别执行3次(-C)readwrite,readonly模式(-m,以","分隔)，一共进行18次压测，每次压测持续20分钟(-time)，压测间隔时间60秒(-s)。指定gobench库(-db)，压测10张表(-tbcount)，每张表1万数据(-tbsize)，readwrite模式(-m)。压测结果保存至当前路径下gobench_3.log中(-l)

    ./gobench -u root -p 123456 -H 192.168.56.2 -P 3310 -db gobench -t 2,4,8 -tbsize 10000 -tbcount 10 -m readwrite,readonly -c run -time 1200 -C 3 -f -s 60 -l ./gobench_3.log

**清除测试环境**

	./gobench -u root -p 123456 -H 192.168.56.2 -P 3310 -db gobench -t 4 -tbsize 10000 -tbcount 10 -c -m readwrite -c cleanup

**Fileio 测试:**

**准备测试环境:**

启动磁盘io压测功能(-fileio),设置并发数4，准备压测文件数10（-filenum）,压测文件总大小10g（-ftotalsize,最好大于内存）,压测文件块大小16k(-fblocksize)。

    ./gobench -fileio -filenum 10 -ftotalsize 10g -fblocksize 16384 -t 4 -c prepare

**一般测试：**

启动磁盘io压测功能(-fileio),设置并发数4(-t)，压测文件数10(-filenum)，压测文件总大小10g(-ftotalsize)，压测文件块大小16k（-fblocksize）,压测模式为顺序写(-ftestmode)，压测时间20分钟(-time)，压测结果保存至当前路径下gobench_fio.log(-l)。

    ./gobench -fileio -filenum 10 -ftotalsize 10g -fblocksize 16384 -ftestmode seqwr -c run -t 4 -time 1200 -l ./gobench_fio.log

**自动化测试：**

**扩展功能1：**

启动磁盘io压测功能(-fileio),设置并发数4(-t)，压测文件数10(-filenum)，压测文件总大小10g(-ftotalsize)，压测文件块大小16k（-fblocksize）,压测模式分别为顺序写，顺序读，随机写，随机读(-ftestmode,以","分隔)，每一个压测模式压测两次(-C)，每次压测时间20分钟(-time)，压测结果保存至当前路径下gobench_fio.log(-l)。

    ./gobench -fileio -filenum 10 -ftotalsize 10g -fblocksize 16384 -ftestmode seqwr,seqrd,rndwr,rndrd -f -C 2 -c run -t 4 -time 1200 -l ./gobench_fio.log

**扩展功能2：**

启动磁盘io压测功能(-fileio),设置并发数分别为4,8,16(-t)，压测文件数10(-filenum)，压测文件总大小10g(-ftotalsize)，压测文件块大小16k（-fblocksize）,每一个并发数下分别执行压测模式:顺序写，顺序读，随机写，随机读(-ftestmode,以","分隔)，每一个压测模式压测3次(-C)，也就是一共会压测36次，每次压测时间20分钟(-time)，压测结果保存至当前路径下gobench_fio.log(-l)。

    ./gobench -fileio -filenum 10 -ftotalsize 10g -fblocksize 16384 -ftestmode seqwr,seqrd,rndwr,rndrd -f -C 3 -c run -t 4,8,16 -time 1200 -l ./gobench_fio.log