package cmd

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-deploy-contract",
	Short: "Go Deploy Smart Contract",
	Long: `
   Go Deploy Smart Contract is a smart contract deployer
   to any given test or main network.
	`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		solCheck := exec.Command("solc", "--version")
		var out bytes.Buffer
		solCheck.Stdout = &out
		err := solCheck.Run()
		if err != nil {
			fmt.Println("solc not installed. Please install solc then try.")
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Parse flags
		abi, _ := cmd.Flags().GetString("abi")
		bin, _ := cmd.Flags().GetString("bin")

		if abi == "" || bin == "" {
			fmt.Println("Please specify abi and bin file")
			return
		}

		//
		abiString, binString, err := ReadABIBIN(abi, bin)
		if err != nil {
			fmt.Println("Unable to read files ", err.Error())
			return
		}

		parsed, err := gethAbi.JSON(strings.NewReader(abiString))
		if err != nil {
			fmt.Println("Unable to parse : ", parsed)
			return
		}

		pb, _ := json.MarshalIndent(parsed, "", " ")
		log.Println(string(pb))
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("Contract needs %d parameters of %s", parsed.Constructor.Inputs.LengthNonIndexed(),
				parsed.Constructor.Sig()),
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Println("Unable to run promt : ", err.Error())
		}
		fmt.Println("Result", result)

		client, err := GetClient("http://127.0.0.1:7545")
		if err != nil {
			fmt.Println("Unable to get client : ", err.Error())
			return
		}

		auth, err := GetAuthFromPrivateKey("37e10f3d5e2a65147b6d43fcb40fd26f359243bf0d12ec7b666d28413554bbd1", client)
		params := strings.Split(result, ",")
		new := make([]interface{}, len(params))
		for i, v := range params {
			switch parsed.Constructor.Inputs[i].Type.T {
			case 1:
				u, _ := strconv.ParseUint(v, 10, 8)
				switch parsed.Constructor.Inputs[i].Type.Size {
				case 8:
					value := reflect.ValueOf(uint8(u))
					new[i] = value.Interface()
				}
			case 3:
				new[i] = v
			}
		}
		address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(binString), client, new...)
		if err != nil {
			fmt.Println("Unable to deploy Contract : ", err.Error())
			return
		}
		fmt.Println(address.Hex())
		fmt.Println(tx.Hash().Hex())

		_ = contract
	},
}

// Execute executes the root command.
func Execute() error {
	rootCmd.PersistentFlags().StringP("abi", "a", "", "abi file location")
	rootCmd.PersistentFlags().StringP("bin", "b", "", "bin file location")
	return rootCmd.Execute()
}

// ReadABIBIN reads abi and bin file content
func ReadABIBIN(abi, bin string) (string, string, error) {
	abiString, err := ioutil.ReadFile(abi)
	if err != nil {
		return "", "", err
	}

	binString, err := ioutil.ReadFile(bin)
	if err != nil {
		return "", "", err
	}

	return string(abiString), string(binString), nil
}

func GetClient(url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}

func GetAuthFromPrivateKey(private string, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Unable to generate public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

/*
func SolidityWatcher(file string, abi bool, bin bool, bindgo bool, dest string) {
	// Generate once before starting Watcher
	err := generator.Generate(file, abi, bin, bindgo, dest)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Generated !!")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					err := generator.Generate(file, abi, bin, bindgo, dest)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
					fmt.Println("Generated !!")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(file)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}*/
