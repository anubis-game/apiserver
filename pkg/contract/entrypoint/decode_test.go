package entrypoint

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_EntryPoint_Decode(t *testing.T) {
	testCases := []struct {
		byt []byte
		ops []UserOperation
		ben common.Address
	}{
		// Case 000, Alchemy, https://sepolia.arbiscan.io/tx/0x310843e119e6c8d574a0ecf591fb5dc92ae3823f2ef5f135fac970e7a1678274
		{
			byt: hexutil.MustDecode("0x765e827f000000000000000000000000000000000000000000000000000000000000004000000000000000000000000005736be876755de230e809784def1937dcb6303e000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000005a68a239acb7f43647f5da19a9904aeec87e73650000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000786f00000000000000000000000000022bb8000000000000000000000000000000000000000000000000000000000000d21d00000000000000000000000007ee38d60000000000000000000000011784c30900000000000000000000000000000000000000000000000000000000000003e000000000000000000000000000000000000000000000000000000000000004a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026447e1da2a000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000010000000000000000000000009632185d3851fd06304c09ba6f1c1308189be12b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000104383cb9270000000000000000000000008988dbc479db976976cd956be4dd74e773b90a060000000000000000000000000000000000000000000000000000000067768c4b000000000000000000000000ad63b2262eb7d1591ee8e6a85959a523dece79830000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004141b627146deeebb99f7eee28d4192e0f028b8c8333d08f430fbdccfc5c9831871b49261f9970eb8bde6c325c140fe6deb01a563e88a0a6d6bc03520fb4d2aae71c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000812cc0c7981d846b9f2a16276556f6e8cb52bfb6330000000000000000000000000000715b00000000000000000000000000000000000000000000000067768ea54551357d9d8c01c7a94285d43a9d7c742641a413851ad73057947efd9edaf8be4bd12942beb1afb73602f31bc2a3eb2e9414ef41fd765c971a17c9e29ec086c91b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000420015c84d293ae94d1363603f7b4b9cb32410fc38a752caf80b55e89895095090ed685adb2146dc576303e9b512bc313ad1e1951c9ded9ef22e55f3a132a0bc6e4d1c000000000000000000000000000000000000000000000000000000000000"),
			ops: []UserOperation{
				{
					Sender:   common.HexToAddress("0x5a68A239acB7F43647f5dA19A9904aEeC87e7365"),
					CallData: hexutil.MustDecode("0x47e1da2a000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000010000000000000000000000009632185d3851fd06304c09ba6f1c1308189be12b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000104383cb9270000000000000000000000008988dbc479db976976cd956be4dd74e773b90a060000000000000000000000000000000000000000000000000000000067768c4b000000000000000000000000ad63b2262eb7d1591ee8e6a85959a523dece79830000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004141b627146deeebb99f7eee28d4192e0f028b8c8333d08f430fbdccfc5c9831871b49261f9970eb8bde6c325c140fe6deb01a563e88a0a6d6bc03520fb4d2aae71c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				},
			},
			ben: common.HexToAddress("0x05736be876755De230e809784DEF1937dCB6303e"),
		},
		// Case 001, Biconomy, https://sepolia.basescan.org/tx/0x480d6e3d721d1d5977545adf5fecf7af5343752ef7b8790d48900dc90a1e8987
		{
			byt: hexutil.MustDecode("0x1fad948c0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000125ad1988cf1ec8c2dd7357621ad82560376836500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000a6e3e87154b0455a936bb4df275c4a9baf901f82000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000180000000000000000000000000000000000000000000000000000000000005dcc20000000000000000000000000000000000000000000000000000000000010fe30000000000000000000000000000000000000000000000000000000000036bfb00000000000000000000000000000000000000000000000000000000000f434100000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000000000000000000000000000000000000000054000000000000000000000000000000000000000000000000000000000000006800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000038400004680000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000484c32b1288a88a48f8e7d20173a1048589df182000000000000000000000000206ab72edea55819a9a90622873976a79d3419e30000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044095ea7b3000000000000000000000000206ab72edea55819a9a90622873976a79d3419e30000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001446e43e82400000000000000000000000000000000000000000000000000062794ddf3c4680000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000006743be0000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000406166386430653663373735633432623633613736303339633333316566356235343437366563313862613463333933636262383632643437343764303931663300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011400000f79b7faf42eebadba19acc07cd08af44789000000000000000000000000efb36b2d443c5a6ff4127cda30944a12b421b9c2000000000000000000000000000000000000000000000000000000006741e5de000000000000000000000000000000000000000000000000000000006741ded600000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000041bfa65c2eedcb36cb4f29957a57a63b712c670a539162f94d46da9353db58cd8721c7e64cf5fe3333bd197900b005d1577dcdf0915106cac468365cb8ed363bc91b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000001c5b32f37f5bea87bdd5374eb2ac54ea8e00000000000000000000000000000000000000000000000000000000000000418479381033b35e9d41d493abffd9e0933781845e50126bf35db97d5bc5cb876b4a3aba61f5cee4b1de830a5d051dc1e702fd13c60fff466a1745d0fa450e7b8d1c00000000000000000000000000000000000000000000000000000000000000"),
			ops: []UserOperation{
				{
					Sender:   common.HexToAddress("0xA6E3e87154B0455A936bb4DF275c4A9bAF901F82"),
					CallData: hexutil.MustDecode("0x00004680000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000484c32b1288a88a48f8e7d20173a1048589df182000000000000000000000000206ab72edea55819a9a90622873976a79d3419e30000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044095ea7b3000000000000000000000000206ab72edea55819a9a90622873976a79d3419e30000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001446e43e82400000000000000000000000000000000000000000000000000062794ddf3c4680000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000006743be0000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000004061663864306536633737356334326236336137363033396333333165663562353434373665633138626134633339336362623836326434373437643039316633000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				},
			},
			ben: common.HexToAddress("0x125Ad1988Cf1Ec8c2Dd7357621ad825603768365"),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var ops []UserOperation
			var ben common.Address
			{
				ops, ben, err = Decode(tc.byt)
				if err != nil {
					t.Fatal(err)
				}
			}

			if len(ops) != len(tc.ops) {
				t.Fatal("expected", len(tc.ops), "got", len(ops))
			}
			if !bytes.Equal(ben.Bytes(), tc.ben.Bytes()) {
				t.Fatal("expected", tc.ben.Hex(), "got", ben.Hex())
			}

			for j := range ops {
				if !bytes.Equal(ops[j].Sender.Bytes(), tc.ops[j].Sender.Bytes()) {
					t.Fatal("expected", tc.ops[j].Sender.Hex(), "got", ops[j].Sender.Hex())
				}
				if !bytes.Equal(ops[j].CallData, tc.ops[j].CallData) {
					t.Fatal("expected", string(tc.ops[j].CallData), "got", string(ops[j].CallData))
				}
			}
		})
	}
}