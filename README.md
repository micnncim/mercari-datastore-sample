# mercari-datastore-sample

A sample of [mercari/datastore](https://github.com/mercari/datastore)

## Try it in local environment

```sh
$ export PROJECT_ID=<YOUR_PROJECT_ID>
$ gcloud beta emulators datastore start
$ $(gcloud beta emulators datastore env-init)
$ go run main.go
```

### routes

- `POST /CreateUser`
- `POST /ListUsers`
- `POST /UpdateUser`
- `POST /DeleteUser`

## LICENSE

[MIT](./LICENSE)
