'use strict';

/* Services */


// Demonstrate how to register services
// In this case it is a simple value service
angular.module('myApp.services', []).
	directive('datetimepicker', function() {
    var datePickertemplate = '<div class="input-append"><input data-format="MM/dd/yyyy HH:mm:ss PP" type="text"></input><span class="add-on"><i data-time-icon="icon-time" data-date-icon="icon-calendar"></i></span></div>';
    return {
        restrict: 'E',
        compile:function(tElement, tAttrs, transclude){            
            var datePickerElement = angular.element(datePickertemplate);            
            tElement.replaceWith(datePickerElement);            
            return  function (scope, element, attr) { 
                 element.datetimepicker({
                      language: 'en',
                      pick12HourFormat: false
                });
            }
        }
    };
	}).
	value('version', '0.1');