package project

/*
各类型所占字节数:

类型	       	  16位	      32 位	     64位
char	    	1	         1	      1
short int		2	         2	      2
int	        	2	         4	      4
unsigned int	2			 4	      4
float			4	         4	      4
double	        8	         8	      8
long	        4	         4	      8
long long	    8	         8	      8
unsigned long	4	         4	      8
------------------------------------------------------------------------------------------------------------------------------------
雪花算法: 用于生成分布式 id, 订单号等, 占8个字节(64 bit 即 64 位), 前41位代表时间戳, 后10位代表工作机器id, 最后12位代表序列号(最大正整数是2^12-1=4095)

uuid缺点, 无法保证按序递增；其次是太长了，存储数据库占用空间比较大，不利于检索和排序

分布式ID雪花算法的原理，即用一个long类型变量存储多个信息。一个long类型长度为8个字节（64bit），最高位是符号位，始终为0，不可用。
雪花算法使用其中41bit记录时间戳，其余bit位存储机房id、机器id、序列号。

Redis的ZSet支持分值为double类型，也是8字节，那么我们也可以使用41位存储时间戳，其实位存储用户的实际积分


在雪花算法中最高位是不用的，目的是不允许生成负数ID，而在实现排行榜中没有这个限制，因为我们最终要的只是用户的积分，而不是加上时间戳的分值
但也要求最高位要么全为0，要么全为1，避免排序错乱。如实现积分倒序排名时可设置最高位全为1，只不过ZSet已经支持倒序获取，不需要多此一举，所以最高位我们依然不使用。
*/
