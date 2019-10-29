package cmds

import (
	"fmt"
	"vault"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.NewVault(encodingKey, secretsPath())
		key, value := args[0], args[1]

		// output := "Key: " + key + " Value: " + value + " Encoding key: " + encodingKey
		// fmt.Println(output)
		// f, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
		// defer f.Close()

		err := v.Set(key, value)
		if err != nil {
			return
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
