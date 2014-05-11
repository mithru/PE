n = 1000
count = 0
for i in range(0,n):
    if (i%3)==0 or (i%5)==0:
        count = count + i

print (count)
