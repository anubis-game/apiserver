package aggregator

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

const (
	Method = "executeBatch"
)

// Decode takes the calldata of a UserOperation object and returns all decoded
// arguments of all batches contained in the provided user operation based on
// the ABI below.
//
//     function executeBatch(address[],uint256[],bytes[]) public
//

func Decode(byt []byte) ([]Transaction, error) {
	var err error
	var exi bool

	var prs abi.ABI
	{
		prs, err = abi.JSON(strings.NewReader(AggregatorBindingABI))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var mth abi.Method
	{
		mth, exi = prs.Methods[Method]
		if !exi {
			return nil, tracer.Mask(fmt.Errorf("method %q not found in aggregator calldata", Method))
		}
	}

	unp := map[string]interface{}{}
	{
		err = mth.Inputs.UnpackIntoMap(unp, byt[4:])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var trg []common.Address
	{
		trg, err = decAdd(unp, "target")
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var val []*big.Int
	{
		val, err = decNum(unp, "value")
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var dat [][]byte
	{
		dat, err = decByt(unp, "callData")
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Verify that all slices have the same length. If there are 2 transactions
	// inside the calldata of one user operation, then we need to end up with 2
	// targets, 2 values and 2 calldatas respectively.
	if len(trg) != len(val) || len(val) != len(dat) {
		return nil, fmt.Errorf("transactions are inconsistent, got target=%d value=%d callData=%d", len(trg), len(val), len(dat))
	}

	// Create the slice of Transaction objects
	var txn []Transaction
	for i := range trg {
		txn = append(txn, Transaction{
			Target:   trg[i],
			Value:    val[i],
			CallData: dat[i],
		})
	}

	return txn, nil
}

func decAdd(unp map[string]interface{}, key string) ([]common.Address, error) {
	val, exi := unp[key]
	if !exi {
		return nil, tracer.Mask(fmt.Errorf("key %q not found in unpacked input map", key))
	}

	add, typ := val.([]common.Address)
	if !typ {
		return nil, tracer.Mask(fmt.Errorf("value %#v is not of type []common.Address", val))
	}

	return add, nil
}

func decNum(unp map[string]interface{}, key string) ([]*big.Int, error) {
	val, exi := unp[key]
	if !exi {
		return nil, tracer.Mask(fmt.Errorf("key %q not found in unpacked input map", key))
	}

	num, typ := val.([]*big.Int)
	if !typ {
		return nil, tracer.Mask(fmt.Errorf("value %#v is not of type []*big.Int", val))
	}

	return num, nil
}

func decByt(unp map[string]interface{}, key string) ([][]byte, error) {
	val, exi := unp[key]
	if !exi {
		return nil, tracer.Mask(fmt.Errorf("key %q not found in unpacked input map", key))
	}

	byt, typ := val.([][]byte)
	if !typ {
		return nil, tracer.Mask(fmt.Errorf("value %#v is not of type [][]byte", val))
	}

	return byt, nil
}
