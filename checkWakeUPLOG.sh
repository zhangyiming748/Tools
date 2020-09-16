!#/bin/bash
# 查看唤醒日志
pmset -g log|grep -e " Wake " -e "Wake Requests" -e "WakeTime"