const express = require('express');
const itemRoutes = require('./controllers/itemController');

const app = express();
const PORT = 8081;

// Middleware
app.use(express.json());

// Routes
app.use('/items', itemRoutes);

// Health check endpoint
app.get('/', (req, res) => {
    res.json({ message: 'Item Service is running!' });
});

app.listen(PORT, () => {
    console.log(`Item Service is running on port ${PORT}`);
});

module.exports = app;
