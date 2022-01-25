# json-cli

Pretty print, Compact & Stringify JSON from command line.

## Build
```bash
go build -o bin/json .
```

## Usage

### Console input
```bash
json -p
{"firstName":"John","lastName":"Doe"}

{
    "firstName": "John",
    "lastName": "Doe"
}
```

### As argument
```bash
json -p '{"firstName":"John","lastName":"Doe"}'

{
    "firstName": "John",
    "lastName": "Doe"
}
```

### With pipe
```bash
echo '{
    "firstName": "John",
    "lastName": "Doe"
}' | json -c

{"firstName":"John","lastName":"Doe"}
```

## Available Options

| Option | Description                                |
|--------|--------------------------------------------|
| -c     | Compact: Compact JSON input in single line |
| -p     | Prettify: Print prettified JSON input      |
| -s     | Stringify: Print stringified JSON string   |
