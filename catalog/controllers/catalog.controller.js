const express = require('express');
const router = express.Router();
const catalogService = require('./catalog.service')

router.get('/', list);
router.get('/stats', stats);
module.exports = router;

function list(req, res, next){
    let list = catalogService.get();
    list.then(function(result){
        res.send(result);
    })
    list.catch(next);
}

function stats(req, res, next){
    let list = catalogService.getStats();
    list.then(function(result){
        res.send(result);
    })
    list.catch(next);
}