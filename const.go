package atgo

const (
	// 1. SMS APIs
	SMSBaseLiveURL = "https://api.africastalking.com/version1/messaging"
	SMSSandboxURL  = "https://api.sandbox.africastalking.com/version1/messaging"

	// 2. Voice APIs
	MakeCallLiveURL    = "https://voice.africastalking.com/call"
	MakeCallSandboxURL = "https://voice.sandbox.africastalking.com/call"

	CallTransferLiveURL    = "https://voice.africastalking.com/callTransfer"
	CallTransferSandboxURL = "https://voice.sandbox.africastalking.com/callTransfer"

	QueuedCallsLiveURL    = "https://voice.africastalking.com/queueStatus"
	QueuedCallsSandboxURL = "https://voice.sandbox.africastalking.com/queueStatus"

	MediaUploadLiveURL    = "https://voice.africastalking.com/mediaUpload"
	MediaUploadSandboxURL = "https://voice.sandbox.africastalking.com/mediaUpload"

	// 3. Airtime API
	AirtimeLiveURL    = "https://api.africastalking.com/version1/airtime/send"
	AirtimeSandboxURL = "https://api.sandbox.africastalking.com/version1/airtime/send"

	// 4. Payments APIs
	MobileCheckoutLiveURL    = "https://payments.africastalking.com/mobile/checkout/request"
	MobileCheckoutSandboxURL = "https://payments.sandbox.africastalking.com/mobile/checkout/request"

	MobileB2CLiveURL = "https://payments.africastalking.com/mobile/b2c/request"
	MobileB2CSandbox = "https://payments.sandbox.africastalking.com/mobile/b2c/request"

	MobileB2BLiveURL    = "https://payments.africastalking.com/mobile/b2b/request"
	MobileB2BSandboxURL = "https://payments.sandbox.africastalking.com/mobile/b2b/request"

	BankCheckoutLiveURL    = "https://payments.africastalking.com/bank/checkout/charge"
	BankCheckoutSandboxURL = "https://payments.sandbox.africastalking.com/bank/checkout/charge"

	BankCheckoutValidateLiveURL    = "https://payments.africastalking.com/bank/checkout/validate"
	BankCheckoutValidateSandboxURL = "https://payments.sandbox.africastalking.com/bank/checkout/validate"

	BankTransferLiveURL    = "https://payments.africastalking.com/bank/transfer"
	BankTransferSandboxURL = "https://payments.sandbox.africastalking.com/bank/transfer"

	CardCheckoutLiveURL    = "https://payments.africastalking.com/card/checkout/charge"
	CardCheckoutSandboxURL = "https://payments.sandbox.africastalking.com/card/checkout/charge"

	CardCheckoutValidateLiveURL    = "https://payments.africastalking.com/card/checkout/validate"
	CardCheckoutValidateSandboxURL = "https://payments.sandbox.africastalking.com/card/checkout/validate"

	WalletTransferLiveURL    = "https://payments.africastalking.com/transfer/wallet"
	WalletTransferSandboxURL = "https://payments.sandbox.africastalking.com/transfer/wallet"

	TopupStashLiveURL    = "https://payments.africastalking.com/topup/stash"
	TopupStashSandboxURL = "https://payments.sandbox.africastalking.com/topup/stash"

	FindTransactionLiveURL    = "https://payments.africastalking.com/query/transaction/find"
	FindTransactionSandboxURL = "https://payments.sandbox.africastalking.com/query/transaction/find"

	FetchProductTransactionsLiveURL    = "https://payments.africastalking.com/query/transaction/fetch"
	FetchProductTransactionsSandboxURL = "https://payments.sandbox.africastalking.com/query/transaction/fetch"

	FetchWalletTransactionsLiveURL    = "https://payments.africastalking.com/query/wallet/fetch"
	FetchWalletTransactionsSandboxURL = "https://payments.sandbox.africastalking.com/query/wallet/fetch"

	FetchWalletBalanceLiveURL    = "https://payments.africastalking.com/query/wallet/balance"
	FetchWalletBalanceSandboxURL = "https://payments.sandbox.africastalking.com/query/wallet/balance"

	// 5. IOT API
	IoTLiveURL = "https://iot.africastalking.com/data/publish"

	// 6. User API
	UserLiveURL    = "https://api.africastalking.com/version1/user"
	UserSandboxURL = "https://api.sandbox.africastalking.com/version1/user"
)
