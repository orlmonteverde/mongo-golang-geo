# Golang and MongoDB Indexes

Introduction to go-mongo-driver . This as an alternative to the well-known mgo package using geospatial queries.

![MongoDB](https://img.shields.io/badge/MongoDB-v4.2.0-brightgreen.svg?logo=mongodb&longCache=true&style=flat) ![Go badge](https://img.shields.io/badge/Go-v1.13.0-blue.svg?logo=go&longCache=true&style=flat)

## Getting Started

This project uses the **Go** programming language (Golang) and the **MongoDB** database engine.

### Prerequisites

[MongoDB](https://www.mongodb.com/) is required in version 3 or higher and [Go](https://golang.org/) at least in version 1.12

### Installing

The dependency required for this project is the MongoDB driver for Go.

* go.mongodb.org/mongo-driver/mongo

#### Using GOPATH

```
go get -u -v go.mongodb.org/mongo-driver/mongo
```

#### Using GOMODULE
```
go build
```

## Deployment

Clone the repository
```
git clone git@github.com:orlmonteverde/mongo-golang-geo.git
```
Enter the repository folder
```
cd mongo-golang-geo
```
Build the binary
```
go build
```
Run the program
```
# In Unix-like OS
./mongo-golang-geo

# In Windows
mongo-golang-geo.exe
```

## Built With

* [go-mongo-driver](https://github.com/mongodb/mongo-go-driver) - The Go driver for MongoDB

## Authors

* **Orlando Monteverde** - *Initial work* - [orlmonteverde](https://github.com/orlmonteverde)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
