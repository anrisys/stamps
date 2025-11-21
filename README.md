# Mini test Stamps

These two small programs are intended to fulfill recruitment process for developer position in Stamps.

## Prerequisites

- Go ^1.21
- API_KEY from OPEN_WEATHER_MAP

## Dependency

This application uses godotenv to simplyfy loading the api key from .env file.

## Directory structure

```
stamps/
├── .example.env                 # env file example
├── foo_bar.go                   # solution for mini test #1
├── main.go                      # entry point for all mini test solutions
├── types.go                     # all types uses for weather forecast
├── weather_forecast.go          # solution for mini test #2
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

## How to install and run the programs

1. Clone the repository

```
git clone https://github.com/anrisys/stamps.git
cd stamps
```

2. Configure environment variables:
   Copy the .example.env file and rename it to .env. Fill in your with your api key from open weather app.

```
cp .example.env .env
```

3. Install dependencies and run the server:

```
go mod tidy
go run .
```

The application will print out the result of both foo bar app and weather forecast in the terminal.
