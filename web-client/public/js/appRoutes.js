// public/js/appRoutes.js
angular.module('appRoutes', []).config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {

    $routeProvider

        // home page
        .when('/', {
            templateUrl: 'views/home.html',
            controller: 'MainController'
        })

        // words page that will use the WordController
        .when('/words', {
            templateUrl: 'views/words.html',
            controller: 'WordController'
        });

    $locationProvider.html5Mode(true);

}]);
