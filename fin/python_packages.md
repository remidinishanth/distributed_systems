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

### Bullish vs Bearish

* Bull attacks with horns from bottom to up, hence market going to upward
* Bear attacks from top to bottom with hands, hence market is going in downward direction
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/d7fcecf8-36b1-48f6-a95c-b96972141004)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/72545224-e4ea-4141-9a5a-4091ae6b5c6c)


### Call and Put Option

<img width="616" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/cd928e2d-5440-4b4d-a016-74efca7c1bca">


<img width="922" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/ea9eef67-2aed-41c6-90a9-db96852f4135">

* The intrinsic value of a call option = Spot price - Strike price
* The intrinsic value of a put option = Strike price - Spot price

#### In the money:

If the intrinsic value of an options contract is a positive number, then the option is considered to be ‘in the money’.

For instance, assume that the spot price of a stock is currently at Rs. 1,600 and the strike price of a call option of that said stock is Rs. 1,500. The intrinsic value of the said call option would be Rs. 100 (Rs. 1,600 - Rs. 1,500). Since the intrinsic value is positive, the call option of the stock is considered to be ‘in the money’.

<img width="687" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/994d7b40-21c2-4095-a116-72078a549057">

#### Pay off

Long Call option - Unlimited Profit, Buying Call option
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/c341527d-d33e-4456-a494-ce049511e273)

Short call option - Limited Profit, Selling call option
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/6968dbc3-61d2-43c2-9bdf-1b2207e5c706)


* Buying a Call option

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/0fec9c82-6746-4916-98f0-cbf482d03a3c)

* Buying a Put option

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/c453e165-5a2f-4ea1-9a08-4a130b368602)

#### Call option value (vs) Stock price

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/67cc129c-1aca-46a5-b795-c8a4ef752713)


### Spreads

When you do this ensure 
* All strikes belong to the same underlying
* Belong to the same expiry series
* Each leg involves the same number of options

#### Bull call spread

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/6d5125d7-fa70-4ad0-ae59-baac0df85adc)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/14e19f14-13d6-410d-9d4a-eec0c542eaec)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/2a10121c-3724-4f16-a9d2-16a72e2090eb)

KEY TAKEAWAYS
* A bull call spread is an options strategy used when a trader is betting that a stock will have a limited increase in its price. 
* The strategy uses two call options to create a range consisting of a lower strike price and an upper strike price.
* The bullish call spread can limit the losses of owning stock, but it also caps the gains.

Ref: https://zerodha.com/varsity/chapter/bull-call-spread/


For Backtesting etc
* https://github.com/jesse-ai/jesse
  - Refer `def _execute(self)` in `Strategy`
  - ![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/c66568d5-1b94-423b-b4aa-3e7937b734a7)

* https://www.backtrader.com/ and https://github.com/kernc/backtesting.py
