const express = require('express');
const app = express();
const cors =require('cors');
const bodyParser = require('body-parser');
const pino = require('pino');
const expressPino = require('express-pino-logger');
const logger = pino({ level: process.env.LOG_LEVEL || 'info' });
const expressLogger = expressPino({ logger });
const jwt = require('./_helpers/jwt');
const errorHandler = require('./_helpers/error-handler');

const port = process.env.SERVICE_PORT || 3000;

app.use(bodyParser.urlencoded({extended: true}));
app.use(bodyParser.json());
app.use(cors());
app.use(jwt());
app.use(expressLogger);
app.use('/catalog', require('./controllers/catalog.controller'));
app.use(errorHandler);

const server = app.listen(port, () => {
    logger.info(`Server started on :${port}`);
});