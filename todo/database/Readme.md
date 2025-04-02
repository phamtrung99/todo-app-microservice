## Create new migration file

Prequisite: [Install go-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#migrate-cli)

Create new: 
``` bash
migrate create -ext sql -dir ./ -seq create_users_table

```