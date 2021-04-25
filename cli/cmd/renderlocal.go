package cmd

import (
	"fmt"
	"os"

	"github.com/Boostport/address"
	"github.com/davecgh/go-spew/spew"
	inputData "github.com/robbyt/coa-nft/inputData/v1"

	"github.com/spf13/cobra"
)

var templatePath string
var validateAddress bool
var id inputData.Data

var artistEmail string
var artistName string
var artistWallet string
var artistWebsite string

var addressLine1 string
var addressLine2 string
var addressCity string
var addressState string
var addressZip string
var addressCountry string

// renderlocalCmd represents the renderlocal command
var renderlocalCmd = &cobra.Command{
	Use:   "renderlocal",
	Short: "Input lots of command line data, and this will render a COA from a template",
	Long: `Example:
	TBD.
`,
	RunE: renderLocal,
}

func renderLocal(cmd *cobra.Command, args []string) error {
	// this prototype only supports a single person
	id.People[0].Email = artistEmail
	id.People[0].Name = artistName
	id.People[0].SetPrimaryETHWallet(artistWallet)
	id.People[0].SetPrimaryWebsite(artistWebsite)

	var mungedAddr []string
	if addressLine1 != "" && addressLine2 != "" {
		mungedAddr = []string{addressLine1, addressLine2}
	} else if addressLine1 != "" && addressLine2 == "" {
		mungedAddr = []string{addressLine1}
	} else {
		fmt.Println("Address format invalid")
		os.Exit(1)
	}

	addr, err := address.NewValid(
		address.WithCountry(addressCountry),
		address.WithName(artistName),
		address.WithStreetAddress(mungedAddr),
		address.WithLocality(addressCity),
		address.WithAdministrativeArea(addressState),
		address.WithPostCode(addressZip),
	)
	if err != nil {
		return err
	}
	id.People[0].SetPrimaryAddress(addr)

	spew.Dump(id)
	return nil
}

func init() {
	id = inputData.DataFactory()

	// using dot notation to access embedded object attributes within lists is jank af
	rootCmd.AddCommand(renderlocalCmd)
	renderlocalCmd.PersistentFlags().StringVarP(&templatePath, "template", "t", "template.tmpl", "This template file will be used to render the COA")
	renderlocalCmd.PersistentFlags().StringVar(&id.WorkTitle, "workTitle", "", "The title of this artwork")
	renderlocalCmd.MarkPersistentFlagRequired("workTitle")

	renderlocalCmd.PersistentFlags().StringVar(&artistEmail, "artistEmail", "", "Artist's Email")
	renderlocalCmd.MarkPersistentFlagRequired("artistEmail")

	renderlocalCmd.PersistentFlags().StringVar(&artistName, "artistName", "", "Artist's Name")
	renderlocalCmd.MarkPersistentFlagRequired("artistName")

	renderlocalCmd.PersistentFlags().StringVar(&artistWallet, "artistWallet", "", "Artist's ETH Wallet")
	renderlocalCmd.MarkPersistentFlagRequired("artistWallet")

	renderlocalCmd.PersistentFlags().StringVar(&artistWebsite, "artistWebsite", "", "Artist's Website")
	renderlocalCmd.MarkPersistentFlagRequired("artistWebsite")

	renderlocalCmd.PersistentFlags().BoolVar(&validateAddress, "skipValidateAddress", false, "Skip address validation (Uses https://chromium-i18n.appspot.com/ssl-address)")

	renderlocalCmd.PersistentFlags().StringVar(&addressLine1, "addressLine1", "", "Address - Line 1")
	renderlocalCmd.MarkPersistentFlagRequired("addressLine1")

	renderlocalCmd.PersistentFlags().StringVar(&addressLine2, "addressLine2", "", "Address - Line 2")

	renderlocalCmd.PersistentFlags().StringVar(&addressCity, "addressCity", "", "Address - City")
	renderlocalCmd.MarkPersistentFlagRequired("addressCity")

	renderlocalCmd.PersistentFlags().StringVar(&addressState, "addressState", "", "Address - State")
	renderlocalCmd.MarkPersistentFlagRequired("addressState")

	renderlocalCmd.PersistentFlags().StringVar(&addressZip, "addressZip", "", "Address - Zip")
	renderlocalCmd.MarkPersistentFlagRequired("addressZip")

	renderlocalCmd.PersistentFlags().StringVar(&addressCountry, "addressCountry", "US", "Address - Country")
}
