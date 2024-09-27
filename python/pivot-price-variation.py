import pandas as pd, numpy as np
from io import StringIO
import os
import sys

df = pd.read_csv(sys.argv[1], sep=",")

print(df)
# restructure dataframe via pivot_table
res = df.pivot_table(index='item.ProductName', columns='data.month', values='item.UnitPrice', margins=True, fill_value=None, aggfunc="max")

print(res)
res.to_csv(sys.argv[2])