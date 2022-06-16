# 基本数据类型

## string
源码位置：`runtime/string.go`

## slice
源码位置：`runtime/slice.go`
``` go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
`slice`是go中的动态数组，支持动态扩容，  
`slice`结构体中包含了三个变量，分别是slice的长度、容量和指向底层数组的指针。

## map
源码位置：`runtime/map.go`
go中的map基于散列表实现，具体结构有两个：`hmap`、`bmap`。  
每个hmap包含大小为2^B的数组，通过每个key计算出的hash值的低B位来定位在数组中的位置。
每个bmap是一个包含8个槽位的桶，每个槽位存储了根据key计算出的hash值的高8位，查找时也通过对比高8位来确定是否是要找的key。

## chan

# 基本语法
## panic&recover

## make&new

## switch

## select
select用于实现同时监听多个channel可读、写，而且也可以实现对channel进行非阻塞的读、写。  



# GC

# GMP
