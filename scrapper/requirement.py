from bs4 import BeautifulSoup
import requests
from util import *


class ScrapableContent:
    def __init__(self, path: str, table_tag: {}):
        self.path = path
        self.table_tag = table_tag


def scrap(base_url):
    scrapable = []
    scrapable.append(ScrapableContent("/w/Attack", {"wikitable align-center-1 align-center-2": 1}))
    scrapable.append(ScrapableContent("/w/Strength", {"wikitable align-center-1": 1}))
    scrapable.append(ScrapableContent("/w/Prayer", {"wikitable align-center-1 align-center-2 align-center-6": 1}))
    scrapable.append(
        ScrapableContent("/w/Runecraft", {"wikitable sortable align-center-1 align-center-2 align-right-4": 2}))
    scrapable.append(ScrapableContent("/w/Fishing", {"wikitable sortable": 1}))

    insert_requirement_values = ["INSERT INTO unlocks (name, level, id_skill, other_requirement)  VALUES"]
    index = 1
    for content in scrapable:
        print(content.path)
        html = requests.get(base_url + content.path).text
        doc = BeautifulSoup(html, "html.parser")
        for key, value in content.table_tag.items():
            requirement_table = doc.find("table", {"class": key})
            rows = requirement_table.find_all("tr")[value:]
            for row in rows:
                cols = row.find_all("td")
                cured_cols = []
                for col in cols:
                    exist = col.has_attr("rowspan")
                    if (exist and col["rowspan"] == '1') or not exist:
                        cured_cols.append(col)
                insert_requirement_values.append(get_requirement_string_format(content.path, cured_cols, index))
        index += 1
    complete_save(insert_requirement_values, "unlocks")


def get_requirement_string_format(url, cols: [], index: int):
    match url:
        case "/w/Attack":
            return requirement_builder(cols[2].text, cols[0].text, index, cols[3].text)
        case "/w/Strength":
            return requirement_builder(cols[1].text, cols[2].text, index, cols[3].text)
        case "/w/Prayer":
            return requirement_builder(cols[2].text, cols[0].text, index, cols[3].text)
        case "/w/Runecraft":
            return requirement_builder(cols[2].text, cols[0].text, index, cols[4].text)
        case "/w/Fishing":
            cols_extracted = fishing_extractor(cols)
            return requirement_builder(cols_extracted[0], cols_extracted[1], index, cols_extracted[2])


def fishing_extractor(cols: []) -> []:
    if len(cols) == 3 or len(cols) == 4:
        return [cols[1].text, cols[2].text, "None"]
    elif len(cols) == 5:
        if cols[0].has_attr("rowspan") and not cols[4].has_attr("rowspan"):
            return [cols[1].text, cols[2].text, cols[4].text]
        else:
            return [cols[2].text, cols[3].text, "None"]
    elif len(cols) == 6:
        if cols[0].has_attr("style"):
            return [cols[3].text, cols[4].text, "None"]
        else:
            return [cols[2].text, cols[3].text, cols[5].text]
    elif len(cols) == 7:
        return [cols[4].text, cols[5].text, "None"]
    elif len(cols) == 8:
        return [cols[4].text, cols[5].text, cols[7].text]
    elif len(cols) == 9:
        return [cols[5].text, cols[6].text, cols[8].text]
    elif len(cols) == 10:
        return [cols[6].text, cols[7].text, cols[9].text]
    else:
        raise Exception(f"fishing_extractor - lenght - {len(cols)}")


def requirement_builder(name_unformatted, level, id_skill, other_requirement):
    name = fix_string_format(name_unformatted)
    print(name)
    return f"('{name}', {fix_int_format(level)}, {id_skill}, '{fix_string_format(other_requirement)}'),"


scrap("https://oldschool.runescape.wiki")
