from bs4 import BeautifulSoup
import requests


def scrap(base_url, urls):
    for url in urls:
        html = requests.get(base_url + url).text
        doc = BeautifulSoup(html, "html.parser")
        requirement_table = doc.find("table", {"class": "wikitable align-center-1 align-center-2"})
        rows = requirement_table.find_all("tr")
        for row in rows:
            cols = row.find_all("td")
            
