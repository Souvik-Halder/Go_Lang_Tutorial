const express = require('express');
const bodyParser = require('body-parser');
const app = express();
app.use(express.json());

app.get('/get', (req, res) => {
  res.send('hello world');
});

app.use(bodyParser.urlencoded({ extended: true }));

app.post('/post', (req, res) => {
  const body = req.body;
  console.log(body);
  res.status(201).json({ body });
});

app.post('/postform', (req, res) => {
  res.status(201).send(JSON.stringify(req.body));
});

const PORT = 5500;

app.listen(PORT, () => {
  console.log(`server is running on port ${PORT}`);
});
