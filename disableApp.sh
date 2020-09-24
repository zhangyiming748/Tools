#!/bin/bash
# adb shell pm disable-user

echo 清除软件更新缓存
adb shell pm clear com.huawei.android.hwouc
echo 停用软件更新
adb shell pm disable-user com.huawei.android.hwouc
echo 清除华为视频缓存
adb shell pm clear com.huawei.himovie
echo 停用华为视频
adb shell pm disable-user com.huawei.himovie
echo 清除华为音乐缓存
adb shell pm clear com.android.mediacenter
echo 停用华为音乐视频
adb shell pm disable-user com.android.mediacenter
echo 清除华为浏览器缓存
adb shell pm clear com.huawei.browser
echo 清除华为浏览器缓存
adb shell pm disable-user com.huawei.browser
echo 清除华为阅读缓存
adb shell pm clear com.huawei.hwireader
echo 清除华为阅读缓存
adb shell pm disable-user com.huawei.hwireader
echo 清除华为主题缓存
adb shell pm clear com.huawei.android.thememanager
echo 清除华为主题缓存
adb shell pm disable-user com.huawei.android.thememanager
