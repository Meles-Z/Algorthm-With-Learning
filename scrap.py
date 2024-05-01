from bs4 import BeautifulSoup
import requests


url = "https://www.tutorialspoint.com/index.htm"
req=requests.get(url)

with open("index.html") as fl:
    soup=BeautifulSoup(fl, 'html.parser')
print(soup)
