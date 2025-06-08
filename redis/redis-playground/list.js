const client = require('./client');

async function implementList() {
    // const res = await client.lpush("messages", "msg1");
    const res = await client.lpop("messages");
    console.log(res)
}
implementList();