#!/bin/env python

import json
import os
import sys
import subprocess


CMD_NAME = sys.argv[1]
CHAT_ID = sys.argv[4]
USER_ID = sys.argv[5]
MSG_ID = sys.argv[6]

API = os.getenv('TG_API')
CFG = os.getenv('TG_CONFIG_FILE')

payload = {
    "type": "plain",
    "user_id": USER_ID,
    "chat_id": CHAT_ID,
    "message_id": MSG_ID,
    "text": "hello",
}

proc = [
    API,
    CFG,
    CMD_NAME,
    'send_text',
    json.dumps(payload),
]

try:
    print("args: " + str(sys.argv))
    ret = subprocess.run(proc, stdout=subprocess.PIPE, text=True)
    print(ret.stdout)
except Exception as e:
    print(f'{CMD_NAME}: error: {e}')
    exit(1)

exit(0)
