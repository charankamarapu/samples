const express = require('express');
const { MongoClient } = require('mongodb');

const app = express();

const client = new MongoClient('mongodb://root:pass@localhost:27017/admin', {
  useUnifiedTopology: true,
});

app.get('/', (req, res) => {
  res.send('Welcome to the world of animals.');
});

app.get('/animals', async (req, res) => {
  try {
    await client.connect();
    const db = client.db('animal_db');
    const animals = await db.collection('animal_tb').find().toArray();
    res.json({ animals });
  } catch (err) {
    console.error(err);
    res.sendStatus(500);
  } finally {
    await client.close();
  }
});

app.listen(5000, () => {
  console.log('Server started on port 5000');
});
