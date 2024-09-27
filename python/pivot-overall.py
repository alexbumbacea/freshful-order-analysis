import pandas as pd, numpy as np
from io import StringIO
import os
import sys

df = pd.read_csv(sys.argv[1], sep=",")

print(df)
# restructure dataframe via pivot_table
res = df.pivot_table(index='data.month', columns='taxon 1', values='item.UnitPrice', margins=True, fill_value=0, aggfunc="sum")

print(res)
res.to_csv(sys.argv[2])