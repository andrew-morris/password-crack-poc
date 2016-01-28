#!/usr/bin/env python

import hashlib
import sys

cipher = sys.argv[1]

print "[+] Cracking..."

for n in range(1, 99999999):
    if hashlib.md5(str(n)).hexdigest() == cipher:
        print "[!] The code has been cracked: %d" % n
        exit()
