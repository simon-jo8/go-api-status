# Go Status API

This API is a work in progress where I experiment in GO

## Getting Started

### Prerequisites
- Go 1.x or higher

### Running the API
```bash
go run cmd/api/main.go
```

## API Endpoints

### Golden Hour Calculation
Get the golden hour (sunset + 1 hour) for a specific location.

```bash
# Example for Paris coordinates (48.8566° N, 2.3522° E)
curl "http://localhost:8080/golden-hour?lat=48.8566&lng=2.3522"
```

### Plus One
Increment a number by one.

```bash
curl "http://localhost:8080/plus-one?number=2024"
```

## Example Coordinates
- Paris, France: 
  - Latitude: 48.8566° N
  - Longitude: 2.3522° E