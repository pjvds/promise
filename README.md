# Promise

I want to play with REST, Go and AngularJS. So Here is a small web application using it.

# Structure

├── controller
│   └── PromiseTicketController.go
├── data
│   ├── JsonMarshaller.go
│   ├── JsonMarshaller_test.go
│   ├── PromiseRepositoryFactory.go
│   ├── PromiseTicketRepository.go
│   ├── memory
│   │   └── InMemoryPromiseTicketRepository.go
│   └── mongo
│       ├── MongoPromiseRepository.go
│       ├── MongoPromiseRepositoryFactory.go
│       └── MongoPromiseSession.go
├── main.go
├── model
│   └── PromiseTicket.go
├── server
└── static
    ├── css
    │   └── app.css
    ├── home.html
    ├── img
    ├── index-async.html
    ├── index.html
    ├── js
    │   ├── app.js
    │   ├── controllers.js
    │   ├── directives.js
    │   ├── filters.js
    │   └── services.js
    ├── lib
    │   └── angular
    │       ├── angular-cookies.js
    │       ├── angular-cookies.min.js
    │       ├── angular-loader.js
    │       ├── angular-loader.min.js
    │       ├── angular-resource.js
    │       ├── angular-resource.min.js
    │       ├── angular-sanitize.js
    │       ├── angular-sanitize.min.js
    │       ├── angular.js
    │       ├── angular.min.js
    │       └── version.txt
    └── partials
        ├── partial1.html
        └── partial2.html

