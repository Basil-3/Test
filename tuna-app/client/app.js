// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_employment_status").hide();
	$("#success_register").hide();
	$("#error_employment_status").hide();
    $("#error_query").hide();
    
    $scope.queryEmployee = function(){

		var id = $scope.employee_id;

		appFactory.queryEmployee(id, function(data){
			$scope.query_employee = data;

			if ($scope.query_employee == "Could not locate employee"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
    }
    
    $scope.queryLabor = function(){

		var id = $scope.labor_id;

		appFactory.queryLabor(id, function(data){
			$scope.query_labor = data;

			if ($scope.query_labor == "Could not locate labor"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
    }
    
    $scope.registerNAV = function(){

		appFactory.registerNAV($scope.nav, function(data){
			$scope.create_nav = data;
			$("#success_create").show();
		});
    }
    
    $scope.registerBank = function(){

		appFactory.registerBank($scope.bank, function(data){
			$scope.create_bank = data;
			$("#success_create").show();
		});
    }
    
    $scope.registerCompany = function(){

		appFactory.registerCompany($scope.msme, function(data){
			$scope.create_msme = data;
			$("#success_create").show();
		});
    }
    
    $scope.registerEmployee = function(){

		appFactory.registerEmployee($scope.employee, function(data){
			$scope.create_employee = data;
			$("#success_create").show();
		});
    }
    
    $scope.registerGovt = function(){

		appFactory.registerGovt($scope.govt, function(data){
			$scope.create_govt = data;
			$("#success_create").show();
		});
    }
    
    $scope.registerLabor = function(){

		appFactory.registerLabor($scope.labor, function(data){
			$scope.create_labor = data;
			$("#success_create").show();
		});
    }
    
    $scope.changeEmploymentStatus = function(){

		appFactory.changeEmploymentStatus($scope.employment_status, function(data){
			$scope.change_employment_status = data;
			if ($scope.change_employment_status == "Error: no employee found"){
				$("#error_employment_status").show();
				$("#success_employment_status").hide();
			} else{
				$("#success_employment_status").show();
				$("#error_employment_status").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
    var factory = {};
    
    factory.queryEmployee = function(id, callback){
    	$http.get('/get_employee/'+id).success(function(output){
			callback(output)
		});
    }
    
    factory.queryLabor = function(id, callback){
    	$http.get('/get_labor/'+id).success(function(output){
			callback(output)
		});
    }
    
    factory.registerNAV = function(data, callback){

		var nav = data.id + "-" + data.NAV_officer_id + "-" + data.NAV_officer_first_name + "-" + data.NAV_officer_last_name ;

    	$http.get('/add_NAV/'+nav).success(function(output){
			callback(output)
		});
    }
    
    factory.registerBank = function(data, callback){

		var bank = data.id + "-" + data.bank_branch_code + "-" + data.bank_name + "-" + data.bank_monetary_capacity + "-" + data.bank_region ;

    	$http.get('/add_Bank/'+bank).success(function(output){
			callback(output)
		});
    }
    
    factory.registerCompany = function(data, callback){

		var msme = data.id + "-" + data.company_registration_number + "-" + data.company_name + "-" + data.company_size + "-" + data.company_region ;

    	$http.get('/add_Company/'+msme).success(function(output){
			callback(output)
		});
    }
    
    factory.registerEmployee = function(data, callback){

		var employee = data.id + "-" + data.employee_id + "-" + data.employee_first_name + "-" + data.employee_last_name + "-" + data.employment_status + "-" + data.employee_age + "-" + data.employee_region ;

    	$http.get('/add_Employee/'+employee).success(function(output){
			callback(output)
		});
    }
    
    factory.registerGovt = function(data, callback){

		var govt = data.id + "-" + data.government_officer_id + "-" + data.government_officer_first_name + "-" + data.government_officer_last_name ;

    	$http.get('/add_Govt/'+govt).success(function(output){
			callback(output)
		});
    }
    
    factory.registerLabor = function(data, callback){

		var labor = data.id + "-" + data.labor_id + "-" + data.labor_first_name + "-" + data.labor_last_name + "-" + data.labor_age + "-" + data.labor_region ;

    	$http.get('/add_Labor/'+labor).success(function(output){
			callback(output)
		});
    }
    
    factory.changeEmploymentStatus = function(data, callback){

		var empstatus = data.id + "-" + data.employment_status;

    	$http.get('/change_employment_status/'+empstatus).success(function(output){
			callback(output)
		});
	}

	return factory;
});