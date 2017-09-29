require('dotenv').config()
var express = require('express');
var bodyParser = require('body-parser');
var methodOverride = require('method-override');

var app = express();

var port = process.env.PROCESS_PORT;


// get all data/stuff of the body (POST) parameters
// parse application/json 
app.use(bodyParser.json());

// parse application/vnd.api+json as json
app.use(bodyParser.json({ type: 'application/vnd.api+json' }));

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: true }));

// override with the X-HTTP-Method-Override header in the request. simulate DELETE/PUT
app.use(methodOverride('X-HTTP-Method-Override'));

// set the static files location /public/img will be /img for users
app.use(express.static(__dirname + '/public'));

// routes ==================================================
require('./app/routes')(app); // configure our routes

// start app ===============================================
app.listen(port);

console.log('Listening on port ' + port);

// expose app           
exports = module.exports = app;

