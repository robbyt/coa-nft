package cmd

import (
	"os"

	"github.com/Boostport/address"
	"github.com/davecgh/go-spew/spew"
	inputData "github.com/robbyt/coa-nft/inputData/v1"
	q "github.com/robbyt/coa-nft/inputData/v1/questions"

	"github.com/spf13/cobra"
	"github.com/tcnksm/go-input"
)

// wizardCmd represents the wizard command
var wizardCmd = &cobra.Command{
	Use:   "wizard",
	Short: "An interactive 'wizard' that asks several questions needed to generate the COA",
	Long: `Example:
	TBD.
`,
	RunE: wizard,
}

func wizard(cmd *cobra.Command, args []string) (err error) {
	id = inputData.DataFactory()

	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	id.WorkTitle, err = ui.Ask(q.WorkTitle, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}

	id.People[0].Name, err = ui.Ask(q.ArtistName, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}

	id.People[0].Email, err = ui.Ask(q.ArtistEmail, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
		ValidateFunc: func(s string) error {
			p := inputData.Person{Email: s}
			return p.ValidEmailFormat()
		},
	})
	if err != nil {
		return err
	}

	artistWallet, err = ui.Ask(q.ArtistWallet, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	id.People[0].SetPrimaryETHWallet(artistWallet)

	artistWebsite, err = ui.Ask(q.ArtistWebsite, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	id.People[0].SetPrimaryWebsite(artistWallet)

	// start address handling
	addressLine1, err = ui.Ask(q.AdressLine1, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	addressLine2, err = ui.Ask(q.AdressLine2, &input.Options{
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	squashedAddress := mungeAddress(addressLine1, addressLine2)

	addressCountry, err = ui.Ask(q.Country, &input.Options{
		Default:   "US", // sorry for the ethnocentrism
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}

	addressCity, err = ui.Ask(q.City, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}

	addressState, err = ui.Ask(q.State, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	addressZip, err = ui.Ask(q.Zip, &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}

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

	rootCmd.AddCommand(wizardCmd)

	wizardCmd.PersistentFlags().StringVarP(&templatePath, "template", "t", "", "This template file will be used to render the COA")
	wizardCmd.PersistentFlags().BoolVar(&validateAddress, "skipValidateAddress", false, "Skip address validation (Uses https://chromium-i18n.appspot.com/ssl-address)")

}
