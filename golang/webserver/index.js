/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require('express')
const app = express()
const port = 8000

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Welcome to LearnCodeonline server")
})

app.get('/get', (req, res) => {
    res.status(200).json({message: "Hello from learnCodeonline.in"})
})


app.post('/post', (req, res) => {
    let myJson = req.body;      // your JSON
	res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})

// http://localhost:8000/postform
// on this route when using postman fill the data inside form-encoded not in raw json or in forms
  

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})

// Our aim here is to handle these backend routes of a server using golang
// consider this mock as our mock server, now we will create requests to this server using golang   

/*
    We will be handling data in 3 different ways.
    1. Get request
    2. Post request with JSON data
    3. Post request with form data -> there are special cases where we would be needing this, like image upload etc.
*/