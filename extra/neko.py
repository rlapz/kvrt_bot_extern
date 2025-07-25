#!/bin/env python

import json
import os
import sys
import subprocess
import requests


CMD_NAME = sys.argv[1]
CHAT_ID = sys.argv[3]
USER_ID = sys.argv[4]
MSG_ID = sys.argv[5]

API = os.getenv('TG_API')
CFG = os.getenv('TG_CONFIG_FILE')

API_URL = 'https://api.nekosia.cat/api/v1/images'
FILTERS = [
    "random", "catgirl", "foxgirl", "wolf-girl", "animal-ears", "tail", "tail-with-ribbon",
    "tail-from-under-skirt", "cute", "cuteness-is-justice", "blue-archive", "girl", "young-girl",
    "maid", "maid-uniform", "vtuber", "w-sitting", "lying-down", "hands-forming-a-heart",
    "wink", "valentine", "headphones", "thigh-high-socks", "knee-high-socks", "white-tights",
    "black-tights", "heterochromia", "uniform", "sailor-uniform", "hoodie", "ribbon", "white-hair",
    "blue-hair", "long-hair", "blonde", "blue-eyes", "purple-eyes",
]


def fetch(filter):
    if len(filter) == 0:
        filter = 'random'
    elif filter not in FILTERS:
        raise 'invalid filter'

    new_url = f'{API_URL}/{filter}'
    resp = requests.get(new_url)
    if resp.status_code != 200:
        raise 'invalid response'

    return resp.json()


def tg_escape(txt) -> str:
    ret = ''
    special = "_*[]()~`>#+-|{}.!"
    for c in txt:
        for s in special:
            if c != s:
                continue
            ret += '\\'
        ret += c

    return ret


def build_text(filter) -> str:
    data = fetch(filter)
    image = data.get('image')
    anime = data.get('anime')
    attr = data.get('attribution')

    comp_url = image.get('compressed').get('url')
    orig_url = image.get('original').get('url')
    anime_char = anime.get('character')
    anime_name = anime.get('title')
    artist = attr.get('artist').get('username', '')
    artist_url = attr.get('artist').get('profile', '')
    source_url = data.get('source').get('url', '')
    category = data.get('category')

    source_url = tg_escape(source_url)
    artist = tg_escape(artist)
    return f'`URL     :`  [Compressed]({comp_url}) \\- [Original]({orig_url})\n' \
        f'`Name    : {anime_char}` from `{anime_name}`\n' \
        f'`Artist  :`  [{artist}]({artist_url})\n' \
        f'`Source  :`  {source_url}\n' \
        f'`Category: {category}`'


try:
    filter = ''

    payload = {
        'type': 'format',
        'user_id': USER_ID,
        'chat_id': CHAT_ID,
        'message_id': MSG_ID,
        'deletable': True,
        'text': build_text(filter),
    }

    proc = [
        API,
        CFG,
        CMD_NAME,
        'send_text',
        json.dumps(payload),
    ]

    subprocess.run(proc)
except Exception as e:
    print(f'{CMD_NAME}: error: {e}')
    exit(1)

exit(0)
