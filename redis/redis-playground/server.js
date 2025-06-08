const express = require('express');
const app = express();
const axios = require('axios').default;
const client = require('./client');

app.get('/', async (req, res) => {
    const cacheValue = await client.get('todos');
    if(cacheValue) {
        console.log('Cache hit');
        return res.json(JSON.parse(cacheValue));
        // parse to convert into JSON object, as the data stored in Redis is string
    }
    const response = await axios.get('https://jsonplaceholder.typicode.com/todos');
    await client.set('todos', JSON.stringify(response.data));
    res.json(response.data);
});

app.listen(4000, () => {   
    console.log('Server is running on port 4000');
});