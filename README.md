<img src="https://raw.githubusercontent.com/scraly/gophers/main/friends.png" alt="friends" width=500>


# Guest list-service
This is a guest list service. You can add, delete, update, get guest info.

## Features
- Add a guest to the guest list
- Get the guest list
- Update the guest info
- Delete guest from list
- Get the arrived guest list
- Get the number of empty seats

## How to use
- use make command to start the docker-compose
```bash
make docker-up
```

## API demo
### Add a guest to the guest list
- `POST` /guest_list/:name
```bash
curl -X POST -H "Content-Type: application/json" -d '{"table" : 1, "accompanying_guests":1}' http://localhost:3000/guest_list/Tom
```

### Get the guest list
- `GET`  /guest_list
```bash
curl -X GET http://localhost:3000/guest_list
```

### Guest Arrives
- `PUT` /guests/:name
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"accompanying_guests":2}' http://localhost:3000/guests/Tom
```

### Guest Leaves
- `DELETE` /guests/:name
```bash
curl -X DELETE http://localhost:3000/guests/Tom
```

### Get arrived guests
- `GET` /guests
```bash
curl -X GET http://localhost:3000/guests
```

### Count number of empty seats
- `GET` /seats_empty
```bash
curl -X GET http://localhost:3000/seats_empty
```