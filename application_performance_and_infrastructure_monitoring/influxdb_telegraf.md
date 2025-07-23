<img width="1625" height="923" alt="image" src="https://github.com/user-attachments/assets/801f3b35-d56e-41ae-ab75-ca51bfebe692" />


TICK stack: Telegraf, InfluxDB, Chronograf, and Kapacitor.

The TICK stack stands for
* Telegraph: a server agent for collecting and reporting metrics
* InfluxDB:  a high-performance time series database which uses SQL like syntax to query the data.
* Chronograph: a user interface to quickly see the data  stored in InfluxDB to build visualization dashboards and alerts
* Kapacitor: a data-processing engine that can process, stream and batch data from InfluxDB


<img width="1305" height="913" alt="image" src="https://github.com/user-attachments/assets/cab75c53-638c-406f-a4e7-6b0a7c3ab886" />


TIG stack: Telegraf, InfluxDB, Grafana

* Telegraf: Collecting data
* Influxdb: Saving data
* Grafana: Displaying data

<img width="2124" height="1400" alt="image" src="https://github.com/user-attachments/assets/a4fe2c28-8979-44ab-80f7-983cfee8f2a7" />


<img width="1200" height="446" alt="image" src="https://github.com/user-attachments/assets/82c65106-c5e4-45dd-9730-aa08edbf3e1f" />

Ref: https://www.influxdata.com/blog/infrastructure-monitoring-basics-telegraf-influxdb-grafana/

## What Is InfluxDB?

InfluxDB is an open source time series database (TSDB). 
This includes APIs for storing and querying data, processing it in the background for ETL, monitoring and alerting purposes, 
user dashboards, visualizing and exploring the data, and more.


Time series databases differ from relational databases because 
they are designed to capture and query continuously updated metrics rather than structured relationships.


InfluxDB is an ideal choice for IoT applications due to the following benefits:

* High-speed ingestion: Handles large volumes of sensor data efficiently.
* Optimized for time-series data: Supports high granularity and retention policies.
* Schema-less design: Easily adapts to various IoT data formats.
* Integrations: Works seamlessly with MQTT, Telegraf, and Grafana.

## Telegraf

<img width="1200" height="482" alt="image" src="https://github.com/user-attachments/assets/e6275969-a384-4d38-b7e0-8ee240ae37e2" />

## MQTT

MQTT (Message Queuing Telemetry Transport) is a lightweight publish/subscribe messaging protocol optimized for machine-to-machine (M2M) communication in IoT environments.

<img width="1200" height="675" alt="image" src="https://github.com/user-attachments/assets/94d9d582-918f-4014-b8f9-427093fdd666" />

* The broker is responsible for filtering messages based on topic, and distributing them to subscribed clients. 
* Using a broker removes the need for direct connections between every client.
