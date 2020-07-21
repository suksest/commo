const express = require('express');
const router = express.Router();
const fetch = require('node-fetch');

router.get('/', list);
module.exports = router;

function list (req, res, next) {
    const rateURL = "https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=8090322b8ffdddb63d5a";
    const url = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list";
    fetch(url)
    .then((resp) => resp.json())
    .then(function(data){
        let addedUSD = data.map(function(item){
            item.priceIDR = Number(item.price);
            item.priceUSD = Number(item.price);
            delete item.price;
            return item;
        });
        fetch(rateURL)
        .then((resp) => resp.json())
        .then(function(data){
            let rate = data.IDR_USD;
            let converted = addedUSD.map(function(item) {
                item.priceUSD *= rate;
                return item;
            })
            res.send(converted);
        })
        .catch(next)
    })
    .catch(next)
}