package commands

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/axengine/utils/crypto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewDecryptCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "decrypt",
		Short: "解密字符串",
		Long:  "AES256 CBC PCSK0",
		RunE:  decrypt,
		PreRun: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().String("input", "", "带解密字符串，支持base64和hex编码")
	cmd.Flags().String("key", "", "密钥，字符串，长度16字节")
	cmd.Flags().String("iv", "", "向量，字符串，长度16字节")
	return cmd
}

func decrypt(cmd *cobra.Command, args []string) error {
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

	var (
		bz  []byte
		err error
	)
	bz, err = hex.DecodeString(input)
	if err != nil {
		bz, err = base64.StdEncoding.DecodeString(input)
	}
	if err != nil {
		return err
	}

	decrypted, err := crypto.AES256CBCPKCS0Decrypt(bz, []byte(iv), []byte(key))
	if err != nil {
		return errors.Wrap(err, "AES256_CBC_PKCS0Encrypt")
	}
	fmt.Println("Result string:", string(decrypted))
	fmt.Println("Result base64:", base64.StdEncoding.EncodeToString(decrypted))
	return nil
}
