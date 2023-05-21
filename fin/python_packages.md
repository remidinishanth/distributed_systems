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

### Concepts
* Price Action
* Support Resistance
* Trend lines
* Trading Uday Zone, see how the market is moving today and draw two lines and based on that, decide when to enter

<img width="922" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/ea9eef67-2aed-41c6-90a9-db96852f4135">

* The intrinsic value of a call option = Spot price - Strike price
* The intrinsic value of a put option = Strike price - Spot price

#### In the money:

If the intrinsic value of an options contract is a positive number, then the option is considered to be ‘in the money’.

For instance, assume that the spot price of a stock is currently at Rs. 1,600 and the strike price of a call option of that said stock is Rs. 1,500. The intrinsic value of the said call option would be Rs. 100 (Rs. 1,600 - Rs. 1,500). Since the intrinsic value is positive, the call option of the stock is considered to be ‘in the money’.
