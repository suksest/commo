const fetch = require("node-fetch");
const sql = require('alasql');

async function getRate(){
    const url = "https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=8090322b8ffdddb63d5a";
    try {
        let response = await fetch(url);
        let rate = response.json();
        return await rate;
    } catch (error) {
        return await error;
    }
}

async function getList(){
    const url = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list";
    try {
        let response = await fetch(url);
        let list = response.json();
        return await list;
    } catch (error) {
        return await error;
    }
}

async function addUSD(arr){
    let rate = await getRate();
    let addedUSD = await arr.map(function(item) {
        item.priceIDR = Number(item.price);
        item.priceUSD = item.priceIDR * rate.IDR_USD;
        delete item.price;
        return item;
    });
    return await addedUSD;
}

exports.get = async function(){
    let list = await getList();
    let addedUSD = await addUSD(list);
    return await addedUSD;
};

exports.getStats = async function(){
    let list = await getList();
    let addedUSD = await addUSD(list);
    let res = sql('SELECT area_provinsi, timestamp, \
    AVG(priceIDR) AS mean_priceIDR, \
    MAX(priceIDR) AS max_priceIDR, \
    MIN(priceIDR) AS min_priceIDR, \
    AVG(priceUSD) AS mean_priceUSD, \
    MAX(priceUSD) AS max_priceUSD, \
    MIN(priceUSD) AS min_priceUSD \
    FROM ? GROUP BY area_provinsi, timestamp', [addedUSD]);
    return res;
}