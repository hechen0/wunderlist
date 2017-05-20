package wunderlist

type List struct {

}

type listService service

//Get all Lists a user has permission to
//
//GET a.wunderlist.com/api/v1/lists
//
//Response
//Status: 200
//
//json
//[
//  {
//      "id": 83526310,
//	"created_at": "2013-08-30T08:29:46.203Z",
//	"title": "Read Later",
//	"list_type": "list",
//	"type": "list",
//	"revision": 10
//  }
//]
func (l *listService) All() []List{
	return []List(nil)
}

//Get a specific List
//
//GET a.wunderlist.com/api/v1/lists/:id
//Response
//
//Status: 200
//
//json
//{
//"id": 83526310,
//"created_at": "2013-08-30T08:29:46.203Z",
//"title": "Read Later",
//"list_type": "list",
//"type": "list",
//"revision": 10
//}
func (l *listService) Get(id int) List{
	return List{}
}

//Create a list
//
//POST a.wunderlist.com/api/v1/lists
//Data
//
//NAME	TYPE	NOTES
//title	string	required. maximum length is 255 characters
//Request body example
//
//json
//{
//"title": "Hallo"
//}
//Response
//
//Status: 201
//
//json
//{
//"id": 83526310,
//"created_at": "2013-08-30T08:29:46.203Z",
//"title": "Read Later",
//"revision": 1000,
//"type": "list"
//}
func (l *listService) Create() (err error){
	return
}

//
//Update a list by overwriting properties
//
//PATCH a.wunderlist.com/api/v1/lists/:id
//Data
//
//NAME	TYPE	NOTES
//revision	integer	required
//title	string	maximum length is 255 characters
//Request body example
//
//json
//{
//"revision": 1000,
//"title": "Hallo"
//}
//Response
//
//Status 200
//
//json
//{
//"id": 409233670,
//"revision": 1001,
//"title": "Hello",
//"type": "list"
//}
func (l *listService) Update() (err error){
	return
}

//Delete a list permanently
//
//DELETE a.wunderlist.com/api/v1/lists/:id
//Params
//
//NAME	TYPE	NOTES
//revision	integer	required
//Response
//
//Status 204
func (l *listService) Delete(id int) (err error){
	return
}