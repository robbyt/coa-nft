package v1

import (
	"time"

	"github.com/Boostport/address"
)

// PersonRole is an Enum of various roles that people could be assigned
type PersonRole int

const (
	Creator PersonRole = iota
	Collaborator
	Mentor
	Assistant
	Engineer
)

// this must be updated when adding more Enum values above
func (s PersonRole) String() string {
	return [...]string{
		"Creator",
		"Collaborator",
		"Mentor",
		"Assistant",
		"Engineer",
	}[s]
}

/* Addr wraps an address.Address objects, and gives it a human-readable Title
 (e.g., Office, Mailing, Gallery, etc.)

Title - The name of this Addr object
Data - address.Address object associated with this Addr object
*/
type Addr struct {
	ID    int
	Title string
	Data  address.Address
}

/* Date wraps a time.Time object, and gives it a human-readable Title (e.g., created, mastered, displayed, etc.)

Title - The name of this data object
Date - time.Time object associated with this Date object
*/
type Date struct {
	ID    int
	Title string
	Date  time.Time
}

/*
Person is a object with several fields that represent a single Person involved in the art

Name - Legal name to be displayed on the COA
Email - Email address associated with this Person
Roles - A slice of this person's roles in the art (they may have several)
PrimaryRoleID - The ID of this person's primary role
Wallets - A slice of this person's crypto currency wallets
PrimaryWalletID - The ID of this person's primary wallet
Websites - A slice of this person's websites
PrimaryWebsiteID - The ID of this person's primary website
PostalAddresses - A slice of address.Address objects, for storing postal addresses
PrimaryPostalAddressesID - The ID of this person's primary postal address
*/
type Person struct {
	ID                       int
	Name                     string
	Email                    string
	Roles                    []PersonRole
	PrimaryRoleID            int
	Wallets                  []Wallet
	PrimaryWalletID          int
	Websites                 []string
	PrimaryWebsiteID         int
	PostalAddresses          []Addr
	PrimaryPostalAddressesID int
}

// SetPrimaryRole initalizes and sets a new primary value
func (p *Person) SetPrimaryRole(primaryRole PersonRole) {
	// initialize the PersonRole slice with this value
	p.Roles = []PersonRole{primaryRole}
	p.PrimaryRoleID = 0
}

// SetPrimaryWebsite initalizes and sets a new primary website
func (p *Person) SetPrimaryWebsite(primaryWS string) {
	p.Websites = []string{primaryWS}
	p.PrimaryWebsiteID = 0
}

// SetPrimaryETHWallet initializes the wallet object (if not already), and sets an ETH value in that object
func (p *Person) SetPrimaryETHWallet(ethVal string) {
	newWal := Wallet{
		Title: "primary",
		Value: ethVal,
		Type:  ETH,
	}

	if p.Wallets == nil {
		p.Wallets = []Wallet{newWal}
		p.PrimaryWalletID = 0
	} else {
		// already initalized, update primary
		p.Wallets[p.PrimaryWalletID].Value = ethVal
	}
}

// SetPrimaryAddress initializes the addresses, and sets this address as primary
func (p *Person) SetPrimaryAddress(primaryAddr address.Address) {
	newAddr := Addr{Title: "primary", Data: primaryAddr}
	p.PostalAddresses = []Addr{newAddr}
	p.PrimaryPostalAddressesID = 0
}

/* ArtMetaData is data about the art

Process - The process for creating the art
Media - The media that has captured this art
Genre - Specific genre for this art
Material - Information about the material used in the creation of this art
Workflow - Note about the creation process of this art
Dimensions - Physical or Digital dimensions
Notes - A slice of any other notes
*/
type Mdata struct {
	Process    string
	Media      string
	Genre      string
	Material   string
	Workflow   string
	Dimensions string
	Notes      []string
}

/* Xaction is data for the NFT transation or signature
 */
type Xaction struct {
}

/*
Data is the main type exported from this package. It has the following fields:

WorkTitle - the name of the work covered by this COA
People - a slice of Person objects, for each of the people involved
PrimaryPersonID - Which of the Person objects is the primary administrator/owner/creator
Dates - a slice of relevant Dates objects
*/
type Data struct {
	WorkTitle       string
	People          []Person
	PrimaryPersonID int
	Dates           []Date
	Metadata        Mdata
	Transaction     Xaction
}

// V1DataFactory returns a very simplified object, useful for prototyping
func DataFactory() Data {
	id := Data{}

	// initialize the People slice with an empty person
	id.People = []Person{Person{}}

	// This verb only supports a single person, as the creator right now
	id.People[0].SetPrimaryRole(Creator)
	id.People[0].SetPrimaryWebsite("")

	id.People[0].Wallets = []Wallet{Wallet{Title: "primary"}}

	return id
}
