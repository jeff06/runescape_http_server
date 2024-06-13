from bs4 import BeautifulSoup
import requests

url = "https://oldschool.runescape.wiki/w/Skills"

html = requests.get(url)
doc = BeautifulSoup(html.text, "html.parser")
skills_table = doc.find_all("table", {"class": "wikitable sortable align-center-1"})
is_member = 0
insert_values = ["INSERT INTO skills (name, description, is_member) VALUES"]
for tbody in skills_table:
    rows = tbody.find_all("tr")[1:]
    for row in rows:
        cols = row.find_all("td")
        insert_values.append(f"('{cols[1].text}', '{cols[2].text.replace("'", "''").replace("\n", "")}', {is_member}),"),
    is_member = 1
l = len(insert_values)
insert_values[l - 1] = insert_values[l - 1][:-1] + ";"
with open('../database/query/skills/insert.sql', 'w') as f:
    for line in insert_values:
        f.write(f"{line}\n")
