// Auto-generated by avdl-compiler v1.3.25 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/stellar1/common.avdl

package stellar1

import (
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type AccountID string

func (o AccountID) DeepCopy() AccountID {
	return o
}

type SecretKey string

func (o SecretKey) DeepCopy() SecretKey {
	return o
}

type TransactionID string

func (o TransactionID) DeepCopy() TransactionID {
	return o
}

type KeybaseTransactionID string

func (o KeybaseTransactionID) DeepCopy() KeybaseTransactionID {
	return o
}

type TimeMs int64

func (o TimeMs) DeepCopy() TimeMs {
	return o
}

type Hash []byte

func (o Hash) DeepCopy() Hash {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type KeybaseRequestID string

func (o KeybaseRequestID) DeepCopy() KeybaseRequestID {
	return o
}

type PaymentID struct {
	TxID TransactionID `codec:"txID" json:"txID"`
}

func (o PaymentID) DeepCopy() PaymentID {
	return PaymentID{
		TxID: o.TxID.DeepCopy(),
	}
}

type Asset struct {
	Type   string `codec:"type" json:"type"`
	Code   string `codec:"code" json:"code"`
	Issuer string `codec:"issuer" json:"issuer"`
}

func (o Asset) DeepCopy() Asset {
	return Asset{
		Type:   o.Type,
		Code:   o.Code,
		Issuer: o.Issuer,
	}
}

type Balance struct {
	Asset  Asset  `codec:"asset" json:"asset"`
	Amount string `codec:"amount" json:"amount"`
	Limit  string `codec:"limit" json:"limit"`
}

func (o Balance) DeepCopy() Balance {
	return Balance{
		Asset:  o.Asset.DeepCopy(),
		Amount: o.Amount,
		Limit:  o.Limit,
	}
}

type AccountReserve struct {
	Amount      string `codec:"amount" json:"amount"`
	Description string `codec:"description" json:"description"`
}

func (o AccountReserve) DeepCopy() AccountReserve {
	return AccountReserve{
		Amount:      o.Amount,
		Description: o.Description,
	}
}

type TransactionStatus int

const (
	TransactionStatus_NONE            TransactionStatus = 0
	TransactionStatus_PENDING         TransactionStatus = 1
	TransactionStatus_SUCCESS         TransactionStatus = 2
	TransactionStatus_ERROR_TRANSIENT TransactionStatus = 3
	TransactionStatus_ERROR_PERMANENT TransactionStatus = 4
)

func (o TransactionStatus) DeepCopy() TransactionStatus { return o }

var TransactionStatusMap = map[string]TransactionStatus{
	"NONE":            0,
	"PENDING":         1,
	"SUCCESS":         2,
	"ERROR_TRANSIENT": 3,
	"ERROR_PERMANENT": 4,
}

var TransactionStatusRevMap = map[TransactionStatus]string{
	0: "NONE",
	1: "PENDING",
	2: "SUCCESS",
	3: "ERROR_TRANSIENT",
	4: "ERROR_PERMANENT",
}

func (e TransactionStatus) String() string {
	if v, ok := TransactionStatusRevMap[e]; ok {
		return v
	}
	return ""
}

type RequestStatus int

const (
	RequestStatus_OK       RequestStatus = 0
	RequestStatus_CANCELED RequestStatus = 1
)

func (o RequestStatus) DeepCopy() RequestStatus { return o }

var RequestStatusMap = map[string]RequestStatus{
	"OK":       0,
	"CANCELED": 1,
}

var RequestStatusRevMap = map[RequestStatus]string{
	0: "OK",
	1: "CANCELED",
}

func (e RequestStatus) String() string {
	if v, ok := RequestStatusRevMap[e]; ok {
		return v
	}
	return ""
}

type PaymentStrategy int

const (
	PaymentStrategy_NONE   PaymentStrategy = 0
	PaymentStrategy_DIRECT PaymentStrategy = 1
	PaymentStrategy_RELAY  PaymentStrategy = 2
)

func (o PaymentStrategy) DeepCopy() PaymentStrategy { return o }

var PaymentStrategyMap = map[string]PaymentStrategy{
	"NONE":   0,
	"DIRECT": 1,
	"RELAY":  2,
}

var PaymentStrategyRevMap = map[PaymentStrategy]string{
	0: "NONE",
	1: "DIRECT",
	2: "RELAY",
}

func (e PaymentStrategy) String() string {
	if v, ok := PaymentStrategyRevMap[e]; ok {
		return v
	}
	return ""
}

type RelayDirection int

const (
	RelayDirection_CLAIM RelayDirection = 0
	RelayDirection_YANK  RelayDirection = 1
)

func (o RelayDirection) DeepCopy() RelayDirection { return o }

var RelayDirectionMap = map[string]RelayDirection{
	"CLAIM": 0,
	"YANK":  1,
}

var RelayDirectionRevMap = map[RelayDirection]string{
	0: "CLAIM",
	1: "YANK",
}

func (e RelayDirection) String() string {
	if v, ok := RelayDirectionRevMap[e]; ok {
		return v
	}
	return ""
}

type PaymentResult struct {
	KeybaseID KeybaseTransactionID `codec:"keybaseID" json:"keybaseID"`
	StellarID TransactionID        `codec:"stellarID" json:"stellarID"`
	Pending   bool                 `codec:"pending" json:"pending"`
}

func (o PaymentResult) DeepCopy() PaymentResult {
	return PaymentResult{
		KeybaseID: o.KeybaseID.DeepCopy(),
		StellarID: o.StellarID.DeepCopy(),
		Pending:   o.Pending,
	}
}

type RelayClaimResult struct {
	ClaimStellarID TransactionID `codec:"claimStellarID" json:"claimStellarID"`
}

func (o RelayClaimResult) DeepCopy() RelayClaimResult {
	return RelayClaimResult{
		ClaimStellarID: o.ClaimStellarID.DeepCopy(),
	}
}

type EncryptedNote struct {
	V         int               `codec:"v" json:"v"`
	E         []byte            `codec:"e" json:"e"`
	N         keybase1.BoxNonce `codec:"n" json:"n"`
	Sender    NoteRecipient     `codec:"sender" json:"sender"`
	Recipient *NoteRecipient    `codec:"recipient,omitempty" json:"recipient,omitempty"`
}

func (o EncryptedNote) DeepCopy() EncryptedNote {
	return EncryptedNote{
		V: o.V,
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		N:      o.N.DeepCopy(),
		Sender: o.Sender.DeepCopy(),
		Recipient: (func(x *NoteRecipient) *NoteRecipient {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Recipient),
	}
}

type NoteRecipient struct {
	User   keybase1.UserVersion          `codec:"user" json:"user"`
	PukGen keybase1.PerUserKeyGeneration `codec:"pukGen" json:"pukGen"`
}

func (o NoteRecipient) DeepCopy() NoteRecipient {
	return NoteRecipient{
		User:   o.User.DeepCopy(),
		PukGen: o.PukGen.DeepCopy(),
	}
}

type NoteContents struct {
	Note      string        `codec:"note" json:"note"`
	StellarID TransactionID `codec:"stellarID" json:"stellarID"`
}

func (o NoteContents) DeepCopy() NoteContents {
	return NoteContents{
		Note:      o.Note,
		StellarID: o.StellarID.DeepCopy(),
	}
}

type EncryptedRelaySecret struct {
	V   int                           `codec:"v" json:"v"`
	E   []byte                        `codec:"e" json:"e"`
	N   keybase1.BoxNonce             `codec:"n" json:"n"`
	Gen keybase1.PerTeamKeyGeneration `codec:"gen" json:"gen"`
}

func (o EncryptedRelaySecret) DeepCopy() EncryptedRelaySecret {
	return EncryptedRelaySecret{
		V: o.V,
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		N:   o.N.DeepCopy(),
		Gen: o.Gen.DeepCopy(),
	}
}

type RelayContents struct {
	StellarID TransactionID `codec:"stellarID" json:"stellarID"`
	Sk        SecretKey     `codec:"sk" json:"sk"`
	Note      string        `codec:"note" json:"note"`
}

func (o RelayContents) DeepCopy() RelayContents {
	return RelayContents{
		StellarID: o.StellarID.DeepCopy(),
		Sk:        o.Sk.DeepCopy(),
		Note:      o.Note,
	}
}

type OutsideCurrencyCode string

func (o OutsideCurrencyCode) DeepCopy() OutsideCurrencyCode {
	return o
}

type OutsideExchangeRate struct {
	Currency OutsideCurrencyCode `codec:"currency" json:"currency"`
	Rate     string              `codec:"rate" json:"rate"`
}

func (o OutsideExchangeRate) DeepCopy() OutsideExchangeRate {
	return OutsideExchangeRate{
		Currency: o.Currency.DeepCopy(),
		Rate:     o.Rate,
	}
}

type CurrencySymbol struct {
	Symbol    string `codec:"symbol" json:"str"`
	Ambigious bool   `codec:"ambigious" json:"ambigious"`
	Postfix   bool   `codec:"postfix" json:"postfix"`
}

func (o CurrencySymbol) DeepCopy() CurrencySymbol {
	return CurrencySymbol{
		Symbol:    o.Symbol,
		Ambigious: o.Ambigious,
		Postfix:   o.Postfix,
	}
}

type OutsideCurrencyDefinition struct {
	Name   string         `codec:"name" json:"name"`
	Symbol CurrencySymbol `codec:"symbol" json:"symbol"`
}

func (o OutsideCurrencyDefinition) DeepCopy() OutsideCurrencyDefinition {
	return OutsideCurrencyDefinition{
		Name:   o.Name,
		Symbol: o.Symbol.DeepCopy(),
	}
}

type StellarServerDefinitions struct {
	Revision   int                                               `codec:"revision" json:"revision"`
	Currencies map[OutsideCurrencyCode]OutsideCurrencyDefinition `codec:"currencies" json:"currencies"`
}

func (o StellarServerDefinitions) DeepCopy() StellarServerDefinitions {
	return StellarServerDefinitions{
		Revision: o.Revision,
		Currencies: (func(x map[OutsideCurrencyCode]OutsideCurrencyDefinition) map[OutsideCurrencyCode]OutsideCurrencyDefinition {
			if x == nil {
				return nil
			}
			ret := make(map[OutsideCurrencyCode]OutsideCurrencyDefinition, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Currencies),
	}
}

type PageCursor struct {
	HorizonCursor string `codec:"horizonCursor" json:"horizonCursor"`
	DirectCursor  string `codec:"directCursor" json:"directCursor"`
	RelayCursor   string `codec:"relayCursor" json:"relayCursor"`
}

func (o PageCursor) DeepCopy() PageCursor {
	return PageCursor{
		HorizonCursor: o.HorizonCursor,
		DirectCursor:  o.DirectCursor,
		RelayCursor:   o.RelayCursor,
	}
}

type CommonInterface interface {
}

func CommonProtocol(i CommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "stellar.1.common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type CommonClient struct {
	Cli rpc.GenericClient
}
