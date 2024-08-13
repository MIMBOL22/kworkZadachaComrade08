rows = []
sorted_rows = {}
with open('../message.txt') as f:
    rows = f.readlines()
for i in rows:
    b = i.replace("\n","").split(" ")
    if len(b) == 2:
      sorted_rows[int(b[0])] = b[1]
max_key = max(list(sorted_rows.keys()))

allowed_keys = []
key = 0
i = 1
while key < max_key:
    key = (i+1)*i//2
    allowed_keys.append(key)
    i+=1
result1 = [];
for i in allowed_keys:
    result1.append(sorted_rows[i]);
print(" ".join(result1))
