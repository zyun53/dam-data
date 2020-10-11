import requests
from bs4 import BeautifulSoup
import re


def top():
    target_url = 'http://damnet.or.jp/cgi-bin/binranA/ZenItiIndex.cgi?zi=zen&sy=go'
    r = requests.get(target_url)
    r.encoding = r.apparent_encoding

    soup = BeautifulSoup(r.text, 'html.parser')

    with open('originDataOld.html', mode='w', encoding = 'utf-8') as fw:
        fw.write(soup.prettify())

    elems = soup.find_all(href=re.compile("TableAllItiran.cgi.*kana"))
    for e in elems:
        get_all(e.get("href"))


def get_all(url):
    r = requests.get(url)
    r.encoding = r.apparent_encoding

    soup = BeautifulSoup(r.text, 'html.parser')

    elems = soup.tr.tr.find_all('tr')
    for e in elems:
        text = e.find_all('b')

        dam_name = text[0].getText().strip()
        dam_address = text[1].getText()
        dam_url = e.find('a').get('href')

        print(dam_name + "," + dam_address + "," + dam_url)

top()
