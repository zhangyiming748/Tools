# excel_u.py
# 批量修改文件为markdown的其中一个步骤
# 导入相应模块
import xlrd
from xlutils.copy import copy

# 打开 excel 文件
readbook = xlrd.open_workbook("7026.xls")

# 复制一份
wb = copy(readbook)

# 选取第一个表单
sh1 = wb.get_sheet(0)

# 在第四行新增写入数据
# sh1.write(3, 0, '王亮')
# sh1.write(3, 1, 59)
for p in range(0, 6483):
    sh1.write(p, 14, '封面</a>')  # n
    sh1.write(p, 20, '截图</a>')  # s
    sh1.write(p, 26, '磁力</a>')  # x
    # sh1.write(p,21,'<a href="')
    # sh1.write(p,23,'">')
    # #sh1.write(p, 11, '<a href="')  # l
    # sh1.write(p, 13, '">')
    # sh1.write(p, 16, '<a href="')  # q
    # sh1.write(p, 21, '<a href="')  # v
    # sh1.write(p, 23, '">')  # x
    # sh1.write(p,0,'<tr>')
    # sh1.write(p,1,'<td>')
    #
    # sh1.write(p,3,'</td>')
    # sh1.write(p,4,'<td>')
    #
    # sh1.write(p, 6, '</td>')
    # sh1.write(p, 7, '<td>')
    #
    # sh1.write(p, 9, '</td>')
    # sh1.write(p, 10, '<td>')
    #
    # sh1.write(p, 12, '</td>')
    # sh1.write(p, 13, '<td>')
    #
    # sh1.write(p, 15, '</td>')
    # sh1.write(p, 16, '<td>')
    #
    # sh1.write(p, 18, '</td>')
    # sh1.write(p, 19, '<td>')
    #
    # sh1.write(p, 21, '</td>')
    # sh1.write(p, 22, '<td>')
    #
    # sh1.write(p, 24, '</td>')
    # sh1.write(p, 25, '<td>')
    #
    # sh1.write(p, 27, '</td>')
    # sh1.write(p, 28, '<td>')
    #
    # sh1.write(p, 30, '</td>')
    # sh1.write(p, 31, '</tr>')

# # 选取第二个表单
# sh1 = wb.get_sheet(1)
#
# # 替换总成绩数据
# sh1.write(1, 0, 246.5)

# 保存
wb.save('7027.xls')
