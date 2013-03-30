'use strict';

/* Controllers */


function MyCtrl1($scope, $http) {

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
MyCtrl1.$inject = ['$scope', '$http'];


function MyCtrl2() {
}
MyCtrl2.$inject = [];
