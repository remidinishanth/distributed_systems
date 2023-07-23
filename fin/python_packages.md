### Ref

* https://v2api.aliceblueonline.com/introduction
* https://github.com/TulipCharts/tulipindicators and https://pypi.org/project/newtulipy/ and https://tulipindicators.org/
* https://technical-analysis-library-in-python.readthedocs.io/en/latest/ta.html
* https://github.com/twopirllc/pandas-ta#quick-start
* https://pandas.pydata.org/Pandas_Cheat_Sheet.pdf

* Awesome resource https://napkinfinance.com/napkins/napkins/

### Also
* https://tradingtick.com/
* https://in.tradingview.com/chart/YQrjA7DA/

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

### Candle Stick Patterns
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/72545224-e4ea-4141-9a5a-4091ae6b5c6c)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/096cd25d-f637-4d77-8b9d-a672464b76d3)

#### Dark cover cloud
The "dark cloud cover" is a candlestick pattern used in technical analysis to predict potential reversals in an upward trend

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/69f837e7-6f93-41a0-83f8-0c5acff2c939)

The bearish candlestick should penetrate at least halfway into the body of the previous bullish candlestick. This shows that the bears are gaining strength and pushing the price lower.

The dark cloud cover pattern suggests a potential reversal in the uptrend because it indicates that bears are starting to exert control over the market. It represents a shift in sentiment from bullishness to bearishness, and traders often interpret it as a signal to consider selling or taking profits.

#### Classic example of Reversal
<img width="383" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/a5370cc6-5f84-4159-b6ea-c68aed0159ff">


### Long vs Short

[long-and-short-trading-term-definitions-1031122_FINAL-1501219fbdac4b0e90bca691294125ae.webm](https://github.com/remidinishanth/distributed_systems/assets/19663316/c419f002-0d55-4165-9b69-b2b2bb746a01)

Thank of Long as buying something, Short as selling something
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/1ea97abe-34f1-4204-bfcd-d971efc1bfe1)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/7b33757b-d9fd-49ff-b4c7-0c8f130eaae4)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/5bdd4a33-fafd-469d-b71d-8de372efa45e)

Long Put is somewhat confusing because the price of the underlying asset actually is going down, hence you buy a put, but it is called Long Put

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/f713bf90-4af2-4fc5-b1fd-fbf5aa586df9)

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
* The bullish call spread can limit the losses of owning stock but also caps the gains.

Ref: https://zerodha.com/varsity/chapter/bull-call-spread/


For Backtesting etc
* https://github.com/jesse-ai/jesse

Refer check function in Strategy
<img width="845" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/b5dc409d-a87d-4833-8208-c3c429872bb2">

> `should_long()` and `should_short()` are for entering trades only. This means they would get called on every new candle only if no position is open, and no order is active.

> If you're looking to close trades dynamically, update_position() is what you're looking for.

> `go_short()` Same as `go_long()` but uses `self.sell` for entry instead of `self.buy`

<img width="1695" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/00498ccc-1f85-4210-af5d-e66b11eac0ea">


  - First you need to define your strategy, Refer `def _execute(self)` in `Strategy`

  - Then based on your route, it does trading
  - ![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/c66568d5-1b94-423b-b4aa-3e7937b734a7)

Not a great design, in go_long we are hard coding, when all to enter and exit to take profits
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/68d05fd8-dcc9-4544-8155-861477fa2c4b)

* `_execute_long` is just looking into `self.buy`, based on the price, it is looking at whether to place (`Stop`, `Limit` or `Market`) order.
* `_detect_and_handle_entry_and_exit_modifications` has logic to handle `self.buy`, `self.stop_loss`, and `self.take_profit` modifications in the program, it will cancel any non-executed orders and will place new ones as per the modifications.
  - Everything in that works using `if not np.array_equal(self.take_profit, self._take_profit)` whether things modified from earlier used values.
  - Broker is handling reducing position as we start taking profit

```python
self._take_profit_orders.append(
    self.broker.reduce_position_at(
        o[0],
        o[1],
        order_roles.CLOSE_POSITION
    )
)
```



* https://www.backtrader.com/ and https://github.com/kernc/backtesting.py
