package entrypoint

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

const (
	Method = "handleOps"
)

// Decode takes the calldata of a types.Transaction object and returns all
// decoded arguments contained in the associated user operations based on the
// implemented ABIs below.
//
//     Alchemy:     function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[], address) public
//     Biconomy:    function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[], address) public
//

func Decode(byt []byte) ([]UserOperation, common.Address, error) {
	var fir error

	var met []*bind.MetaData
	{
		met = []*bind.MetaData{
			AlchemyBindingMetaData,
			BiconomyBindingMetaData,
		}
	}

	for _, x := range met {
		var err error
		var ops []UserOperation
		var ben common.Address
		{
			ops, ben, err = decode(x, byt)
			if err != nil {
				if fir == nil {
					fir = err
				}

				continue
			}
		}

		return ops, ben, nil
	}

	return nil, common.Address{}, tracer.Mask(fir)
}

func decode(met *bind.MetaData, byt []byte) ([]UserOperation, common.Address, error) {
	var err error
	var exi bool

	var prs abi.ABI
	{
		prs, err = abi.JSON(strings.NewReader(met.ABI))
		if err != nil {
			return nil, common.Address{}, tracer.Mask(err)
		}
	}

	var mth abi.Method
	{
		mth, exi = prs.Methods[Method]
		if !exi {
			return nil, common.Address{}, tracer.Mask(fmt.Errorf("method %q not found in entrypoint calldata", Method))
		}
	}

	unp := map[string]interface{}{}
	{
		err = mth.Inputs.UnpackIntoMap(unp, byt[4:])
		if err != nil {
			return nil, common.Address{}, tracer.Mask(err)
		}
	}

	var ops []UserOperation
	{
		ops, err = decOps(unp, "ops")
		if err != nil {
			return nil, common.Address{}, tracer.Mask(err)
		}
	}

	var ben common.Address
	{
		ben, err = decAdd(unp, "beneficiary")
		if err != nil {
			return nil, common.Address{}, tracer.Mask(err)
		}
	}

	return ops, ben, nil
}

func decOps(unp map[string]interface{}, key string) ([]UserOperation, error) {
	var err error

	var val interface{}
	var exi bool
	{
		val, exi = unp[key]
		if !exi {
			return nil, tracer.Mask(fmt.Errorf("key %q not found in unpacked input map", key))
		}
	}

	// Note that we are marshalling and unmarshalling the available interface{}
	// type in order to receive a proper []UserOperation type. There wasn't any
	// other better way to cast, convert, decode or parse the underlying structure
	// as provided by the ABI specific Go bindings, because the underlying type of
	// the available interface{} is some arbitrary anonymous inline struct that
	// cannot be easily type asserted for some obscure reason.

	var byt []byte
	{
		byt, err = json.Marshal(val)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var ops []UserOperation
	{
		err = json.Unmarshal(byt, &ops)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return ops, nil
}

func decAdd(unp map[string]interface{}, key string) (common.Address, error) {
	val, exi := unp[key]
	if !exi {
		return common.Address{}, tracer.Mask(fmt.Errorf("key %q not found in unpacked input map", key))
	}

	add, typ := val.(common.Address)
	if !typ {
		return common.Address{}, tracer.Mask(fmt.Errorf("value %#v is not of type common.Address", val))
	}

	return add, nil
}
