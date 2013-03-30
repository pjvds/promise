'use strict';

/* Controllers */


function HomeController($scope, $http) {
	$http.get('/api/v1/promise').
		success(function(data){
			$scope.promises = data
		});

  	$scope.send = function() {
		var promise = { 'Name': $scope.name };
	    $http.post('/api/v1/promise',promise).
	    	success(function(data){
	        	$scope.success = true;
	        	$scope.msg = {"text": "done"};
	      	}).
	      	error(function(data){
	        	$scope.httpError = true;
	        	$scope.msg = data;
	      	});
	}
}
HomeController.$inject = ['$scope', '$http'];


function MyCtrl2() {
}
MyCtrl2.$inject = [];
