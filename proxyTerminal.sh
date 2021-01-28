#!/bin/bash
# 设置当前的终端走代理
export https_proxy=http://127.0.0.1:8889
export http_proxy=http://127.0.0.1:8889
# export all_proxy=socks5://127.0.0.1:1089
curl cip.cc
