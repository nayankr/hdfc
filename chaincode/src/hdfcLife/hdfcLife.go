package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// HDFC is a high level smart contract that HDFCs together business artifact based smart contracts
type HDFC struct {

}

// Application is for storing retreived Application

type Application struct{	
	ApplicationId string `json:"applicationId"`
	Status string `json:"status"`
	Title string `json:"title"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Gender string `json:"gender"`
	Dob string `json:"dob"`
	Age string `json:"age"`
	MartialStatus string `json:"martialStatus"`
	FatherName string `json:"fatherName"`
	MotherName string `json:"motherName"`
	Nationality string `json:"nationality"`
	ResidentialStatus string `json:"residentialStatus"`
	PlaceOfBirth string `json:"placeOfBirth"`
	PanNumber string `json:"panNumber"`
	AadharNumber string `json:"aadharNumber"`
	EducationalQualification string `json:"educationalQualification"`
	PoliticallyExposed string `json:"politicallyExposed"`
	DisablePersonPolicy string `json:"disablePersonPolicy"`
	AnyCriminalProceeding string `json:"anyCriminalProceeding"`
	LifeApprovalStatus string `json:"lifeApprovalStatus"`
	HealthApprovalStatus string `json:"healthApprovalStatus"`
	

}

// ListApplication is for storing retreived Application list with status
type ListApplication struct{	
	ApplicationId string `json:"applicationId"`
	Status string `json:"status"`
}

// CountApplication is for storing retreived Application count
type CountApplication struct{	
	Count int `json:"count"`
}

// Init initializes the smart contracts
func (t *HDFC) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("ApplicationTable")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ApplicationTable", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "applicationId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "age", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "martialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "fatherName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "motherName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nationality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "residentialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "placeOfBirth", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "panNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "aadharNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "educationalQualification", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "politicallyExposed", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "disablePersonPolicy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "anyCriminalProceeding", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeApprovalStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthApprovalStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
	return nil, nil
}


func (t *HDFC) getNumApplications(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0.")
	}

	var columns []shim.Column

	contractCounter := 0

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	for row := range rows {
		if len(row.Columns) != 0 {
			contractCounter++
		}
	}

	res2E := CountApplication{}
	res2E.Count = contractCounter
	mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil
}


func (t *HDFC) UpdateStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	newStatus := args[1]

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	//applicationId := row.Columns[0].GetString_()
	status := newStatus
	title := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	gender := row.Columns[5].GetString_()
	dob := row.Columns[6].GetString_()
	age := row.Columns[7].GetString_()
	martialStatus := row.Columns[8].GetString_()
	fatherName := row.Columns[9].GetString_()
	motherName := row.Columns[10].GetString_()
	nationality := row.Columns[11].GetString_()
	residentialStatus := row.Columns[12].GetString_()
	placeOfBirth := row.Columns[13].GetString_()
	panNumber := row.Columns[14].GetString_()
	aadharNumber := row.Columns[15].GetString_()
	educationalQualification := row.Columns[16].GetString_()
	politicallyExposed := row.Columns[17].GetString_()
	disablePersonPolicy := row.Columns[18].GetString_()
	anyCriminalProceeding := row.Columns[19].GetString_()
	lifeApprovalStatus:=row.Columns[20].GetString_()
	healthApprovalStatus:=row.Columns[21].GetString_()
	
	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: healthApprovalStatus}},
		}})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}


func (t *HDFC) UpdateStatusLife(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	newStatus := args[1]

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	//applicationId := row.Columns[0].GetString_()
	status := row.Columns[1].GetString_()
	title := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	gender := row.Columns[5].GetString_()
	dob := row.Columns[6].GetString_()
	age := row.Columns[7].GetString_()
	martialStatus := row.Columns[8].GetString_()
	fatherName := row.Columns[9].GetString_()
	motherName := row.Columns[10].GetString_()
	nationality := row.Columns[11].GetString_()
	residentialStatus := row.Columns[12].GetString_()
	placeOfBirth := row.Columns[13].GetString_()
	panNumber := row.Columns[14].GetString_()
	aadharNumber := row.Columns[15].GetString_()
	educationalQualification := row.Columns[16].GetString_()
	politicallyExposed := row.Columns[17].GetString_()
	disablePersonPolicy := row.Columns[18].GetString_()
	anyCriminalProceeding := row.Columns[19].GetString_()
	lifeApprovalStatus:=newStatus
	healthApprovalStatus:=row.Columns[21].GetString_()
	
	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: healthApprovalStatus}},
		}})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}



func (t *HDFC) UpdateStatusHealth(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	newStatus := args[1]

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	//applicationId := row.Columns[0].GetString_()
	status := row.Columns[1].GetString_()
	title := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	gender := row.Columns[5].GetString_()
	dob := row.Columns[6].GetString_()
	age := row.Columns[7].GetString_()
	martialStatus := row.Columns[8].GetString_()
	fatherName := row.Columns[9].GetString_()
	motherName := row.Columns[10].GetString_()
	nationality := row.Columns[11].GetString_()
	residentialStatus := row.Columns[12].GetString_()
	placeOfBirth := row.Columns[13].GetString_()
	panNumber := row.Columns[14].GetString_()
	aadharNumber := row.Columns[15].GetString_()
	educationalQualification := row.Columns[16].GetString_()
	politicallyExposed := row.Columns[17].GetString_()
	disablePersonPolicy := row.Columns[18].GetString_()
	anyCriminalProceeding := row.Columns[19].GetString_()
	lifeApprovalStatus:=row.Columns[20].GetString_()
	healthApprovalStatus:=newStatus
	
	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: healthApprovalStatus}},
		}})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}

func (t *HDFC) getApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
	}

	applicationId := args[0]
	

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + applicationId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + applicationId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := Application{}
	
	res2E.ApplicationId = row.Columns[0].GetString_()
	res2E.Status = row.Columns[1].GetString_()
	res2E.Title = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	res2E.LastName = row.Columns[4].GetString_()
	res2E.Gender = row.Columns[5].GetString_()
	res2E.Dob = row.Columns[6].GetString_()
	res2E.Age = row.Columns[7].GetString_()
	res2E.MartialStatus = row.Columns[8].GetString_()
	res2E.FatherName = row.Columns[9].GetString_()
	res2E.MotherName = row.Columns[10].GetString_()
	res2E.Nationality = row.Columns[11].GetString_()
	res2E.ResidentialStatus = row.Columns[12].GetString_()
	res2E.PlaceOfBirth = row.Columns[13].GetString_()
	res2E.PanNumber = row.Columns[14].GetString_()
	res2E.AadharNumber = row.Columns[15].GetString_()
	res2E.EducationalQualification = row.Columns[16].GetString_()
	res2E.PoliticallyExposed = row.Columns[17].GetString_()
	res2E.DisablePersonPolicy = row.Columns[18].GetString_()
	res2E.AnyCriminalProceeding = row.Columns[19].GetString_()
	res2E.LifeApprovalStatus = row.Columns[20].GetString_()
	res2E.HealthApprovalStatus = row.Columns[21].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}



func (t *HDFC) listAllApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0.")
	}

	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	res2E:= []*ListApplication{}	
	
	for row := range rows {
		newApp:= new(ListApplication)
		newApp.ApplicationId = row.Columns[0].GetString_()
		newApp.Status = row.Columns[1].GetString_()
		res2E=append(res2E,newApp)
	}
	
	res2F, _ := json.Marshal(res2E)
    fmt.Println(string(res2F))
	return res2F, nil

}


func (t *HDFC) getApplicationByPanNumber(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	lastPan := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	res2E := Application{}
	
	for row := range rows {
		fetchedPan := row.Columns[14].GetString_()
		
		if fetchedPan == lastPan{
			res2E.ApplicationId = row.Columns[0].GetString_()
			res2E.Status = row.Columns[1].GetString_()
			res2E.Title = row.Columns[2].GetString_()
			res2E.FirstName = row.Columns[3].GetString_()
			res2E.LastName = row.Columns[4].GetString_()
			res2E.Gender = row.Columns[5].GetString_()
			res2E.Dob = row.Columns[6].GetString_()
			res2E.Age = row.Columns[7].GetString_()
			res2E.MartialStatus = row.Columns[8].GetString_()
			res2E.FatherName = row.Columns[9].GetString_()
			res2E.MotherName = row.Columns[10].GetString_()
			res2E.Nationality = row.Columns[11].GetString_()
			res2E.ResidentialStatus = row.Columns[12].GetString_()
			res2E.PlaceOfBirth = row.Columns[13].GetString_()
			res2E.PanNumber = row.Columns[14].GetString_()
			res2E.AadharNumber = row.Columns[15].GetString_()
			res2E.EducationalQualification = row.Columns[16].GetString_()
			res2E.PoliticallyExposed = row.Columns[17].GetString_()
			res2E.DisablePersonPolicy = row.Columns[18].GetString_()
			res2E.AnyCriminalProceeding = row.Columns[19].GetString_()
			res2E.LifeApprovalStatus = row.Columns[20].GetString_()
			res2E.HealthApprovalStatus = row.Columns[21].GetString_()
			
			mapB, _ := json.Marshal(res2E)
			fmt.Println(string(mapB))
			
			return mapB, nil			
		}
	}
	
	return nil, errors.New("There is no application with the specified Pan Number.")

}



func (t *HDFC) listAllApplicationByStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	argStatus := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	res2E:= []*ListApplication{}	
	
	for row := range rows {
		fetchedStatus := row.Columns[1].GetString_()
		
		if fetchedStatus == argStatus{
			newApp:= new(ListApplication)
			newApp.ApplicationId = row.Columns[0].GetString_()
			newApp.Status = row.Columns[1].GetString_()
			res2E=append(res2E,newApp)
		}
	}
	
	res2F, _ := json.Marshal(res2E)
    fmt.Println(string(res2F))
	return res2F, nil

}



func (t *HDFC) listAllApplicationByLastName(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	argLastName := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	res2E:= []*ListApplication{}	
	
	for row := range rows {
		fetchedLastName := row.Columns[4].GetString_()
		
		if fetchedLastName == argLastName{
			newApp:= new(ListApplication)
			newApp.ApplicationId = row.Columns[0].GetString_()
			newApp.Status = row.Columns[1].GetString_()
			res2E=append(res2E,newApp)
		}
	}
	
	res2F, _ := json.Marshal(res2E)
    fmt.Println(string(res2F))
	return res2F, nil

}



// Invoke invokes the chaincode
func (t *HDFC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "submitApplication" {
		if len(args) != 22 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 22. Got: %d.", len(args))
		}

		applicationId := args[0]
		status := args[1]
		title := args[2]
		firstName := args[3]
		lastName := args[4]
		gender := args[5]
		dob := args[6]
		age := args[7]
		martialStatus := args[8]
		fatherName := args[9]
		motherName := args[10]
		nationality := args[11]
		residentialStatus := args[12]
		placeOfBirth := args[13]
		panNumber := args[14]
		aadharNumber := args[15]
		educationalQualification := args[16]
		politicallyExposed := args[17]
		disablePersonPolicy := args[18]
		anyCriminalProceeding := args[19]
		lifeApprovalStatus := args[20]
		healthApprovalStatus := args[21]

		// Insert a row
		ok, err := stub.InsertRow("ApplicationTable", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: healthApprovalStatus}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		return nil, err
	} else if function == "updateApplicationStatus" { 
		t := HDFC{}
		return t.UpdateStatus(stub, args)
	} else if function == "UpdateStatusLife" { 
		t := HDFC{}
		return t.UpdateStatusLife(stub, args)
	} else if function == "UpdateStatusHealth" { 
		t := HDFC{}
		return t.UpdateStatusHealth(stub, args)
	} 

	return nil, errors.New("Invalid invoke function name.")

}

func (t *HDFC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getApplication" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
		}
		t := HDFC{}
		return t.getApplication(stub, args)		
	}else if function == "listAllApplication" { 
		t := HDFC{}
		return t.listAllApplication(stub, args)
	}else if function == "getNumApplications" { 
		t := HDFC{}
		return t.getNumApplications(stub, args)
	}else if function == "getApplicationByPanNumber" { 
		t := HDFC{}
		return t.getApplicationByPanNumber(stub, args)
	}else if function == "listAllApplicationByStatus" { 
		t := HDFC{}
		return t.listAllApplicationByStatus(stub, args)
	}else if function == "listAllApplicationByLastName" { 
		t := HDFC{}
		return t.listAllApplicationByLastName(stub, args)
	}
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(HDFC))
	if err != nil {
		fmt.Printf("Error starting HDFC: %s", err)
	}
} 