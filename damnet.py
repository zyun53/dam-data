import requests
import time
from bs4 import BeautifulSoup
import re
import csv


def get_kana_list():
    """
    五十音別全項目表インデックスからURLを取得します。
    """
    target_url = 'http://damnet.or.jp/cgi-bin/binranA/ZenItiIndex.cgi?zi=zen&sy=go'
    r = requests.get(target_url)
    r.encoding = r.apparent_encoding

    soup = BeautifulSoup(r.text, 'html.parser')

    elems = soup.find_all(href=re.compile("TableAllItiran.cgi.*kana"))

    urls = []
    for e in elems:
       urls.append(e.get("href"))

    return urls

def get_dam_from_kana(url):
    """
    仮名一覧からダムのデータを取得します。
    """
    r = requests.get(url)
    r.encoding = r.apparent_encoding

    soup = BeautifulSoup(r.text, 'html.parser')

    elems = soup.tr.tr.find_all('tr')

    dams = []
    for e in elems:
        text = e.find_all('b')

        dam_name = text[0].getText().strip()
        dam_address = text[1].getText().strip()
        dam_url = e.find('a').get('href')

        dams.append({'name': dam_name, 'address': dam_address, 'url': dam_url})
    return dams

def main():
    """
    ダムの一覧をcsvに出力します。
    """
    kana = get_kana_list()

    dams = []

    for i in kana:
        time.sleep(1)
        dams += get_dam_from_kana(i)

    ans = sorted(dams, key=lambda x:x['name'])

    with open('damnet.csv', 'w') as f:
        field_names = ['name', 'address', 'url']
        writer = csv.DictWriter(f, fieldnames=field_names)
        writer.writerows(ans)

main()
