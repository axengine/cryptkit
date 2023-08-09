package commands

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/axengine/utils/crypto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewEncryptCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "encrypt",
		Short: "加密字符串",
		Long:  "AES256 CBC PCSK0",
		RunE:  encrypt,
		PreRun: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().String("input", "", "待加密，字符串")
	cmd.Flags().String("key", "", "密钥，字符串，长度16字节")
	cmd.Flags().String("iv", "", "向量，字符串，长度16字节")
	return cmd
}

func encrypt(cmd *cobra.Command, args []string) error {
	input, _ := cmd.Flags().GetString("input")
	key, _ := cmd.Flags().GetString("key")
	iv, _ := cmd.Flags().GetString("iv")
	if len(input) == 0 {
		return errors.Errorf("bad input:%s", input)
	}
	if len(key) == 0 {
		return errors.Errorf("bad key:%s", key)
	}
	if len(iv) == 0 {
		return errors.Errorf("bad iv:%s", iv)
	}
	encrypted, err := crypto.AES256CBCPKCS0Encrypt([]byte(input), []byte(iv), []byte(key))
	if err != nil {
		return errors.Wrap(err, "AES256_CBC_PKCS0Encrypt")
	}
	fmt.Println("Result base64:", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Println("Result hex:", hex.EncodeToString(encrypted))
	return nil
}
