# Go basic CRUD API with Fiber
This mini-project is a simple CRUD (Create, Read, Update, Delete) API built using Golang-Fiber. <br> Connetcing to MySQL database by using Fiber, following a tutorial from freeCodeCamp.org.

## Features
__Create__: Add a new infomation into the database.<br>
__Read__: Get all infomation or a specific company by ID.<br>
__Update__: Modify an existing company's details.<br>
__Delete__: Flag delete a company from the collection in the database.<br>

## API Endpoints
POST api/v1/lead - Add a new infomation to the database.<br>
GET api/v1/lead - Retrieve all company.<br>
GET api/v1/lead/{id} - Retrieve a company by its ID.<br>
PUT api/v1/lead{id} - Update an existing company by its ID.<br>
DELETE /api/v1/lead/{id} - Flag delete a company by its ID.<br>

## Tutorial Reference
This project is based on the freeCodeCamp.org tutorial:[⌨️ CRM with Golang Fiber](https://www.youtube.com/watch?v=jFfo23yIWac) 

## Problem
- Error : Incorrect table name '' when doing put method.
- Solution : using c.BodyParser before update data.
