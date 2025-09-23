"""
fetching and parsing country info taken from https://github.com/samayo/country-json/
"""
from git import Repo # pip install gitpython
import json
import os
from pathlib import Path
import shutil
from sys import platform
from subprocess import check_output
import stat
import base64

repo_download_path = "country-json"
update = False

data_path = "data"
flags_path = "flags"
country_data_filenames = ["country-by-name.json", "country-by-capital-city.json", "country-by-abbreviation.json", "country-by-continent.json", "country-by-geo-coordinates.json", "country-by-flag.json"]

countries = []
capitals = [] #245
abbrvs = [] #245
continents = [] #244
coords = [] #244
flags = [] #245

def main():
    global countries
    global capitals
    global abbrvs
    global continents
    global coords
    global flags

    if update:
        get_country_data()

    fnum = 0

    with open(os.path.join(data_path, country_data_filenames[fnum]), encoding='utf-8') as f:
        countries = json.load(f)
    fnum += 1

    # get country and capital
    with open(os.path.join(data_path, country_data_filenames[fnum]), encoding='utf-8') as f:
        capitals = json.load(f)
    fnum += 1

    # get abbreviations
    with open(os.path.join(data_path, country_data_filenames[fnum]), encoding='utf-8') as f:
        abbrvs = json.load(f)
    fnum += 1

    # get continents
    with open(os.path.join(data_path, country_data_filenames[fnum]), encoding='utf-8') as f:
        continents = json.load(f)
    fnum += 1

    # get geographical coordinates
    with open(os.path.join(data_path, country_data_filenames[fnum]), encoding='utf-8') as f:
        coords = json.load(f)
    fnum += 1

    # get flags
    with open(os.path.join(data_path, country_data_filenames[fnum]), encoding='utf-8') as f:
        flags = json.load(f)
            
    # rename 'country' key to 'name'
    for i in range(len(countries)):
        countries[i].update({"name": countries[i]["country"]})
        countries[i].pop("country", None)

    clean_country_data()

    # add capitals
    for i in range(len(countries)):
        if countries[i]["name"] == capitals[i]["country"]: # double check if countries match
            #print(countries[i])
            countries[i].update({"capital": capitals[i]["city"]})
            #print(countries[i]) 
        else:
            print(f"capital country mismatch at index {i}: {countries[i]["name"]} and {capitals[i]["country"]}")
            break

    # add abbrvs
    for i in range(len(countries)):
        if countries[i]["name"] == abbrvs[i]["country"]: # double check if countries match
            #print(countries[i])
            countries[i].update({"abbreviation": abbrvs[i]["abbreviation"]})
            #print(countries[i]) 
        else:
            print(f"abbrv country mismatch at index {i}: {countries[i]["name"]} and {abbrvs[i]["country"]}")
            break
    
    # add continents
    for i in range(len(countries)):
        if countries[i]["name"] == continents[i]["country"]: # double check if countries match
            #print(countries[i])
            countries[i].update({"continent": continents[i]["continent"]})
            #print(countries[i]) 
        else:
            print(f"continent country mismatch at index {i}: {countries[i]["name"]} and {continents[i]["country"]}")
            break

    # add flags
    for i in range(len(countries)):
        if countries[i]["name"] == flags[i]["country"]: # double check if countries match
            #print(countries[i])
            countries[i].update({"flag": flags[i]["flag_base64"].split(",")[1] if not flags[i]["flag_base64"] is None else flags[i]["flag_base64"]})
            #print(countries[i]) 
        else:
            print(f"flags country mismatch at index {i}: {countries[i]["name"]} and {flags[i]["country"]}")
            break

    #print(json.dumps(countries, indent=3))
    countries_filename = "countries.json"
    with open(countries_filename, 'w', encoding='utf-8') as f:
        json.dump(countries, f, ensure_ascii=False, indent=4)

    # copy json to api
    api_data_path = "../api/data"
    if not Path(api_data_path).is_dir():
        os.mkdir(api_data_path)

    shutil.copyfile(f"./{countries_filename}", f"{api_data_path}/{countries_filename}")


def clean_country_data():
    global countries
    global capitals
    global abbrvs
    global continents
    global flags

    # Ivory Coast wrong place
    ivory_coast = countries.pop(52)
    countries.insert(109, ivory_coast)

    # England wrong place
    england = countries.pop(63)
    countries.insert(64, england)

    # DRC wrong place and add 2nd "the"
    drc = countries.pop(49)
    drc["name"] = "The Democratic Republic of the Congo"
    countries.insert(218, drc)
    capitals[215]["country"] = "The Democratic Republic of the Congo"
    abbrvs[216]["country"] = "The Democratic Republic of the Congo"
    continents[215]["country"] = "The Democratic Republic of the Congo"
    flags[216]["country"] = "The Democratic Republic of the Congo"


    # missing Guernsey capital and continent
    guernsey = {"country": "Guernsey", "city": "Saint Peter Port"}
    capitals.insert(89, guernsey)
    guernsey2 = {"country": "Guernsey", "continent": "Europe"}
    continents.insert(89, guernsey2)

    # missing Isle of Man capital and continent
    iom = {"country": "Isle of Man", "city": "Douglas"}
    capitals.insert(106, iom)
    iom_2 = {"country": "Isle of Man", "continent": "Europe"}
    continents.insert(106, iom_2)

    # missing Jersey capital, continent, flag
    jersey = {"country": "Jersey", "city": "St Helier"}
    capitals.insert(111, jersey)
    jersey_2 = {"country": "Jersey", "continent": "Europe"}
    continents.insert(111, jersey_2)
    with open(os.path.join(flags_path, "jersey.txt"), encoding='utf-8') as f:
        base64_str = f.read()
        jersey_3 = {"country": "Jersey", "flag_base64": base64_str}
        flags.insert(111, jersey_3)

    # missing Montenegro flag
    with open(os.path.join(flags_path, "montenegro.txt"), encoding='utf-8') as f:
        base64_str = f.read()
        montenegro = {"country": "Montenegro", "flag_base64": base64_str}
        flags.insert(145, montenegro)

    # Montserrat wrong place
    montserrat = countries.pop(145)
    countries.insert(146, montserrat)

    # missing Timor-Leste capital, continent, flag
    t_l = {"country": "Timor-Leste", "city": "Dili"}
    capitals.insert(219, t_l)
    t_l_2 = {"country": "Timor-Leste", "continent": "Asia"}
    continents.insert(219, t_l_2)
    with open(os.path.join(flags_path, "timor-leste.txt"), encoding='utf-8') as f:
        base64_str = f.read()
        montenegro = {"country": "Timor-Leste", "flag_base64": base64_str}
        flags.insert(219, montenegro)

    #remove Holy See + readdas Vatican City
    vc = countries.pop(95)
    vc_cap = capitals.pop(95)
    vc_abbr = abbrvs.pop(94)
    vc_cont = continents.pop(95)
    vc_flag = flags.pop(95)
    
    #add and rename Vatican City
    vc["name"] = "Vatican City"
    countries.insert(237, vc)
    capitals.pop(238)

    vc_cap["country"] = "Vatican City"
    capitals.insert(237, vc_cap)

    vc_abbr["country"] = "Vatican City"
    abbrvs.insert(235, vc_abbr)

    vc_cont["country"] = "Vatican City"
    continents.insert(237, vc_cont)
    
    vc_flag["country"] = "Vatican City"
    flags.insert(237, vc_flag)

    # Israel wrong order
    is_1 = countries.pop(104)
    countries.insert(105, is_1)
    is_2 = capitals.pop(104)
    capitals.insert(105, is_2)
    is_3 = continents.pop(104)
    continents.insert(105, is_3)

    # missing England, Scotland, and Wales abbrv
    en = {"country": "England", "abbreviation": "GB"}
    abbrvs.insert(63, en)
    sc = {"country": "Scotland", "abbreviation": "GB"}
    abbrvs.insert(192, sc)
    w = {"country": "Wales", "abbreviation": "GB"}
    abbrvs.insert(242, w)

    # missing British Indian Ocean Territory capital
    capitals[30]["city"] = "Diego Garcia"

    # missing South Georgia and the South Sandwich Islands capital
    capitals[203]["city"] = "King Edward Point"


# get country data from repo and delete unneeded files
def get_country_data():

    if not Path(f"./{repo_download_path}").is_dir():
        Repo.clone_from("https://github.com/samayo/country-json/", f"./{repo_download_path}")

    #keep req files
    if not Path(data_path).is_dir():
        os.mkdir(data_path)

    for filename in country_data_filenames:
        shutil.copyfile(f"./{repo_download_path}/src/{filename}", f"./{data_path}/{filename}")

    # delete repo
    if Path(f"./{repo_download_path}").is_dir():
        for root, dirs, files in os.walk(repo_download_path):  
            for dir in dirs:
                os.chmod(os.path.join(root, dir), stat.S_IRWXU)
            for file in files:
                os.chmod(os.path.join(root, file), stat.S_IRWXU)
        shutil.rmtree(repo_download_path)

if __name__ == '__main__':
    main()
