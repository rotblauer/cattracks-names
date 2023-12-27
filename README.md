# CatTracks Names

Repository contains canonical cat names and their respective matching patterns.
A CLI tool is included to make working with GeoJSON cat tracks convenient.

## Build CLI

```sh
go build -o catnames-cli ./cli
```

# CLI Usage

```sh
# Print aliases (canonical cat names).
catnames-cli aliases

# Modify incoming GeoJSON, eg. changing cat names (eg. `properties.Name` values) to their aliases.
zcat ~/tdata/edge.json.gz | catnames-cli modify
```
