<img width="863" height="543" alt="image" src="https://github.com/user-attachments/assets/1fa3d529-146a-4520-89ff-8c749e437056" />

<img width="1105" height="461" alt="image" src="https://github.com/user-attachments/assets/3364dd2e-23af-45c6-a51e-937734e2b1b0" />

## Architecture

<img width="2210" height="1212" alt="image" src="https://github.com/user-attachments/assets/7b6b7beb-c61e-4f00-8eff-4eff736dab49" />

<img width="2024" height="838" alt="image" src="https://github.com/user-attachments/assets/f4b4a716-1352-4695-bf89-e42dde82e654" />

Client libraries help in pulling metrics from the applications and send it to the Prometheus Server as direct instrumentation.

#### Exporters

<img width="1948" height="730" alt="image" src="https://github.com/user-attachments/assets/ce0b01b5-3a00-4812-8aa0-6c53e7ea01ab" />

Exporter - A Prometheus exporter is any application that exposes metric data in a format that can be collected (or
"scraped") by the Prometheus server.

For example, Node Exporter runs on a Linux machine and collects a variety of system metrics. It then exposes them to
Prometheus.

<img width="808" height="624" alt="image" src="https://github.com/user-attachments/assets/12145d8f-4b3e-4dc0-8260-53104f499ef2" />

By default these converted metrics are exposed by the exporter on /metrics path(HTTPS endpoint) of the target.

<img width="1364" height="758" alt="image" src="https://github.com/user-attachments/assets/4eb88768-a20b-40d7-b91d-7df2366b278e" />


<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/01e3ddab-3782-4831-87f8-c652ce5e864a" />

## Data Model

<img width="1820" height="1162" alt="image" src="https://github.com/user-attachments/assets/fc094b75-ecf6-47bb-9dd1-b996f74b8360" />

<img width="2100" height="1104" alt="image" src="https://github.com/user-attachments/assets/56935c13-dd3f-4e93-b74e-e94920ae151b" />

The system collects data samples over time, each of which is a tuple of an UNIX epoch and a float64 value.

```
(timestamp1, value1), (timestamp2, value2), ...
```

<img width="1594" height="442" alt="image" src="https://github.com/user-attachments/assets/58f09269-a5e6-4488-9932-4150daa26a7a" />

