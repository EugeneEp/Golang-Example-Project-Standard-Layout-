package user

const (
	productPath    = `SOFTWARE\ProductName`
	Path           = productPath + `\` + Key
	Key            = `USERS`
	CollectionName = `Users`

	FieldDisplayName   = `Display Name`
	DefaultDisplayName = ``

	FieldCreatedAt   = `Created At`
	DefaultCreatedAt = 0
)
