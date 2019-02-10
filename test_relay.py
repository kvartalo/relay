#!/usr/bin/env python3
"""Test for kvartalo/relay
"""

import json
import requests
#  import datetime
import provoj

URL = "http://127.0.0.1:3000"
ADDR = "0xbc88fcc53af747D30F1e17729659001d1129ddd7"

T = provoj.NewTest("Relay")

R = requests.get(URL + "/balance/"+ADDR)
T.rStatus("get balanceOf", R)
jsonR = R.json()
print(jsonR)

R = requests.get(URL + "/tx/nonce/"+ADDR)
T.rStatus("get nonceOf", R)

R = requests.post(URL + "/tx", json={'a': 'b'})
T.rStatus("post tx", R)

T.printScores()
