// public/js/controllers/WordCtrl.js
angular.module('WordCtrl', ['WordService']).controller('WordController', function ($scope, Word) {

    $scope.tagline = 'Words are very important. Given your profession you should care about them more.';

    someData = Word.get();

    $scope.someData = someData.$$state;

    console.log($scope.someData);

});