### Ref

* https://v2api.aliceblueonline.com/introduction
* https://github.com/TulipCharts/tulipindicators and https://pypi.org/project/newtulipy/ and https://tulipindicators.org/
* https://technical-analysis-library-in-python.readthedocs.io/en/latest/ta.html
* https://github.com/twopirllc/pandas-ta#quick-start
* https://pandas.pydata.org/Pandas_Cheat_Sheet.pdf


Converting minute candles to hourly candles
```python
import pandas as pd
import datetime

df = pd.read_csv('test_data.csv')

df['datetime'] = pd.to_datetime(df['datetime'])

# Clean any data which is not within market time
df = df[(df.datetime.dt.time > datetime.time(9, 15)) & (df.datetime.dt.time < datetime.time(15, 31))]

df['calender_date'] = df.datetime.dt.date
day_wise_data = df.groupby(df['calender_date'])


combined_data = []
for calender_date, day_data in day_wise_data:
    # print(calender_date, val)
    day_data.set_index('datetime', inplace=True)
    day_data = day_data.resample('60T', origin='start').agg({'open':'first', 'high': 'max', 'low':'min', 'close':'last'})
    day_data.reset_index(inplace=True)
    combined_data.append(day_data)

resampled_df = pd.concat(combined_data, ignore_index=True)

print(resampled_df)
```

Can also try `df.resample('1H', origin=origin_start)`
