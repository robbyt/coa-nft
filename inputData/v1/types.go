package v1

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Boostport/address"
)

// PersonRole is an Enum of various roles that people could be assigned
type PersonRole int

const (
	Creator PersonRole = iota
	Collaborator
	Curator
	Mentor
	Assistant
	Engineer
	Gallerist
)

// this must be updated when adding more Enum values above
func (s PersonRole) String() string {
	return [...]string{
		"Creator",
		"Collaborator",
		"Curator",
		"Mentor",
		"Assistant",
		"Engineer",
		"Gallerist",
	}[s]
}

// Addr wraps an address.Address objects, and gives it a human-readable Title (e.g., Office, Mailing, Gallery, etc.)
type Addr struct {
	ID    int
	Title string          // Name of this addr object
	Data  address.Address // object associated with this Addr object
}

// Date wraps a time.Time object, and gives it a human-readable Title (e.g., created, mastered, displayed, etc.)
type Date struct {
	ID    int
	Title string    // The name of this data object
	Date  time.Time // Actual object associated with this Date object
}

// Person is a object with several fields that represent a single Person involved in the art
type Person struct {
	ID                       int
	Name                     string       // Legal name to be displayed on the COA
	Email                    string       // Email address associated with this Person
	Roles                    []PersonRole // A slice of this person's roles in the art (they may have several)
	PrimaryRoleID            int          // The ID of this person's primary role
	Wallets                  []Wallet     // A slice of this person's crypto currency wallets
	PrimaryWalletID          int          // The ID of this person's primary wallet
	Websites                 []string     // A slice of this person's websites
	PrimaryWebsiteID         int          // The ID of this person's primary website
	PostalAddresses          []Addr       // A slice of address.Address objects, for storing postal addresses
	PrimaryPostalAddressesID int          // The ID of this person's primary postal address
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

// ValidEmailFormat checks the length and format of an email, returns an error if the format or size is invalid
func (p *Person) ValidEmailFormat() error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(p.Email) < 3 {
		return fmt.Errorf("Email address too short")
	}
	if len(p.Email) > 254 {
		return fmt.Errorf("Email address too long")
	}
	if !emailRegex.MatchString(p.Email) {
		return fmt.Errorf("Invalid email address format")
	}
	return nil
}

// SetPrimaryAddress initializes the addresses, and sets this address as primary
func (p *Person) SetPrimaryAddress(primaryAddr address.Address) {
	newAddr := Addr{Title: "primary", Data: primaryAddr}
	p.PostalAddresses = []Addr{newAddr}
	p.PrimaryPostalAddressesID = 0
}

// ArtMetaData is data about the art
type Mdata struct {
	Process    string   // The process for creating the art
	Media      string   // The media that has captured this art
	Genre      string   // Specific genre for this art
	Material   string   // Information about the material used in the creation of this art
	Workflow   string   // Note about the creation process of this art
	Dimensions string   // Physical or Digital dimensions
	Format     string   // File format, or other physical info
	Notes      []string // A slice of any other notes
}

// Xaction is data for the NFT transation or signature
type Xaction struct {
}

// Data is the main type exported from this package. It has the following fields:
type Data struct {
	WorkTitle       string   // The name of the work covered by this COA
	People          []Person // A slice of Person objects, for each of the people involved
	PrimaryPersonID int      // Which of the Person objects is the primary administrator/owner/creator
	Dates           []Date   // A slice of relevant Dates objects
	Metadata        Mdata    // Additional metadata
	Transaction     Xaction  // Unused info about the transaction
}

// V1DataFactory returns a very simplified object, mostly useful for prototyping.
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
