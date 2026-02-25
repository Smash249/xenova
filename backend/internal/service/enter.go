package service

var ServicesApp = new(servicesApp)

type servicesApp struct {
	UserService
	ProductService
	JournalismService
	SoftwareService
	HonorService
	SystemService
}
