TICK stack: Telegraf, InfluxDB, Chronograf, and Kapacitor.

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

## MQTT

MQTT (Message Queuing Telemetry Transport) is a lightweight publish/subscribe messaging protocol optimized for machine-to-machine (M2M) communication in IoT environments.

<img width="1200" height="675" alt="image" src="https://github.com/user-attachments/assets/94d9d582-918f-4014-b8f9-427093fdd666" />

* The broker is responsible for filtering messages based on topic, and distributing them to subscribed clients. 
* Using a broker removes the need for direct connections between every client.
