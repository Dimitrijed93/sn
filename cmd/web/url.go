package main

import "os"

// GET_ALL api gateway url to retrieve all snippets
const getAll = "https://hdxa1sl70b.execute-api.eu-central-1.amazonaws.com/dev/snippets"
const insert = "https://hdxa1sl70b.execute-api.eu-central-1.amazonaws.com/dev/snippets"

var cloud = ":" + os.Getenv("PORT")

const local = ":4000"
