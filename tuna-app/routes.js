//SPDX-License-Identifier: Apache-2.0

var abcp = require('./controller.js');

module.exports = function(app){

  app.get('/get_employee/:id', function(req, res){
    abcp.get_employee(req, res);
  });
  app.get('/get_labor/:id', function(req, res){
    abcp.get_labor(req, res);
  });
  app.get('/add_NAV/:nav', function(req, res){
    abcp.add_NAV(req, res);
  });
  app.get('/add_Bank/:bank', function(req, res){
    abcp.add_Bank(req, res);
  });
  app.get('/add_Company/:msme', function(req, res){
    abcp.add_Company(req, res);
  });
  app.get('/add_Employee/:employee', function(req, res){
    abcp.add_Employee(req, res);
  });
  app.get('/add_Govt/:govt', function(req, res){
    abcp.add_Govt(req, res);
  });
  app.get('/add_Labor/:labor', function(req, res){
    abcp.add_Labor(req, res);
  });
  app.get('/change_employment_status/:empstatus', function(req, res){
    tuna.change_employment_status(req, res);
  });
}
