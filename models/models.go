package models

type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"db_cluster_identifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserEuid"`
}

type Category struct {
	CategId   int    `json:"Categ_Id"`
	CategName string `json:"Categ_Name"`
	CategPath string `json:"Categ_Path"`
}
