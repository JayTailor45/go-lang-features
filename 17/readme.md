# Event Management GO API

## Features

- user can signup
- user can create, update, delete, list the events
- user can register to an event
- user can cancel thier registraion to an event
- only authorized user can create, update, delete events
- non authorized user can see events and event details
- authorized user can also see all event registrations
- user can update and delete events they created

## API details

| Method   | API endpoint          | Auth     | Description                          | Payload                                                                                                                                                 |
| -------- | --------------------- | -------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `POST`   | /signup               |          | create a user                        | { "email": "jay@tailor.com", "password": "Password", }                                                                                                  |
| `POST`   | /login                |          | login and returns token              | { "email": "jay@tailor.com", "password": "Password", }                                                                                                  |
| `GET`    | /events               |          | return list of events                |                                                                                                                                                         |
| `GET`    | /events/:id           |          | return details of event              |                                                                                                                                                         |
| `POST`   | /events               | Required | create event                         | { "Name": "Go meetup", "Description": "A Go meetup to network with like minded people", "Location": "Surat", "DateTime": "2026-04-17T15:30:00.000Z" }   |
| `PUT`    | /events/:id           | Required | update event                         | { "Name": "Go meetup", "Description": "A Go meetup to network with like minded people", "Location": "Bardoli", "DateTime": "2026-04-17T15:30:00.000Z" } |
| `DELETE` | /events/:id           | Required | delete event                         |                                                                                                                                                         |
| `POST`   | /events/:id/register  | Required | register user to event               |                                                                                                                                                         |
| `DELETE` | /events/:id/cancel    | Required | cancel user registertaion to a event |                                                                                                                                                         |
| `GET`    | /events/registrations | Required | returns all registrations            |                                                                                                                                                         |

## Run the project

1. `go mod tidy`
2. `go run .`

## How to build and run the project

1. `go mod tidy`
2. `go build`
3. `./event-management.exe`
