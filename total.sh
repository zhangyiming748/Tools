#!/opt/homebrew/bin/bash
#!/bin/bash
# 统计代码行数
# shellcheck disable=SC2046
wc -l `find .  ! -path "./vendor/*"  -name '*.go'`