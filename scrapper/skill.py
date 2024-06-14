from requirement import *
from util import *

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
        print(cols[1].text)
        insert_skills_values.append(
            f"('{cols[1].text}', '{fix_string_format(cols[2].text)}', {is_member}),"),
    is_member = 1

complete_save(insert_skills_values, "skills")

scrap(base_url)
