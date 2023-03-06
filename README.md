# Historical OHLC Price Data

We have just purchased a large amount of historical OHLC price data that were shared to us in CSV files format. We need to start centralising and digitalising those data. These files can be ranging from a few GBs to a couple of TBs.
 

## Quickstart

### Deploy using Docker

1. Make sure you have Docker installed.
2. Clone this repos to your local machine

```bash
git clone https://github.com/vandanapareek/OHLC-Price-Data.git
```

3. Build or rebuild the services

```bash
docker compose build
```
4. Start a container that serves the development version of the app

```bash
# Open your browser at http://localhost:8080 to access the app
docker compose up
```
