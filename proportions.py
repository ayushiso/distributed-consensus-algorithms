# plots proportions of each decision value (actual decision values not plotted) before consensus is reached

import matplotlib.pyplot as plt 
import numpy as np
import sys

infile = sys.argv[1]

with open(infile, 'r') as in_file:
    header = False
    i = 1
    for line in in_file:
        print("reading line", i)
        if header == True:
            vals = line.rstrip().split('\t')
            row = np.array(vals)

            # checks if consensus has been reached/all processors have dropped
            if len(np.unique(row)) > 1: 
                matrix = np.vstack((matrix, vals))
            else:
                print("skipped!")
                break
        else:
            # header line contains initial decision value of each processor
            vals = line.rstrip().split('\t')
            matrix = np.array(vals)
            header = True
        i+=1
            
matrix = matrix.astype(np.float)

# get unique values from the first round (second row)
vals = np.unique(matrix[0])
zero = np.array([0])
vals = np.setdiff1d(vals, zero)
vals = vals.astype(int)
#print("Unique values=", vals)

counts = vals

# loop through and count processors having a particular decision value
for row in matrix:
    rowcounts = np.zeros(len(vals), dtype=int)
    for i in range(len(vals)):
        count = np.count_nonzero(row==vals[i])
        rowcounts[i] = count

    counts = np.vstack((counts, rowcounts))

# remove actual decision values (comment out for debugging)
counts = counts[1:]
#print(counts)

row_sums = counts.sum(axis=1)
new_matrix = counts / row_sums[:, np.newaxis]

title = infile.split('-')
n = title[1]
t = title[-1].split('.')[0]

cat_1 = "total"
cat_2 = "dropped"

#print(counts)
plt.plot(new_matrix)
plt.xlabel("Rounds")
plt.ylabel("Counts")
plt.title("{} = {}, {} = {}".format(cat_1, cat_2, n ,t))
plt.savefig(infile[:-3] + "png")