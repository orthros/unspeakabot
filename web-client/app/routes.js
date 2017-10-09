module.exports = function (app) {

    // server routes ===========================================================
    // handle things like api calls
    // authentication routes
    app.get('/api/words', function (req, res) {
        res.status(200).json({ word: ["Hello world", "Asdf", "QWERTY"] });
    })
    // frontend routes =========================================================
    // route to handle all angular requests
    app.get('*', function (req, res) {
        res.sendfile('./public/index.html');
    });

};