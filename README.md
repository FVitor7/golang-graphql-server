# Running the project

1. Create and start container:
```
task run
```
or

```
docker-compose up -d
```

# Accessing graphql

1. Access:
```
localhost:8080/graphql
```


# Mutations Example

1. Create Brand:
```
mutation createBrand {
  createBrand(input: {name: "Fiat"}) {
    name
  }
}
```

2. Create Car:
```
mutation createCar {
  createCar(input: {name: "Punto", brandID: $brandID}) {
    name
  }
}

```

# Query Example

1. List all Brands:
```
query listBrands{
  brands {
    id
    name
  }
}
```

2. List all Cars:
```
query listCars {
  cars {
    id
    name
    brand {
      name
      id
    }
  }
}

```

# Stopping the project

```
task stop
```
or

```
docker-compose down -v
```