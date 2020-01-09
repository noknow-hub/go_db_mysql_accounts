//////////////////////////////////////////////////////////////////////
// insert.go
//////////////////////////////////////////////////////////////////////
package accounts

import (
    "log"
    "strconv"
    "time"
)


//////////////////////////////////////////////////////////////////////
// Insert into accounts table.
//////////////////////////////////////////////////////////////////////
func (c *AccountsColumns) Insert() (*AccountsColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // ID
    queryCols = append(queryCols, COL_ACCOUNTS_ID...)
    queryVals = append(queryVals, "?"...)
    if c.ID != "" {
        args = append(args, c.ID)
    } else {
        id := time.Now().UnixNano()
        for {
            isUnique, err := IsUniqueAccountsID(strconv.FormatInt(id, 10))
            if err != nil {
                return c, err
            }
            if isUnique {
                break
            }
            id = id + 1
        }
        args = append(args, id)
        c.ID = strconv.FormatInt(id, 10)
    }
    // FirstName
    if c.FirstName != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_FIRST_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.FirstName)
    }
    // LastName
    if c.LastName != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_LAST_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.LastName)
    }
    // MiddleName
    if c.MiddleName != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_MIDDLE_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.MiddleName)
    }
    // Name
    if c.Name != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Name)
    }
    // NickName
    if c.NickName != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_NICK_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.NickName)
    }
    // Nationality
    if c.Nationality != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_NATIONALITY...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Nationality)
    }
    // Email
    if c.Email != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_EMAIL...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Email)
    }
    // PhoneNumber
    if c.PhoneNumber != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_PHONE_NUMBER...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.PhoneNumber)
    }
    // Password
    if c.Password != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_PASSWORD...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Password)
    }
    // Status
    if c.Status != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_STATUS...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Status)
    }
    // PublishableToken
    if c.PublishableToken != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_PUBLISHABLE_TOKEN...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.PublishableToken)
    }
    // SecretToken
    if c.SecretToken != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_SECRET_TOKEN...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.SecretToken)
    }
    // AddressCountry
    if c.AddressCountry != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_ADDRESS_COUNTRY...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressCountry)
    }
    // AddressCity
    if c.AddressCity != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_ADDRESS_CITY...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressCity)
    }
    // AddressZipCode
    if c.AddressZipCode != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_ADDRESS_ZIP_CODE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressZipCode)
    }
    // Address
    if c.Address != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_ADDRESS...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Address)
    }
    // AddressOption
    if c.AddressOption != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNTS_ADDRESS_OPTION...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressOption)
    }
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNTS + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}


//////////////////////////////////////////////////////////////////////
// Insert into account_meta table.
//////////////////////////////////////////////////////////////////////
func (c *AccountMetaColumns) Insert() (*AccountMetaColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // AccountID
    queryCols = append(queryCols, COL_ACCOUNTS_ID...)
    queryVals = append(queryVals, "?"...)
    args = append(args, c.AccountID)
    // MetaKey
    queryCols = append(queryCols, "," + COL_ACCOUNT_META_KEY...)
    queryVals = append(queryVals, ",?"...)
    args = append(args, c.MetaKey)
    // MetaValue
    queryCols = append(queryCols, "," + COL_ACCOUNT_META_VALUE...)
    queryVals = append(queryVals, ",?"...)
    args = append(args, c.MetaValue)
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNT_META + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}


//////////////////////////////////////////////////////////////////////
// Insert into account_images table.
//////////////////////////////////////////////////////////////////////
func (c *AccountImagesColumns) Insert() (*AccountImagesColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // ID
    queryCols = append(queryCols, COL_ACCOUNT_IMAGES_ID...)
    queryVals = append(queryVals, "?"...)
    if c.ID != "" {
        args = append(args, c.ID)
    } else {
        id := time.Now().UnixNano()
        for {
            isUnique, err := IsUniqueAccountImagesID(strconv.FormatInt(id, 10))
            if err != nil {
                return c, err
            }
            if isUnique {
                break
            }
            id = id + 1
        }
        args = append(args, id)
        c.ID = strconv.FormatInt(id, 10)
    }
    // Name
    if c.Name != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Name)
    }
    // Size
    if c.Size != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_SIZE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Size)
    }
    // Path
    if c.Path != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_PATH...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Path)
    }
    // URI
    if c.URI != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_URI...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.URI)
    }
    // Alt
    if c.Alt != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_ALT...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Alt)
    }
    // Type
    if c.Type != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_TYPE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Type)
    }
    // Link
    if c.Link != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_LINK...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Link)
    }
    // MimeType
    if c.MimeType != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_MIME_TYPE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.MimeType)
    }
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNT_IMAGES + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}


//////////////////////////////////////////////////////////////////////
// Insert into account_image_map table.
//////////////////////////////////////////////////////////////////////
func (c *AccountImageMapColumns) Insert() (*AccountImageMapColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // AccountID
    queryCols = append(queryCols, COL_ACCOUNTS_ID...)
    queryVals = append(queryVals, "?"...)
    args = append(args, c.AccountID)
    // ImageID
    queryCols = append(queryCols, "," + COL_ACCOUNT_IMAGES_ID...)
    queryVals = append(queryVals, ",?"...)
    args = append(args, c.ImageID)
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNT_IMAGE_MAP + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}


//////////////////////////////////////////////////////////////////////
// Insert into account_companies table.
//////////////////////////////////////////////////////////////////////
func (c *AccountCompaniesColumns) Insert() (*AccountCompaniesColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // ID
    queryCols = append(queryCols, COL_ACCOUNT_COMPANIES_ID...)
    queryVals = append(queryVals, "?"...)
    if c.ID != "" {
        args = append(args, c.ID)
    } else {
        id := time.Now().UnixNano()
        for {
            isUnique, err := IsUniqueAccountCompaniesID(strconv.FormatInt(id, 10))
            if err != nil {
                return c, err
            }
            if isUnique {
                break
            }
            id = id + 1
        }
        args = append(args, id)
        c.ID = strconv.FormatInt(id, 10)
    }
    // Name
    if c.Name != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Name)
    }
    // NameEn
    if c.NameEn != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_NAME_EN...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.NameEn)
    }
    // FoundationDate
    if !c.FoundationDate.IsZero() {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_FOUNDATION_DATE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.FoundationDate)
    }
    // AddressCountry
    if c.AddressCountry != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressCountry)
    }
    // AddressCity
    if c.AddressCity != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_ADDRESS_CITY...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressCity)
    }
    // AddressZipCode
    if c.AddressZipCode != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressZipCode)
    }
    // Address
    if c.Address != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_ADDRESS...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Address)
    }
    // AddressOption
    if c.AddressOption != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_ADDRESS_OPTION...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.AddressOption)
    }
    // Industry
    if c.Industry != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_INDUSTRY...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Industry)
    }
    // Email
    if c.Email != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_EMAIL...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Email)
    }
    // PhoneNumber
    if c.PhoneNumber != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_PHONE_NUMBER...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.PhoneNumber)
    }
    // Website
    if c.Website != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_WEBSITE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Website)
    }
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNT_COMPANIES + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}


//////////////////////////////////////////////////////////////////////
// Insert into account_company_map table.
//////////////////////////////////////////////////////////////////////
func (c *AccountCompanyMapColumns) Insert() (*AccountCompanyMapColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // AccountID
    queryCols = append(queryCols, COL_ACCOUNTS_ID...)
    queryVals = append(queryVals, "?"...)
    args = append(args, c.AccountID)
    // CompanyID
    queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANIES_ID...)
    queryVals = append(queryVals, ",?"...)
    args = append(args, c.CompanyID)
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNT_COMPANY_MAP + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}


//////////////////////////////////////////////////////////////////////
// Insert into account_company_images table.
//////////////////////////////////////////////////////////////////////
func (c *AccountCompanyImagesColumns) Insert() (*AccountCompanyImagesColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // ID
    queryCols = append(queryCols, COL_ACCOUNT_COMPANY_IMAGES_ID...)
    queryVals = append(queryVals, "?"...)
    if c.ID != "" {
        args = append(args, c.ID)
    } else {
        id := time.Now().UnixNano()
        for {
            isUnique, err := IsUniqueAccountCompanyImagesID(strconv.FormatInt(id, 10))
            if err != nil {
                return c, err
            }
            if isUnique {
                break
            }
            id = id + 1
        }
        args = append(args, id)
        c.ID = strconv.FormatInt(id, 10)
    }
    // Name
    if c.Name != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_NAME...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Name)
    }
    // Size
    if c.Size != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_SIZE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Size)
    }
    // Path
    if c.Path != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_PATH...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Path)
    }
    // URI
    if c.URI != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_URI...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.URI)
    }
    // Alt
    if c.Alt != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_ALT...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Alt)
    }
    // Type
    if c.Type != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_TYPE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Type)
    }
    // Link
    if c.Link != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_LINK...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.Link)
    }
    // MimeType
    if c.MimeType != "" {
        queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE...)
        queryVals = append(queryVals, ",?"...)
        args = append(args, c.MimeType)
    }
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNT_COMPANY_IMAGES + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}


//////////////////////////////////////////////////////////////////////
// Insert into account_company_image_map table.
//////////////////////////////////////////////////////////////////////
func (c *AccountCompanyImageMapColumns) Insert() (*AccountCompanyImageMapColumns, error) {
    queryCols := make([]byte, 0)
    queryVals := make([]byte, 0)
    args := make([]interface{}, 0)
    // CompanyID
    queryCols = append(queryCols, COL_ACCOUNT_COMPANIES_ID...)
    queryVals = append(queryVals, "?"...)
    args = append(args, c.CompanyID)
    // CompanyImageID
    queryCols = append(queryCols, "," + COL_ACCOUNT_COMPANY_IMAGES_ID...)
    queryVals = append(queryVals, ",?"...)
    args = append(args, c.CompanyImageID)
    // SQL
    query := "INSERT INTO " + TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP + "(" + string(queryCols[:]) + ") VALUES (" + string(queryVals[:]) + ")"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return c, err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return c, err
    }
    return c, nil
}



