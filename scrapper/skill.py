from bs4 import BeautifulSoup
import requests
from requirement import *

base_url = "https://oldschool.runescape.wiki"
page_path = "/w/Skills"

html = requests.get(base_url + page_path).text
doc = BeautifulSoup(html, "html.parser")
skills_table = doc.find_all("table", {"class": "wikitable sortable align-center-1"})
is_member = 0
skills_page_path = []
insert_skills_values = ["INSERT INTO skills (name, description, is_member) VALUES"]
for tbody in skills_table:
    rows = tbody.find_all("tr")[1:]
    for row in rows:
        cols = row.find_all("td")
        skills_page_path.append(cols[1].find("a")["href"])
        insert_skills_values.append(
            f"('{cols[1].text}', '{cols[2].text.replace("'", "''").replace("\n", "")}', {is_member}),"),
    is_member = 1

l = len(insert_skills_values)
insert_skills_values[l - 1] = insert_skills_values[l - 1][:-1] + ";"
with open('../database/query/skills/insert.sql', 'w') as f:
    for line in insert_skills_values:
        f.write(f"{line}\n")

scrap(base_url, skills_page_path)
