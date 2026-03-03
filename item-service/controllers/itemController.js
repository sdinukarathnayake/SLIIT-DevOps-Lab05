const express = require('express');
const router = express.Router();

let items = ['Book', 'Laptop', 'Phone'];

router.get('/', (req, res) => {
    res.json(items);
});

router.post('/', (req, res) => {
    const item = req.body.item || req.body;
    
    if (!item) {
        return res.status(400).json({ error: 'Item is required' });
    }

    const itemToAdd = typeof item === 'string' ? item : item.toString();    
    items.push(itemToAdd);
    res.status(201).json({ message: `Item added: ${itemToAdd}` });
});

router.get('/:id', (req, res) => {
    const id = parseInt(req.params.id);

    if (isNaN(id)) {
        return res.status(400).json({ error: 'Invalid item ID' });
    }    
    if (id < 0 || id >= items.length) {
        return res.status(404).json({ error: 'Item not found' });
    }
    
    res.json({ item: items[id] });
});

module.exports = router;