# MIGRATIONS GO VOTHINGS

Description migrations

### Installation

```
$ go get github.com/gobuffalo/pop/...
```

```
$ go install github.com/gobuffalo/pop/v6/soda@latest
```

## How to Use

```
$ soda -c configs/database.yml create -a
```

```
$ soda generate sql --config configs/database.yml CreateVothingTable
```

```
$ soda -c configs/database.yml migate
```
