package db

// Landlord represents a property owner (tenant)
type Landlord struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	TenantID  string  `json:"tenant_id"`
}

// Property represents a property owned by a landlord
type Property struct {
	ID        string `json:"id"`
	LandlordID string `json:"landlord_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Type      string  `json:"single_type,multi-type"`
	Description string `json:"description"`
}

// Unit represents a unit within a property
type Unit struct {
	ID         string `json:"id"`
	PropertyID string `json:"property_id"`
	LandlordID string `json:"landlord_id"`
	Name       string `json:"name"`
	Description string `json:"description"`
	Status     string `json:"status"` // e.g., occupied, vacant, under-repair
}

// Renter represents a client occupying a unit
type Renter struct {
	ID       string `json:"id"`
	UnitID   string `json:"unit_id"`
	LandlordID string `json:"landlord_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	NationalID  string `json:"national_id"`
	LeaseStart string `json:"lease_start"`
	LeaseEnd   string `json:"lease_end"`
}


type complaint struct {
    Title   string   `json:"Title"`
	Date   string    `json:"date"`
	Description   string   `json:"description"`
	Status      string     `json:"status`
}