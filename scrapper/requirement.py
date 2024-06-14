from bs4 import BeautifulSoup
import requests
from util import *


class Requirement:
    def __init__(self, url, table_tag, parse, cols, index):
        self.url = url
        self.table_tag = table_tag
        self.parse = parse
        self.cols = cols
        self.index = index


def scrap(base_url):
    dict = {
        "/w/Attack": "wikitable align-center-1 align-center-2",
        "/w/Strength": "wikitable align-center-1",
        "/w/Prayer": "wikitable align-center-1 align-center-2 align-center-6",
        #"/w/Runecraft": "wikitable sortable align-center-1 align-center-2 align-right-4 jquery-tablesorter"
    }
    insert_requirement_values = ["INSERT INTO unlocks (name, level, id_skill, other_requirement)  VALUES"]
    index = 1
    for url in dict:
        print(url)
        html = requests.get(base_url + url).text
        doc = BeautifulSoup(html, "html.parser")
        requirement_table = doc.find("table", {"class": dict[url]})
        rows = requirement_table.find_all("tr")[1:]
        for row in rows:
            cols = row.find_all("td")
            insert_requirement_values.append(get_requirement_string_format(url, cols, index))
        index += 1
    complete_save(insert_requirement_values, "unlocks")


def get_requirement_string_format(url, cols: [], index: int):
    match url:
        case "/w/Attack":
            return attack_requirement_builder(cols, index)
        case "/w/Strength":
            return strength_requirement_builder(cols, index)
        case "/w/Prayer":
            return prayer_requirement_builder(cols, index)
        case "/w/Runecraft":
            return runecraft_requirement_builder(cols, index)


def attack_requirement_builder(cols: [], index: int):
    name = fix_string_format(cols[2].text)
    print(name)
    return f"('{name}', {fix_int_format(cols[0].text)}, {index}, '{fix_string_format(cols[3].text)}'),"


def strength_requirement_builder(cols: [], index: int):
    name = fix_string_format(cols[1].text)
    print(name)
    return f"('{name}', {fix_int_format(cols[2].text)}, {index}, '{cols[4].text}'),"


def prayer_requirement_builder(cols: [], index: int):
    name = fix_string_format(cols[2].text)
    print(name)
    return f"('{name}', {fix_int_format(cols[0].text)}, {index}, '{fix_string_format(cols[3].text)}'),"


def runecraft_requirement_builder(cols: [], index: int):
    name = fix_string_format(cols[2].text)
    print(name)
    return f"('{name}', {fix_int_format(cols[0].text)}, {index}, '{fix_string_format(cols[4].text)}'),"
