# Analyzer 🕵️‍♀️
Analyzer is a simple web application written in Go. It allows you to gather some basic analytics about any given website. Supported analytics:
1. **Document Type Definition** (extracting it from the HTML if possible)
2. **Website's title**
3. **Details about headings**
4. **Whether website contains login form** (searching for input types of password)
5. **List of links on the website**
    1. URL
    2. Is it an internal or external link?
    3. Number of occurrences of this link
    4. Whether it's reachable (`HEAD` request must return 200)
    5. HTTP status code for reachability check

# How to build and run 🛠
It's a very simple application, so not much is required to build and run this project. Below you can find necessary steps.
1. Clone this repository to your local machine.
2. Navigate to project's root directory.
3. Build and run application using:
    ```
    go run main.go
    ```
4. Open browser and navigate to `localhost:8080`
5. Have fun 😉

# Project details 🏗
This project used [go modules](https://golang.org/ref/mod) as it's dependency management tool. There's just one 3rd party library (not counting `golang.org/x/net`) called [goquery](https://github.com/PuerkitoBio/goquery). It's used for easy jQuery-like traversal of HTML.

## Directories structure:
```
.
├── LICENSE
├── README.md
├── config
│   └── tmpl.go
├── go.mod
├── go.sum
├── internal
│   ├── handler
│   │   └── handlers.go
│   ├── models
│   │   ├── error.go
│   │   └── report.go
│   └── website
│       ├── analyze.go
│       ├── analyze_test.go
│       └── fetch.go
├── main.go
└── web
    └── template
        ├── index.gohtml
        └── report.gohtml
```

# Possible suggestions and improvements 🩹
What's available right now is a nice starter with many possibilities for future improvements and new features. This section is split into two separate categories - one for technological improvements and second one with a product related improvements.
## Technology
1. Dockerize this app.
2. Replace `goquery` with `golang.org/x/net/html`, so that we can get rid of 3rd party dependencies.
3. Better modularization.
4. Improving test coverage and adding benchmark tests.
5. Creating a better UI and UX.
6. Meaningful logs.
7. Better routing.

## Product
1. Gathering more analytics
    + Whether website uses JS etc.
2. Extended login form checking to also include social network login buttons.
3. Create a DTD extractor which would be able to determine HTML version instead of just returning whole `<!DOCTYPE>`.
4. Improving URL parsing (for links), so that it takes into account local paths, removes whitespaces etc.
5. A CLI tool or just plain API instead of being a web app with a frontend.