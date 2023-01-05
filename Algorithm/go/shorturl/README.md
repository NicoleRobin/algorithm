# 短连接系统设计

## 要点
1、长URL到短URL的映射，需要确保唯一，而hash是无法确保唯一的，因此可以考虑通过给每一个长URL指定一个整数ID来实现，并通过将整数ID转换为62进制的值来实现减小
短URL的长度，10位的短URL支持62^10个URL，6千2百亿，绝对满足要求。
2、短URL到长URL的映射需要保存下来，因为当后续通过短URL进行访问时需要能返回长URL并返回HTTP status code指示客户端跳转。  
3、HTTP status code返回301还是302？301：永久重定向；302：临时重定向，因此选择301还是302取决于是否需要统计短URL的每次访问。  
4、  

## 参考资料
1、http://n00tc0d3r.blogspot.com/  
2、https://stackoverflow.com/questions/742013/how-do-i-create-a-url-shortener  