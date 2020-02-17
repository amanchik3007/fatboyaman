package models

type Ogpo struct{
	Name 						string 		`json:"name"`
	LastName 					string 		`json:"last_name"`
	FirstName 					string 		`json:"first_name"`
	FatherName 					string 		`json:"father_name"`
	City 						string 		`json:"city"`
	PhoneNumber 				string 		`json:"phone_number"`
	Email 						string 		`json:"email"`
	IIN 						string 		`json:"iin"`
	ExpAge 						string 		`json:"exp_age"`
	VehicleType 				string 		`json:"vehicle_type"`
	RegistrationVehLoc		    string 		`json:"registration_veh_loc"`
	UsageDuration 				string 		`json:"usage_duration"`
	TotalPolicyPrice 			float64 	`json:"total_policy_price"`
	Language 					string 		`json:"language"`
}
