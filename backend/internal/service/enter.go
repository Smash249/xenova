package service

var ServicesApp = new(servicesApp)

type servicesApp struct {
	UserService
	ProductService
	AccessoryService
	JournalismService
	SoftwareService
	HonorService
	SystemService
}
