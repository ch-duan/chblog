import os

src_path = './f429_app_test1.bin'
des_path = './4.txt'


def ReadFile():
    binfile = open(src_path, 'rb')  # 打开二进制文件
    size = os.path.getsize(src_path)  # 获得文件大小
    Note = open(des_path, mode='w')
    print(size)  # 打印文件大小
    for i in range(size):  # 遍历输出文件内容
        data = binfile.read(1)  # 每次输出一个字节
        n = int.from_bytes(data, byteorder='big')
        # print('{:02X} '.format(n), end="")
        Note.write('{:02X} '.format(n))
        if (i+1) % 1024 == 0:
            # print("")
            Note.write("\n")
    binfile.close()  # close文件
    Note.close()


if __name__ == '__main__':
    ReadFile()
