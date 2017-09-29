// public/js/services/WordService.js
angular.module('WordService', []).factory('Word', ['$http', function ($http) {

    return {
        // call to get all nerds
        get: function () {
            return $http.get('/api/words');
        },

        // these will work when more API routes are defined on the Node side of things
        // call to POST and create a new nerd
        create: function (wordData) {
            return $http.post('/api/words', wordData);
        },

        // call to DELETE a nerd
        delete: function (word) {
            return $http.delete('/api/words/' + word);
        }
    }

}]);
