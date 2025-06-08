const client = require('./client');

async function init() {
    await client.set('user:5', "this is user 5")
    await client.expire('user:5', 5)
    const result = await client.get('user:5')
    console.log("Result: ", result);
} 
init();