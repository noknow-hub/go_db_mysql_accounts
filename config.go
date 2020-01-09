//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package accounts

import (
    "time"
)

const (
    CHARSET = "utf8mb4"
    TABLE_NAME_ACCOUNTS = "accounts"
    TABLE_NAME_ACCOUNT_META = "account_meta"
    TABLE_NAME_ACCOUNT_IMAGES = "account_images"
    TABLE_NAME_ACCOUNT_IMAGE_MAP = "account_image_map"
    TABLE_NAME_ACCOUNT_COMPANIES = "account_companies"
    TABLE_NAME_ACCOUNT_COMPANY_MAP = "account_company_map"
    TABLE_NAME_ACCOUNT_COMPANY_IMAGES = "account_company_images"
    TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP = "account_company_image_map"
    COL_ACCOUNTS_ID = "id"
    COL_ACCOUNTS_FIRST_NAME = "first_name"
    COL_ACCOUNTS_LAST_NAME = "last_name"
    COL_ACCOUNTS_MIDDLE_NAME = "middle_name"
    COL_ACCOUNTS_NAME = "name"
    COL_ACCOUNTS_NICK_NAME = "nick_name"
    COL_ACCOUNTS_NATIONALITY = "nationality"
    COL_ACCOUNTS_EMAIL = "email"
    COL_ACCOUNTS_PHONE_NUMBER = "phone_number"
    COL_ACCOUNTS_PASSWORD = "password"
    COL_ACCOUNTS_STATUS = "status"
    COL_ACCOUNTS_PUBLISHABLE_TOKEN = "publishable_token"
    COL_ACCOUNTS_SECRET_TOKEN = "secret_token"
    COL_ACCOUNTS_ADDRESS_COUNTRY = "address_country"
    COL_ACCOUNTS_ADDRESS_CITY = "address_city"
    COL_ACCOUNTS_ADDRESS_ZIP_CODE = "address_zip_code"
    COL_ACCOUNTS_ADDRESS = "address"
    COL_ACCOUNTS_ADDRESS_OPTION = "address_option"
    COL_ACCOUNTS_CREATED_AT = "created_at"
    COL_ACCOUNTS_UPDATED_AT = "updated_at"
    COL_ACCOUNT_META_KEY = "meta_key"
    COL_ACCOUNT_META_VALUE = "meta_value"
    COL_ACCOUNT_IMAGES_ID = "image_id"
    COL_ACCOUNT_IMAGES_NAME = "image_name"
    COL_ACCOUNT_IMAGES_SIZE = "image_size"
    COL_ACCOUNT_IMAGES_PATH = "image_path"
    COL_ACCOUNT_IMAGES_URI = "image_uri"
    COL_ACCOUNT_IMAGES_ALT = "image_alt"
    COL_ACCOUNT_IMAGES_TYPE = "image_type"
    COL_ACCOUNT_IMAGES_LINK = "image_link"
    COL_ACCOUNT_IMAGES_MIME_TYPE = "image_mime_type"
    COL_ACCOUNT_IMAGES_DATE = "image_date"
    COL_ACCOUNT_COMPANIES_ID = "company_id"
    COL_ACCOUNT_COMPANIES_NAME = "company_name"
    COL_ACCOUNT_COMPANIES_NAME_EN = "company_name_en"
    COL_ACCOUNT_COMPANIES_FOUNDATION_DATE = "company_foundation_date"
    COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY = "company_address_country"
    COL_ACCOUNT_COMPANIES_ADDRESS_CITY = "company_address_city"
    COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE = "company_address_zip_code"
    COL_ACCOUNT_COMPANIES_ADDRESS = "company_address"
    COL_ACCOUNT_COMPANIES_ADDRESS_OPTION = "company_address_option"
    COL_ACCOUNT_COMPANIES_INDUSTRY = "company_industry"
    COL_ACCOUNT_COMPANIES_EMAIL = "company_email"
    COL_ACCOUNT_COMPANIES_PHONE_NUMBER = "company_phone_number"
    COL_ACCOUNT_COMPANIES_WEBSITE = "company_website"
    COL_ACCOUNT_COMPANIES_CREATED_AT = "company_created_at"
    COL_ACCOUNT_COMPANIES_UPDATED_AT = "company_updated_at"
    COL_ACCOUNT_COMPANY_IMAGES_ID = "company_image_id"
    COL_ACCOUNT_COMPANY_IMAGES_NAME = "company_image_name"
    COL_ACCOUNT_COMPANY_IMAGES_SIZE = "company_image_size"
    COL_ACCOUNT_COMPANY_IMAGES_PATH = "company_image_path"
    COL_ACCOUNT_COMPANY_IMAGES_URI = "company_image_uri"
    COL_ACCOUNT_COMPANY_IMAGES_ALT = "company_image_alt"
    COL_ACCOUNT_COMPANY_IMAGES_TYPE = "company_image_type"
    COL_ACCOUNT_COMPANY_IMAGES_LINK = "company_image_link"
    COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE = "company_image_mime_type"
    COL_ACCOUNT_COMPANY_IMAGES_DATE = "company_image_date"
)


type AccountsColumns struct {
    ID string
    FirstName string
    LastName string
    MiddleName string
    Name string
    NickName string
    Nationality string
    Email string
    PhoneNumber string
    Password string
    Status string
    PublishableToken string
    SecretToken string
    AddressCountry string
    AddressCity string
    AddressZipCode string
    Address string
    AddressOption string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type SelectAccountsColumns struct {
    ID []string
    FirstName []string
    LastName []string
    MiddleName []string
    Name []string
    NickName []string
    Nationality []string
    Email []string
    PhoneNumber []string
    Password []string
    Status []string
    PublishableToken []string
    SecretToken []string
    AddressCountry []string
    AddressCity []string
    AddressZipCode []string
    Address []string
    AddressOption []string
    CreatedAtBefore time.Time
    CreatedAtAfter time.Time
    UpdatedAtBefore time.Time
    UpdatedAtAfter time.Time
}

type UpdateAccountsColumns struct {
    Values AccountsColumns
    Where AccountsColumns
}

type AccountMetaColumns struct {
    AccountID string
    MetaKey string
    MetaValue string
}

type SelectAccountMetaColumns struct {
    AccountID []string
    MetaKey []string
    MetaValue []string
}

type UpdateAccountMetaColumns struct {
    Values AccountMetaColumns
    Where AccountMetaColumns
}

type AccountImagesColumns struct {
    ID string
    Name string
    Size string
    Path string
    URI string
    Alt string
    Type string
    Link string
    MimeType string
    Date time.Time
}

type SelectAccountImagesColumns struct {
    ID []string
    Name []string
    SizeOrMore string
    SizeOrLess string
    Path []string
    URI []string
    Alt []string
    Type []string
    Link []string
    MimeType []string
    DateBefore time.Time
    DateAfter time.Time
}

type UpdateAccountImagesColumns struct {
    Values AccountImagesColumns
    Where AccountImagesColumns
}

type AccountImageMapColumns struct {
    AccountID string
    ImageID string
}

type SelectAccountImageMapColumns struct {
    AccountID []string
    ImageID []string
}

type UpdateAccountImageMapColumns struct {
    Values AccountImageMapColumns
    Where AccountImageMapColumns
}

type AccountCompaniesColumns struct {
    ID string
    Name string
    NameEn string
    FoundationDate time.Time
    AddressCountry string
    AddressCity string
    AddressZipCode string
    Address string
    AddressOption string
    Industry string
    Email string
    PhoneNumber string
    Website string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type SelectAccountCompaniesColumns struct {
    ID []string
    Name []string
    NameEn []string
    FoundationDate time.Time
    AddressCountry []string
    AddressCity []string
    AddressZipCode []string
    Address []string
    AddressOption []string
    Industry []string
    Email []string
    PhoneNumber []string
    Website []string
    CreatedAtBefore time.Time
    CreatedAtAfter time.Time
    UpdatedAtBefore time.Time
    UpdatedAtAfter time.Time
}

type UpdateAccountCompaniesColumns struct {
    Values AccountCompaniesColumns
    Where AccountCompaniesColumns
}

type AccountCompanyMapColumns struct {
    AccountID string
    CompanyID string
}

type SelectAccountCompanyMapColumns struct {
    AccountID []string
    CompanyID []string
}

type UpdateAccountCompanyMapColumns struct {
    Values AccountCompanyMapColumns
    Where AccountCompanyMapColumns
}

type AccountCompanyImagesColumns struct {
    ID string
    Name string
    Size string
    Path string
    URI string
    Alt string
    Type string
    Link string
    MimeType string
    Date time.Time
}

type SelectAccountCompanyImagesColumns struct {
    ID []string
    Name []string
    SizeOrMore string
    SizeOrLess string
    Path []string
    URI []string
    Alt []string
    Type []string
    Link []string
    MimeType []string
    DateBefore time.Time
    DateAfter time.Time
}

type UpdateAccountCompanyImagesColumns struct {
    Values AccountCompanyImagesColumns
    Where AccountCompanyImagesColumns
}

type AccountCompanyImageMapColumns struct {
    CompanyID string
    CompanyImageID string
}

type SelectAccountCompanyImageMapColumns struct {
    CompanyID []string
    CompanyImageID []string
}

type UpdateAccountCompanyImageMapColumns struct {
    Values AccountCompanyImageMapColumns
    Where AccountCompanyImageMapColumns
}

