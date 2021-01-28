# coding=UTF-8

import requests
from bs4 import BeautifulSoup
import codecs
import time
# from from_python_to_mysql.connect import mysqlto

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

def parse_html(html):
    soup = BeautifulSoup(html, features="html.parser")
    imglist=soup.find_all('img')
    for i in imglist:
        print(i)
    divlist = soup.find('div', attrs={'id': 'waterfall'})
    for i in divlist:
        print(i)
        src=i.get('src')
        print(src)
    movie_name_list = []
    title_list = []
    movie_id_list = []
    for movie_li in divlist.find_all('li'):
        detail = movie_li.find('div', attrs={'class': 'hd'})
        movie_url = detail.find('a')['href']
        movie_id = movie_url[-8:-2]
        time.sleep(1)  # 延时函数，再被封就不止一个小时了
        # movie_title, movie_director, movie_date, movie_time, movie_introduction = parse_detail_html(movie_url)
        time.sleep(1)  # 延时函数，再被封就不止一个小时了
        # if movie_id == '129512':
        # mysqlto(movie_id, movie_title, movie_director, movie_date, movie_time, movie_introduction)
        # connect.mysqlto(movie_id, movie_title, 'dsddsds', 'dsddsds', 'dsddsds', 'dsddsds')
        time.sleep(1)  # 延时函数，再被封就不止一个小时了
        title = detail.find('span', attrs={'class': 'title'}).getText()
        print('completely')
        movie_name_list.append(title)
        title_list.append(movie_title)
        movie_id_list.append(movie_id)

        time.sleep(1)

    next_page = soup.find('span', attrs={'class': 'next'}).find('a')
    if next_page:
        return movie_name_list, DOWNLOAD_URL + next_page['href']
    return movie_name_list, None


def parse_detail_html(movie_url):
    html = download_page(movie_url)
    soup = BeautifulSoup(html, features="html.parser")
    movie_title = soup.find('span', attrs={'property': 'v:itemreviewed'}).get_text()
    movie_title = movie_title.replace("'", "\\'")#辛德勒的名单
    time.sleep(1)#延时函数，再被封就不止一个小时了
    movie_director = soup.find('a', attrs={"rel": "v:directedBy"}).get_text()
    time.sleep(1)#延时函数，再被封就不止一个小时了
    movie_introduction = soup.find('span', attrs={"property": "v:summary"}).get_text()
    time.sleep(1)#延时函数，再被封就不止一个小时了
    movie_date = soup.find('span', attrs={"property": "v:initialReleaseDate"}).get_text()
    # 延时函数，再被封就不止一个小时了
    movie_time = soup.find('span', attrs={"property": "v:runtime"}).get_text()
    time.sleep(1)#延时函数，再被封就不止一个小时了

    # connect.mysqlto(movie_id=movie_id, movie_title=movie_title, movie_director=movie_director, movie_date=movie_date,movie_time=movie_time, movie_introduction=movie_introduction)
    return movie_title, movie_director, movie_date, movie_time, movie_introduction


if __name__ == '__main__':
    url = DOWNLOAD_URL

    with codecs.open('movies.txt', 'wb', encoding='utf-8') as fp:
        while url:
            html = download_page(url)
            movies, url = parse_html(html)
            fp.write('\n'.join(movies))
            # fp.write('\n'.join(m_id))
            time.sleep(1)