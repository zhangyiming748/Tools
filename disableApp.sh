#!/bin/bash
# adb shell pm disable-user

# 停用软件更新
adb shell pm disable-user com.huawei.android.hwouc
# 停用华为视频
adb shell pm disable-user com.huawei.himovie
# 停用华为音乐
adb shell pm disable-user com.android.mediacenter
# 停用华为浏览器
adb shell pm disable-user com.huawei.browser
