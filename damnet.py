import requests

url = 'http://damnet.or.jp/cgi-bin/binranA/ZenItiIndex.cgi?zi=iti&sy=go'

r = requests.get(url)

print(r)
