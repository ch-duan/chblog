# centos cmd

##### 给某个用户文件权限

u代表用户，g代表组
setfacl -R -m u:username:rwx your_file_or_fir

##### 查看文件每个用户权限

getfacl your_file_or_fir

