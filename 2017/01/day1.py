#!/usr/bin/env python
l=raw_input()
n=len(l)
print sum([int(l[i]) for i in range(n) if l[i]==l[(i+n/2)%n]])
