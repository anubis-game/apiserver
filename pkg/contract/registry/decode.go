package registry

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

// Decode takes the input data of of a types.Transaction object and returns
// parameters used to conduct the associated contract write on the
// Registry.request function based on the ABI below.
//
//     function request(address grd, uint64 tim, address wal, bytes memory sgn) public
//

func (r *Registry) Decode(byt []byte) (common.Address, time.Time, common.Address, []byte, error) {
	var err error

	var reg abi.ABI
	{
		reg, err = abi.JSON(strings.NewReader(RegistryBindingABI))
		if err != nil {
			return common.Address{}, time.Time{}, common.Address{}, nil, err
		}
	}

	var met *abi.Method
	{
		met, err = reg.MethodById(byt[:4])
		if err != nil {
			return common.Address{}, time.Time{}, common.Address{}, nil, err
		}
	}

	unp := map[string]interface{}{}
	{
		err = met.Inputs.UnpackIntoMap(unp, byt[4:])
		if err != nil {
			return common.Address{}, time.Time{}, common.Address{}, nil, err
		}
	}

	var grd common.Address
	{
		grd, err = decAdd(unp, "grd")
		if err != nil {
			return common.Address{}, time.Time{}, common.Address{}, nil, err
		}
	}

	var tim time.Time
	{
		tim, err = decTim(unp, "tim")
		if err != nil {
			return common.Address{}, time.Time{}, common.Address{}, nil, err
		}
	}

	var wal common.Address
	{
		wal, err = decAdd(unp, "wal")
		if err != nil {
			return common.Address{}, time.Time{}, common.Address{}, nil, err
		}
	}

	var sgn []byte
	{
		sgn, err = decByt(unp, "sgn")
		if err != nil {
			return common.Address{}, time.Time{}, common.Address{}, nil, err
		}
	}

	return grd, tim, wal, sgn, nil
}

func decByt(unp map[string]interface{}, key string) ([]byte, error) {
	val, exi := unp[key]
	if !exi {
		return nil, tracer.Mask(fmt.Errorf("key %q not found in unpacked input map", key))
	}

	byt, typ := val.([]byte)
	if !typ {
		return nil, tracer.Mask(fmt.Errorf("value %#v is not of type []byte", val))
	}

	return byt, nil
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

func decTim(unp map[string]interface{}, key string) (time.Time, error) {
	val, exi := unp[key]
	if !exi {
		return time.Time{}, tracer.Mask(fmt.Errorf("key %q not found in unpacked input map", key))
	}

	unt, typ := val.(uint64)
	if !typ {
		return time.Time{}, tracer.Mask(fmt.Errorf("value %#v is not of type uint64", val))
	}

	return time.Unix(int64(unt), 0).UTC(), nil
}
