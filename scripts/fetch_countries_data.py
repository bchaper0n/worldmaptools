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

repo_download_path = "country-json"
update = False

data_path = "data"
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

    with open(os.path.join(data_path, country_data_filenames[fnum])) as f:
        countries = json.load(f)
    fnum += 1

    # get country and capital
    with open(os.path.join(data_path, country_data_filenames[fnum])) as f:
        capitals = json.load(f)
    fnum += 1

    # get abbreviations
    with open(os.path.join(data_path, country_data_filenames[fnum])) as f:
        abbrvs = json.load(f)
    fnum += 1

    # get continents
    with open(os.path.join(data_path, country_data_filenames[fnum])) as f:
        continents = json.load(f)
    fnum += 1

    # get geographical coordinates
    with open(os.path.join(data_path, country_data_filenames[fnum])) as f:
        coords = json.load(f)
    fnum += 1

    # get flags
    with open(os.path.join(data_path, country_data_filenames[fnum])) as f:
        flags = json.load(f)


    #print(f"{len(countries)} vs {len(abbrvs)}")

    clean_country_data()

    #print(f"{len(countries)} vs {len(abbrvs)}")

    # add capitals
    for i in range(len(countries)):
        if countries[i]["country"] == capitals[i]["country"]: # double check if countries match
            #print(countries[i])
            countries[i].update({"capital city": capitals[i]["city"]})
            #print(countries[i]) 
        else:
            print(f"country mismatch at index {i}: {countries[i]["country"]} and {capitals[i]["country"]}")
            break

    # add abbrvs
    for i in range(len(countries)):
        if countries[i]["country"] == "Holy See (Vatican City State)":
            print(i)

        if countries[i]["country"] == abbrvs[i]["country"]: # double check if countries match
            #print(countries[i])
            countries[i].update({"abbreviation": abbrvs[i]["abbreviation"]})
            #print(countries[i]) 
        else:
            print(f"country mismatch at index {i}: {countries[i]["country"]} and {abbrvs[i]["country"]}")
            break

    #print(json.dumps(countries, indent=3))

def clean_country_data():
    global countries
    global capitals
    global abbrvs

    # Ivory Coast wrong place
    ivory_coast = countries.pop(52)
    countries.insert(109, ivory_coast)

    # England wrong place
    england = countries.pop(63)
    countries.insert(64, england)

    # DRC wrong place and add 2nd "the"
    drc = countries.pop(49)
    drc["country"] = "The Democratic Republic of the Congo"
    countries.insert(218, drc)
    capitals[215]["country"] = "The Democratic Republic of the Congo"
    abbrvs[216]["country"] = "The Democratic Republic of the Congo"

    # missing Guernsey capital
    guernsey = {"country": "Guernsey", "city": "Saint Peter Port"}
    capitals.insert(89, guernsey)

    # missing Isle of Man capital
    iom = {"country": "Isle of Man", "city": "Douglas"}
    capitals.insert(106, iom)
    
    # missing Jersey capital
    jersey = {"country": "Jersey", "city": "St Helier"}
    capitals.insert(111, jersey)

    # Montserrat wrong place
    montserrat = countries.pop(145)
    countries.insert(146, montserrat)

    # missing Timor-Leste capital
    t_l = {"country": "Timor-Leste", "city": "Dili"}
    capitals.insert(219, t_l)

    #add and rename Vatican City
    vc = {"country": "Vatican City"}
    countries.insert(239, vc)
    capitals[239]["country"] = "Vatican City"
    capitals[239]["city"] = "Vatican City"
    vc_cap = {"country": "Vatican City", "abbreviation": "VA"}
    abbrvs.insert(237, vc_cap)

    # Israel wrong order
    is_1 = countries.pop(105)
    countries.insert(106, is_1)
    is_2 = capitals.pop(105)
    capitals.insert(106, is_2)

    # missing England, Scotland, and Wales abbrv
    en = {"country": "England", "abbreviation": "GB"}
    abbrvs.insert(63, en)
    sc = {"country": "Scotland", "abbreviation": "GB"}
    abbrvs.insert(193, sc)
    w = {"country": "Wales", "abbreviation": "GB"}
    abbrvs.insert(243, w)

    #remove Holy See
    countries.pop(95)
    capitals.pop(95)
    abbrvs.pop(95)

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
