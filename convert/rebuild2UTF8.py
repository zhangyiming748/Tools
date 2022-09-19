import chardet
import os


# 批量重命名文本文档并且统一文本文档编码为utf-8
def getTxt():
    path = os.listdir('.')
    for p in path:
        if not os.path.isdir(p):
             print(p)
        if p.split('.')[-1] == 'txt':
            readAndWrite(p)
        print(p)


def readAndWrite(file):
    with open(file, 'rb') as f:
        text = f.read()
        res = chardet.detect(text)
        print(res['encoding'])

    with open(file, 'r', encoding=res['encoding'], errors='replace') as f:
        content = f.read()
        content = content.replace('?', '')
        content = content.replace('�', '')
        content = content.replace('U', '')
        content = content.replace('@', '')
        content = content.replace('*', '')
        print(content)

#os.remove(file)

    with open(file, 'a+', encoding='utf8')as f:
        f.write(content)


def renameTxt():
    path = os.listdir('.')
    for p in path:
        if not os.path.isdir(p):
            pass# print(p)
        if p.split('.')[-1] == 'txt':
            pass# readAndWrite(p)
        print(p)
        old = p
        new = myReplace(p)
        os.rename(old, new)


def myReplace(p):
    p = p.replace('（', '(')
    p = p.replace('）', ')')
    p = p.replace('【', '')
    p = p.replace('】', '')
    p = p.replace('「', '{')
    p = p.replace('」', '}')
    p = p.replace('，', ',')
    p = p.replace('。', '')
    p = p.replace('！', '!')
    p = p.replace('？', '?')
    p = p.replace('《', '')
    p = p.replace('》', '')
    p = p.replace('~', '')
    p = p.replace('～', '')
    p = p.replace('－', '-')
    p = p.replace('_', '')
    p = p.replace('：', ':')
    p = p.replace('；', ';')
    p = p.replace('—', '')
    p = p.replace('０', '0')
    p = p.replace('１', '1')
    p = p.replace('２', '2')
    p = p.replace('３', '3')
    p = p.replace('４', '4')
    p = p.replace('５', '5')
    p = p.replace('６', '6')
    p = p.replace('７', '7')
    p = p.replace('８', '8')
    p = p.replace('９', '9')
    return p


if __name__ == '__main__':
# f = '单纯的.txt'
# readAndWrite(f)
    renameTxt()
    getTxt()
