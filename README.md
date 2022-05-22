# ULIDs ![goreportcard](https://goreportcard.com/badge/github.com/fahmifan/ulids)

> ULIDs support null ULID & store ULID as string to DB. 

By default, the `ulid.ULID` from [github.com/oklog/ulid](https://github.com/oklog/ulid) is stored as binary, and it is hard to query it in DBMS, that's why I make this little helper that wraped the original `ulid.ULID` to store it as string to DB. Also, I'd like to avoid pointer in struct if possible, insipred by the [github.com/guregu/null](https://github.com/guregu/null) library, I added a `Null` type that can store `ulid.ULID` as `NULL` in DB and also supported null `JSON`.

## Install
```
go get github.com/fahmifan/ulids
```

## Example

Usage example can be found [here](https://github.com/fahmifan/shortly/blob/74e8dbfb30fbcb25bcc0c4c6840bbbfabf2b7ad9/repository/sqlite/url.go#L17)