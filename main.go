package main

/*
#cgo LDFLAGS: -lktoblzcheck
#include <ktoblzcheck.h>
*/
import "C"
import (
	"fmt"
	"strconv"
)

// LibraryVersion returns the value of ktoblzcheck's configuration variable VERSION as a string
func LibraryVersion() string {
	return C.GoString(C.AccountNumberCheck_libraryVersion())
}

// BankDataDir returns the directory where the bankdata file is stored
func BankDataDir() string {
	return C.GoString(C.AccountNumberCheck_bankdata_dir())
}

// StringEncoding returns the character encoding that is used when strings are returned
func StringEncoding() string {
	return C.GoString(C.AccountNumberCheck_stringEncoding())
}

// AccountNumberCheck wraps ktoblzcheck AccountNumberCheck instance
type AccountNumberCheck struct {
	ptr *C.struct_AccountNumberCheck
}

// CheckResult indicates a bank lookup result
type CheckResult int

const (
	// Ok means that account and bank match
	Ok CheckResult = iota
	// Unknown means that validation is not available because the validation algorithm is unknown/unimplemented in ktoblzcheck
	Unknown
	// Error means that the result of the validation algorithm is that the account and bank probably do not match
	Error
	// BankNotKnown indicates an unknown bankID
	BankNotKnown
)

// Check tests if bankId and accountId from a valid combination
func (check *AccountNumberCheck) Check(bankID, accountID string) CheckResult {
	return CheckResult(C.AccountNumberCheck_check(check.ptr, C.CString(bankID), C.CString(accountID)))
}

// NewDefaultAccountNumberCheck returns a new AccountNumberCheck instance initialized with the default ktoblzcheck data
func NewDefaultAccountNumberCheck() AccountNumberCheck {
	return AccountNumberCheck{
		ptr: C.AccountNumberCheck_new(),
	}
}

// NewAccountNumberCheck returns a new AccountNumberCheck instance initialized with a custom ktoblzcheck data
func NewAccountNumberCheck(dataDir string) AccountNumberCheck {
	return AccountNumberCheck{
		ptr: C.AccountNumberCheck_new_file(C.CString(dataDir)),
	}
}

// Free destructs the underlying ktoblzcheck AccountNumberCheck instance
func (check *AccountNumberCheck) Free() {
	C.AccountNumberCheck_delete(check.ptr)
}

// BankCount returns the number of bank-records currently loaded
func (check *AccountNumberCheck) BankCount() uint {
	return uint(C.AccountNumberCheck_bankCount(check.ptr))
}

// Record represents a bank as defined in ktoblzcheck
type Record struct {
	BankID   string
	Name     string
	Location string
}

// UnknownRecord is used to indicate that a record requested via FindBank is unknown
type UnknownRecord string

func (f UnknownRecord) Error() string {
	return fmt.Sprintf("unknown BLZ: %v", string(f))
}

// FindBank looks up a bank by its given bankId
func (check *AccountNumberCheck) FindBank(bankID string) (Record, error) {
	bank := C.AccountNumberCheck_findBank(check.ptr, C.CString(bankID))
	if bank == nil {
		return Record{}, UnknownRecord(bankID)
	}
	var record = Record{
		BankID:   strconv.Itoa(int(C.AccountNumberCheck_Record_bankId(bank))),
		Name:     C.GoString(C.AccountNumberCheck_Record_bankName(bank)),
		Location: C.GoString(C.AccountNumberCheck_Record_location(bank)),
	}
	// C.AccountNumberCheck_Record_delete(bank)
	return record, nil
}
