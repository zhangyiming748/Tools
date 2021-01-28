#!/bin/bash
# 在终端中执行,使得 go get 命令可用
export https_proxy=http://127.0.0.1:8889
export http_proxy=http://127.0.0.1:8889
git config --global http.proxy 'http://127.0.0.1:8889'
git config --global https.proxy 'http://127.0.0.1:8889'