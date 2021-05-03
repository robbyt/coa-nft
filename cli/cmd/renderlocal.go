package cmd

import (
	"fmt"
	"os"

	"github.com/Boostport/address"
	"github.com/davecgh/go-spew/spew"
	inputData "github.com/robbyt/coa-nft/inputData/v1"

	"github.com/spf13/cobra"
)

// renderlocalCmd represents the renderlocal command
var renderlocalCmd = &cobra.Command{
	Use:   "renderlocal",
	Short: "Input lots of command line data, and this will render a COA from a template",
	Long: `Example:
	TBD.
`,
	RunE: renderLocal,
}

// Does some basic error handling for the addresses
func mungeAddress(a1, a2 string) []string {
	var mungedAddr []string
	if addressLine1 != "" && addressLine2 != "" {
		mungedAddr = []string{addressLine1, addressLine2}
	} else if addressLine1 != "" && addressLine2 == "" {
		mungedAddr = []string{addressLine1}
	} else {
		fmt.Println("Address format invalid")
		os.Exit(1)
	}
	return mungedAddr
}
func renderLocal(cmd *cobra.Command, args []string) error {
	id = inputData.DataFactory()

	// this prototype only supports a single person
	id.WorkTitle = workTitle
	id.People[0].Name = artistName
	id.People[0].Email = artistEmail
	id.People[0].SetPrimaryETHWallet(artistWallet)
	id.People[0].SetPrimaryWebsite(artistWebsite)

	// input address may have one or two lines, handle both here
	squashedAddress := mungeAddress(addressLine1, addressLine2)

	var addr address.Address
	if validateAddress {
		var err error
		addr, err = address.NewValid(
			address.WithCountry(addressCountry),
			address.WithName(artistName),
			address.WithStreetAddress(squashedAddress),
			address.WithLocality(addressCity),
			address.WithAdministrativeArea(addressState),
			address.WithPostCode(addressZip),
		)
		if err != nil {
			return err
		}
	} else {
		addr = address.New(
			address.WithCountry(addressCountry),
			address.WithName(artistName),
			address.WithStreetAddress(squashedAddress),
			address.WithLocality(addressCity),
			address.WithAdministrativeArea(addressState),
			address.WithPostCode(addressZip),
		)
	}
	id.People[0].SetPrimaryAddress(addr)

	spew.Dump(id)
	return nil
}

func init() {
	// using dot notation to access embedded object attributes within lists is jank af
	rootCmd.AddCommand(renderlocalCmd)
	renderlocalCmd.PersistentFlags().StringVarP(&templatePath, "template", "t", "", "This template file will be used to render the COA")

	renderlocalCmd.PersistentFlags().StringVar(&workTitle, "workTitle", "", "The title of this artwork")
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
