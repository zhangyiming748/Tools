# coding=UTF-8

import requests
from bs4 import BeautifulSoup
import codecs
import time
import os

DOWNLOAD_URL = 'https://images.cveoy.com/'


def download_page(url):
    headers = {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36'
    }
    proxy = {
        'http': '127.0.0.1:1080',
        'https': '127.0.0.1:1080'
    }
    return requests.get(url, headers=headers).content


#

def parse_html(initPage, html):
    soup = BeautifulSoup(html, features="html.parser")
    imglist = soup.find_all('img')
    if initPage == 1:
        pass

    nextPage(1,soup)
    time.sleep(300)
    for i in imglist:
        print(i)
        src = i.get('src')
        # print(src)
        surffix = src.split('\\', 3)
        second = surffix[1] + '/' + surffix[2]
        first = 'https://images.cveoy.com/images/'
        full = first + second
        print('full url is %s' % full)
        # for j in surffix:
        #     print('后缀%s'%j)
        recodeDownlink(full)
        time.sleep(10)


def nextPage(current, soup):
    rightNum = current + 1
    btnBar = soup.find('div', attrs={'class': 'text-center hidden-xs'})
    btns = btnBar.find('ul', attrs={'class': 'pagination pagination-lg'})
    hrefs = btns.find_all('a')
    for i in hrefs:
        Num = i.get_text()
        if int(Num) == rightNum:
            next = i['href']
            perfix='https://images.cveoy.com/'
            perfix='https://images.cveoy.com/'
            nextLink=perfix+next
            print('完整的下一页链接%s'%nextLink)


        print('按钮超链接%s' % i)
        no = i['href']
        print('page : %s' % no)
    # .get_text()


def recodeDownlink(link):
    with open('./link.txt', 'a+', encoding='utf-8')as f:
        f.write(link)
        f.write('\n')


if __name__ == '__main__':
    url = DOWNLOAD_URL
    html = download_page(url)
    parse_html(1, html)
