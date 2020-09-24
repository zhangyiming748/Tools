#!/bin/bash
# adb shell pm disable-user

# 停用软件更新
adb shell pm clear com.huawei.android.hwouc
adb shell pm disable-user com.huawei.android.hwouc
# 停用华为视频
adb shell pm clear com.huawei.himovie
adb shell pm disable-user com.huawei.himovie
# 停用华为音乐
adb shell pm clear com.android.mediacenter
adb shell pm disable-user com.android.mediacenter
# 停用华为浏览器
adb shell pm clear com.huawei.browser
adb shell pm disable-user com.huawei.browser
# 华为阅读
adb shell pm clear com.huawei.hwireader
adb shell pm disable-user com.huawei.hwireader
# 华为主题
adb shell pm clear com.huawei.android.thememanager
adb shell pm disable-user com.huawei.android.thememanager
