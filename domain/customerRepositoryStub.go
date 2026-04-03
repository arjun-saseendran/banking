package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Maria", "Bridgen", "149344", "2000-01-01", "1"},
		{"1001", "Devika", "Mambala", "189344", "1998-02-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
