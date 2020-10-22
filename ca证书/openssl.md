# OPENSSl

#### 对称加密

对称加密需要使用的标准命令为 enc ，用法如下：

```go
openssl enc -ciphername [-in filename] [-out filename] [-pass arg] [-e] [-d] [-a/-base64]
       [-A] [-k password] [-kfile filename] [-K key] [-iv IV] [-S salt] [-salt] [-nosalt] [-z] [-md]
       [-p] [-P] [-bufsize number] [-nopad] [-debug] [-none] [-engine id]
```

常用选项有：
-in filename：指定要加密的文件存放路径  
-out filename：指定加密后的文件存放路径  
-salt：自动插入一个随机数作为文件内容加密，默认选项  
-e：可以指明一种加密算法，若不指的话将使用默认加密算法  
-d：解密，解密时也可以指定算法，若不指定则使用默认算法，但一定要与加密时的算法一致  
-a/-base64：使用-base64位编码格式  

1 |示例：  
2 |加密：openssl enc -e -des3 -a -salt -in fstab -out jiami  
3 |解密：openssl enc -d -des3 -a -salt -in fstab -out jiami  

###### 生成key

openssl genrsa -out ca.key 2048  

##### 生成证书

openssl req -x509 -new -nodes -key ca_key.pem -sha256 -days 3650 -subj "/CN=www.emqx.io" -out ca.pem  

##### 生成请求

req -new -key server_key.pem -out server.csr -subj "/CN=mq.letsit.vip"

##### 签名

openssl x509 -req -in ./server.csr -CA ca.pem -CAkey ca_key.pem -CAcreateserial -out server.pem -days 3650 -sha256

openssl x509 -extensions v3_ca -req -in ./server.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out server.pem -days 3650 -sha256

##### 生成能签发子证书的二级证书

openssl ca -extensions v3_ca -in roota.csr -out roota.crt -cert CARoot.crt -keyfile CARoot.key -days 1826 -policy policy_anything
