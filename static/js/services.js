'use strict';

/* Services */


// Demonstrate how to register services
// In this case it is a simple value service
angular.module('myApp.services', []).
	directive('datetimepicker', function() {
    return {
      restrict: 'E',
      require: 'ngModel',
      scope: { 'ngModel': '=ngModel' },
      priority: 100,
      template: '<div class="input-append">'+
      					   '<input data-format="MM/dd/yyyy HH:mm:ss PP" data-ng-model="ngModel" type="text"></input>'+
      					   '<span class="add-on">'+
      					     '<i data-time-icon="icon-time" data-date-icon="icon-calendar"></i>'+
      					   '</span>'+
      					'</div>',
      replace: true,
      link:function($scope, elem, attr, ctrl) {
				var updateModel;
        updateModel = function(ev) {
        	var putObject;
        	putObject = function (path, object, value) {
				    var modelPath = path.split(".");
				 
				    function fill(object, elements, depth, value) {
				        var hasNext = ((depth + 1) < elements.length);
				        if(depth < elements.length && hasNext) {
				            if(!object.hasOwnProperty(modelPath[depth])) {
				                object[modelPath[depth]] = {};
				            }
				            fill(object[modelPath[depth]], elements, ++depth, value);
				        } else {
				            object[modelPath[depth]] = value;
				        }
				    }
				    fill(object, modelPath, 0, value);
					};

          $scope.ngModel = ev.date;
          $scope.$apply();
        };
        if (ctrl != null) {
          ctrl.$render = function() {
            elem.datetimepicker().data().datetimepicker.date = ctrl.$viewValue;
            elem.datetimepicker('setValue');
            elem.datetimepicker('update');
            return ctrl.$viewValue;
          };
        }
        return attr.$observe('datetimepicker', function(value) {
          var options;
          options = {};
          if (angular.isObject(value)) {
            options = value;
          }
          if (typeof(value) === "string" && value.length > 0) {
            options = angular.fromJson(value);
          }
          return elem.datetimepicker(options).on('changeDate', updateModel);
        });
      },
    };
	}).
	value('version', '0.1');